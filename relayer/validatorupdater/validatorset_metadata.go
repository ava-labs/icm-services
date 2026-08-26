// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validatorupdater

import (
	"bytes"
	"sort"

	"github.com/ava-labs/avalanchego/ids"
)

// Validator matches [github.com/ava-labs/avalanchego/vms/platformvm/warp/message.Validator]
// on branches that define it; kept here so icm-services can stay on a release
// avalanchego version while producing the same shard bytes.
type Validator struct {
	UncompressedPublicKeyBytes [96]byte `serialize:"true"`
	Weight                     uint64   `serialize:"true"`
}

// ValidatorSetMetadata is the warp payload for subset validator set registration.
// Wire layout (linear codec type ID 4) matches platformvm/warp/message on
// avalanchego branches that register SubnetToL1Conversion … ValidatorSetMetadata
// in that order.
//
// Nothing builds or sends this payload any more, but the type is still registered
// by [merkleCommitmentCodec] to hold codec position 4. See [ValidatorSetDiff] for
// why those positions must not shift.
type ValidatorSetMetadata struct {
	BlockchainID    ids.ID   `serialize:"true" json:"blockchainID"`
	PChainHeight    uint64   `serialize:"true" json:"pChainHeight"`
	PChainTimestamp uint64   `serialize:"true" json:"pChainTimestamp"`
	ShardHashes     []ids.ID `serialize:"true" json:"shardHashes"`
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
