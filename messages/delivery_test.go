// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package messages

import (
	"context"
	"errors"
	"testing"

	"github.com/ava-labs/avalanchego/utils/logging"
	mock_evm "github.com/ava-labs/icm-services/vms/evm/mocks"
	mock_vms "github.com/ava-labs/icm-services/vms/mocks"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/crypto"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

var (
	senderA = common.HexToAddress("0x11")
	senderB = common.HexToAddress("0x22")
	otherC  = common.HexToAddress("0x33")
)

// rpcError mirrors the error libevm's RPC client returns for a JSON-RPC error object, carrying
// the revert output in the data field.
type rpcError struct {
	msg  string
	data any
}

func (e *rpcError) Error() string {
	return e.msg
}

func (e *rpcError) ErrorData() any {
	return e.data
}

// revertWithReason is the error a node returns for a call that reverted with Error(string).
func revertWithReason(t *testing.T, reason string) error {
	t.Helper()
	stringType, err := abi.NewType("string", "", nil)
	require.NoError(t, err)
	packed, err := abi.Arguments{{Type: stringType}}.Pack(reason)
	require.NoError(t, err)
	data := append(crypto.Keccak256([]byte("Error(string)"))[:4], packed...)
	return &rpcError{msg: "execution reverted: " + reason, data: hexutil.Encode(data)}
}

func TestIsAllowedRelayer(t *testing.T) {
	require.True(t, IsAllowedRelayer(nil, senderA))
	require.True(t, IsAllowedRelayer([]common.Address{senderA}, senderA))
	require.False(t, IsAllowedRelayer([]common.Address{otherC}, senderA))
}

func TestContainsAllowedRelayer(t *testing.T) {
	senders := []common.Address{senderA, senderB}
	require.True(t, ContainsAllowedRelayer(nil, senders))
	require.True(t, ContainsAllowedRelayer([]common.Address{otherC, senderB}, senders))
	require.False(t, ContainsAllowedRelayer([]common.Address{otherC}, senders))
	require.False(t, ContainsAllowedRelayer([]common.Address{otherC}, nil))
}

// Checks that an allowed relayer address is picked when the message restricts relayers, and
// the first sender otherwise.
func TestSelectSenderAddress(t *testing.T) {
	senders := []common.Address{senderA, senderB}
	require.Equal(t, senderA, SelectSenderAddress(senders, nil))
	require.Equal(t, senderB, SelectSenderAddress(senders, []common.Address{senderB}))
	// No eligible sender: fall back to the first, for callers that only need a plausible From.
	require.Equal(t, senderA, SelectSenderAddress(senders, []common.Address{otherC}))
	require.Equal(t, common.Address{}, SelectSenderAddress(nil, nil))
}

// Checks that a delivery that ran out of gas is classified as non-retryable, while other
// reverts stay retryable.
func TestCheckReceiptStatus(t *testing.T) {
	const gasLimit = uint64(1_000_000)
	txHash := common.HexToHash("0xabc")

	require.NoError(t, CheckReceiptStatus(
		&types.Receipt{Status: types.ReceiptStatusSuccessful, TxHash: txHash}, gasLimit))

	err := CheckReceiptStatus(
		&types.Receipt{Status: types.ReceiptStatusFailed, GasUsed: gasLimit, TxHash: txHash}, gasLimit)
	require.ErrorIs(t, err, ErrDeliveryOutOfGas)

	err = CheckReceiptStatus(
		&types.Receipt{Status: types.ReceiptStatusFailed, GasUsed: gasLimit / 2, TxHash: txHash}, gasLimit)
	require.Error(t, err)
	require.NotErrorIs(t, err, ErrNonRetryable)
}

// Checks that gas estimation buffers a successful estimate, abandons simulated reverts without
// a fallback, and falls back to the block gas limit only on non-revert failures.
func TestEstimateGasLimit(t *testing.T) {
	const blockGasLimit = uint64(15_000_000)
	to := common.HexToAddress("0xbb")

	newClients := func(t *testing.T, estimated uint64, estimateErr error) *mock_vms.MockDestinationClient {
		ctrl := gomock.NewController(t)
		ethClient := mock_evm.NewMockClient(ctrl)
		ethClient.EXPECT().
			EstimateGas(gomock.Any(), gomock.Any()).
			DoAndReturn(func(_ context.Context, _ ethereum.CallMsg) (uint64, error) {
				return estimated, estimateErr
			}).
			AnyTimes()
		destinationClient := mock_vms.NewMockDestinationClient(ctrl)
		destinationClient.EXPECT().Client().Return(ethClient).AnyTimes()
		destinationClient.EXPECT().BlockGasLimit().Return(blockGasLimit).AnyTimes()
		return destinationClient
	}

	t.Run("buffers successful estimate", func(t *testing.T) {
		gasLimit, err := EstimateGasLimit(
			context.Background(), logging.NoLog{}, newClients(t, 1_000_000, nil), senderA, to, nil,
		)
		require.NoError(t, err)
		require.Equal(t, uint64(1_250_000), gasLimit)
	})

	t.Run("caps buffered estimate at block gas limit", func(t *testing.T) {
		gasLimit, err := EstimateGasLimit(
			context.Background(), logging.NoLog{}, newClients(t, 14_000_000, nil), senderA, to, nil,
		)
		require.NoError(t, err)
		require.Equal(t, blockGasLimit, gasLimit)
	})

	t.Run("abandons simulated revert", func(t *testing.T) {
		_, err := EstimateGasLimit(
			context.Background(), logging.NoLog{},
			newClients(t, 0, revertWithReason(t, "unauthorized relayer")), senderA, to, nil,
		)
		require.ErrorIs(t, err, ErrDeliveryReverts)
		require.ErrorContains(t, err, "unauthorized relayer")
	})

	t.Run("falls back to block gas limit on non-revert failure", func(t *testing.T) {
		gasLimit, err := EstimateGasLimit(
			context.Background(), logging.NoLog{},
			newClients(t, 0, errors.New("connection refused")), senderA, to, nil,
		)
		require.NoError(t, err)
		require.Equal(t, blockGasLimit, gasLimit)
	})
}
