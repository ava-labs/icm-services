// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package messages

// This file holds delivery helpers shared by the message handlers, covering relayer
// eligibility, receipt-status classification, and gas estimation. They live next to the
// delivery errors they wrap, so that handlers for different protocols make the same
// send, skip, and abandon decisions.

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/icm-services/vms"
	"github.com/ava-labs/icm-services/vms/evm"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"go.uber.org/zap"
)

const (
	// gasLimitBufferNumerator/Denominator adds a safety margin on top of the estimated gas
	// to absorb state drift between estimation and inclusion.
	gasLimitBufferNumerator   = 5
	gasLimitBufferDenominator = 4

	// MaxDeliveryTxSize is the largest transaction, in bytes, that the EVM transaction pool
	// accepts (4 slots of 32KiB). Larger transactions are refused before execution.
	MaxDeliveryTxSize = 4 * 32 * 1024

	// minTxEnvelopeSize is a lower bound on the non-calldata bytes of a signed transaction
	// (type byte, list header, nonce, fees, gas, recipient and signature). It is deliberately
	// low so that the pre-check never rejects a transaction the node would accept; anything it
	// misses is caught by [WrapSendError].
	minTxEnvelopeSize = 100
)

// nodeRefusalMarkers are fragments of the errors a node returns when it refuses a transaction
// for its size or gas limits. These are properties of the transaction, so a resend is refused
// identically.
var nodeRefusalMarkers = []string{
	"oversized data",
	"exceeds block gas limit",
	"intrinsic gas too low",
}

// CheckDeliverySize returns [ErrDeliveryTooLarge] if a transaction carrying [callData] exceeds
// [MaxDeliveryTxSize], so the destination node would refuse it on every attempt. Gas limits are
// not checked here: gas estimation covers them, and [WrapSendError] catches node refusals.
func CheckDeliverySize(callData []byte) error {
	if size := len(callData) + minTxEnvelopeSize; size > MaxDeliveryTxSize {
		return fmt.Errorf("%w: transaction size at least %d bytes, limit %d",
			ErrDeliveryTooLarge, size, MaxDeliveryTxSize)
	}
	return nil
}

// WrapSendError marks an error from broadcasting the delivery as [ErrDeliveryTooLarge] when the
// node refused the transaction for its size or gas, which no retry can change. Other errors are
// returned unchanged and stay retryable.
func WrapSendError(err error) error {
	if err == nil {
		return nil
	}
	msg := err.Error()
	for _, marker := range nodeRefusalMarkers {
		if strings.Contains(msg, marker) {
			return fmt.Errorf("%w: node refused transaction: %w", ErrDeliveryTooLarge, err)
		}
	}
	return err
}

// IsAllowedRelayer reports whether the EOA may deliver a message restricted to
// [allowedRelayers]. An empty restriction allows every relayer.
func IsAllowedRelayer(allowedRelayers []common.Address, eoa common.Address) bool {
	if len(allowedRelayers) == 0 {
		return true
	}
	return slices.Contains(allowedRelayers, eoa)
}

// ContainsAllowedRelayer reports whether any of the EOAs may deliver a message restricted to
// [allowedRelayers]. An empty restriction allows every relayer.
func ContainsAllowedRelayer(allowedRelayers []common.Address, eoas []common.Address) bool {
	if len(allowedRelayers) == 0 {
		return true
	}
	return slices.ContainsFunc(eoas, func(eoa common.Address) bool {
		return slices.Contains(allowedRelayers, eoa)
	})
}

// SelectSenderAddress picks a sender EOA eligible to deliver the message, to use as the sender
// of simulated deliveries and gas estimates. Any eligible EOA gives the same result, since the
// only sender-dependent check the messenger makes is the allowed-relayer one.
func SelectSenderAddress(senders []common.Address, allowedRelayers []common.Address) common.Address {
	for _, sender := range senders {
		if IsAllowedRelayer(allowedRelayers, sender) {
			return sender
		}
	}
	if len(senders) > 0 {
		return senders[0]
	}
	return common.Address{}
}

// CheckReceiptStatus maps a failed delivery receipt to a delivery error. Running out of gas is
// non-retryable, since a resend would burn the same gas again. Other reverts stay retryable.
func CheckReceiptStatus(receipt *types.Receipt, gasLimit uint64) error {
	if receipt.Status == types.ReceiptStatusSuccessful {
		return nil
	}
	if receipt.GasUsed >= gasLimit {
		return fmt.Errorf("%w: used %d of %d", ErrDeliveryOutOfGas, receipt.GasUsed, gasLimit)
	}
	return fmt.Errorf("transaction %s reverted", receipt.TxHash.Hex())
}

// EstimateGasLimit estimates the gas for delivering [callData] to [to] and applies a safety
// buffer, capped at the destination's block gas limit. A delivery that reverts in estimation
// is abandoned with [ErrDeliveryReverts] instead of being broadcast, since the broadcast would
// mine the same revert. Only non-revert estimation failures fall back to the block gas limit.
func EstimateGasLimit(
	ctx context.Context,
	logger logging.Logger,
	destinationClient vms.DestinationClient,
	from common.Address,
	to common.Address,
	callData []byte,
) (uint64, error) {
	blockGasLimit := destinationClient.BlockGasLimit()
	estimated, err := destinationClient.Client().EstimateGas(ctx, ethereum.CallMsg{
		From: from,
		To:   &to,
		Data: callData,
	})
	if err == nil {
		buffered := estimated * gasLimitBufferNumerator / gasLimitBufferDenominator
		if blockGasLimit != 0 && buffered > blockGasLimit {
			buffered = blockGasLimit
		}
		return buffered, nil
	}

	if revertErr := evm.AsRevertError(err); revertErr != nil {
		// The node simulated the delivery and it reverted. Broadcasting it would pay for a
		// mined revert on every retry.
		return 0, fmt.Errorf("%w: delivery simulation reverted: %s", ErrDeliveryReverts, revertErr)
	}

	// A non-revert estimation failure, such as a flaky node, does not prove the delivery
	// fails. Fall back to the block gas limit rather than dropping the message.
	logger.Warn("Gas estimation failed, falling back to block gas limit",
		zap.Error(err),
		zap.Uint64("blockGasLimit", blockGasLimit),
	)
	if blockGasLimit == 0 {
		return 0, fmt.Errorf("failed to estimate gas and no block gas limit configured: %w", err)
	}
	return blockGasLimit, nil
}
