// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proofs

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/ava-labs/libevm/common"
	ssz "github.com/ferranbt/fastssz"
	"github.com/stretchr/testify/require"
)

// executionFixture mirrors the executionProof field of the fixture object
// from the fixture generation script tool.
// See: scripts/tools/fixture-gen/generate_fixture.mts
type executionFixture struct {
	AnchorBeaconBlockRoot string `json:"anchorBeaconBlockRoot"`
	ExecutionProof        struct {
		AnchorSlot                 uint64   `json:"anchorSlot"`
		TargetSlot                 uint64   `json:"targetSlot"`
		AnchorBeaconStateRoot      string   `json:"anchorBeaconStateRoot"`
		AnchorBeaconStateProof     []string `json:"anchorBeaconStateProof"`
		TargetBeaconStateRoot      string   `json:"targetBeaconStateRoot"`
		TargetBeaconStateProof     []string `json:"targetBeaconStateProof"`
		TargetExecutionHeaderRoot  string   `json:"targetExecutionHeaderRoot"`
		TargetExecutionHeaderProof []string `json:"targetExecutionHeaderProof"`
		TargetReceiptsRoot         string   `json:"targetReceiptsRoot"`
		TargetReceiptsProof        []string `json:"targetReceiptsProof"`
	} `json:"executionProof"`
}

// loadFixture reads the mainnet-generated fixture from testdata.
func loadFixture(t *testing.T) *executionFixture {
	t.Helper()
	data, err := os.ReadFile("testdata/ethereum_fixture.json")
	require.NoError(t, err)
	var fixture executionFixture
	require.NoError(t, json.Unmarshal(data, &fixture))
	return &fixture
}

// fixtureToProof assembles an ssz.Proof from the fixture's hex encoding:
// the proven position (gIndex), the value claimed there (leaf), and the
// sibling hashes needed to reconstruct the root.
func fixtureToProof(t *testing.T, gIndex int, leaf string, siblings []string) *ssz.Proof {
	t.Helper()
	hashes := make([][]byte, len(siblings))
	for i, s := range siblings {
		hash := common.FromHex(s)
		require.Len(t, hash, 32, "sibling %d", i)
		hashes[i] = hash
	}
	return &ssz.Proof{
		Index:  gIndex,
		Leaf:   common.HexToHash(leaf).Bytes(),
		Hashes: hashes,
	}
}

// verifyProof asserts one link of the chain of trust, which is that the proof verifies
// against expectedRoot.
func verifyProof(t *testing.T, expectedRoot string, proof *ssz.Proof) {
	t.Helper()
	ok, err := ssz.VerifyProof(common.HexToHash(expectedRoot).Bytes(), proof)
	require.NoError(t, err)
	require.True(t, ok)
}

// newTree builds a synthetic SSZ tree over numLeaves amount of zero chunks with
// value at leafIndex.
func newTree(t *testing.T, numLeaves int, leafIndex int, value common.Hash) *ssz.Node {
	t.Helper()
	chunks := make([][]byte, numLeaves)
	for i := range chunks {
		chunks[i] = make([]byte, 32)
	}
	copy(chunks[leafIndex], value.Bytes())
	tree, err := ssz.TreeFromChunks(chunks)
	require.NoError(t, err)
	return tree
}

// Verify the fixture is valid, meaning the chain of trust verifies correctly.
// Checks that the GIndex constants in execution_proofs.go (11, 70, 13, 88, 35) are
// correct for the Ethereum mainnet beacon chain. This is anchor beacon block root
// -> anchor beacon state -> target state -> execution header -> receipts root.
func TestExecutionFixtureChainOfTrust(t *testing.T) {
	fixture := loadFixture(t)
	proof := fixture.ExecutionProof

	// Proof 1: anchor beacon block root -> anchor beacon state root.
	verifyProof(t,
		fixture.AnchorBeaconBlockRoot,
		// gIndexBlockStateRoot is from execution_proofs.go.
		fixtureToProof(t, gIndexBlockStateRoot, proof.AnchorBeaconStateRoot, proof.AnchorBeaconStateProof),
	)

	// Proof 2: anchor beacon state root -> target beacon state root.
	vectorIndex := int(proof.TargetSlot % StateRootsVectorSize)
	targetGIndex := (gIndexBaseStateRoots << stateRootsDepth) + vectorIndex
	verifyProof(t,
		proof.AnchorBeaconStateRoot,
		// targetGIndex is from execution_proofs.go.
		fixtureToProof(t, targetGIndex, proof.TargetBeaconStateRoot, proof.TargetBeaconStateProof),
	)

	// Proof 3: target beacon state root -> execution payload header root.
	verifyProof(t,
		proof.TargetBeaconStateRoot,
		// gIndexExecPayloadHeader is from execution_proofs.go.
		fixtureToProof(t, gIndexExecPayloadHeader, proof.TargetExecutionHeaderRoot, proof.TargetExecutionHeaderProof),
	)

	// Proof 4: execution payload header root -> receipts root.
	verifyProof(t,
		proof.TargetExecutionHeaderRoot,
		// gIndexReceiptsRoot is from execution_proofs.go.
		fixtureToProof(t, gIndexReceiptsRoot, proof.TargetReceiptsRoot, proof.TargetReceiptsProof),
	)
}

