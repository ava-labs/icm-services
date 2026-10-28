// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"testing"

	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/crypto"
	"github.com/stretchr/testify/require"
)

// rpcError is the shape of the error libevm's RPC client returns for a JSON-RPC error object: the
// message, plus whatever the node put in the "data" field.
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

// errorStringRevertData ABI-encodes a Solidity Error(string) revert.
func errorStringRevertData(t *testing.T, reason string) []byte {
	t.Helper()
	stringType, err := abi.NewType("string", "", nil)
	require.NoError(t, err)
	packed, err := abi.Arguments{{Type: stringType}}.Pack(reason)
	require.NoError(t, err)
	return append(crypto.Keccak256([]byte("Error(string)"))[:4], packed...)
}

// panicRevertData ABI-encodes a Solidity Panic(uint256) revert.
func panicRevertData(t *testing.T, code uint64) []byte {
	t.Helper()
	uintType, err := abi.NewType("uint256", "", nil)
	require.NoError(t, err)
	packed, err := abi.Arguments{{Type: uintType}}.Pack(new(big.Int).SetUint64(code))
	require.NoError(t, err)
	return append(crypto.Keccak256([]byte("Panic(uint256)"))[:4], packed...)
}

func TestAsRevertError(t *testing.T) {
	requireReason := errorStringRevertData(t, "TeleporterMessenger: unauthorized relayer")
	panicReason := panicRevertData(t, 0x11)
	customError := []byte{0xde, 0xad, 0xbe, 0xef, 0x01}

	tests := []struct {
		name     string
		err      error
		expected *RevertError
	}{
		{
			name: "require revert with data",
			err: &rpcError{
				msg:  "execution reverted: TeleporterMessenger: unauthorized relayer",
				data: hexutil.Encode(requireReason),
			},
			expected: &RevertError{
				Reason: "TeleporterMessenger: unauthorized relayer",
				Data:   requireReason,
			},
		},
		{
			name: "panic revert with data",
			err:  &rpcError{msg: "execution reverted", data: hexutil.Encode(panicReason)},
			expected: &RevertError{
				Reason: "arithmetic underflow or overflow",
				Data:   panicReason,
			},
		},
		{
			name: "custom error is reported as hex",
			err:  &rpcError{msg: "execution reverted", data: hexutil.Encode(customError)},
			expected: &RevertError{
				Reason: hexutil.Encode(customError),
				Data:   customError,
			},
		},
		{
			name:     "revert without data",
			err:      &rpcError{msg: "execution reverted"},
			expected: &RevertError{Reason: "execution reverted"},
		},
		{
			name:     "revert with data of an unexpected shape",
			err:      &rpcError{msg: "execution reverted", data: map[string]any{"x": 1}},
			expected: &RevertError{Reason: "execution reverted"},
		},
		{
			name:     "out of gas",
			err:      &rpcError{msg: "out of gas"},
			expected: &RevertError{Reason: "out of gas", OutOfGas: true},
		},
		{
			name: "wrapped revert",
			err: fmt.Errorf(
				"failed to estimate gas: %w",
				&rpcError{msg: "execution reverted", data: hexutil.Encode(panicReason)},
			),
			expected: &RevertError{
				Reason: "arithmetic underflow or overflow",
				Data:   panicReason,
			},
		},
		{
			name: "transport failure",
			err:  errors.New("connection refused"),
		},
		{
			name: "timeout",
			err:  context.DeadlineExceeded,
		},
		{
			name: "estimate exceeds the node's gas cap",
			err:  &rpcError{msg: "gas required exceeds allowance (50000000)"},
		},
		{
			name: "nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			revert := AsRevertError(tt.err)
			if tt.expected == nil {
				require.Nil(t, revert)
				return
			}
			require.NotNil(t, revert)
			require.Equal(t, tt.expected.Reason, revert.Reason)
			require.Equal(t, tt.expected.Data, revert.Data)
			require.Equal(t, tt.expected.OutOfGas, revert.OutOfGas)
			// The original error stays reachable for callers that inspect it.
			require.ErrorIs(t, revert, tt.err)
		})
	}
}
