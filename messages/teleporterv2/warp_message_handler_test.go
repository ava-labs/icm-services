// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporterv2

import (
	"context"
	"errors"
	"testing"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/evm/predicate"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	gasUtils "github.com/ava-labs/icm-services/icm-contracts/utils/gas-utils"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/utils"
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

// expectedWarpGasLimit is the gas limit SendMessage computes for delivering [signedMessage].
func expectedWarpGasLimit(t *testing.T, handler *warpMessageHandler, signedMessage *warp.Message) uint64 {
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
	return gasLimit
}

// A delivery that reverts in simulation, such as one carrying a receipt the messenger cannot
// account for, is reproducible from the message itself. It must be abandoned as non-retryable
// without broadcasting anything: no SendTx expectation is registered, so the mock fails the test
// if the handler sends the transaction anyway.
func TestWarpSendMessageSimulatedRevertIsNonRetryable(t *testing.T) {
	ctrl := gomock.NewController(t)
	handler, handlerMocks := newTestWarpHandler(t, ctrl)
	signedMessage := testSignedWarpMessage(t, handler)

	var simulated []ethereum.CallMsg
	handlerMocks.routeCalls(t, callRoutes{
		delivered: false,
		simulate: func(call ethereum.CallMsg) ([]byte, error) {
			simulated = append(simulated, call)
			return nil, panicRevert(t, 0x11)
		},
	})

	_, err := handler.SendMessage(context.Background(), signedMessage)
	require.ErrorIs(t, err, messages.ErrDeliveryReverts)
	require.ErrorIs(t, err, messages.ErrNonRetryable)
	require.ErrorContains(t, err, "arithmetic underflow or overflow")

	// The simulation must be the transaction that would have been sent: same sender, target,
	// gas limit and signed message in the access list, so that whatever it reports holds for
	// the real delivery.
	require.Len(t, simulated, 1)
	call := simulated[0]
	require.Equal(t, testRelayerAddress, call.From)
	require.Equal(t, testTeleporterAddress, *call.To)
	require.Equal(t, expectedWarpGasLimit(t, handler, signedMessage), call.Gas)
	require.Equal(t, utils.SignedWarpMessageToAccessList(signedMessage), call.AccessList)
}

// A simulated revert because another relayer delivered the message first is not a failure: there
// is nothing left to send.
func TestWarpSendMessageSimulatedRevertAlreadyDelivered(t *testing.T) {
	ctrl := gomock.NewController(t)
	handler, handlerMocks := newTestWarpHandler(t, ctrl)
	signedMessage := testSignedWarpMessage(t, handler)

	handlerMocks.routeCalls(t, callRoutes{
		delivered: true,
		simulate: func(ethereum.CallMsg) ([]byte, error) {
			return nil, revertWithReason(t, "TeleporterMessenger: message already received")
		},
	})

	txHash, err := handler.SendMessage(context.Background(), signedMessage)
	require.NoError(t, err)
	require.Equal(t, common.Hash{}, txHash)
}

// If the simulation cannot be run at all, nothing is known about the delivery: it is neither
// abandoned nor broadcast blind, and the failure stays retryable.
func TestWarpSendMessageSimulationFailureIsRetryable(t *testing.T) {
	ctrl := gomock.NewController(t)
	handler, handlerMocks := newTestWarpHandler(t, ctrl)
	signedMessage := testSignedWarpMessage(t, handler)

	handlerMocks.routeCalls(t, callRoutes{
		simulate: func(ethereum.CallMsg) ([]byte, error) {
			return nil, errors.New("connection refused")
		},
	})

	_, err := handler.SendMessage(context.Background(), signedMessage)
	require.Error(t, err)
	require.NotErrorIs(t, err, messages.ErrNonRetryable)
}

// Running out of gas in simulation cannot be told apart from the node capping the call's gas, so
// the delivery is sent and the receipt decides.
func TestWarpSendMessageSimulatedOutOfGasIsSent(t *testing.T) {
	ctrl := gomock.NewController(t)
	handler, handlerMocks := newTestWarpHandler(t, ctrl)
	signedMessage := testSignedWarpMessage(t, handler)

	handlerMocks.routeCalls(t, callRoutes{
		simulate: func(ethereum.CallMsg) ([]byte, error) {
			return nil, &rpcError{msg: "out of gas"}
		},
	})
	expectedTxHash := common.HexToHash("0x01")
	handlerMocks.destinationClient.EXPECT().
		SendTx(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&types.Receipt{Status: types.ReceiptStatusSuccessful, TxHash: expectedTxHash}, nil).
		Times(1)

	txHash, err := handler.SendMessage(context.Background(), signedMessage)
	require.NoError(t, err)
	require.Equal(t, expectedTxHash, txHash)
}

// A delivery that simulates successfully is broadcast exactly as simulated.
func TestWarpSendMessageSimulationPassesThenSends(t *testing.T) {
	ctrl := gomock.NewController(t)
	handler, handlerMocks := newTestWarpHandler(t, ctrl)
	signedMessage := testSignedWarpMessage(t, handler)

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
			require.Equal(t, simulated.AccessList, accessList)
			require.Equal(t, set.Of(testRelayerAddress), deliverers)
			require.Equal(t, testTeleporterAddress, toAddress)
			require.Equal(t, simulated.Gas, gasLimit)
			require.Equal(t, simulated.Data, callData)
			return &types.Receipt{Status: types.ReceiptStatusSuccessful, TxHash: expectedTxHash}, nil
		}).
		Times(1)

	txHash, err := handler.SendMessage(context.Background(), signedMessage)
	require.NoError(t, err)
	require.Equal(t, expectedTxHash, txHash)
}

