// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package validatorupdater

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/ava-labs/avalanchego/utils/logging"
	zkadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKAdapter"
	mock_evm "github.com/ava-labs/icm-services/vms/evm/mocks"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/crypto"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
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

	// A proof the contract rejects must not block a valid proof from the same state indexed
	// after it.
	t.Run("failed proof is skipped for the same state", func(t *testing.T) {
		alternative := testProof(t, startingState, testState(103), 101*slotsPerEpoch)
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

const testLastAppliedSlot = uint64(100)

var testZKAdapterAddress = common.HexToAddress("0x00000000000000000000000000000000000000bb")

// newTestConsensusUpdater returns an updater that assumes the contract is at
// testLastAppliedSlot and whose transactions are followed by [tracker]. It has
// no subgraph client or contract, so a test in which the updater reaches for one
// of them panics: the paths that must not apply a proof are held to that by
// construction.
func newTestConsensusUpdater(tracker *fakeTxTracker) *ConsensusUpdater {
	return &ConsensusUpdater{
		logger:          logging.NoLog{},
		txTracker:       tracker,
		maxGasPriceWei:  new(big.Int),
		lastAppliedSlot: testLastAppliedSlot,
	}
}

// newZKAdapterBackend returns a chain client on which the ZKAdapter's consensus
// position reads as [slot] and [state], which is what the updater syncs to
// before it applies a proof.
func newZKAdapterBackend(t *testing.T, slot uint64, state zkadapter.ConsensusState) *mock_evm.MockClient {
	t.Helper()
	adapterABI, err := zkadapter.ZKAdapterMetaData.GetAbi()
	require.NoError(t, err)

	backend := mock_evm.NewMockClient(gomock.NewController(t))
	backend.EXPECT().
		CallContract(gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, call ethereum.CallMsg, _ *big.Int) ([]byte, error) {
			method, err := adapterABI.MethodById(call.Data)
			if err != nil {
				return nil, err
			}
			switch method.Name {
			case "lastConfirmedSlot":
				return method.Outputs.Pack(slot)
			case "getCurrentState":
				return method.Outputs.Pack(state)
			default:
				return nil, fmt.Errorf("unexpected call to %s", method.Name)
			}
		}).
		AnyTimes()
	return backend
}

// While a transition's transaction is unconfirmed, the poll does nothing but
// follow it up: it applies no further proof and leaves its position alone.
// Replacing the transaction is subject to the updater's gas price gate.
func TestConsensusUpdaterFollowsUpUnconfirmedTx(t *testing.T) {
	tracker := &fakeTxTracker{unconfirmed: newUnconfirmedTx()}
	u := newTestConsensusUpdater(tracker)

	require.NoError(t, u.checkAndUpdate(context.Background()))
	require.NotNil(t, tracker.Unconfirmed())
	require.Equal(t, testLastAppliedSlot, u.lastAppliedSlot)
	// No max gas price is configured, so the gate lets a replacement through.
	require.True(t, tracker.gateAsked)
	require.True(t, tracker.gateAllowed)

	tracker.resolveErr = errors.New("rpc down")
	require.ErrorIs(t, u.checkAndUpdate(context.Background()), tracker.resolveErr)
	require.Equal(t, testLastAppliedSlot, u.lastAppliedSlot)
}

// Once the transaction is settled, the updater cannot know whether the
// transition was applied, so it re-reads the contract's position and asks the
// subgraph for the proofs after that, not after the position it had assumed.
func TestConsensusUpdaterResyncsAfterSettledTx(t *testing.T) {
	tracker := &fakeTxTracker{unconfirmed: newUnconfirmedTx(), settled: true}
	u := newTestConsensusUpdater(tracker)

	// The contract has confirmed no beacon block: its storage holds slot 0 and the
	// zero state.
	backend := newZKAdapterBackend(t, 0, zkadapter.ConsensusState{})
	contract, err := zkadapter.NewZKAdapter(testZKAdapterAddress, backend)
	require.NoError(t, err)
	u.contract = contract
	server, lastRequest := newSubgraphServer(t, http.StatusOK, `{"data": {"signalEthereumProofs": []}}`)
	u.subgraphClient = NewSubgraphClient(server.URL)

	require.NoError(t, u.checkAndUpdate(context.Background()))
	require.Nil(t, tracker.Unconfirmed())
	require.Zero(t, u.lastAppliedSlot)
	require.Contains(t, *lastRequest, `finalizedSlot_gt: \"0\"`)
}

// A transition whose transaction is not mined in time has not failed. The
// updater keeps its position and does not mark the proof as failed, and the
// proofs after it, which build on it, are not sent behind it.
func TestConsensusUpdaterStopsAtUnconfirmedTransition(t *testing.T) {
	tracker := &fakeTxTracker{notMined: true}
	u := newTestConsensusUpdater(tracker)
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	u.txOpts, err = bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	require.NoError(t, err)

	// The contract is where the updater assumes it is, and the subgraph has two
	// proofs that extend its state, the second building on the first.
	startState, firstPost := testState(3), testState(4)
	first := testProof(t, startState, firstPost, 132)
	second := testProof(t, firstPost, testState(5), 164)

	backend := newZKAdapterBackend(t, testLastAppliedSlot, startState)
	// What the binding asks of the node to price, size and number the transaction.
	backend.EXPECT().
		HeaderByNumber(gomock.Any(), gomock.Any()).
		Return(&types.Header{BaseFee: big.NewInt(1)}, nil).
		AnyTimes()
	backend.EXPECT().SuggestGasTipCap(gomock.Any()).Return(big.NewInt(1), nil).AnyTimes()
	backend.EXPECT().PendingCodeAt(gomock.Any(), gomock.Any()).Return([]byte{0x01}, nil).AnyTimes()
	backend.EXPECT().EstimateGas(gomock.Any(), gomock.Any()).Return(uint64(100_000), nil).AnyTimes()
	backend.EXPECT().PendingNonceAt(gomock.Any(), gomock.Any()).Return(uint64(7), nil).AnyTimes()
	// Only the first transition is broadcast.
	backend.EXPECT().SendTransaction(gomock.Any(), gomock.Any()).Return(nil).Times(1)
	contract, err := zkadapter.NewZKAdapter(testZKAdapterAddress, backend)
	require.NoError(t, err)
	u.contract = contract
	server, _ := newSubgraphServer(t, http.StatusOK, fmt.Sprintf(`{"data": {"signalEthereumProofs": [
		{"finalizedSlot": "%d", "journalData": %q, "seal": "0xbeef"},
		{"finalizedSlot": "%d", "journalData": %q, "seal": "0xf00d"}
	]}}`,
		first.FinalizedSlot, hexutil.Encode(first.JournalData),
		second.FinalizedSlot, hexutil.Encode(second.JournalData),
	))
	u.subgraphClient = NewSubgraphClient(server.URL)

	require.NoError(t, u.checkAndUpdate(context.Background()))
	require.NotNil(t, tracker.Unconfirmed())
	require.Equal(t, testLastAppliedSlot, u.lastAppliedSlot)
	require.Equal(t, startState, u.currentState)
	require.Empty(t, u.failedProofs)
}
