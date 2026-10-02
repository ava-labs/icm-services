// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package beacon

import (
	"fmt"

	electra "github.com/attestantio/go-eth2-client/spec/electra"
	ssz "github.com/ferranbt/fastssz"
)

// numBlockFieldLeaves pads the BeaconBlock's 5 fields to the next power of 2,
// which is 8.
// https://github.com/ethereum/consensus-specs/blob/master/specs/phase0/beacon-chain.md#beaconblock
const numBlockFieldLeaves = 8

// ParseBlockTree deserializes an SSZ encoded SignedBeaconBlock and builds the
// Merkle tree of the inner BeaconBlock message.
//
// The block's state_root field in raw bytes is also returned, which is the expected root
// the anchor state tree is verified against.
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
