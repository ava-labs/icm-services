// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package proofs

import (
	"context"
	"errors"
	"testing"

	"github.com/OffchainLabs/go-bitfield"
	"github.com/attestantio/go-eth2-client/spec/altair"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/capella"
	"github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/electra"
	"github.com/attestantio/go-eth2-client/spec/fulu"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ava-labs/libevm/common"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
)

// mockBeaconClient serves SSZ-encoded mock beacon blocks and states by slot
// and records which slots were fetched.
type mockBeaconClient struct {
	blocks  map[uint64][]byte
	states  map[uint64][]byte
	fetched []uint64
}

func (m *mockBeaconClient) Block(_ context.Context, slot uint64) ([]byte, error) {
	m.fetched = append(m.fetched, slot)
	b, ok := m.blocks[slot]
	if !ok {
		return nil, errors.New("no block at slot")
	}
	return b, nil
}

func (m *mockBeaconClient) State(_ context.Context, slot uint64) ([]byte, error) {
	m.fetched = append(m.fetched, slot)
	s, ok := m.states[slot]
	if !ok {
		return nil, errors.New("no state at slot")
	}
	return s, nil
}

// mockFixture is a synthetic anchor block, anchor state, and target state,
// linked the way real beacon data is. The anchor block commits to the anchor
// state's root, and the anchor state's state_roots vector holds the target
// state's root at the target slot. The mock client serves their SSZ
// encodings, and the roots are stored for asserting against the built proof.
type mockFixture struct {
	client          *mockBeaconClient
	anchorSlot      uint64
	targetSlot      uint64
	anchorBlockRoot common.Hash
	anchorStateRoot common.Hash
	targetStateRoot common.Hash
	receiptsRoot    common.Hash
}

// Follows the established chain of trust block -> anchor state -> target state -> receipts.
func TestBuildExecutionProofForSlots(t *testing.T) {
	f := newMockFixture(t)

	proof, err := BuildExecutionProofForSlots(
		context.Background(), f.client, f.anchorSlot, f.targetSlot, f.anchorBlockRoot)
	require.NoError(t, err)

	require.Equal(t, f.anchorSlot, proof.AnchorSlot)
	require.Equal(t, f.targetSlot, proof.TargetSlot)
	require.Equal(t, f.anchorStateRoot, common.Hash(proof.AnchorBeaconStateRoot))
	require.Equal(t, f.targetStateRoot, common.Hash(proof.TargetBeaconStateRoot))
	require.Equal(t, f.receiptsRoot, common.Hash(proof.TargetReceiptsRoot))

	// Exactly the block and two states are fetched, anchor first.
	require.Equal(t, []uint64{f.anchorSlot, f.anchorSlot, f.targetSlot}, f.client.fetched)
}

// A confirmed anchor root the fetched block does not hash to must fail at
// the chain's first link.
func TestBuildExecutionProofForSlotsWrongAnchorRoot(t *testing.T) {
	f := newMockFixture(t)

	_, err := BuildExecutionProofForSlots(
		context.Background(), f.client, f.anchorSlot, f.targetSlot, common.HexToHash("0xbad"))
	require.ErrorContains(t, err, "does not verify against expected root")
}

// Invalid slot windows must be rejected before anything is fetched.
func TestBuildExecutionProofForSlotsRejectsWindowBeforeFetching(t *testing.T) {
	client := &mockBeaconClient{}

	_, err := BuildExecutionProofForSlots(context.Background(), client, 100, 100, common.Hash{})
	require.ErrorContains(t, err, "must be before")

	_, err = BuildExecutionProofForSlots(context.Background(), client, 10_000, 100, common.Hash{})
	require.ErrorContains(t, err, "state_roots window")

	require.Empty(t, client.fetched)
}

// Fetch failures must surface with the slot that failed.
func TestBuildExecutionProofForSlotsFetchFailure(t *testing.T) {
	f := newMockFixture(t)
	delete(f.client.states, f.targetSlot)

	_, err := BuildExecutionProofForSlots(
		context.Background(), f.client, f.anchorSlot, f.targetSlot, f.anchorBlockRoot)
	require.ErrorContains(t, err, "failed to fetch target state at slot 100")
}

func newMockFixture(t *testing.T) *mockFixture {
	t.Helper()
	const anchorSlot, targetSlot = uint64(200), uint64(100)

	targetState := minimalBeaconState(t, targetSlot)
	targetStateRoot, err := targetState.HashTreeRoot()
	require.NoError(t, err)

	anchorState := minimalBeaconState(t, anchorSlot)
	anchorState.StateRoots[targetSlot%StateRootsVectorSize] = targetStateRoot
	anchorStateRoot, err := anchorState.HashTreeRoot()
	require.NoError(t, err)

	block := minimalSignedBeaconBlock(t, anchorSlot, anchorStateRoot)
	anchorBlockRoot, err := block.Message.HashTreeRoot()
	require.NoError(t, err)

	blockSSZ, err := block.MarshalSSZ()
	require.NoError(t, err)
	anchorSSZ, err := anchorState.MarshalSSZ()
	require.NoError(t, err)
	targetSSZ, err := targetState.MarshalSSZ()
	require.NoError(t, err)

	return &mockFixture{
		client: &mockBeaconClient{
			blocks: map[uint64][]byte{anchorSlot: blockSSZ},
			states: map[uint64][]byte{anchorSlot: anchorSSZ, targetSlot: targetSSZ},
		},
		anchorSlot:      anchorSlot,
		targetSlot:      targetSlot,
		anchorBlockRoot: common.Hash(anchorBlockRoot),
		anchorStateRoot: common.Hash(anchorStateRoot),
		targetStateRoot: common.Hash(targetStateRoot),
		receiptsRoot:    common.Hash(targetState.LatestExecutionPayloadHeader.ReceiptsRoot),
	}
}

