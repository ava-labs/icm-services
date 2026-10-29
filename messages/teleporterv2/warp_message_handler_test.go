// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporterv2

import (
	"context"
	"testing"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	warpadapter "github.com/ava-labs/icm-services/abi-bindings/go/teleporterV2/WarpAdapter"
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

// expectedWarpCallData is the receiveCrossChainMessage calldata SendMessage packs for delivering
// [signedMessage].
func expectedWarpCallData(t *testing.T, handler *warpMessageHandler, signedMessage *warp.Message) []byte {
	t.Helper()
	callData, err := warpadapter.PackReceiveCrossChainMessage(
		*handler.teleporterMessage,
		signedMessage.SourceChainID,
		0,
		testRewardAddress,
	)
	require.NoError(t, err)
	return callData
}

// The handler must have the node estimate, and then broadcast, the delivery of the signed message
// to the TeleporterMessengerV2: the relayer as sender, the WarpAdapter's receiveCrossChainMessage
// calldata and the signed message as the predicate access list, bounded by the block gas limit.
// The transaction carries the estimate with the sender's margin.
func TestWarpSendMessageEstimatesThenSends(t *testing.T) {
	ctrl := gomock.NewController(t)
	handler, handlerMocks := newTestWarpHandler(t, ctrl)
	signedMessage := testSignedWarpMessage(t, handler)
	expectedCallData := expectedWarpCallData(t, handler, signedMessage)
	expectedAccessList := utils.SignedWarpMessageToAccessList(signedMessage)

	var estimated ethereum.CallMsg
	handlerMocks.routeCalls(t, callRoutes{
		estimate: func(call ethereum.CallMsg) (uint64, error) {
			estimated = call
			return 1_000_000, nil
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
			require.Equal(t, uint64(1_250_000), gasLimit)
			require.Equal(t, expectedCallData, callData)
			return &types.Receipt{Status: types.ReceiptStatusSuccessful, TxHash: expectedTxHash}, nil
		}).
		Times(1)

	txHash, err := handler.SendMessage(context.Background(), signedMessage)
	require.NoError(t, err)
	require.Equal(t, expectedTxHash, txHash)

	require.Equal(t, testRelayerAddress, estimated.From)
	require.Equal(t, testTeleporterAddress, *estimated.To)
	require.Equal(t, testBlockGasLimit, estimated.Gas)
	require.Equal(t, expectedCallData, estimated.Data)
	require.Equal(t, expectedAccessList, estimated.AccessList)
}

// A delivery that reverts when the node executes it for the estimate, such as one carrying a
// receipt the messenger cannot account for, is abandoned as non-retryable without being
// broadcast: no SendTx expectation is registered, so the mock fails the test if the handler
// sends it anyway.
func TestWarpSendMessageEstimationRevertIsNonRetryable(t *testing.T) {
	ctrl := gomock.NewController(t)
	handler, handlerMocks := newTestWarpHandler(t, ctrl)
	signedMessage := testSignedWarpMessage(t, handler)

	handlerMocks.routeCalls(t, callRoutes{
		delivered: false,
		estimate: func(ethereum.CallMsg) (uint64, error) {
			return 0, &evm.RevertError{Reason: "arithmetic underflow or overflow"}
		},
	})

	_, err := handler.SendMessage(context.Background(), signedMessage)
	require.ErrorIs(t, err, messages.ErrDeliveryReverts)
	require.ErrorIs(t, err, messages.ErrNonRetryable)
	require.ErrorContains(t, err, "arithmetic underflow or overflow")
}
