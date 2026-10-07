// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package proofs

import (
	"context"
	"fmt"

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

// BuildExecutionProofForSlots builds the complete execution proof linking
// the targetSlot's receipts root to the confirmed anchor beacon block root at
// anchorSlot.
//
// In more detail, this function performs the following tasks:
// 1. Fetches the anchor beacon block and both beacon states (anchor and target)
// from the beaconClient. Note the anchorBlockRoot is the root ZKAdapter has confirmed for the
// anchorSlot.
// 2. Parse the anchor beacon block into a regular Merkle tree, and the two beacon states into
// lite Merkle trees (to save memory).
// 3. Use the execution proof builder to create a proof from the parsed trees and provided slots.
// The execution proof builder will verify each tree against the expected chain of trust.
//
// TODO: We can cache the anchor's parsed lite state. The reason is that a single anchor beacon state
// may be used to verify multiple target slots within the 8192-slot window of the anchor state's
// state_roots vector. Issue: https://github.com/ava-labs/icm-services/issues/1542
func BuildExecutionProofForSlots(
	ctx context.Context,
	client beaconClient,
	anchorSlot uint64,
	targetSlot uint64,
	anchorBlockRoot common.Hash,
) (*zkadapter.ExecutionProof, error) {
	// Safety check
	if err := validateSlotWindow(anchorSlot, targetSlot); err != nil {
		return nil, err
	}

	// Get the SSZ-encoded anchor beacon block from the beaconClient and parse it into a tree.
	blockSSZ, err := client.Block(ctx, anchorSlot)
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

	// Get the SSZ-encoded anchor beacon state from the beaconClient and parse it into a lite tree.
	anchorStateSSZ, err := client.State(ctx, anchorSlot)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch anchor state at slot %d: %w", anchorSlot, err)
	}
	anchorLite, err := beacon.ParseLiteBeaconState(anchorStateSSZ)
	if err != nil {
		return nil, fmt.Errorf("failed to parse anchor state: %w", err)
	}
	anchorStateTree, err := anchorLite.AnchorStateTree(anchorStateRoot)
	if err != nil {
		return nil, fmt.Errorf("anchor state at slot %d: %w", anchorSlot, err)
	}

	// Get the SSZ-encoded target beacon state from the beaconClient and parse it into a lite tree.
	targetStateSSZ, err := client.State(ctx, targetSlot)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch target state at slot %d: %w", targetSlot, err)
	}
	targetLite, err := beacon.ParseLiteBeaconState(targetStateSSZ)
	if err != nil {
		return nil, fmt.Errorf("failed to parse target state: %w", err)
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