// TODO: This is duplicated code from lite_state_test.go. To resolve this,
// we should flatten the proofs and beacon package into a single package,
// and move the test helpers into a shared testutil package.
// Issue: https://github.com/ava-labs/icm-services/issues/1541

// minimalBeaconState builds the smallest fulu.BeaconState that attestantio
// will serialize and hash
func minimalBeaconState(t *testing.T, slot uint64) *fulu.BeaconState {
	t.Helper()

	syncCommittee := &altair.SyncCommittee{}
	for i := range syncCommittee.Pubkeys {
		syncCommittee.Pubkeys[i] = phase0.BLSPubKey{byte(i)}
	}

	state := &fulu.BeaconState{
		GenesisTime:           1,
		GenesisValidatorsRoot: phase0.Root{0x01},
		Slot:                  phase0.Slot(slot),
		Fork: &phase0.Fork{
			PreviousVersion: phase0.Version{0x01},
			CurrentVersion:  phase0.Version{0x02},
			Epoch:           1,
		},
		LatestBlockHeader: &phase0.BeaconBlockHeader{
			Slot:       phase0.Slot(slot) - 1,
			ParentRoot: phase0.Root{0x02},
			StateRoot:  phase0.Root{0x03},
			BodyRoot:   phase0.Root{0x04},
		},
		BlockRoots:      make([]phase0.Root, 8192),
		StateRoots:      make([]phase0.Root, 8192),
		HistoricalRoots: []phase0.Root{},
		ETH1Data: &phase0.ETH1Data{
			DepositRoot:  phase0.Root{0x05},
			DepositCount: 1,
			BlockHash:    make([]byte, 32),
		},
		ETH1DataVotes:                 []*phase0.ETH1Data{},
		Validators:                    []*phase0.Validator{},
		Balances:                      []phase0.Gwei{},
		RANDAOMixes:                   make([]phase0.Root, 65536),
		Slashings:                     make([]phase0.Gwei, 8192),
		PreviousEpochParticipation:    []altair.ParticipationFlags{},
		CurrentEpochParticipation:     []altair.ParticipationFlags{},
		JustificationBits:             make([]byte, 1),
		PreviousJustifiedCheckpoint:   &phase0.Checkpoint{Epoch: 1, Root: phase0.Root{0x06}},
		CurrentJustifiedCheckpoint:    &phase0.Checkpoint{Epoch: 2, Root: phase0.Root{0x07}},
		FinalizedCheckpoint:           &phase0.Checkpoint{Epoch: 3, Root: phase0.Root{0x08}},
		InactivityScores:              []uint64{},
		CurrentSyncCommittee:          syncCommittee,
		NextSyncCommittee:             syncCommittee,
		LatestExecutionPayloadHeader:  minimalExecHeader(),
		NextWithdrawalIndex:           1,
		NextWithdrawalValidatorIndex:  2,
		DepositRequestsStartIndex:     3,
		DepositBalanceToConsume:       4,
		ExitBalanceToConsume:          5,
		EarliestExitEpoch:             6,
		ConsolidationBalanceToConsume: 7,
		EarliestConsolidationEpoch:    8,
		ProposerLookahead:             make([]phase0.ValidatorIndex, 64),
	}
	return state
}

// minimalSignedBeaconBlock builds the smallest signed block attestantio will
// serialize and hash, at the given slot and with the given state_root.
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

// minimalExecHeader builds a fully-populated execution payload header.
func minimalExecHeader() *deneb.ExecutionPayloadHeader {
	return &deneb.ExecutionPayloadHeader{
		ParentHash:       phase0.Hash32{0x01},
		FeeRecipient:     [20]byte{0x02},
		StateRoot:        phase0.Root{0x03},
		ReceiptsRoot:     phase0.Root{0x04},
		LogsBloom:        [256]byte{0x05},
		PrevRandao:       [32]byte{0x06},
		BlockNumber:      7,
		GasLimit:         8,
		GasUsed:          9,
		Timestamp:        10,
		ExtraData:        []byte{0x0b, 0x0c},
		BaseFeePerGas:    uint256.NewInt(13),
		BlockHash:        phase0.Hash32{0x0e},
		TransactionsRoot: phase0.Root{0x0f},
		WithdrawalsRoot:  phase0.Root{0x10},
		BlobGasUsed:      17,
		ExcessBlobGas:    18,
	}
}
