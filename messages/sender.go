// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package messages

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/icm-services/vms"
	"github.com/ava-labs/icm-services/vms/evm"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"go.uber.org/zap"
)

const (
	// gasLimitBufferNumerator/Denominator is the safety margin added to the estimated gas. The
	// estimate is the least gas at which the delivery does not revert, and Teleporter does not
	// revert when the message's own execution fails: it checks that the message's required gas is
	// still available, forwards it to the destination less the 63/64 cut of EIP-150, and records a
	// failed execution if that was not enough. At the bare estimate the execution can therefore
	// come up short. The margin covers that, and state drift between estimation and inclusion.
	gasLimitBufferNumerator   = 5
	gasLimitBufferDenominator = 4
)

// Sender delivers one message to a destination contract with a transaction. It decides, before
// and after the transaction, whether the message is still worth paying for: a delivery that
// reverts when the node executes it to estimate its gas, for a reason the message itself causes,
// is abandoned without being broadcast, and so is one that does not fit in a block; a message
// another relayer delivered in the meantime is left alone; and a mined delivery that ran out of
// gas is abandoned rather than retried. Abandoned deliveries are reported with errors wrapping
// [ErrNonRetryable], so that the relayer checkpoints past the message instead of exiting and
// replaying it on every restart.
type Sender struct {
	Logger            logging.Logger
	DestinationClient vms.DestinationClient
	// To is the contract the delivery transaction calls.
	To common.Address
	// AllowedRelayers restricts which relayer EOAs may deliver the message. Empty allows any.
	AllowedRelayers []common.Address
	// Delivered reports whether the destination contract has already received the message.
	Delivered func() (bool, error)
}

// Send has the destination node estimate the gas for the delivery, which executes it against
// the chain's latest state, and unless that execution shows the delivery cannot be mined,
// broadcasts it with the buffered estimate and resolves the receipt. Returns the transaction hash
// of a successful delivery, or the zero hash if nothing was sent because the message turned out
// to be delivered already.
func (s *Sender) Send(
	ctx context.Context,
	accessList types.AccessList,
	callData []byte,
) (common.Hash, error) {
	gasLimit, err := s.EstimateGas(ctx, accessList, callData)
	var revert *evm.RevertError
	switch {
	case err == nil:
	case errors.As(err, &revert):
		return s.AbandonRevert(revert)
	default:
		return common.Hash{}, err
	}
	return s.SendTxAndConfirm(accessList, gasLimit, callData)
}

// EstimateGas has the destination node estimate the gas for the delivery that SendTxAndConfirm
// would broadcast, and returns the estimate with a safety margin, capped at the configured block
// gas limit. Estimation executes the delivery against the chain's latest state without
// broadcasting anything, so it is also the simulation that vets the delivery: it returns an
// [*evm.RevertError] if the node ran the delivery and it reverted,
// [ErrDeliveryExceedsBlockGasLimit] if the delivery succeeds with no gas the node may try, and
// any other error if it could not be estimated.
//
// The estimated call carries the same sender, calldata and access list as the transaction,
// bounded by the configured block gas limit (zero leaves the bound to the node). For a delivery
// that carries a signed Warp message as a predicate in [accessList], the node does not verify the
// predicate for a call: getVerifiedWarpMessage reports the message as valid and the execution
// runs every check the destination contract makes after verification against the message
// itself. A revert is therefore reproducible from the message content, independent of the
// signature, and would be mined identically on every retry. The verification failures a retry
// with a fresh signature can fix are exactly the ones estimation cannot produce.
//
// Nodes cap the gas they will try at their RPC gas cap, so the destination node must allow calls
// at least the configured block gas limit for every delivery that fits in a block to be
// estimated (subnet-evm and geth default to 50M, against a 12M default block gas limit here).
func (s *Sender) EstimateGas(
	ctx context.Context,
	accessList types.AccessList,
	callData []byte,
) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.DefaultRPCTimeout)
	defer cancel()
	blockGasLimit := s.DestinationClient.BlockGasLimit()
	estimated, err := s.DestinationClient.Client().EstimateGas(ctx, ethereum.CallMsg{
		From:       s.SenderAddress(),
		To:         &s.To,
		Gas:        blockGasLimit,
		Data:       callData,
		AccessList: accessList,
	})
	if err != nil {
		if revert := evm.AsRevertError(err); revert != nil {
			return 0, revert
		}
		if evm.ExceedsGasAllowance(err) {
			s.Logger.Warn(
				"Delivery does not fit in the destination block gas limit, not sending",
				zap.Uint64("blockGasLimit", blockGasLimit),
				zap.Error(err),
			)
			return 0, fmt.Errorf("%w: %w", ErrDeliveryExceedsBlockGasLimit, err)
		}
		s.Logger.Error("Failed to estimate delivery gas", zap.Error(err))
		return 0, fmt.Errorf("failed to estimate delivery gas: %w", err)
	}

	gasLimit := estimated * gasLimitBufferNumerator / gasLimitBufferDenominator
	if blockGasLimit != 0 && gasLimit > blockGasLimit {
		gasLimit = blockGasLimit
	}
	s.Logger.Debug(
		"Estimated delivery gas",
		zap.Uint64("estimatedGas", estimated),
		zap.Uint64("gasLimit", gasLimit),
	)
	return gasLimit, nil
}

