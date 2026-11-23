// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package validatorupdater

import (
	"math/big"
	"testing"

	"github.com/ava-labs/avalanchego/utils/logging"
	zkadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKAdapter"
	"github.com/ava-labs/libevm/common"
	"github.com/stretchr/testify/require"
)

// slotsPerEpoch matches ZKStateManager's SLOT_PER_EPOCH.
const slotsPerEpoch = 32

// testState returns a distinct consensus state whose head is the first slot of epoch.
func testState(epoch uint64) zkadapter.ConsensusState {
	return zkadapter.ConsensusState{
		Head:              epoch * slotsPerEpoch,
		Header:            [32]byte{byte(epoch)},
		SyncCommitteeHash: [32]byte{0xc0, byte(epoch)},
	}
}

// testOutputs returns public values for a transition from pre to post.
func testOutputs(pre, post zkadapter.ConsensusState) proofOutputs {
	return proofOutputs{
		PrevHeader:            pre.Header,
		PrevHead:              new(big.Int).SetUint64(pre.Head),
		PrevSyncCommitteeHash: pre.SyncCommitteeHash,
		NewHead:               new(big.Int).SetUint64(post.Head),
		NewHeader:             post.Header,
		ExecutionStateRoot:    [32]byte{0xee},
		ExecutionBlockNumber:  big.NewInt(1),
		SyncCommitteeHash:     post.SyncCommitteeHash,
		NextSyncCommitteeHash: [32]byte{0xc1},
		StorageSlots:          []storageSlot{{Key: [32]byte{1}, Value: [32]byte{2}, ContractAddress: common.Address{3}}},
	}
}

// testProof returns a proof whose public values transition pre to post.
func testProof(t *testing.T, pre, post zkadapter.ConsensusState) *ConsensusProof {
	t.Helper()
	publicValues, err := proofOutputsArguments.Pack(testOutputs(pre, post))
	require.NoError(t, err)
	return &ConsensusProof{NewHead: post.Head, PublicValues: publicValues}
}

func TestDecodePublicValues(t *testing.T) {
	pre, post := testState(100), testState(101)
	proof := testProof(t, pre, post)

	transition, err := decodePublicValues(proof.PublicValues)
	require.NoError(t, err)
	require.Equal(t, pre, transition.PreState)
	require.Equal(t, post, transition.PostState)

	_, err = decodePublicValues([]byte{1, 2, 3})
	require.Error(t, err)

	// The contract reverts on a head that does not fit in uint64, so such a proof is undecodable.
	outputs := testOutputs(pre, post)
	outputs.NewHead = new(big.Int).Lsh(big.NewInt(1), 64)
	publicValues, err := proofOutputsArguments.Pack(outputs)
	require.NoError(t, err)
	_, err = decodePublicValues(publicValues)
	require.ErrorContains(t, err, "uint64")
}

// The proof source may hold proofs of other states, which can never be applied. nextProof must
// pick the proof extending the contract's actual state, then chain from its post-state.
func TestNextProof(t *testing.T) {
	startingState := testState(100)
	firstPost := testState(101)
	secondPost := testState(102)

	unrelated := testProof(t, testState(5), testState(6))
	first := testProof(t, startingState, firstPost)
	second := testProof(t, firstPost, secondPost)
	undecodable := &ConsensusProof{NewHead: 1, PublicValues: []byte{1, 2, 3}}

	t.Run("chains from the current state", func(t *testing.T) {
		u := &ConsensusUpdater{logger: logging.NoLog{}, currentState: startingState}
		proofs := []*ConsensusProof{undecodable, unrelated, first, second}

		proof, transition := u.nextProof(proofs)
		require.Same(t, first, proof)
		require.Equal(t, firstPost, transition.PostState)

		// As after applying it: the next proof must extend the first proof's post-state.
		u.currentState, u.lastAppliedSlot = transition.PostState, proof.NewHead
		proof, _ = u.nextProof(proofs)
		require.Same(t, second, proof)
	})

	t.Run("matching state at a stale slot is skipped", func(t *testing.T) {
		u := &ConsensusUpdater{
			logger:          logging.NoLog{},
			currentState:    startingState,
			lastAppliedSlot: firstPost.Head,
		}
		proof, transition := u.nextProof([]*ConsensusProof{first})
		require.Nil(t, proof)
		require.Nil(t, transition)
	})

	t.Run("no proof extends the state", func(t *testing.T) {
		u := &ConsensusUpdater{logger: logging.NoLog{}, currentState: startingState}
		proof, _ := u.nextProof([]*ConsensusProof{unrelated, second})
		require.Nil(t, proof)
	})

	// A proof the contract rejects must not block a valid proof from the same state indexed
	// after it.
	t.Run("failed proof is skipped for the same state", func(t *testing.T) {
		alternative := testProof(t, startingState, testState(103))
		u := &ConsensusUpdater{logger: logging.NoLog{}, currentState: startingState}
		proofs := []*ConsensusProof{first, alternative}

		u.markProofFailed(first)
		proof, _ := u.nextProof(proofs)
		require.Same(t, alternative, proof)

		// Once the contract's state moves, earlier failures no longer apply.
		u.currentState = testState(99)
		u.markProofFailed(alternative)
		u.currentState = startingState
		proof, _ = u.nextProof(proofs)
		require.Same(t, first, proof)
	})
}
