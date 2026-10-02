// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proofs

import (
	"context"
	"fmt"

	zkadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKAdapter"
	"github.com/ava-labs/icm-services/messages/zk/proofs/beacon"
	"github.com/ava-labs/libevm/common"
)

// beaconDataSource specifies that whatever beacon client is passed in
// must implement the Block and State methods, which return the SSZ encoded
// SignedBeaconBlock and BeaconState respectively. See /beacon/beacon_client.go
// for the implementation.
type beaconDataSource interface {
	Block(ctx context.Context, slot uint64) ([]byte, error)
	State(ctx context.Context, slot uint64) ([]byte, error)
}

// BuildExecutionProofForSlots builds the complete execution proof linking
// targetSlot's receipts root to the confirmed anchor beacon block root at
// anchorSlot.
//
// It fetches the anchor block and both beacon states (anchor and target), reduces
// the states to lite trees, and runs the proof builder over them. anchorBlockRoot
// is the root the ZKAdapter has confirmed for anchorSlot. Every tree is
// verified against this chain of trust as it is built, so wrong or
// inconsistent beacon data fails here rather than on-chain.
//
// The two state fetches are ~100-200MB each.
//
// TODO: cache the anchor's parsed lite state keyed by slot. Follow up work.
func BuildExecutionProofForSlots(
	ctx context.Context,
	client beaconDataSource,
	anchorSlot uint64,
	targetSlot uint64,
	anchorBlockRoot common.Hash,
) (*zkadapter.ExecutionProof, error) {
	// Reject invalid slot windows.
	if targetSlot >= anchorSlot {
		return nil, fmt.Errorf("target slot %d must be before anchor slot %d", targetSlot, anchorSlot)
	}

	// Reject an out of bounds target slot in the anchor state.
	if anchorSlot-targetSlot > StateRootsVectorSize {
		return nil, fmt.Errorf(
			"target slot %d is outside the anchor slot %d's state_roots window (%d slots)",
			targetSlot, anchorSlot, StateRootsVectorSize)
	}

	// Get the SSZ-encoded anchor beacon block and parse it into a tree.
	blockSSZ, err := client.Block(ctx, anchorSlot)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch anchor block at slot %d: %w", anchorSlot, err)
	}
	anchorBlockTree, anchorStateRoot, err := beacon.ParseBlockTree(blockSSZ)
	if err != nil {
		return nil, err
	}

	// Get the SSZ-encoded anchor becon state and parse it into a lite tree.
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

	// Get the SSZ-encoded target beacon state and parse it into a lite tree.
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
