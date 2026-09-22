// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proofs

import (
	"fmt"

	zkadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKAdapter"
	"github.com/ava-labs/libevm/common"
	ssz "github.com/ferranbt/fastssz"
)

// Fulu generalized indices into the beacon block state tree.
// These constants must match the contract-level deployed config set in ZKStateManager during construction,
// otherwise proofs built here verify client-side but revert on-chain.
//
// See: https://github.com/ava-labs/icm-services/blob/41c5606ff428c2110b243e281d774c7fb51c256a/icm-contracts/avalanche/verifiers/ethereum/ZKAdapter.sol#L42
//
//nolint:lll
const (
	// gIndexBlockStateRoot locates state_root within a BeaconBlock.
	gIndexBlockStateRoot = 11
	// gIndexBaseStateRoots locates the state_roots vector within a BeaconState.
	gIndexBaseStateRoots = 70
	// stateRootsDepth is the depth of the state_roots vector's own subtree,
	// 8192 roots = 2^13 leaves.
	stateRootsDepth = 13
	// gIndexExecPayloadHeader locates latest_execution_payload_header within a
	// BeaconState.
	gIndexExecPayloadHeader = 88
	// gIndexReceiptsRoot locates receipts_root within an ExecutionPayloadHeader.
	gIndexReceiptsRoot = 35
	// StateRootsVectorSize is the beacon state_roots vector length: an anchor
	// can prove any slot within this many slots behind it.
	StateRootsVectorSize = 8192
)

