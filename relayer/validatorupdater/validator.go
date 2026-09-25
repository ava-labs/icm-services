// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validatorupdater

import (
	"bytes"
	"context"
	"fmt"
	"sort"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/icm-services/peers/clients"
)

// Validator matches [github.com/ava-labs/avalanchego/vms/platformvm/warp/message.Validator]
// on branches that define it; kept here so icm-services can stay on a release
// avalanchego version while producing the same shard bytes.
type Validator struct {
	UncompressedPublicKeyBytes [96]byte `serialize:"true"`
	Weight                     uint64   `serialize:"true"`
}

// SortValidators sorts validators in-place by ascending lexicographic order of their
// uncompressed BLS public key bytes. This matches the canonical order required
// by both the contracts and the signature aggregator.
func SortValidators(validators []*Validator) {
	sort.Slice(validators, func(i, j int) bool {
		return bytes.Compare(
			validators[i].UncompressedPublicKeyBytes[:],
			validators[j].UncompressedPublicKeyBytes[:],
		) < 0
	})
}

// ValidatorsFromWarpSet converts a canonical warp validator set into the Merkle
// leaf representation, preserving the canonical ordering.
//
// The canonical set (see [validators.FlattenValidatorSet]) omits validators
// without a registered BLS key and merges validators that share a BLS key into
// a single entry with their weights summed. This is the index space that the
// signature aggregator's BitSetSignature refers to, so the Merkle tree and the
// attestation must be built over exactly this list. Building them from the
// per-node P-chain view instead would misalign the signer bitset as soon as
// two nodes share a key or a node has no key.
func ValidatorsFromWarpSet(warpSet validators.WarpSet) []*Validator {
	vdrs := make([]*Validator, len(warpSet.Validators))
	for i, vdr := range warpSet.Validators {
		vdrs[i] = &Validator{
			UncompressedPublicKeyBytes: [96]byte(vdr.PublicKey.Serialize()),
			Weight:                     vdr.Weight,
		}
	}
	return vdrs
}

// FetchCanonicalValidators returns [subnetID]'s canonical warp validator set at
// [pChainHeight] as Merkle leaves, in canonical order. The per-node validator
// set is flattened with the same [validators.FlattenValidatorSet] the P-chain
// uses to serve canonical validator sets to the signature aggregator.
func FetchCanonicalValidators(
	ctx context.Context,
	pChainClient clients.CanonicalValidatorState,
	subnetID ids.ID,
	pChainHeight uint64,
) ([]*Validator, error) {
	nodeValidators, err := pChainClient.GetValidatorsAt(ctx, subnetID, pChainHeight)
	if err != nil {
		return nil, fmt.Errorf("failed to get validators for subnet %s at height %d: %w",
			subnetID, pChainHeight, err)
	}
	warpSet, err := validators.FlattenValidatorSet(nodeValidators)
	if err != nil {
		return nil, fmt.Errorf("failed to flatten validators for subnet %s at height %d: %w",
			subnetID, pChainHeight, err)
	}
	return ValidatorsFromWarpSet(warpSet), nil
}
