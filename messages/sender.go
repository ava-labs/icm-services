// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package messages

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/icm-services/vms"
	"github.com/ava-labs/icm-services/vms/evm"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"go.uber.org/zap"
)

// estimateGasTimeout bounds one gas estimation. Nodes binary-search the gas, executing the
// delivery many times, which can far exceed the timeout of an ordinary RPC.
const estimateGasTimeout = 30 * time.Second

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
// of at least the configured block gas limit for its verdict to apply to the transaction
// (subnet-evm and geth default to 50M, against a 12M default block gas limit here). The
// destination client checks that when it connects, see [evm.CheckRPCGasCap]. A lower bound
// reported here is therefore a node whose cap has dropped since; it is logged, and the delivery
// is still abandoned.
func (s *Sender) EstimateGas(
	ctx context.Context,
	accessList types.AccessList,
	callData []byte,
) (uint64, error) {
	// Estimation is a binary search over many executions, so it gets its own budget instead of
	// the single-RPC timeout, and does not inherit what is left of the caller's deadline, which
	// signature aggregation may have used up.
	ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), estimateGasTimeout)
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
		if allowance, exceeds := evm.ExceedsGasAllowance(err); exceeds {
			if allowance != 0 && allowance < blockGasLimit {
				s.Logger.Error(
					"Node capped the estimate below the block gas limit, not sending: its RPC gas cap dropped since startup",
					zap.Uint64("allowance", allowance),
					zap.Uint64("blockGasLimit", blockGasLimit),
				)
			} else {
				s.Logger.Warn(
					"Delivery does not fit in the destination block gas limit, not sending",
					zap.Uint64("blockGasLimit", blockGasLimit),
					zap.Error(err),
				)
			}
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

	// Out of gas is non-retryable: the gas limit was estimated against the same state, with a
	// margin on top, so a retry would fail identically and only bill the relayer for a second
	// reverted transaction. A revert that left gas unspent is a logic failure, such as
	// verification against a validator set that has churned since the signature was aggregated,
	// which a fresh attempt may resolve.
	err = CheckReceiptStatus(receipt, gasLimit)
	log.Error(
		"Transaction failed",
		zap.Uint64("gasUsed", receipt.GasUsed),
		zap.Uint64("gasLimit", gasLimit),
		zap.Error(err),
	)
	return common.Hash{}, err
}

// SenderAddress picks a relayer EOA eligible to deliver the message, as the sender of estimated
// deliveries.
func (s *Sender) SenderAddress() common.Address {
	return SelectSenderAddress(s.DestinationClient.SenderAddresses(), s.AllowedRelayers)
}
