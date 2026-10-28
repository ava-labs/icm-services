// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporterv2

import (
	"context"
	"math/big"
	"testing"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/evm/predicate"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	warpadapter "github.com/ava-labs/icm-services/abi-bindings/go/teleporterV2/WarpAdapter"
	gasUtils "github.com/ava-labs/icm-services/icm-contracts/utils/gas-utils"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/icm-services/vms/evm"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func newTestWarpHandler(t *testing.T, ctrl *gomock.Controller) (*warpMessageHandler, *handlerMocks) {
	t.Helper()
	handler, handlerMocks := newTestHandler(t, ctrl, VerifierTypeWarp)
	warpHandler, ok := handler.(*warpMessageHandler)
	require.True(t, ok)
	return warpHandler, handlerMocks
}

// testSignedWarpMessage signs the handler's message with an empty signature, which is all the
// warp path needs to build the delivery: the signature is verified on chain, not by the relayer.
func testSignedWarpMessage(t *testing.T, handler *warpMessageHandler) *warp.Message {
	t.Helper()
	signedMessage, err := warp.NewMessage(handler.unsignedMessage, &warp.BitSetSignature{})
	require.NoError(t, err)
	return signedMessage
}

// expectedWarpDelivery is the gas limit and calldata SendMessage computes for delivering
// [signedMessage].
func expectedWarpDelivery(
	t *testing.T,
	handler *warpMessageHandler,
	signedMessage *warp.Message,
) (uint64, []byte) {
	t.Helper()
	numSigners, err := signedMessage.Signature.NumSigners()
	require.NoError(t, err)
	gasLimit, err := gasUtils.CalculateReceiveMessageGasLimit(
		numSigners,
		handler.teleporterMessage.RequiredGasLimit,
		len(predicate.New(signedMessage.Bytes())),
		len(signedMessage.Payload),
		len(handler.teleporterMessage.Receipts),
	)
	require.NoError(t, err)
	callData, err := warpadapter.PackReceiveCrossChainMessage(
		*handler.teleporterMessage,
		signedMessage.SourceChainID,
		0,
		testRewardAddress,
	)
	require.NoError(t, err)
	return gasLimit, callData
}

// The handler must simulate, and then broadcast, the delivery of the signed message to the
// TeleporterMessengerV2: the relayer as sender, the computed gas limit, the WarpAdapter's
// receiveCrossChainMessage calldata and the signed message as the predicate access list.
func TestWarpSendMessageSimulatesThenSends(t *testing.T) {
	ctrl := gomock.NewController(t)
	handler, handlerMocks := newTestWarpHandler(t, ctrl)
	signedMessage := testSignedWarpMessage(t, handler)
	expectedGasLimit, expectedCallData := expectedWarpDelivery(t, handler, signedMessage)
	expectedAccessList := utils.SignedWarpMessageToAccessList(signedMessage)

	var simulated ethereum.CallMsg
	handlerMocks.routeCalls(t, callRoutes{
		simulate: func(call ethereum.CallMsg) ([]byte, error) {
			simulated = call
			return []byte{}, nil
		},
	})
	expectedTxHash := common.HexToHash("0x01")
	handlerMocks.destinationClient.EXPECT().
		SendTx(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(
			_ logging.Logger,
			accessList types.AccessList,
			deliverers set.Set[common.Address],
			toAddress common.Address,
			gasLimit uint64,
			callData []byte,
		) (*types.Receipt, error) {
			require.Equal(t, expectedAccessList, accessList)
			require.Equal(t, set.Of(testRelayerAddress), deliverers)
			require.Equal(t, testTeleporterAddress, toAddress)
			require.Equal(t, expectedGasLimit, gasLimit)
			require.Equal(t, expectedCallData, callData)
			return &types.Receipt{Status: types.ReceiptStatusSuccessful, TxHash: expectedTxHash}, nil
		}).
		Times(1)

	txHash, err := handler.SendMessage(context.Background(), signedMessage)
	require.NoError(t, err)
	require.Equal(t, expectedTxHash, txHash)

	require.Equal(t, testRelayerAddress, simulated.From)
	require.Equal(t, testTeleporterAddress, *simulated.To)
	require.Equal(t, expectedGasLimit, simulated.Gas)
	require.Equal(t, expectedCallData, simulated.Data)
	require.Equal(t, expectedAccessList, simulated.AccessList)
}

// A delivery whose computed gas exceeds the destination block gas limit is abandoned as
// non-retryable before anything is simulated or sent. The message's own required gas fits the
// block here; the verification overhead added on top of it is what does not. No CallContract or
// SendTx expectation is registered, so the mock fails the test on either.
func TestWarpSendMessageExceedingBlockGasLimitIsNonRetryable(t *testing.T) {
	ctrl := gomock.NewController(t)
	handler, _ := newTestWarpHandler(t, ctrl)
	handler.teleporterMessage.RequiredGasLimit = new(big.Int).SetUint64(testBlockGasLimit)
	signedMessage := testSignedWarpMessage(t, handler)

	_, err := handler.SendMessage(context.Background(), signedMessage)
	require.ErrorIs(t, err, messages.ErrDeliveryExceedsBlockGasLimit)
	require.ErrorIs(t, err, messages.ErrNonRetryable)
}

// A delivery that reverts in simulation, such as one carrying a receipt the messenger cannot
// account for, is abandoned as non-retryable without being broadcast: no SendTx expectation is
// registered, so the mock fails the test if the handler sends it anyway.
func TestWarpSendMessageSimulatedRevertIsNonRetryable(t *testing.T) {
	ctrl := gomock.NewController(t)
	handler, handlerMocks := newTestWarpHandler(t, ctrl)
	signedMessage := testSignedWarpMessage(t, handler)

	handlerMocks.routeCalls(t, callRoutes{
		delivered: false,
		simulate: func(ethereum.CallMsg) ([]byte, error) {
			return nil, &evm.RevertError{Reason: "arithmetic underflow or overflow"}
		},
	})

	_, err := handler.SendMessage(context.Background(), signedMessage)
	require.ErrorIs(t, err, messages.ErrDeliveryReverts)
	require.ErrorIs(t, err, messages.ErrNonRetryable)
	require.ErrorContains(t, err, "arithmetic underflow or overflow")
}
