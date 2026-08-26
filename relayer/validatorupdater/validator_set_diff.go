// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validatorupdater

import "github.com/ava-labs/avalanchego/ids"

// ValidatorChange is a single validator addition, removal, or modification in a
// ValidatorSetDiff. Weight is the current weight (0 for removals).
// Wire layout matches platformvm/warp/message.ValidatorChange on avalanchego
// branches that define it.
type ValidatorChange struct {
	UncompressedPublicKeyBytes [96]byte `serialize:"true" json:"publicKey"`
	Weight                     uint64   `serialize:"true" json:"weight"`
}

// ValidatorSetDiff is the warp payload for validator set diff updates.
// Wire layout (linear codec type ID 5) matches platformvm/warp/message on
// avalanchego branches that register … ValidatorSetDiff in sixth position.
//
// Nothing builds or sends this payload any more, but the type is still
// registered by [merkleCommitmentCodec] to hold codec position 5, which is what
// keeps ValidatorSetMerkleCommitment at type ID 6. Removing it would silently
// renumber the merkle payload and break the wire format.
type ValidatorSetDiff struct {
	BlockchainID      ids.ID            `serialize:"true" json:"blockchainID"`
	PreviousHeight    uint64            `serialize:"true" json:"previousHeight"`
	PreviousTimestamp uint64            `serialize:"true" json:"previousTimestamp"`
	CurrentHeight     uint64            `serialize:"true" json:"currentHeight"`
	CurrentTimestamp  uint64            `serialize:"true" json:"currentTimestamp"`
	Changes           []ValidatorChange `serialize:"true" json:"changes"`
	NumAdded          uint32            `serialize:"true" json:"numAdded"`
}
