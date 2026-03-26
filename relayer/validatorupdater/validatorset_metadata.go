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
// serialized holds the full codec output after [initializeValidatorSetMetadata]
// (same role as message's unexported payload slice; not part of the struct encoding).
type ValidatorSetMetadata struct {
	serialized []byte `serialize:"false"`

	BlockchainID    ids.ID   `serialize:"true" json:"blockchainID"`
	PChainHeight    uint64   `serialize:"true" json:"pChainHeight"`
	PChainTimestamp uint64   `serialize:"true" json:"pChainTimestamp"`
	ShardHashes     []ids.ID `serialize:"true" json:"shardHashes"`
}

func (v *ValidatorSetMetadata) Bytes() []byte {
	return v.serialized
}

func (v *ValidatorSetMetadata) initialize(b []byte) {
	v.serialized = b
}

type validatorSetMetadataPayload interface {
	Bytes() []byte
	initialize([]byte)
}

const validatorSetMetadataCodecVersion = 0

// metadataCodec registers the same first five warp message types as current
// avalanchego development branches, so ValidatorSetMetadata keeps type ID 4.
var metadataCodec codec.Manager

func init() {
	metadataCodec = codec.NewManager(math.MaxInt)
	lc := linearcodec.NewDefault()
	err := errors.Join(
		lc.RegisterType(&message.SubnetToL1Conversion{}),
		lc.RegisterType(&message.RegisterL1Validator{}),
		lc.RegisterType(&message.L1ValidatorRegistration{}),
		lc.RegisterType(&message.L1ValidatorWeight{}),
		lc.RegisterType(&ValidatorSetMetadata{}),
		metadataCodec.RegisterCodec(validatorSetMetadataCodecVersion, lc),
	)
	if err != nil {
		panic(err)
	}
}

// NewValidatorSetMetadata builds and serializes a ValidatorSetMetadata payload
// (same role as message.NewValidatorSetMetadata on newer avalanchego).
func NewValidatorSetMetadata(
	blockchainID ids.ID,
	pChainHeight uint64,
	pChainTimestamp uint64,
	shardHashes []ids.ID,
) (*ValidatorSetMetadata, error) {
	msg := &ValidatorSetMetadata{
		BlockchainID:    blockchainID,
		PChainHeight:    pChainHeight,
		PChainTimestamp: pChainTimestamp,
		ShardHashes:     shardHashes,
	}
	return msg, initializeValidatorSetMetadata(msg)
}

func initializeValidatorSetMetadata(v *ValidatorSetMetadata) error {
	var p validatorSetMetadataPayload = v
	bytes, err := metadataCodec.Marshal(validatorSetMetadataCodecVersion, &p)
	if err != nil {
		return fmt.Errorf("couldn't marshal ValidatorSetMetadata payload: %w", err)
	}
	v.initialize(bytes)
	return nil
}
