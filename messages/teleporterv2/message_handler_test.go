// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporterv2

import (
	"context"
	"errors"
	"testing"

	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/crypto/bls/signer/localsigner"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	merkleregistry "github.com/ava-labs/icm-services/abi-bindings/go/MerkleValidatorSetRegistry"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/relayer/validatorupdater"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func newTestMerkleHandler(t *testing.T, ctrl *gomock.Controller) (*messageHandler, *handlerMocks) {
	t.Helper()
	handler, handlerMocks := newTestHandler(t, ctrl, VerifierTypeMerkle)
	merkleHandler, ok := handler.(*messageHandler)
	require.True(t, ok)
	return merkleHandler, handlerMocks
}

// testMerkleSignature signs the handler's message with a single-validator set and returns the
// signed message together with that set, which is what SendMessage builds the attestation from.
func testMerkleSignature(
	t *testing.T,
	handler *messageHandler,
) (*warp.Message, []*validatorupdater.Validator) {
	t.Helper()
	signer, err := localsigner.New()
	require.NoError(t, err)
	sig, err := signer.Sign(handler.unsignedMessage.Bytes())
	require.NoError(t, err)
	signedMessage, err := warp.NewMessage(handler.unsignedMessage, &warp.BitSetSignature{
		Signers:   set.NewBits(0).Bytes(),
		Signature: [bls.SignatureLen]byte(bls.SignatureToBytes(sig)),
	})
	require.NoError(t, err)
	validators := []*validatorupdater.Validator{
		{
			UncompressedPublicKeyBytes: [96]byte(signer.PublicKey().Serialize()),
			Weight:                     10,
		},
	}
	return signedMessage, validators
}

// testCommitment is the registry commitment the test attestations are built against.
func testCommitment() merkleregistry.ValidatorSetMerkleCommitment {
	return merkleregistry.ValidatorSetMerkleCommitment{
		AvalancheBlockchainID: testSourceBlockchainID,
		Root:                  [32]byte{0xaa},
		TotalWeight:           10,
		PChainHeight:          100,
		PChainTimestamp:       1_700_000_000,
	}
}

// Gas estimation executes the delivery, so a revert there is the delivery reverting. Against an
// unchanged commitment it is reproducible and the message must be abandoned as non-retryable
// without broadcasting: no SendTx expectation is registered, so the mock fails the test if the
// handler falls back to sending at the block gas limit.
func TestMerkleSendMessageEstimationRevertIsNonRetryable(t *testing.T) {
	ctrl := gomock.NewController(t)
	handler, handlerMocks := newTestMerkleHandler(t, ctrl)
	signedMessage, validators := testMerkleSignature(t, handler)
	commitment := testCommitment()

	handlerMocks.ethClient.EXPECT().
		EstimateGas(gomock.Any(), gomock.Any()).
		Return(uint64(0), panicRevert(t, 0x11)).
		Times(1)
	handlerMocks.routeCalls(t, callRoutes{delivered: false, commitment: &commitment})

	_, err := handler.SendMessage(context.Background(), signedMessage, validators, commitment)
	require.ErrorIs(t, err, messages.ErrDeliveryReverts)
	require.ErrorIs(t, err, messages.ErrNonRetryable)
	require.ErrorContains(t, err, "arithmetic underflow or overflow")
}

// If the registry moved to a newer commitment while the message was being signed, the attestation
// is merely stale: the failure stays retryable so that the next attempt rebuilds it, and nothing
// is broadcast in the meantime.
func TestMerkleSendMessageEstimationRevertAfterCommitmentChangeIsRetryable(t *testing.T) {
	ctrl := gomock.NewController(t)
	handler, handlerMocks := newTestMerkleHandler(t, ctrl)
	signedMessage, validators := testMerkleSignature(t, handler)
	commitment := testCommitment()
	current := commitment
	current.Root = [32]byte{0xbb}
	current.PChainHeight++

	handlerMocks.ethClient.EXPECT().
		EstimateGas(gomock.Any(), gomock.Any()).
		Return(uint64(0), revertWithReason(t, "TeleporterMessenger: message verification failed")).
		Times(1)
	handlerMocks.routeCalls(t, callRoutes{delivered: false, commitment: &current})

	_, err := handler.SendMessage(context.Background(), signedMessage, validators, commitment)
	require.Error(t, err)
	require.NotErrorIs(t, err, messages.ErrNonRetryable)
	require.ErrorContains(t, err, "commitment changed")
}

// A revert during estimation because another relayer delivered the message first is not a
// failure: there is nothing left to send.
func TestMerkleSendMessageEstimationRevertAlreadyDelivered(t *testing.T) {
	ctrl := gomock.NewController(t)
	handler, handlerMocks := newTestMerkleHandler(t, ctrl)
	signedMessage, validators := testMerkleSignature(t, handler)
	commitment := testCommitment()

	handlerMocks.ethClient.EXPECT().
		EstimateGas(gomock.Any(), gomock.Any()).
		Return(uint64(0), revertWithReason(t, "TeleporterMessenger: message already received")).
		Times(1)
	handlerMocks.routeCalls(t, callRoutes{delivered: true, commitment: &commitment})

	txHash, err := handler.SendMessage(context.Background(), signedMessage, validators, commitment)
	require.NoError(t, err)
	require.Equal(t, common.Hash{}, txHash)
}

// An estimation that could not be run says nothing about the delivery. It used to fall back to
// broadcasting at the block gas limit; now the failure stays retryable and nothing is sent.
func TestMerkleSendMessageEstimationFailureIsNotBroadcast(t *testing.T) {
	ctrl := gomock.NewController(t)
	handler, handlerMocks := newTestMerkleHandler(t, ctrl)
	signedMessage, validators := testMerkleSignature(t, handler)
	commitment := testCommitment()

	handlerMocks.ethClient.EXPECT().
		EstimateGas(gomock.Any(), gomock.Any()).
		Return(uint64(0), errors.New("connection refused")).
		Times(1)

	_, err := handler.SendMessage(context.Background(), signedMessage, validators, commitment)
	require.Error(t, err)
	require.NotErrorIs(t, err, messages.ErrNonRetryable)
}

// A successful estimate is sent with a safety buffer, capped at the block gas limit.
func TestMerkleSendMessageGasLimitIsBufferedAndCapped(t *testing.T) {
	tests := []struct {
		name             string
		estimated        uint64
		expectedGasLimit uint64
	}{
		{name: "buffered", estimated: 1_000_000, expectedGasLimit: 1_250_000},
		{name: "capped at block gas limit", estimated: 14_000_000, expectedGasLimit: testBlockGasLimit},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			handler, handlerMocks := newTestMerkleHandler(t, ctrl)
			signedMessage, validators := testMerkleSignature(t, handler)

			handlerMocks.ethClient.EXPECT().
				EstimateGas(gomock.Any(), gomock.Any()).
				Return(tt.estimated, nil).
				Times(1)
			expectedTxHash := common.HexToHash("0x01")
			handlerMocks.destinationClient.EXPECT().
				SendTx(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				DoAndReturn(func(
					_ logging.Logger,
					accessList types.AccessList,
					_ set.Set[common.Address],
					toAddress common.Address,
					gasLimit uint64,
					_ []byte,
				) (*types.Receipt, error) {
					require.Nil(t, accessList)
					require.Equal(t, testTeleporterAddress, toAddress)
					require.Equal(t, tt.expectedGasLimit, gasLimit)
					return &types.Receipt{Status: types.ReceiptStatusSuccessful, TxHash: expectedTxHash}, nil
				}).
				Times(1)

			txHash, err := handler.SendMessage(context.Background(), signedMessage, validators, testCommitment())
			require.NoError(t, err)
			require.Equal(t, expectedTxHash, txHash)
		})
	}
}
