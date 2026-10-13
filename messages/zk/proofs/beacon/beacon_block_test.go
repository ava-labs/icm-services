// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package beacon

import (
	"testing"

	"github.com/OffchainLabs/go-bitfield"
	"github.com/attestantio/go-eth2-client/spec/altair"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/capella"
	"github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/electra"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
)

const testGIndexBlockStateRoot = 11

// The block tree built from SSZ bytes must hash to the same root attestantio
// computes for the block.
func TestParseBlockTreeMatchesTypedRoot(t *testing.T) {
	signedBlock := minimalSignedBeaconBlock(t, 100, phase0.Root{0xaa})
	expectedRoot, err := signedBlock.Message.HashTreeRoot()
	require.NoError(t, err)

	blockSSZ, err := signedBlock.MarshalSSZ()
	require.NoError(t, err)

	tree, stateRoot, err := ParseBlockTree(blockSSZ)
	require.NoError(t, err)
	require.Equal(t, expectedRoot[:], tree.Hash())
	require.Equal(t, [32]byte(signedBlock.Message.StateRoot), stateRoot)

	leaf, err := tree.Get(testGIndexBlockStateRoot)
	require.NoError(t, err)
	require.Equal(t, signedBlock.Message.StateRoot[:], leaf.Hash())
}

// Malformed bytes must be rejected at deserialization.
func TestParseBlockTreeRejectsGarbage(t *testing.T) {
	_, _, err := ParseBlockTree([]byte{0x01, 0x02, 0x03})
	require.ErrorContains(t, err, "failed to deserialize beacon block")
}

// minimalSignedBeaconBlock builds the smallest signed block attestantio will
// serialize and hash, at the given slot and with the given state_root.
// TODO: This is duplicated in the proofs package's tests pending package consolidation.
// See Issue: https://github.com/ava-labs/icm-services/issues/1541
func minimalSignedBeaconBlock(t *testing.T, slot uint64, stateRoot phase0.Root) *electra.SignedBeaconBlock {
	t.Helper()
	return &electra.SignedBeaconBlock{
		Message: &electra.BeaconBlock{
			Slot:          phase0.Slot(slot),
			ProposerIndex: 7,
			ParentRoot:    phase0.Root{0x01},
			StateRoot:     stateRoot,
			Body: &electra.BeaconBlockBody{
				ETH1Data: &phase0.ETH1Data{
					DepositRoot: phase0.Root{0x02},
					BlockHash:   make([]byte, 32),
				},
				ProposerSlashings: []*phase0.ProposerSlashing{},
				AttesterSlashings: []*electra.AttesterSlashing{},
				Attestations:      []*electra.Attestation{},
				Deposits:          []*phase0.Deposit{},
				VoluntaryExits:    []*phase0.SignedVoluntaryExit{},
				SyncAggregate: &altair.SyncAggregate{
					SyncCommitteeBits: bitfield.NewBitvector512(),
				},
				ExecutionPayload: &deneb.ExecutionPayload{
					BaseFeePerGas: uint256.NewInt(1),
					Transactions:  []bellatrix.Transaction{},
					Withdrawals:   []*capella.Withdrawal{},
				},
				BLSToExecutionChanges: []*capella.SignedBLSToExecutionChange{},
				BlobKZGCommitments:    []deneb.KZGCommitment{},
				ExecutionRequests: &electra.ExecutionRequests{
					Deposits:       []*electra.DepositRequest{},
					Withdrawals:    []*electra.WithdrawalRequest{},
					Consolidations: []*electra.ConsolidationRequest{},
				},
			},
		},
	}
}
