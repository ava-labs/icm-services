// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package validatorupdater

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	zkadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKAdapter"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/libevm/ethclient"
	"go.uber.org/zap"
)

const (
	// maxTransitionsPerTick caps consensus transitions on destination
	// chain applied per poll interval
	maxTransitionsPerTick = 4

	// proofPageSize is how many proofs one proof source query fetches. It exceeds
	// maxTransitionsPerTick because the source may hold several proofs over the same slots
	// (e.g. from other provers, or extending other states), of which only the one whose
	// previous values match the contract's current state can be applied.
	proofPageSize = 100
)

// ConsensusUpdater keeps the destination ZKAdapter's view of Ethereum
// consensus up to sync by applying SP1 Helios light client proofs of beacon
// chain finality, fetched from a ProofSource.
//
// The ZKStateManager imposes two properties this updater is built around:
//
//  1. Transitions are sequential: each proof's previous head, header, and sync
//     committee must match the contract's stored state. The updater therefore
//     selects proofs by state, not by slot alone: it applies the proof whose
//     pre-state equals the contract's current state, then the one extending
//     that proof's post-state, and so on, as the source may hold proofs of
//     other states.
//
//  2. Transitions are time-bounded: a single transition may span at most
//     permissibleTimespan of chain time, and each must be verified by the
//     sync committee of the current head's period, so this updater must run
//     continuously regardless of message traffic. A stalled updater bricks
//     the contract into admin-only (manualTransition) recovery mode.
type ConsensusUpdater struct {
	logger          logging.Logger
	proofSource     ProofSource
	ethClient       *ethclient.Client
	contract        *zkadapter.ZKAdapter
	contractAddress common.Address
	txOpts          *bind.TransactOpts

	pollInterval time.Duration
	// maxGasPriceWei is the maximum suggested gas price (in wei) at which a
	// transition transaction will be submitted. If the destination network's
	// suggested gas price exceeds this value, transitions are deferred until
	// the next poll. A zero value disables gas gating.
	maxGasPriceWei *big.Int
	// lastAppliedSlot is the contract's lastConfirmedSlot. It is read from the contract at the
	// start of each tick and advances with each transition this updater applies. It is also
	// used as the proof source query lower bound.
	lastAppliedSlot uint64
	// currentState is the contract's consensus state: the preState the next applied proof must
	// carry. It is read from the contract at the start of each tick, so a change made outside
	// this updater (another submitter, or an admin manualTransition) is picked up, and advances
	// to each applied proof's postState.
	currentState zkadapter.ConsensusState
	// failedProofs holds the proofs, keyed by proofKey, whose transition failed from
	// failedProofsState. nextProof skips them, so a proof the contract rejects cannot block a
	// valid proof indexed after it. It is cleared when the contract's state moves away from
	// failedProofsState, and when every matching proof has failed, so a proof that failed for a
	// transient reason is retried rather than skipped for good.
	failedProofs      map[common.Hash]struct{}
	failedProofsState zkadapter.ConsensusState
}

func NewConsensusUpdater(
	logger logging.Logger,
	proofSource ProofSource,
	ethClient *ethclient.Client,
	contract *zkadapter.ZKAdapter,
	contractAddress common.Address,
	txOpts *bind.TransactOpts,
	pollInterval time.Duration,
	maxGasPriceGwei uint64,
) *ConsensusUpdater {
	maxGasPriceWei := new(big.Int).Mul(new(big.Int).SetUint64(maxGasPriceGwei), weiPerGwei)
	return &ConsensusUpdater{
		logger:          logger,
		proofSource:     proofSource,
		ethClient:       ethClient,
		contract:        contract,
		contractAddress: contractAddress,
		txOpts:          txOpts,
		pollInterval:    pollInterval,
		maxGasPriceWei:  maxGasPriceWei,
	}
}