// BuildExecutionProof builds the four SSZ Merkle proofs that the ZKAdapter's
// Execution.verify consumes, returning the generated binding struct
// zkadapter.ExecutionProof directly so any contract level changes surface as compile
// errors here. These steps exactly mirror the contract-side verification flow:
//
//	Chain of trust: Trusted Beacon Block -> Anchor Beacon State -> Target Inter-Epoch Beacon State -> Execution Header -> Receipts Root
//	1. Anchor check: Validates an anchor beacon state against a trusted block root.
//	2. History check: Validates the target beacon state against the anchor beacon state's `state_roots` history vector.
//	3. Execution check: Validates the execution payload header against the target beacon state.
//	4. Receipts check: Validates the receipts root against the execution payload header.
//
// See: https://github.com/ava-labs/icm-services/blob/41c5606ff428c2110b243e281d774c7fb51c256a/icm-contracts/avalanche/verifiers/ethereum/StateManagerLibrary.sol
//
//nolint:lll
func BuildExecutionProof(
	anchorBlock *ssz.Node,
	anchorState *ssz.Node,
	targetState *ssz.Node,
	anchorBlockRoot common.Hash,
	anchorSlot uint64,
	targetSlot uint64,
) (*zkadapter.ExecutionProof, error) {
	// Safety checks
	if targetSlot >= anchorSlot {
		return nil, fmt.Errorf("target slot %d must be before anchor slot %d", targetSlot, anchorSlot)
	}
	if anchorSlot-targetSlot > StateRootsVectorSize {
		return nil, fmt.Errorf(
			"target slot (%d) is outside the anchor slot's (%d) state_roots window (%d slots)",
			targetSlot, anchorSlot, StateRootsVectorSize)
	}

	// 1. Anchor state proof: trusted anchor beacon block root -> anchor beacon state root, where the beacon block
	// root is stored on-chain and considered trusted.
	anchorStateProof, err := proveAgainst(anchorBlock, gIndexBlockStateRoot, anchorBlockRoot)
	if err != nil {
		return nil, fmt.Errorf("anchor state proof generation failed: %w", err)
	}
	anchorStateRoot := common.BytesToHash(anchorStateProof.Leaf)

	// 2. History proof: anchor beacon state root -> target beacon state root. Prove that the target beacon state root is
	// in the anchor beacon state's history using the G-index and an SSZ Merkle proof. This is possible since beacon states contain
	// a vector of historical state roots `state_roots` referencing the last 8192 slots.

	// First, calculate the specific G-Index for 'state_roots[targetSlot]' within the beacon state SSZ structure.
	vectorIndex := int(targetSlot % StateRootsVectorSize)
	targetGIndex := (gIndexBaseStateRoots << stateRootsDepth) + vectorIndex
	targetStateProof, err := proveAgainst(anchorState, targetGIndex, anchorStateRoot)
	if err != nil {
		return nil, fmt.Errorf("history proof generation failed: %w", err)
	}
	targetStateRoot := common.BytesToHash(targetStateProof.Leaf)

	// 3. Execution proof: target beacon state root -> execution payload header root.
	// Verify that the execution header root is in the target beacon state.
	execHeaderProof, err := proveAgainst(targetState, gIndexExecPayloadHeader, targetStateRoot)
	if err != nil {
		return nil, fmt.Errorf("execution header proof generation failed: %w", err)
	}
	execHeaderRoot := common.BytesToHash(execHeaderProof.Leaf)

	// 4. Receipts proof: execution payload header root -> receipts root.

	// First, get the execution header node from the target beacon state tree.
	execHeaderNode, err := targetState.Get(gIndexExecPayloadHeader)
	if err != nil {
		return nil, fmt.Errorf("failed to extract execution payload header subtree: %w", err)
	}
	receiptsProof, err := proveAgainst(execHeaderNode, gIndexReceiptsRoot, execHeaderRoot)
	if err != nil {
		return nil, fmt.Errorf("receipts root proof generation failed: %w", err)
	}

	// Convert proof siblings to the contract's bytes32[] type representation.
	anchorProofHashes, err := toHash32Slice(anchorStateProof.Hashes)
	if err != nil {
		return nil, fmt.Errorf("anchor state proof: %w", err)
	}
	targetProofHashes, err := toHash32Slice(targetStateProof.Hashes)
	if err != nil {
		return nil, fmt.Errorf("history proof: %w", err)
	}
	execProofHashes, err := toHash32Slice(execHeaderProof.Hashes)
	if err != nil {
		return nil, fmt.Errorf("execution header proof: %w", err)
	}
	receiptsProofHashes, err := toHash32Slice(receiptsProof.Hashes)
	if err != nil {
		return nil, fmt.Errorf("receipts root proof: %w", err)
	}

	return &zkadapter.ExecutionProof{
		AnchorSlot:                 anchorSlot,
		TargetSlot:                 targetSlot,
		AnchorBeaconStateRoot:      anchorStateRoot,
		AnchorBeaconStateProof:     anchorProofHashes,
		TargetBeaconStateRoot:      targetStateRoot,
		TargetBeaconStateProof:     targetProofHashes,
		TargetExecutionHeaderRoot:  execHeaderRoot,
		TargetExecutionHeaderProof: execProofHashes,
		TargetReceiptsRoot:         common.BytesToHash(receiptsProof.Leaf),
		TargetReceiptsProof:        receiptsProofHashes,
	}, nil
}

// toHash32Slice converts proof siblings to fixed 32-byte hashes, mirroring the
// contract's bytes32[] proof representation.
func toHash32Slice(hashes [][]byte) ([][32]byte, error) {
	out := make([][32]byte, len(hashes))
	for i, h := range hashes {
		if len(h) != 32 {
			return nil, fmt.Errorf("proof sibling %d has length %d, want 32", i, len(h))
		}
		copy(out[i][:], h)
	}
	return out, nil
}

// proveAgainst computes the SSZ Merkle proof for the provided gIndex
// and verifies it against the supplied expectedRoot.
func proveAgainst(node *ssz.Node, gIndex int, expectedRoot common.Hash) (*ssz.Proof, error) {
	proof, err := node.Prove(gIndex)
	if err != nil {
		return nil, fmt.Errorf("failed to prove gindex %d: %w", gIndex, err)
	}
	ok, err := ssz.VerifyProof(expectedRoot.Bytes(), proof)
	if err != nil {
		return nil, fmt.Errorf("failed to verify proof for gindex %d: %w", gIndex, err)
	}
	if !ok {
		return nil, fmt.Errorf(
			"proof for gindex %d does not verify against expected root %s (wrong or mismatched input tree)",
			gIndex, expectedRoot.Hex())
	}
	return proof, nil
}