// BuildExecutionProof must build and verify the 4 proof chain of trust. The trees
// are synthetic but shaped so the real GIndices resolve as leaves, with each
// tree's root planted as the correct leaf of its parent. This is the same
// as the real beacon structures.
//
// This test is built bottom-up:
// execution header -> target beacon state -> anchor beacon state -> anchor beacon block.
func TestBuildExecutionProofSynthetic(t *testing.T) {
	const (
		anchorSlot = uint64(200)
		targetSlot = uint64(100)

		// 0-indexed depths for GIndices taken from execution_proofs.go.
		blockDepth      = 3                            // gindex 11 (state_root)
		stateDepth      = 6                            // gindex 88 (latest_execution_payload_header)
		execHeaderDepth = 5                            // gindex 35 (receipts_root)
		anchorDepth     = stateDepth + stateRootsDepth // history vector under the state
	)
	receiptsRoot := common.HexToHash("0x1111111111111111111111111111111111111111111111111111111111111111")

	// Leaf index of gindex g in a tree whose leaves sit at depth d.
	leafIdx := func(g, d int) int { return g - (1 << d) }

	// The idea is we will build 4 synthetic trees, each with a single non-zero leaf at
	// the GIndex we want to prove against.

	// Build tree for execution payload header, holding the receipts root.
	execHeader := newTree(t, 1<<execHeaderDepth, leafIdx(gIndexReceiptsRoot, execHeaderDepth), receiptsRoot)
	execHeaderRoot := common.BytesToHash(execHeader.Hash())

	// Build tree for target beacon state, committing to the header.
	targetState := newTree(t, 1<<stateDepth, leafIdx(gIndexExecPayloadHeader, stateDepth), execHeaderRoot)
	targetStateRoot := common.BytesToHash(targetState.Hash())

	// Build tree for anchor beacon state, holding the target state root in its history.
	historyGIndex := (gIndexBaseStateRoots << stateRootsDepth) + int(targetSlot%StateRootsVectorSize)
	anchorState := newTree(t, 1<<anchorDepth, leafIdx(historyGIndex, anchorDepth), targetStateRoot)
	anchorStateRoot := common.BytesToHash(anchorState.Hash())

	// Build tree for anchor block, committing to the anchor state.
	anchorBlock := newTree(t, 1<<blockDepth, leafIdx(gIndexBlockStateRoot, blockDepth), anchorStateRoot)
	anchorBlockRoot := common.BytesToHash(anchorBlock.Hash())

	// Build the execution proof, which internally verifies each step of the chain of trust.
	execProof, err := BuildExecutionProof(
		anchorBlock, anchorState, targetState, execHeader, anchorBlockRoot, anchorSlot, targetSlot)
	require.NoError(t, err)

	// The returned roots must match the values the trees were built with,
	// confirming each proof resolved to the right position.
	require.Equal(t, anchorSlot, execProof.AnchorSlot)
	require.Equal(t, targetSlot, execProof.TargetSlot)
	require.Equal(t, anchorStateRoot, common.Hash(execProof.AnchorBeaconStateRoot))
	require.Equal(t, targetStateRoot, common.Hash(execProof.TargetBeaconStateRoot))
	require.Equal(t, receiptsRoot, common.Hash(execProof.TargetReceiptsRoot))

	// Sanity check: given a wrong anchor beacon block root, the proof must fail to verify.
	wrongRoot := common.HexToHash("0xbad")
	_, err = BuildExecutionProof(
		anchorBlock, anchorState, targetState, execHeader, wrongRoot, anchorSlot, targetSlot)
	require.ErrorContains(t, err, "does not verify against expected root")
}

// Corrupted proof must fail to verify.
func TestExecutionFixtureTamperedProofFails(t *testing.T) {
	fixture := loadFixture(t)
	proof := fixture.ExecutionProof

	corrupted := fixtureToProof(
		t,
		gIndexBlockStateRoot,
		proof.AnchorBeaconStateRoot,
		proof.AnchorBeaconStateProof,
	)
	corrupted.Hashes[0][0] ^= 0xff

	ok, err := ssz.VerifyProof(
		common.HexToHash(fixture.AnchorBeaconBlockRoot).Bytes(), corrupted)
	require.NoError(t, err)
	require.False(t, ok)
}

func TestProveAgainst(t *testing.T) {
	// builds small test tree beacon state tree with 8 leaves,
	// proving against gIndexBlockStateRoot=11
	chunks := make([][]byte, 8)
	for i := range chunks {
		chunk := make([]byte, 32)
		chunk[0] = byte(i + 1)
		chunks[i] = chunk
	}
	gIndexBlockStateRoot := 11
	leafIndex := 3 // Gindex 11 in an 8-leaf tree is leaf index 3.
	tree, err := ssz.TreeFromChunks(chunks)
	require.NoError(t, err)
	root := common.BytesToHash(tree.Hash())

	// happy path
	proof, err := proveAgainst(tree, gIndexBlockStateRoot, root)
	require.NoError(t, err)
	require.Equal(t, chunks[leafIndex], proof.Leaf)

	// A mismatched expected root must be rejected.
	wrongRoot := common.HexToHash("0xcafebabe")
	_, err = proveAgainst(tree, gIndexBlockStateRoot, wrongRoot)
	require.ErrorContains(t, err, "does not verify against expected root")
}

func TestToHash32Slice(t *testing.T) {
	valid := [][]byte{make([]byte, 32), make([]byte, 32)}
	valid[0][31] = 0xaa
	out, err := toHash32Slice(valid)
	require.NoError(t, err)
	require.Len(t, out, 2)
	require.Equal(t, byte(0xaa), out[0][31])

	_, err = toHash32Slice([][]byte{make([]byte, 31)})
	require.ErrorContains(t, err, "length 31, want 32")
}
