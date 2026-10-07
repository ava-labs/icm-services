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

// mockBeaconClient serves SSZ-encoded mock beacon blocks and states by slot.
type mockBeaconClient struct {
	blocks map[uint64][]byte
	states map[uint64][]byte
}

func (m *mockBeaconClient) Block(_ context.Context, slot uint64) ([]byte, error) {
	b, ok := m.blocks[slot]
	if !ok {
		return nil, errors.New("no block at slot")
	}
	return b, nil
}

func (m *mockBeaconClient) State(_ context.Context, slot uint64) ([]byte, error) {
	s, ok := m.states[slot]
	if !ok {
		return nil, errors.New("no state at slot")
	}
	return s, nil
}

// mockFixture is a synthetic anchor block, anchor state, and target state,
// chained the way real beacon data is. The mock client serves their SSZ
// encodings and the roots are stored for asserting against the built proof.
type mockFixture struct {
	client          *mockBeaconClient
	anchorSlot      uint64
	targetSlot      uint64
	anchorBlockRoot common.Hash
	anchorStateRoot common.Hash
	targetStateRoot common.Hash
	receiptsRoot    common.Hash
}

// Builds and verifies an execution proof for the given mock fixture.
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
}

// A non-confirmed anchor beacon block root must be rejected.
func TestBuildExecutionProofForSlotsWrongAnchorRoot(t *testing.T) {
	f := newMockFixture(t)

	_, err := BuildExecutionProofForSlots(
		context.Background(), f.client, f.anchorSlot, f.targetSlot, common.HexToHash("0xbad"))
	require.ErrorContains(t, err, "expected confirmed root")
}

// Invalid slot windows must be rejected.
func TestBuildExecutionProofForSlotsRejectsInvalidWindow(t *testing.T) {
	client := &mockBeaconClient{}

	// Target slot at or after the anchor slot.
	_, err := BuildExecutionProofForSlots(context.Background(), client, 100, 100, common.Hash{})
	require.ErrorContains(t, err, "must be before")

	// Target slot outside the anchor's state_roots window.
	_, err = BuildExecutionProofForSlots(context.Background(), client, 10_000, 100, common.Hash{})
	require.ErrorContains(t, err, "state_roots window")
}

// Fetch failures must surface with the slot that failed.
func TestBuildExecutionProofForSlotsFetchFailure(t *testing.T) {
	f := newMockFixture(t)
	delete(f.client.states, f.targetSlot)

	_, err := BuildExecutionProofForSlots(
		context.Background(), f.client, f.anchorSlot, f.targetSlot, f.anchorBlockRoot)
	require.ErrorContains(t, err, "failed to fetch target state at slot 100")
}

// An anchor state whose history vector does not hold the target state's root
// must fail when the target state tree is verified against it.
func TestBuildExecutionProofForSlotsBrokenHistoryLink(t *testing.T) {
	f := newMockFixture(t)

	// Rebuild the anchor state without the target root planted.
	anchorState := minimalBeaconState(t, f.anchorSlot)
	anchorStateRoot, err := anchorState.HashTreeRoot()
	require.NoError(t, err)
	block := minimalSignedBeaconBlock(t, f.anchorSlot, anchorStateRoot)
	anchorBlockRoot, err := block.Message.HashTreeRoot()
	require.NoError(t, err)

	f.client.states[f.anchorSlot], err = anchorState.MarshalSSZ()
	require.NoError(t, err)
	f.client.blocks[f.anchorSlot], err = block.MarshalSSZ()
	require.NoError(t, err)

	// Since the minimalBeaconState contains a state_roots vector of all zeroes, the
	// target state root will not be found in the anchor state, and the proof will fail.
	_, err = BuildExecutionProofForSlots(
		context.Background(), f.client, f.anchorSlot, f.targetSlot, common.Hash(anchorBlockRoot))
	require.ErrorContains(t, err, "lite state root mismatch")
}

// A block fetch failure must surface with the slot that failed.
func TestBuildExecutionProofForSlotsBlockFetchFailure(t *testing.T) {
	f := newMockFixture(t)
	delete(f.client.blocks, f.anchorSlot)

	_, err := BuildExecutionProofForSlots(
		context.Background(), f.client, f.anchorSlot, f.targetSlot, f.anchorBlockRoot)
	require.ErrorContains(t, err, "failed to fetch anchor block at slot 200")
}

// An anchor state that is not the one the anchor block commits to must fail
// when the anchor state tree is verified against the block's state_root.
func TestBuildExecutionProofForSlotsAnchorStateMismatch(t *testing.T) {
	f := newMockFixture(t)

	// Serve a different anchor state than the block's state_root commits to.
	otherState := minimalBeaconState(t, f.anchorSlot)
	otherState.GenesisTime = 999
	var err error
	f.client.states[f.anchorSlot], err = otherState.MarshalSSZ()
	require.NoError(t, err)

	_, err = BuildExecutionProofForSlots(
		context.Background(), f.client, f.anchorSlot, f.targetSlot, f.anchorBlockRoot)
	require.ErrorContains(t, err, "root mismatch")
}

// The state_roots window is inclusive, matching the contract: a gap of
// exactly StateRootsVectorSize slots succeeds, one more fails.
func TestBuildExecutionProofForSlotsWindowBoundary(t *testing.T) {
	const targetSlot = uint64(100)

	f := newMockFixtureWithSlots(t, targetSlot+StateRootsVectorSize, targetSlot)
	_, err := BuildExecutionProofForSlots(
		context.Background(), f.client, f.anchorSlot, f.targetSlot, f.anchorBlockRoot)
	require.NoError(t, err)

	_, err = BuildExecutionProofForSlots(
		context.Background(), &mockBeaconClient{}, targetSlot+StateRootsVectorSize+1, targetSlot, common.Hash{})
	require.ErrorContains(t, err, "state_roots window")
}

// newMockFixture builds the default fixture: anchor slot 200, target slot 100.
func newMockFixture(t *testing.T) *mockFixture {
	return newMockFixtureWithSlots(t, 200, 100)
}

// newMockFixture builds a mock fixture for a mock beacon client and the synthetic beacon data it serves.
func newMockFixtureWithSlots(t *testing.T, anchorSlot, targetSlot uint64) *mockFixture {
	t.Helper()

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
