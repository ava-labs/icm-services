// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package proofs

import (
	"context"
	"fmt"

	"github.com/ava-labs/avalanchego/cache/lru"
	zkadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKAdapter"
	"github.com/ava-labs/icm-services/messages/zk/proofs/beacon"
	"github.com/ava-labs/libevm/common"
)

// Overview: orchestrator.go connects the three main individual components of the proofs package.
// 1) the beacon client, which fetches SSZ-encoded beacon blocks and states by slot
// 2) the lite beacon tree builder, which parses the SSZ into lite (memory-efficient) Merkle trees
// 3) the execution proof builder, which produces a proof from the anchor block root to the target receipts root
//
// BuildExecutionProofForSlots is the main entry point for this orchestrator, which takes in a beacon client,
// an anchor slot, target slot, and the anchor block root, and returns an execution proof ready to be sent to the
// contract.
// TODO: Export in future PR.
type beaconClient interface {
	Block(ctx context.Context, slot uint64) ([]byte, error)
	State(ctx context.Context, slot uint64) ([]byte, error)
}

// liteStateCacheSize is the number of parsed lite beacon states kept in memory. A single anchor
// beacon state may be used to verify multiple target slots within the 8192-slot window of the
// anchor state's state_roots vector, so caching avoids refetching ~200MB state.
const liteStateCacheSize = 16

// ProofBuilder builds execution proofs from beacon data fetched via the beaconClient and
// cached parsed lite beacon states by slot.
type ProofBuilder struct {
	client     beaconClient
	liteStates *lru.Cache[uint64, *beacon.LiteBeaconState]
}

func NewProofBuilder(client beaconClient) *ProofBuilder {
	return &ProofBuilder{
		client:     client,
		liteStates: lru.NewCache[uint64, *beacon.LiteBeaconState](liteStateCacheSize),
	}
}

// BuildExecutionProofForSlots builds the complete execution proof linking
// the targetSlot's receipts root to the confirmed anchor beacon block root at
// anchorSlot.
//
// In more detail, this function performs the following tasks:
// 1. Fetches the anchor beacon block and both beacon states (anchor and target)
// from the beaconClient. Note the anchorBlockRoot is the root ZKAdapter has confirmed for the
// anchorSlot. Beacon states are served from the lite state cache when present.
// 2. Parse the anchor beacon block into a regular Merkle tree, and the two beacon states into
// lite Merkle trees (to save memory).
// 3. Use the execution proof builder to create a proof from the parsed trees and provided slots.
// The execution proof builder will verify each tree against the expected chain of trust.
//
// TODO: We can cache the anchor's parsed lite state. The reason is that a single anchor beacon state
// may be used to verify multiple target slots within the 8192-slot window of the anchor state's
// state_roots vector. Issue: https://github.com/ava-labs/icm-services/issues/1542
func (b *ProofBuilder) BuildExecutionProofForSlots(
	ctx context.Context,
	anchorSlot uint64,
	targetSlot uint64,
	anchorBlockRoot common.Hash,
) (*zkadapter.ExecutionProof, error) {
	// Safety check
	if err := validateSlotWindow(anchorSlot, targetSlot); err != nil {
		return nil, err
	}

	// Get the SSZ-encoded anchor beacon block from the beaconClient and parse it into a tree.
	blockSSZ, err := b.client.Block(ctx, anchorSlot)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch anchor block at slot %d: %w", anchorSlot, err)
	}
	anchorBlockTree, anchorStateRoot, err := beacon.ParseBlockTree(blockSSZ)
	if err != nil {
		return nil, err
	}

	// Verify the fetched block is the confirmed anchor before fetching ~200MB
	// of state against it.
	if got := common.BytesToHash(anchorBlockTree.Hash()); got != anchorBlockRoot {
		return nil, fmt.Errorf("anchor block at slot %d has root %s, expected confirmed root %s",
			anchorSlot, got, anchorBlockRoot)
	}

	// Get the SSZ-encoded anchor beacon state from the beaconClient or the cache
	// and parse it into a lite tree.
	anchorLite, err := b.getLiteState(ctx, anchorSlot)
	if err != nil {
		return nil, fmt.Errorf("anchor state: %w", err)
	}

	anchorStateTree, err := anchorLite.AnchorStateTree(anchorStateRoot)
	if err != nil {
		return nil, fmt.Errorf("anchor state at slot %d: %w", anchorSlot, err)
	}

	// Get the SSZ-encoded target beacon state from the beaconClient or the cache
	// and parse it into a lite tree.
	targetLite, err := b.getLiteState(ctx, targetSlot)
	if err != nil {
		return nil, fmt.Errorf("target state: %w", err)
	}

	targetStateRoot := anchorLite.StateRootAt(targetSlot)
	targetStateTree, execHeaderTree, err := targetLite.TargetStateTree(targetStateRoot)
	if err != nil {
		return nil, fmt.Errorf("target state at slot %d: %w", targetSlot, err)
	}

	// Use the assembled trees to produce an execution proof from the beacon block root
	// down to the receipts root via the chain of trust.
	return BuildExecutionProof(
		anchorBlockTree,
		anchorStateTree,
		targetStateTree,
		execHeaderTree,
		anchorBlockRoot,
		anchorSlot,
		targetSlot,
	)
}

// liteState returns the parsed lite beacon state at the given slot, fetching the SSZ-encoded
// beacon state from the beaconClient and parsing it on a cache miss.
func (b *ProofBuilder) getLiteState(ctx context.Context, slot uint64) (*beacon.LiteBeaconState, error) {
	if lite, ok := b.liteStates.Get(slot); ok {
		return lite, nil
	}
	stateSSZ, err := b.client.State(ctx, slot)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch beacon state at slot %d: %w", slot, err)
	}
	lite, err := beacon.ParseLiteBeaconState(stateSSZ)
	if err != nil {
		return nil, fmt.Errorf("failed to parse beacon state at slot %d: %w", slot, err)
	}
	b.liteStates.Put(slot, lite)
	return lite, nil
}
