// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package messages

import (
	"testing"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/stretchr/testify/require"
)

var (
	senderA = common.HexToAddress("0x11")
	senderB = common.HexToAddress("0x22")
	otherC  = common.HexToAddress("0x33")
)

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
