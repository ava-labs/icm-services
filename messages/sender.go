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

// Sender delivers one message to a destination contract with a transaction. It decides, before
// and after the transaction, whether the message is still worth paying for: a delivery that
// reverts in simulation for a reason the message itself causes is abandoned without being
// broadcast, a message another relayer delivered in the meantime is left alone, and a mined
// delivery that ran out of gas is abandoned rather than retried. Abandoned deliveries are
// reported with errors wrapping [ErrNonRetryable], so that the relayer checkpoints past the
// message instead of exiting and replaying it on every restart.
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

// Send checks that the delivery fits in a block, simulates it and, unless the simulation shows
// it would revert, broadcasts it and resolves the receipt. Returns the transaction hash of a
// successful delivery, or the zero hash if nothing was sent because the message turned out to be
// delivered already.
func (s *Sender) Send(
	ctx context.Context,
	accessList types.AccessList,
	gasLimit uint64,
	callData []byte,
) (common.Hash, error) {
	// A delivery that cannot fit in a block is refused by the node at submission, which would
	// look like a transient failure and be retried. The simulation cannot catch it, because a
	// call is bounded by the node's RPC gas cap rather than by the block gas limit. A zero limit
	// means the destination's is not configured, and the node decides.
	if blockGasLimit := s.DestinationClient.BlockGasLimit(); blockGasLimit != 0 && gasLimit > blockGasLimit {
		s.Logger.Warn(
			"Delivery gas limit exceeds destination block gas limit, not sending",
			zap.Uint64("gasLimit", gasLimit),
			zap.Uint64("blockGasLimit", blockGasLimit),
		)
		return common.Hash{}, fmt.Errorf(
			"%w: %d exceeds %d", ErrDeliveryExceedsBlockGasLimit, gasLimit, blockGasLimit,
		)
	}

	err := s.Simulate(ctx, accessList, gasLimit, callData)
	var revert *evm.RevertError
	switch {
	case err == nil:
	case !errors.As(err, &revert):
		s.Logger.Error("Failed to simulate delivery", zap.Error(err))
		return common.Hash{}, err
	case revert.OutOfGas:
		// Running out of gas in simulation is most likely deterministic too, but it is also what a
		// node's RPC gas cap looks like when it is below the gas limit. Let the chain decide: if
		// the transaction runs out of gas there as well, SendTxAndConfirm abandons the message
		// after this one attempt.
		s.Logger.Warn(
			"Delivery ran out of gas in simulation, sending anyway",
			zap.Uint64("gasLimit", gasLimit),
			zap.String("reason", revert.Reason),
		)
	default:
		return s.AbandonRevert(revert)
	}
	return s.SendTxAndConfirm(accessList, gasLimit, callData)
}

// Simulate executes the delivery that SendTxAndConfirm would broadcast as an eth_call against
// the destination chain's latest state, without broadcasting anything. It returns nil if the call
// would succeed, an [*evm.RevertError] if the node executed the call and it failed, and any other
// error if the call could not be simulated.
//
// The call carries the same sender, calldata, access list and gas limit as the transaction. For
// a delivery that carries a signed Warp message as a predicate in [accessList], the node does not
// verify the predicate for a call: getVerifiedWarpMessage reports the message as valid and the
// simulation runs every check the destination contract makes after verification against the
// message itself. A revert is therefore reproducible from the message content, independent of
// the signature, and would be mined identically on every retry. The verification failures a
// retry with a fresh signature can fix are exactly the ones the simulation cannot produce.
//
// Nodes silently cap the gas of a call at their RPC gas cap, so the destination node must allow
// calls at least the configured block gas limit for the simulation to be faithful (subnet-evm
// and geth default to 50M, against a 12M default block gas limit here).
func (s *Sender) Simulate(
	ctx context.Context,
	accessList types.AccessList,
	gasLimit uint64,
	callData []byte,
) error {
	ctx, cancel := context.WithTimeout(ctx, utils.DefaultRPCTimeout)
	defer cancel()
	_, err := s.DestinationClient.Client().CallContract(ctx, ethereum.CallMsg{
		From:       s.SenderAddress(),
		To:         &s.To,
		Gas:        gasLimit,
		Data:       callData,
		AccessList: accessList,
	}, nil)
	if err == nil {
		return nil
	}
	if revert := evm.AsRevertError(err); revert != nil {
		return revert
	}
	return fmt.Errorf("failed to simulate delivery: %w", err)
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

	// A revert that consumed the entire gas limit is an out-of-gas failure. The gas limit is
	// computed from the message, or estimated against the same state, so a retry would fail
	// identically and only bill the relayer for a second reverted transaction. Reverts that left
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

// SenderAddress picks a relayer EOA eligible to deliver the message, as the sender of simulated
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
