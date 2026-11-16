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
	"github.com/ava-labs/libevm/ethclient"
	"go.uber.org/zap"
)

const (
	// maxTransitionsPerTick caps consensus transitions on destination
	// chain applied per poll interval
	maxTransitionsPerTick = 4

	// subgraphPageSize is how many proofs one subgraph query fetches. It exceeds
	// maxTransitionsPerTick because the subgraph may index several proofs over the same slots
	// (e.g. from other provers, or extending other states), of which only the one whose
	// preState matches the contract's current state can be applied.
	subgraphPageSize = 100

	// slotsPerEpoch matches ZKStateManager's SLOT_PER_EPOCH.
	slotsPerEpoch = 32
)

// ConsensusUpdater keeps the destination ZKAdapter's view of Ethereum
// consensus up to sync by applying Boundless proofs of beacon chain
// finality, fetched from the Boundless subgraph tool.
//
// The ZKStateManager imposes two properties this updater is built around:
//
//  1. Transitions are sequential: each proof's journal preState must match
//     the contract's stored state. The updater therefore selects proofs by
//     state, not by slot alone: it applies the proof whose preState equals the
//     contract's current state, then the one extending that proof's postState,
//     and so on. Slot order alone is not enough, as lastConfirmedSlot is zero on
//     a fresh deployment and the subgraph may index proofs of other states.
//
//  2. Transitions are time-bounded: a single transition may span at most
//     permissibleTimespan of chain time, so this updater must run
//     continuously regardless of message traffic. The permissible timespan
//     is also enforced at the contract level. A stalled updater bricks
//     the contract into admin-only (manualTransition) recovery mode.
type ConsensusUpdater struct {
	logger          logging.Logger
	subgraphClient  *SubgraphClient
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
	// lastAppliedSlot is the highest finalizedSlot this updater has successfully
	// transitioned the contract to. It is also used as the subgraph query lower
	// bound.
	lastAppliedSlot uint64
	// currentState is the contract's consensus state: the preState the next applied proof must
	// carry. It advances to each applied proof's postState.
	currentState zkadapter.ConsensusState
	// synced reports whether the contract's latest applied slot and current
	// state have been read into lastAppliedSlot and currentState. Set false on
	// boot and after a transition revert; while false, the next tick reads them
	// from the contract before applying proofs.
	synced bool
}

func NewConsensusUpdater(
	logger logging.Logger,
	subgraphClient *SubgraphClient,
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
		subgraphClient:  subgraphClient,
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
	if !u.synced {
		if err := u.initializeLocalState(ctx); err != nil {
			return err
		}
	}

	// Fetch the indexed proofs that may extend the contract's current state.
	proofs, err := u.subgraphClient.ConsensusProofsAfterSlot(ctx, u.queryLowerBound(), subgraphPageSize)
	if err != nil {
		return fmt.Errorf("failed to query subgraph: %w", err)
	}

	proof, journal := u.nextProof(proofs)
	if proof == nil {
		logFn := u.logger.Debug
		if len(proofs) > 0 {
			// Proofs exist past the contract's position, yet none extends its state: the subgraph
			// is missing the next link of the chain.
			logFn = u.logger.Warn
		}
		logFn("No indexed consensus proof extends the contract's current state",
			zap.Uint64("lastAppliedSlot", u.lastAppliedSlot),
			zap.Uint64("finalizedEpoch", u.currentState.FinalizedCheckpoint.Epoch),
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
			// This means the relayer's lastAppliedSlot has diverged from
			// the contract. In this case, we re-sync the contract and retry on the
			// next tick.
			u.logger.Warn("Transition failed, resyncing local state from contract",
				zap.Uint64("finalizedSlot", proof.FinalizedSlot),
				zap.Error(err),
			)
			// We set this to false because the transition fails if the contract's state is ahead
			// of the lastAppliedSlot. This can happen if another party has updated the contract
			// with a proof before the relayer did. Thus, we set the synced flag to false and read
			// the contract's current state so on the next tick we are back in sync with the contract.
			u.synced = false
			return nil
		}
		u.currentState = journal.PostState
		proof, journal = u.nextProof(proofs)
	}
	return nil
}

// queryLowerBound returns the slot after which the subgraph is queried for proofs extending
// currentState. Such a proof finalizes a slot after lastAppliedSlot and after the state's own
// finalized checkpoint. The checkpoint bound is what matters on a fresh deployment, where
// lastAppliedSlot is still zero and querying from it would return the oldest indexed proofs.
// It is taken one epoch early, as a checkpoint's block may precede its epoch's first slot when
// slots were skipped.
func (u *ConsensusUpdater) queryLowerBound() uint64 {
	bound := u.lastAppliedSlot
	if epoch := u.currentState.FinalizedCheckpoint.Epoch; epoch > 1 {
		bound = max(bound, (epoch-1)*slotsPerEpoch)
	}
	return bound
}

// nextProof returns the first of proofs, with its decoded journal, that extends currentState
// and finalizes a slot after lastAppliedSlot, or nil if none does. The contract rejects any
// other proof, so the rest are not applied. Proofs whose journal cannot be decoded are skipped.
func (u *ConsensusUpdater) nextProof(proofs []*ConsensusProof) (*ConsensusProof, *consensusJournal) {
	for _, proof := range proofs {
		journal, err := decodeJournal(proof.JournalData)
		if err != nil {
			u.logger.Warn("Skipping consensus proof with undecodable journal",
				zap.Uint64("finalizedSlot", proof.FinalizedSlot),
				zap.Error(err),
			)
			continue
		}
		if journal.PreState == u.currentState && journal.FinalizedSlot > u.lastAppliedSlot {
			return proof, journal
		}
	}
	return nil, nil
}

// applyTransition submits one consensus transition and waits for it to mine.
func (u *ConsensusUpdater) applyTransition(ctx context.Context, proof *ConsensusProof) error {
	tx, err := u.contract.Transition(u.txOpts, zkadapter.ConsensusData{
		JournalData: proof.JournalData,
		Seal:        proof.Seal,
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

	u.lastAppliedSlot = proof.FinalizedSlot
	u.logger.Info("Consensus transition applied",
		zap.Uint64("finalizedSlot", proof.FinalizedSlot),
		zap.String("txHash", tx.Hash().Hex()),
	)
	return nil
}

// initializeLocalState reconstructs the contract's current consensus position.
// The contract itself is the checkpoint so no local persistence is needed.
func (u *ConsensusUpdater) initializeLocalState(ctx context.Context) error {
	slot, err := u.latestConfirmedSlot(ctx)
	if err != nil {
		return fmt.Errorf("failed to read contract consensus position: %w", err)
	}
	state, err := u.contract.GetCurrentState(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("failed to read contract consensus state: %w", err)
	}

	u.lastAppliedSlot = slot
	u.currentState = state
	u.synced = true
	u.logger.Info("Synced the consensus state from contract",
		zap.Uint64("lastAppliedSlot", slot),
		zap.Uint64("finalizedEpoch", state.FinalizedCheckpoint.Epoch),
	)
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
