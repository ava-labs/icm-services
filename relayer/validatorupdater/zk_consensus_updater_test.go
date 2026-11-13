// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validatorupdater

import (
	"context"
	"errors"
	"math/big"
	"net/http"
	"testing"

	"github.com/ava-labs/avalanchego/utils/logging"
	zkadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKAdapter"
	mock_evm "github.com/ava-labs/icm-services/vms/evm/mocks"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/crypto"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

const testLastAppliedSlot = uint64(100)

var testZKAdapterAddress = common.HexToAddress("0x00000000000000000000000000000000000000bb")

// newTestConsensusUpdater returns a synced updater whose transactions are
// followed by [tracker]. It has no subgraph client or contract, so a test in
// which the updater reaches for one of them panics: the paths that must not
// apply a proof are held to that by construction.
func newTestConsensusUpdater(tracker *fakeTxTracker) *ConsensusUpdater {
	return &ConsensusUpdater{
		logger:          logging.NoLog{},
		txTracker:       tracker,
		maxGasPriceWei:  new(big.Int),
		lastAppliedSlot: testLastAppliedSlot,
		synced:          true,
	}
}

// While a transition's transaction is unconfirmed, the poll does nothing but
// follow it up: it applies no further proof and leaves its position alone.
// Replacing the transaction is subject to the updater's gas price gate.
func TestConsensusUpdaterFollowsUpUnconfirmedTx(t *testing.T) {
	tracker := &fakeTxTracker{unconfirmed: newUnconfirmedTx()}
	u := newTestConsensusUpdater(tracker)

	require.NoError(t, u.checkAndUpdate(context.Background()))
	require.NotNil(t, tracker.Unconfirmed())
	require.True(t, u.synced)
	// No max gas price is configured, so the gate lets a replacement through.
	require.True(t, tracker.gateAsked)
	require.True(t, tracker.gateAllowed)

	tracker.resolveErr = errors.New("rpc down")
	require.ErrorIs(t, u.checkAndUpdate(context.Background()), tracker.resolveErr)
	require.True(t, u.synced)
}

// Once the transaction is settled, the updater cannot know whether the
// transition was applied, so it re-reads the contract's position and asks the
// subgraph for the proofs after that, not after the position it had assumed.
func TestConsensusUpdaterResyncsAfterSettledTx(t *testing.T) {
	tracker := &fakeTxTracker{unconfirmed: newUnconfirmedTx(), settled: true}
	u := newTestConsensusUpdater(tracker)

	// The contract has confirmed no beacon block, whichever way its position is
	// read: it has emitted no ConfirmedBeaconBlock event and its storage holds slot 0.
	backend := mock_evm.NewMockClient(gomock.NewController(t))
	backend.EXPECT().FilterLogs(gomock.Any(), gomock.Any()).Return(nil, nil).AnyTimes()
	backend.EXPECT().
		CallContract(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(make([]byte, 32), nil).
		AnyTimes()
	contract, err := zkadapter.NewZKAdapter(testZKAdapterAddress, backend)
	require.NoError(t, err)
	u.contract = contract
	server, lastRequest := newSubgraphServer(t, http.StatusOK, `{"data": {"signalEthereumProofs": []}}`)
	u.subgraphClient = NewSubgraphClient(server.URL)

	require.NoError(t, u.checkAndUpdate(context.Background()))
	require.Nil(t, tracker.Unconfirmed())
	require.True(t, u.synced)
	require.Zero(t, u.lastAppliedSlot)
	require.Contains(t, *lastRequest, `finalizedSlot_gt: \"0\"`)
}

// A transition whose transaction is not mined in time has not failed. The
// updater keeps its position without re-syncing, and the proofs after it, which
// build on it, are not sent behind it.
func TestConsensusUpdaterStopsAtUnconfirmedTransition(t *testing.T) {
	tracker := &fakeTxTracker{notMined: true}
	u := newTestConsensusUpdater(tracker)
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	u.txOpts, err = bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	require.NoError(t, err)

	// What the binding asks of the node to price, size and number the transaction.
	backend := mock_evm.NewMockClient(gomock.NewController(t))
	backend.EXPECT().
		HeaderByNumber(gomock.Any(), gomock.Any()).
		Return(&types.Header{BaseFee: big.NewInt(1)}, nil).
		AnyTimes()
	backend.EXPECT().SuggestGasTipCap(gomock.Any()).Return(big.NewInt(1), nil).AnyTimes()
	backend.EXPECT().PendingCodeAt(gomock.Any(), gomock.Any()).Return([]byte{0x01}, nil).AnyTimes()
	backend.EXPECT().EstimateGas(gomock.Any(), gomock.Any()).Return(uint64(100_000), nil).AnyTimes()
	backend.EXPECT().PendingNonceAt(gomock.Any(), gomock.Any()).Return(uint64(7), nil).AnyTimes()
	// The subgraph has two proofs, but only the first transition is broadcast.
	backend.EXPECT().SendTransaction(gomock.Any(), gomock.Any()).Return(nil).Times(1)
	contract, err := zkadapter.NewZKAdapter(testZKAdapterAddress, backend)
	require.NoError(t, err)
	u.contract = contract
	server, _ := newSubgraphServer(t, http.StatusOK, `{"data": {"signalEthereumProofs": [
		{"finalizedSlot": "132", "journalData": "0xdead", "seal": "0xbeef"},
		{"finalizedSlot": "164", "journalData": "0xcafe", "seal": "0xf00d"}
	]}}`)
	u.subgraphClient = NewSubgraphClient(server.URL)

	require.NoError(t, u.checkAndUpdate(context.Background()))
	require.NotNil(t, tracker.Unconfirmed())
	require.True(t, u.synced)
	require.Equal(t, testLastAppliedSlot, u.lastAppliedSlot)
}
