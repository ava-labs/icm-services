// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package beacon

import (
	"crypto/sha256"
	"testing"

	"github.com/attestantio/go-eth2-client/spec/altair"
	"github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/fulu"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	ssz "github.com/ferranbt/fastssz"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
)

// Gindices the proofs walk to, duplicated here because the beacon package
// does not import proofs: state_roots[i] sits at
// (gIndexBaseStateRoots << stateRootsDepth) + i in the state tree, and
// receipts_root at gindex 35 in the header subtree. The proofs package's
// fixture tests pin these values against mainnet data.
const (
	testGIndexStateRootsBase = 70 << 13
	testGIndexReceiptsRoot   = 35
)

func TestUint64Root(t *testing.T) {
	root := uint64Root(0x0102030405060708)
	// SSZ packs basic values little-endian into the chunk's first bytes.
	require.Equal(t, byte(0x08), root[0])
	require.Equal(t, byte(0x01), root[7])
	require.Equal(t, make([]byte, 24), root[8:32])
}

func TestPackUint64s(t *testing.T) {
	// Five uint64s pack four per chunk: two chunks, second mostly zero.
	chunks := packUint64s([]uint64{1, 2, 3, 4, 5})
	require.Len(t, chunks, 2)
	require.Equal(t, byte(1), chunks[0][0])
	require.Equal(t, byte(4), chunks[0][24])
	require.Equal(t, byte(5), chunks[1][0])
}

// rootsVectorRoot must agree with an independent merkleization of the same
// chunks (fastssz's tree builder as the oracle).
func TestRootsVectorRoot(t *testing.T) {
	roots := make([]phase0.Root, 4)
	for i := range roots {
		roots[i][0] = byte(i + 1)
	}

	tree, err := ssz.TreeFromChunks(toByteSlices(rootsToFixed(roots)))
	require.NoError(t, err)

	require.Equal(t, [32]byte(tree.Hash()), rootsVectorRoot(roots))
}

// rootsListRoot must equal merkleize(chunks padded to limit) mixed with the
// length, computed here from first principles with sha256.
func TestRootsListRoot(t *testing.T) {
	var r phase0.Root
	r[0] = 0xaa

	// One root, limit 2: chunksRoot = sha256(r ++ zero), then mix in the
	// length: sha256(chunksRoot ++ uint256(1)).
	chunksRoot := sha256.Sum256(append(r[:], make([]byte, 32)...))
	lengthChunk := uint64Root(1)
	expected := sha256.Sum256(append(chunksRoot[:], lengthChunk[:]...))

	require.Equal(t, expected, rootsListRoot([]phase0.Root{r}, 2))
}

// merkleizeToLimit must virtually pad to the limit's depth: a single chunk
// with limit 4 hashes against zero at level 0 and the zero-pair hash at
// level 1.
func TestMerkleizeToLimit(t *testing.T) {
	var chunk [32]byte
	chunk[0] = 0xaa

	zero := [32]byte{}
	level0 := sha256.Sum256(append(chunk[:], zero[:]...))
	zeroPair := sha256.Sum256(append(zero[:], zero[:]...))
	expected := sha256.Sum256(append(level0[:], zeroPair[:]...))

	require.Equal(t, expected, merkleizeToLimit([][32]byte{chunk}, 4))

	// An empty list at limit 4 is the pure zero tree of depth 2.
	expectedEmpty := sha256.Sum256(append(zeroPair[:], zeroPair[:]...))
	require.Equal(t, expectedEmpty, merkleizeToLimit(nil, 4))
}

// The per-field decomposition of the execution payload header, merkleized,
// must equal the container's own hash tree root: attestantio's typed hashing
// is the oracle for our field layout.
func TestExecHeaderFieldRootsMatchesTypedRoot(t *testing.T) {
	header := testExecHeader()

	fieldRoots := execHeaderFieldRoots(header)
	require.Len(t, fieldRoots, numExecHeaderLeaves)

	// receipts_root must land at field 3 (gindex 35 in the header subtree).
	require.Equal(t, [32]byte(header.ReceiptsRoot), fieldRoots[3])

	require.Equal(t,
		mustRoot(header.HashTreeRoot()),
		merkleizeChunks(fieldRoots),
	)
}

