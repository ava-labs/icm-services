// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"errors"
	"regexp"
	"strconv"
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
// eth_call or eth_estimateGas request or one wrapping a [*RevertError], or nil if [err] is not
// one. Transport failures, timeouts and errors the node raised without executing the call are
// not execution failures.
func AsRevertError(err error) *RevertError {
	if err == nil {
		return nil
	}
	// An error that already carries a classified failure keeps it.
	var revert *RevertError
	if errors.As(err, &revert) {
		return revert
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

// gasAllowanceMessage is the error geth-derived nodes return when gas estimation runs out of gas
// at the most they would try; gasAllowancePattern captures the bound they state with it.
const gasAllowanceMessage = "gas required exceeds allowance"

var gasAllowancePattern = regexp.MustCompile(gasAllowanceMessage + ` \((\d+)\)`)

// ExceedsGasAllowance reports whether [err] is a node's answer that gas estimation found no gas
// limit at which the call succeeds, up to the most the node would try, and returns that bound:
// the lower of the caller's gas limit and the node's RPC gas cap. The call ran out of gas at that
// bound; a call that reverts there for any other reason is reported as a revert instead. The
// bound is 0 if the node did not state it.
func ExceedsGasAllowance(err error) (uint64, bool) {
	if err == nil {
		return 0, false
	}
	msg := err.Error()
	if !strings.Contains(msg, gasAllowanceMessage) {
		return 0, false
	}
	match := gasAllowancePattern.FindStringSubmatch(msg)
	if match == nil {
		return 0, true
	}
	allowance, err := strconv.ParseUint(match[1], 10, 64)
	if err != nil {
		return 0, true
	}
	return allowance, true
}
