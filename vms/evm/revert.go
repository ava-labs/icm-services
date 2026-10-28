// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"errors"
	"strings"

	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/core/vm"
	"github.com/ava-labs/libevm/rpc"
)

// RevertError reports that the node executed a simulated call (eth_call or eth_estimateGas) and
// the call failed inside the EVM. It is distinct from the request failing to reach the node or to
// be served by it: against the same state, the same call fails the same way again.
type RevertError struct {
	// Reason is the decoded Error(string) or Panic(uint256) revert reason when the call returned
	// one, the hex-encoded revert data when it returned some other (custom) error, and the node's
	// error message when the call failed without returning data.
	Reason string
	// Data is the raw revert output, if the call returned any.
	Data []byte
	// OutOfGas reports that the call ran out of gas rather than reverting.
	OutOfGas bool

	err error
}

func (e *RevertError) Error() string {
	if e.OutOfGas {
		return "call ran out of gas: " + e.Reason
	}
	return "call reverted: " + e.Reason
}

func (e *RevertError) Unwrap() error {
	return e.err
}

// AsRevertError returns the EVM execution failure described by [err], an error returned by an
// eth_call or eth_estimateGas request, or nil if [err] is not one. Transport failures, timeouts
// and errors the node raised without executing the call are not execution failures.
func AsRevertError(err error) *RevertError {
	if err == nil {
		return nil
	}
	// geth-style nodes attach the output of a reverted call to the error as hex data.
	var dataErr rpc.DataError
	if errors.As(err, &dataErr) {
		if data := revertData(dataErr.ErrorData()); len(data) > 0 {
			return &RevertError{Reason: revertReason(data), Data: data, err: err}
		}
	}
	// A call that reverted without returning data, or that ran out of gas, only carries the EVM
	// error message.
	msg := err.Error()
	switch {
	case strings.Contains(msg, vm.ErrExecutionReverted.Error()):
		return &RevertError{Reason: msg, err: err}
	case strings.Contains(msg, vm.ErrOutOfGas.Error()):
		return &RevertError{Reason: msg, OutOfGas: true, err: err}
	}
	return nil
}

// revertData decodes the revert output a node attached to an error. Nodes encode it as a hex
// string; any other shape is treated as no data.
func revertData(data any) []byte {
	hexData, ok := data.(string)
	if !ok {
		return nil
	}
	decoded, err := hexutil.Decode(hexData)
	if err != nil {
		return nil
	}
	return decoded
}

// revertReason renders revert output for logs and error messages: the text of an Error(string)
// or Panic(uint256) revert, or the hex encoding of anything else, such as a custom error.
func revertReason(data []byte) string {
	if reason, err := abi.UnpackRevert(data); err == nil {
		return reason
	}
	return hexutil.Encode(data)
}