// A mined delivery that reverted having consumed its entire gas limit ran out of gas. The gas
// limit is computed from the message, so a retry would fail identically and only pay for a second
// reverted transaction: the message must be abandoned.
func TestWarpSendMessageOutOfGasReceiptIsNonRetryable(t *testing.T) {
	ctrl := gomock.NewController(t)
	handler, handlerMocks := newTestWarpHandler(t, ctrl)
	signedMessage := testSignedWarpMessage(t, handler)

	handlerMocks.routeCalls(t, callRoutes{
		delivered: false,
		simulate: func(ethereum.CallMsg) ([]byte, error) {
			return []byte{}, nil
		},
	})
	// Echo the gas limit the handler chose back as gas used: that is what an out-of-gas revert
	// looks like on chain.
	handlerMocks.destinationClient.EXPECT().
		SendTx(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(
			_ logging.Logger,
			_ types.AccessList,
			_ set.Set[common.Address],
			_ common.Address,
			gasLimit uint64,
			_ []byte,
		) (*types.Receipt, error) {
			return &types.Receipt{Status: types.ReceiptStatusFailed, GasUsed: gasLimit}, nil
		}).
		Times(1)

	_, err := handler.SendMessage(context.Background(), signedMessage)
	require.ErrorIs(t, err, messages.ErrDeliveryOutOfGas)
	require.ErrorIs(t, err, messages.ErrNonRetryable)
}

// A mined revert that left gas unspent after a successful simulation is a verification failure,
// which a fresh signature may resolve, so it stays retryable.
func TestWarpSendMessageRevertWithGasLeftIsRetryable(t *testing.T) {
	ctrl := gomock.NewController(t)
	handler, handlerMocks := newTestWarpHandler(t, ctrl)
	signedMessage := testSignedWarpMessage(t, handler)

	handlerMocks.routeCalls(t, callRoutes{
		delivered: false,
		simulate: func(ethereum.CallMsg) ([]byte, error) {
			return []byte{}, nil
		},
	})
	handlerMocks.destinationClient.EXPECT().
		SendTx(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&types.Receipt{Status: types.ReceiptStatusFailed, GasUsed: 21_000}, nil).
		Times(1)

	_, err := handler.SendMessage(context.Background(), signedMessage)
	require.Error(t, err)
	require.NotErrorIs(t, err, messages.ErrNonRetryable)
}

// A mined revert because another relayer delivered the message first counts as delivered.
func TestWarpSendMessageRevertedReceiptAlreadyDelivered(t *testing.T) {
	ctrl := gomock.NewController(t)
	handler, handlerMocks := newTestWarpHandler(t, ctrl)
	signedMessage := testSignedWarpMessage(t, handler)

	handlerMocks.routeCalls(t, callRoutes{
		delivered: true,
		simulate: func(ethereum.CallMsg) ([]byte, error) {
			return []byte{}, nil
		},
	})
	expectedTxHash := common.HexToHash("0x01")
	handlerMocks.destinationClient.EXPECT().
		SendTx(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&types.Receipt{Status: types.ReceiptStatusFailed, GasUsed: 21_000, TxHash: expectedTxHash}, nil).
		Times(1)

	txHash, err := handler.SendMessage(context.Background(), signedMessage)
	require.NoError(t, err)
	require.Equal(t, expectedTxHash, txHash)
}