// Start runs the polling loop that applies new consensus proofs to the contract.
func (u *ConsensusUpdater) Start(ctx context.Context) error {
	u.logger.Info("Starting ConsensusUpdater",
		zap.Stringer("contractAddress", u.contractAddress),
		zap.Duration("pollInterval", u.pollInterval),
	)

	if err := u.checkAndUpdate(ctx); err != nil {
		u.logger.Error("Initial update failed", zap.Error(err))
	}

	ticker := time.NewTicker(u.pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			u.logger.Info("ConsensusUpdater stopping")
			// Being asked to stop is a clean exit, not a failure: returning ctx.Err() would make
			// the relayer's errgroup treat every graceful shutdown as an error.
			return nil
		case <-ticker.C:
			if err := u.checkAndUpdate(ctx); err != nil {
				u.logger.Error("Update check failed", zap.Error(err))
			}
		}
	}
}

func (u *ConsensusUpdater) checkAndUpdate(ctx context.Context) error {
	// Re-read the contract's position every tick rather than trusting the cached one. Both are
	// view calls, and without this a state change made outside this updater would leave no
	// proof matching the cached state, stalling the updater for good.
	if err := u.syncLocalState(ctx); err != nil {
		return err
	}

	// Fetch the proofs that may extend the contract's current state. The contract's head is its
	// lastConfirmedSlot, so any such proof finalizes a later slot.
	proofs, err := u.proofSource.ConsensusProofsAfterSlot(ctx, u.lastAppliedSlot, proofPageSize)
	if err != nil {
		return fmt.Errorf("failed to query proof source: %w", err)
	}

	proof, transition := u.nextProof(proofs)
	if proof == nil && len(u.failedProofs) > 0 {
		// Every proof extending the current state has failed. Retry them from the first, as
		// their failures may have been transient.
		u.failedProofs = nil
		proof, transition = u.nextProof(proofs)
	}
	if proof == nil {
		logFn := u.logger.Debug
		if len(proofs) > 0 {
			// Proofs exist past the contract's position, yet none extends its state: the proof
			// source is missing the next link of the chain.
			logFn = u.logger.Warn
		}
		logFn("No consensus proof extends the contract's current state",
			zap.Uint64("lastAppliedSlot", u.lastAppliedSlot),
			zap.Uint64("head", u.currentState.Head),
			zap.Int("candidates", len(proofs)),
		)
		return nil
	}

	ok, err := u.gasPriceWithinBounds(ctx)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	for applied := 0; proof != nil && applied < maxTransitionsPerTick; applied++ {
		if err := u.applyTransition(ctx, proof); err != nil {
			// The next tick re-reads the contract's state. If it moved (e.g. another party
			// applied a proof first), the updater continues from there.
			u.logger.Warn("Transition failed, retrying on the next tick",
				zap.Uint64("newHead", proof.NewHead),
				zap.Error(err),
			)
			// The contract may also have rejected the proof itself, in which case the next tick
			// finds the state unchanged and must try another proof instead of this one again.
			u.markProofFailed(proof)
			return nil
		}
		u.currentState = transition.PostState
		proof, transition = u.nextProof(proofs)
	}
	return nil
}

// nextProof returns the first of proofs, with its decoded transition, that extends currentState
// and finalizes a slot after lastAppliedSlot, or nil if none does. The contract rejects any
// other proof, so the rest are not applied. Proofs whose public values cannot be decoded, and
// proofs that already failed from currentState, are skipped.
func (u *ConsensusUpdater) nextProof(proofs []*ConsensusProof) (*ConsensusProof, *consensusTransition) {
	for _, proof := range proofs {
		if _, failed := u.failedProofs[proofKey(proof)]; failed {
			continue
		}
		transition, err := decodePublicValues(proof.PublicValues)
		if err != nil {
			u.logger.Warn("Skipping consensus proof with undecodable public values",
				zap.Uint64("newHead", proof.NewHead),
				zap.Error(err),
			)
			continue
		}
		if transition.PreState == u.currentState && transition.PostState.Head > u.lastAppliedSlot {
			return proof, transition
		}
	}
	return nil, nil
}

