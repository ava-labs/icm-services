// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package beacon

import (
	"fmt"

	electra "github.com/attestantio/go-eth2-client/spec/electra"
	ssz "github.com/ferranbt/fastssz"
)

// numBlockFieldLeaves pads the BeaconBlock's 5 fields to 2^3 = 8 leaves.
// https://github.com/ethereum/consensus-specs/blob/master/specs/phase0/beacon-chain.md#beaconblock
const numBlockFieldLeaves = 8

// ParseBlockTree deserializes an SSZ-encoded SignedBeaconBlock and builds the
// Merkle tree of its BeaconBlock. Note that a "lite tree" (a memory-efficient Merkle tree)
// is not needed here because we are not expanding any of the anchor beacon block's subtrees.
//
// Note the block's state_root field, which is the Merkle root of the anchor beacon state,
// is also returned.
func ParseBlockTree(blockSSZ []byte) (*ssz.Node, [32]byte, error) {
	var signedBlock electra.SignedBeaconBlock
	if err := signedBlock.UnmarshalSSZ(blockSSZ); err != nil {
		return nil, [32]byte{}, fmt.Errorf("failed to deserialize beacon block: %w", err)
	}
	block := signedBlock.Message

	bodyRoot, err := block.Body.HashTreeRoot()
	if err != nil {
		return nil, [32]byte{}, fmt.Errorf("failed to hash block body: %w", err)
	}

	fieldRoots := make([][32]byte, numBlockFieldLeaves)
	fieldRoots[0] = uint64Root(uint64(block.Slot))
	fieldRoots[1] = uint64Root(uint64(block.ProposerIndex))
	fieldRoots[2] = block.ParentRoot
	fieldRoots[3] = block.StateRoot
	fieldRoots[4] = bodyRoot

	tree, err := ssz.TreeFromChunks(toByteSlices(fieldRoots))
	if err != nil {
		return nil, [32]byte{}, fmt.Errorf("failed to build block tree: %w", err)
	}
	return tree, block.StateRoot, nil
}
