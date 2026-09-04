// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validatorupdater

import (
	"bytes"
	"sort"
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