// AbandonRevert reports a delivery that reverts in simulation. The one benign cause is that
// another relayer delivered the message in the meantime, in which case there is nothing left to
// do and the zero hash is returned. Any other revert is reproducible from the message, so the
// returned error is non-retryable: broadcasting the delivery would only mine the same revert, and
// retrying it would mine it again.
func (s *Sender) AbandonRevert(revert *evm.RevertError) (common.Hash, error) {
	delivered, err := s.Delivered()
	if err != nil {
		s.Logger.Error(
			"Failed to check if message has been delivered to destination chain.",
			zap.Error(err),
		)
		return common.Hash{}, fmt.Errorf("failed to check if message has been delivered: %w", err)
	}
	if delivered {
		s.Logger.Info("Delivery reverts in simulation: message already delivered to destination.")
		return common.Hash{}, nil
	}

	s.Logger.Error("Delivery reverts in simulation, not sending", zap.String("revertReason", revert.Reason))
	return common.Hash{}, fmt.Errorf("%w: %s", ErrDeliveryReverts, revert.Reason)
}

// SendTxAndConfirm broadcasts the delivery and resolves the receipt. A reverted transaction is
// not necessarily a failure: another relayer may have delivered the message first, so delivery
// status is checked before the revert is reported as an error.
func (s *Sender) SendTxAndConfirm(
	accessList types.AccessList,
	gasLimit uint64,
	callData []byte,
) (common.Hash, error) {
	receipt, err := s.DestinationClient.SendTx(
		s.Logger,
		accessList,
		set.Of(s.AllowedRelayers...),
		s.To,
		gasLimit,
		callData,
	)
	if err != nil {
		s.Logger.Error("Failed to send tx.", zap.Error(err))
		return common.Hash{}, err
	}

	txHash := receipt.TxHash
	log := s.Logger.With(zap.Stringer("txID", txHash))
	if receipt.Status == types.ReceiptStatusSuccessful {
		log.Info("Delivered message to destination chain")
		return txHash, nil
	}

	delivered, err := s.Delivered()
	if err != nil {
		log.Error("Transaction failed", zap.Error(err))
		return common.Hash{}, fmt.Errorf("transaction failed with status: %d", receipt.Status)
	}
	if delivered {
		log.Info("Execution reverted: message already delivered to destination.")
		return txHash, nil
	}

	// A revert that consumed the entire gas limit is an out-of-gas failure. The gas limit was
	// estimated against the same state, with a margin on top, so a retry would fail identically
	// and only bill the relayer for a second reverted transaction. Reverts that left
	// gas unspent are logic failures, such as verification against a validator set that has
	// churned since the signature was aggregated, which a fresh attempt may resolve.
	if receipt.GasUsed >= gasLimit {
		log.Error(
			"Transaction ran out of gas",
			zap.Uint64("gasUsed", receipt.GasUsed),
			zap.Uint64("gasLimit", gasLimit),
		)
		return common.Hash{}, fmt.Errorf(
			"%w: used %d of %d",
			ErrDeliveryOutOfGas, receipt.GasUsed, gasLimit,
		)
	}

	log.Error("Transaction failed")
	return common.Hash{}, fmt.Errorf("transaction failed with status: %d", receipt.Status)
}

// SenderAddress picks a relayer EOA eligible to deliver the message, as the sender of estimated
// deliveries. Any eligible EOA gives the same result, since the only sender-dependent check a
// destination contract makes is the allowed-relayer one.
func (s *Sender) SenderAddress() common.Address {
	senders := s.DestinationClient.SenderAddresses()
	for _, sender := range senders {
		if IsAllowedRelayer(s.AllowedRelayers, sender) {
			return sender
		}
	}
	if len(senders) > 0 {
		return senders[0]
	}
	return common.Address{}
}

// IsAllowedRelayer reports whether [eoa] may deliver a message that restricts its relayers to
// [allowedRelayers]. An empty list allows anyone.
func IsAllowedRelayer(allowedRelayers []common.Address, eoa common.Address) bool {
	if len(allowedRelayers) == 0 {
		return true
	}
	return slices.Contains(allowedRelayers, eoa)
}

// ContainsAllowedRelayer reports whether any of [eoas] may deliver a message that restricts its
// relayers to [allowedRelayers].
func ContainsAllowedRelayer(allowedRelayers []common.Address, eoas []common.Address) bool {
	if len(allowedRelayers) == 0 {
		return true
	}
	return slices.ContainsFunc(eoas, func(eoa common.Address) bool {
		return IsAllowedRelayer(allowedRelayers, eoa)
	})
}
