// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

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

// maxTransitionsPerTick caps consensus transitions on destination
// chain applied per poll interval
const maxTransitionsPerTick = 4

// ConsensusUpdater keeps the destination ZKAdapter's view of Ethereum
// consensus up to sync by applying Boundless proofs of beacon chain
// finality, fetched from the Boundless subgraph tool.
//
// The ZKStateManager imposes two properties this updater is built around:
//
//  1. Transitions are sequential: each proof's journal preState must match
//     the contract's stored state, so proofs are applied in ascending
//     finalizedSlot order from the contract's current position.
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
	// synced reports whether the contract's latest applied slot has been read
	// into lastAppliedSlot. Set false on boot and after a transition revert;
	// while false, the next tick reads it from the contract before applying
	// proofs.
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
			return ctx.Err()
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

	// Fetch new consensus proofs from subgraph tool after the last applied slot.
	proofs, err := u.subgraphClient.ConsensusProofsAfterSlot(ctx, u.lastAppliedSlot, maxTransitionsPerTick)
	if err != nil {
		return fmt.Errorf("failed to query subgraph: %w", err)
	}
	if len(proofs) == 0 {
		u.logger.Debug("No new consensus proofs",
			zap.Uint64("lastAppliedSlot", u.lastAppliedSlot),
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

	for _, proof := range proofs {
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
	}
	return nil
}

// applyTransition submits one consensus transition and waits for it to mine.
func (u *ConsensusUpdater) applyTransition(ctx context.Context, proof *ConsensusProof) error {
	tx, err := u.contract.Transition(u.txOpts, zkadapter.ConsensusData{
		JournalData:   proof.JournalData,
		Seal:          proof.Seal,
		FinalizedSlot: proof.FinalizedSlot,
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

	u.lastAppliedSlot = slot
	u.synced = true
	u.logger.Info("Synced the consensus state from contract",
		zap.Uint64("lastAppliedSlot", slot),
	)
	return nil
}

// latestConfirmedSlot returns the highest beacon slot the contract has
// confirmed by scanning its ConfirmedBeaconBlock events for the most recent one.
func (u *ConsensusUpdater) latestConfirmedSlot(ctx context.Context) (uint64, error) {
	iter, err := u.contract.FilterConfirmedBeaconBlock(
		&bind.FilterOpts{Context: ctx},
		nil,
		nil,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to filter ConfirmedBeaconBlock events: %w", err)
	}
	defer iter.Close()

	var latest uint64
	for iter.Next() {
		if slot := iter.Event.Slot; slot > latest {
			latest = slot
		}
	}
	if err := iter.Error(); err != nil {
		return 0, err
	}
	return latest, nil
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