// markProofFailed records that proof failed to transition the contract from currentState.
func (u *ConsensusUpdater) markProofFailed(proof *ConsensusProof) {
	if u.failedProofs == nil || u.failedProofsState != u.currentState {
		u.failedProofs = make(map[common.Hash]struct{})
		u.failedProofsState = u.currentState
	}
	u.failedProofs[proofKey(proof)] = struct{}{}
}

// proofKey identifies a proof by its public values and proof bytes, which together are what the
// contract verifies.
func proofKey(proof *ConsensusProof) common.Hash {
	return crypto.Keccak256Hash(proof.PublicValues, proof.Proof)
}

// applyTransition submits one consensus transition and waits for it to mine.
func (u *ConsensusUpdater) applyTransition(ctx context.Context, proof *ConsensusProof) error {
	tx, err := u.contract.Transition(u.txOpts, zkadapter.ConsensusData{
		PublicValues: proof.PublicValues,
		Proof:        proof.Proof,
	})
	if err != nil {
		return fmt.Errorf("failed to submit transition: %w", err)
	}
	receipt, err := bind.WaitMined(ctx, u.ethClient, tx)
	if err != nil {
		return fmt.Errorf("waiting for transition tx: %w", err)
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return fmt.Errorf("transition tx reverted: %s", tx.Hash().Hex())
	}

	u.lastAppliedSlot = proof.NewHead
	u.logger.Info("Consensus transition applied",
		zap.Uint64("newHead", proof.NewHead),
		zap.String("txHash", tx.Hash().Hex()),
	)
	return nil
}

// syncLocalState reads the contract's current consensus position into lastAppliedSlot and
// currentState. The contract itself is the checkpoint so no local persistence is needed.
func (u *ConsensusUpdater) syncLocalState(ctx context.Context) error {
	slot, err := u.latestConfirmedSlot(ctx)
	if err != nil {
		return fmt.Errorf("failed to read contract consensus position: %w", err)
	}
	state, err := u.contract.GetCurrentState(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("failed to read contract consensus state: %w", err)
	}

	// Log only a change from the cached position, as this runs every tick. On boot the cache is
	// empty, so the first read is always logged.
	if slot != u.lastAppliedSlot || state != u.currentState {
		u.logger.Info("Synced the consensus state from contract",
			zap.Uint64("lastAppliedSlot", slot),
			zap.Uint64("head", state.Head),
		)
	}
	u.lastAppliedSlot = slot
	u.currentState = state
	// Failures from a state the contract has since left say nothing about the proofs that
	// extend its new state.
	if state != u.failedProofsState {
		u.failedProofs = nil
	}
	return nil
}

// latestConfirmedSlot returns the contract's lastConfirmedSlot. It is read from storage rather
// than from ConfirmedBeaconBlock events, so it needs no unbounded eth_getLogs query and reflects
// a manualTransition that rewinds the slot.
func (u *ConsensusUpdater) latestConfirmedSlot(ctx context.Context) (uint64, error) {
	slot, err := u.contract.LastConfirmedSlot(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, fmt.Errorf("failed to read lastConfirmedSlot: %w", err)
	}
	return slot, nil
}

// gasPriceWithinBounds reports whether the destination's suggested gas price
// is at or below the configured threshold. A zero threshold disables the
// check. A false result means "defer until the next poll".
func (u *ConsensusUpdater) gasPriceWithinBounds(ctx context.Context) (bool, error) {
	if u.maxGasPriceWei.Sign() == 0 {
		return true, nil
	}
	suggested, err := u.ethClient.SuggestGasPrice(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to get suggested gas price: %w", err)
	}
	if suggested.Cmp(u.maxGasPriceWei) > 0 {
		u.logger.Info("Skipping validator set update: suggested gas price above threshold",
			zap.String("suggestedGasPriceWei", suggested.String()),
			zap.String("maxGasPriceWei", u.maxGasPriceWei.String()),
		)
		return false, nil
	}
	u.logger.Debug("Suggested gas price within threshold",
		zap.String("suggestedGasPriceWei", suggested.String()),
		zap.String("maxGasPriceWei", u.maxGasPriceWei.String()),
	)
	return true, nil
}
