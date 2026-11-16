// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package validatorupdater

import (
	"encoding/binary"
	"encoding/json"
	"os"
	"strconv"
	"testing"

	"github.com/ava-labs/avalanchego/utils/logging"
	zkadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKAdapter"
	"github.com/ava-labs/libevm/common"
	"github.com/stretchr/testify/require"
)

// stateFromABI decodes an ABI-encoded Consensus.State: four 32-byte words holding the current
// justified checkpoint's epoch and root, then the finalized checkpoint's epoch and root.
func stateFromABI(t *testing.T, encoded []byte) zkadapter.ConsensusState {
	t.Helper()
	require.Len(t, encoded, 4*32)
	word := func(i int) []byte { return encoded[i*32 : (i+1)*32] }
	checkpoint := func(i int) zkadapter.ConsensusCheckpoint {
		return zkadapter.ConsensusCheckpoint{
			Epoch: binary.BigEndian.Uint64(word(i)[24:]),
			Root:  [32]byte(word(i + 1)),
		}
	}
	return zkadapter.ConsensusState{
		CurrentJustifiedCheckpoint: checkpoint(0),
		FinalizedCheckpoint:        checkpoint(2),
	}
}

// Decodes a real Boundless proof's journal and checks it against the fixture's separately
// listed pre-state, post-state, and finalized slot.
func TestDecodeJournal(t *testing.T) {
	raw, err := os.ReadFile("../../tests/testdata/boundless_fixture.json")
	require.NoError(t, err)
	var fixture struct {
		FinalizedSlot string `json:"finalizedSlot"`
		PreState      string `json:"preState"`
		PostState     string `json:"postState"`
		JournalData   string `json:"journalData"`
	}
	require.NoError(t, json.Unmarshal(raw, &fixture))
	finalizedSlot, err := strconv.ParseUint(fixture.FinalizedSlot, 10, 64)
	require.NoError(t, err)

	journal, err := decodeJournal(common.FromHex(fixture.JournalData))
	require.NoError(t, err)
	require.Equal(t, stateFromABI(t, common.FromHex(fixture.PreState)), journal.PreState)
	require.Equal(t, stateFromABI(t, common.FromHex(fixture.PostState)), journal.PostState)
	require.Equal(t, finalizedSlot, journal.FinalizedSlot)

	_, err = decodeJournal([]byte{1, 2, 3})
	require.Error(t, err)
}

// testState returns a distinct consensus state whose finalized checkpoint is at epoch.
func testState(epoch uint64) zkadapter.ConsensusState {
	return zkadapter.ConsensusState{
		CurrentJustifiedCheckpoint: zkadapter.ConsensusCheckpoint{Epoch: epoch + 1, Root: [32]byte{byte(epoch + 1)}},
		FinalizedCheckpoint:        zkadapter.ConsensusCheckpoint{Epoch: epoch, Root: [32]byte{byte(epoch)}},
	}
}

// testProof returns a proof whose journal transitions pre to post, finalizing slot.
func testProof(t *testing.T, pre, post zkadapter.ConsensusState, slot uint64) *ConsensusProof {
	t.Helper()
	journalData, err := journalArguments.Pack(consensusJournal{PreState: pre, PostState: post, FinalizedSlot: slot})
	require.NoError(t, err)
	return &ConsensusProof{FinalizedSlot: slot, JournalData: journalData}
}

// On a fresh deployment lastConfirmedSlot is zero, so slot order alone would pick the oldest
// indexed proofs, which extend some other state and can never be applied. nextProof must pick
// the proof extending the contract's actual state, then chain from its post-state.
func TestNextProof(t *testing.T) {
	startingState := testState(100)
	firstPost := testState(101)
	secondPost := testState(102)

	unrelated := testProof(t, testState(5), testState(6), 6*slotsPerEpoch)
	first := testProof(t, startingState, firstPost, 101*slotsPerEpoch)
	second := testProof(t, firstPost, secondPost, 102*slotsPerEpoch)
	undecodable := &ConsensusProof{FinalizedSlot: 1, JournalData: []byte{1, 2, 3}}

	t.Run("fresh deployment chains from the starting state", func(t *testing.T) {
		u := &ConsensusUpdater{logger: logging.NoLog{}, currentState: startingState}
		proofs := []*ConsensusProof{undecodable, unrelated, first, second}

		proof, journal := u.nextProof(proofs)
		require.Same(t, first, proof)
		require.Equal(t, firstPost, journal.PostState)

		// As after applying it: the next proof must extend the first proof's post-state.
		u.currentState, u.lastAppliedSlot = journal.PostState, proof.FinalizedSlot
		proof, _ = u.nextProof(proofs)
		require.Same(t, second, proof)
	})

	t.Run("matching state at a stale slot is skipped", func(t *testing.T) {
		u := &ConsensusUpdater{
			logger:          logging.NoLog{},
			currentState:    startingState,
			lastAppliedSlot: 101 * slotsPerEpoch,
		}
		proof, journal := u.nextProof([]*ConsensusProof{first})
		require.Nil(t, proof)
		require.Nil(t, journal)
	})

	t.Run("no proof extends the state", func(t *testing.T) {
		u := &ConsensusUpdater{logger: logging.NoLog{}, currentState: startingState}
		proof, _ := u.nextProof([]*ConsensusProof{unrelated, second})
		require.Nil(t, proof)
	})
}

func TestQueryLowerBound(t *testing.T) {
	testCases := []struct {
		name            string
		lastAppliedSlot uint64
		finalizedEpoch  uint64
		expected        uint64
	}{
		{
			// The epoch bound skips past every indexed proof older than the starting state.
			name:           "fresh deployment",
			finalizedEpoch: 100,
			expected:       99 * slotsPerEpoch,
		},
		{
			name:            "last applied slot is later",
			lastAppliedSlot: 100*slotsPerEpoch + 5,
			finalizedEpoch:  100,
			expected:        100*slotsPerEpoch + 5,
		},
		{
			name:           "epoch zero does not underflow",
			finalizedEpoch: 0,
			expected:       0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			u := &ConsensusUpdater{lastAppliedSlot: tc.lastAppliedSlot, currentState: testState(tc.finalizedEpoch)}
			require.Equal(t, tc.expected, u.queryLowerBound())
		})
	}
}