// The lite trees must reproduce the state root that attestantio's typed
// hashing computes, and must expose the expanded subtrees at the gindices
// the proofs walk to.
func TestLiteStateTreesMatchTypedRoot(t *testing.T) {
	state := minimalBeaconState(t)
	expectedRoot, err := state.HashTreeRoot()
	require.NoError(t, err)

	stateSSZ, err := state.MarshalSSZ()
	require.NoError(t, err)
	lite, err := ParseLiteBeaconState(stateSSZ)
	require.NoError(t, err)

	// Anchor role: state_roots expanded. Assembly includes the root check,
	// so success means the lite tree reproduces the typed root exactly.
	anchorTree, err := lite.AnchorStateTree(expectedRoot)
	require.NoError(t, err)
	leaf, err := anchorTree.Get(testGIndexStateRootsBase + 5)
	require.NoError(t, err)
	require.Equal(t, state.StateRoots[5][:], leaf.Hash())

	// Target role: execution payload header expanded, receipts root at its
	// gindex within the subtree.
	_, execHeaderTree, err := lite.TargetStateTree(expectedRoot)
	require.NoError(t, err)
	receiptsLeaf, err := execHeaderTree.Get(testGIndexReceiptsRoot)
	require.NoError(t, err)
	require.Equal(t, state.LatestExecutionPayloadHeader.ReceiptsRoot[:], receiptsLeaf.Hash())
}

// A wrong expected root must be rejected by assembly's root check.
func TestLiteStateRootMismatch(t *testing.T) {
	state := minimalBeaconState(t)
	stateSSZ, err := state.MarshalSSZ()
	require.NoError(t, err)
	lite, err := ParseLiteBeaconState(stateSSZ)
	require.NoError(t, err)

	var wrong [32]byte
	wrong[0] = 0xff
	_, err = lite.AnchorStateTree(wrong)
	require.ErrorContains(t, err, "root mismatch")
}

// testExecHeader builds a fully-populated execution payload header.
func testExecHeader() *deneb.ExecutionPayloadHeader {
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

// minimalBeaconState builds the smallest fulu.BeaconState that attestantio
// will serialize and hash: fixed-size vectors at spec length, lists empty,
// containers populated.
func minimalBeaconState(t *testing.T) *fulu.BeaconState {
	t.Helper()

	syncCommittee := &altair.SyncCommittee{}
	for i := range syncCommittee.Pubkeys {
		syncCommittee.Pubkeys[i] = phase0.BLSPubKey{byte(i)}
	}

	state := &fulu.BeaconState{
		GenesisTime:           1,
		GenesisValidatorsRoot: phase0.Root{0x01},
		Slot:                  100,
		Fork: &phase0.Fork{
			PreviousVersion: phase0.Version{0x01},
			CurrentVersion:  phase0.Version{0x02},
			Epoch:           1,
		},
		LatestBlockHeader: &phase0.BeaconBlockHeader{
			Slot:       99,
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
		LatestExecutionPayloadHeader:  testExecHeader(),
		NextWithdrawalIndex:           1,
		NextWithdrawalValidatorIndex:  2,
		HistoricalSummaries:           nil,
		DepositRequestsStartIndex:     3,
		DepositBalanceToConsume:       4,
		ExitBalanceToConsume:          5,
		EarliestExitEpoch:             6,
		ConsolidationBalanceToConsume: 7,
		EarliestConsolidationEpoch:    8,
		PendingDeposits:               nil,
		PendingPartialWithdrawals:     nil,
		PendingConsolidations:         nil,
		ProposerLookahead:             make([]phase0.ValidatorIndex, 64),
	}
	state.StateRoots[5] = phase0.Root{0x02}
	return state
}
