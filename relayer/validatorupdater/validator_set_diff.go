// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validatorupdater

import (
	"errors"
	"fmt"
	"math"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/codec/linearcodec"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/message"
)

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
// serialized holds the full codec output after [initializeValidatorSetDiff].
type ValidatorSetDiff struct {
	serialized []byte `serialize:"false"`

	BlockchainID      ids.ID `serialize:"true" json:"blockchainID"`
	PreviousHeight    uint64 `serialize:"true" json:"previousHeight"`
	PreviousTimestamp uint64 `serialize:"true" json:"previousTimestamp"`
	CurrentHeight     uint64 `serialize:"true" json:"currentHeight"`
	CurrentTimestamp  uint64 `serialize:"true" json:"currentTimestamp"`
	Changes           []ValidatorChange `serialize:"true" json:"changes"`
	NumAdded          uint32            `serialize:"true" json:"numAdded"`
}

func (v *ValidatorSetDiff) Bytes() []byte {
	return v.serialized
}

func (v *ValidatorSetDiff) initialize(b []byte) {
	v.serialized = b
}

type validatorSetDiffPayload interface {
	Bytes() []byte
	initialize([]byte)
}

const validatorSetDiffCodecVersion = 0

// validatorSetDiffCodec registers the same first six warp message types as
// current avalanchego development branches, so ValidatorSetDiff keeps type ID 5.
var validatorSetDiffCodec codec.Manager

func init() {
	validatorSetDiffCodec = codec.NewManager(math.MaxInt)
	lc := linearcodec.NewDefault()
	err := errors.Join(
		lc.RegisterType(&message.SubnetToL1Conversion{}),
		lc.RegisterType(&message.RegisterL1Validator{}),
		lc.RegisterType(&message.L1ValidatorRegistration{}),
		lc.RegisterType(&message.L1ValidatorWeight{}),
		lc.RegisterType(&ValidatorSetMetadata{}),
		lc.RegisterType(&ValidatorSetDiff{}),
		validatorSetDiffCodec.RegisterCodec(validatorSetDiffCodecVersion, lc),
	)
	if err != nil {
		panic(err)
	}
}

// NewValidatorSetDiff builds and serializes a ValidatorSetDiff payload (same
// role as message.NewValidatorSetDiff on newer avalanchego).
func NewValidatorSetDiff(
	blockchainID ids.ID,
	previousHeight uint64,
	previousTimestamp uint64,
	currentHeight uint64,
	currentTimestamp uint64,
	changes []ValidatorChange,
	numAdded uint32,
) (*ValidatorSetDiff, error) {
	msg := &ValidatorSetDiff{
		BlockchainID:      blockchainID,
		PreviousHeight:    previousHeight,
		PreviousTimestamp: previousTimestamp,
		CurrentHeight:     currentHeight,
		CurrentTimestamp:  currentTimestamp,
		Changes:           changes,
		NumAdded:          numAdded,
	}
	return msg, initializeValidatorSetDiff(msg)
}

func initializeValidatorSetDiff(v *ValidatorSetDiff) error {
	var p validatorSetDiffPayload = v
	bytes, err := validatorSetDiffCodec.Marshal(validatorSetDiffCodecVersion, &p)
	if err != nil {
		return fmt.Errorf("couldn't marshal ValidatorSetDiff payload: %w", err)
	}
	v.initialize(bytes)
	return nil
}
