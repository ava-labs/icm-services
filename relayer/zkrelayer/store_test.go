package zkrelayer

import (
	"testing"

	"github.com/ava-labs/icm-services/database"
	"github.com/ava-labs/libevm/common"
	"github.com/stretchr/testify/require"
)

// memDatabase is a minimal in-memory RelayerDatabase for tests.
type memDatabase struct {
	values map[common.Hash]map[database.DataKey][]byte
}

func newMemDatabase() *memDatabase {
	return &memDatabase{values: map[common.Hash]map[database.DataKey][]byte{}}
}

func (m *memDatabase) Get(relayerID common.Hash, key database.DataKey) ([]byte, error) {
	data, ok := m.values[relayerID][key]
	if !ok {
		return nil, database.ErrKeyNotFound
	}
	return data, nil
}

func (m *memDatabase) Put(relayerID common.Hash, key database.DataKey, value []byte) error {
	if m.values[relayerID] == nil {
		m.values[relayerID] = map[database.DataKey][]byte{}
	}
	m.values[relayerID][key] = value
	return nil
}

func (m *memDatabase) Close() error { return nil }

// newTestStore returns a DatabaseStore backed by an in-memory database.
func newTestStore(t *testing.T) *DatabaseStore {
	t.Helper()
	return NewDatabaseStore(newMemDatabase(), common.HexToHash("0x01"))
}

// sampleState returns a populated RelayerState for round-trip tests.
func sampleState(txHash common.Hash) *RelayerState {
	return &RelayerState{
		LastScannedBlock: 12_245_678,
		LastAppliedSlot:  123_345,
		ConfirmedAnchors: []uint64{200_000, 115_000},
		PendingMessages: map[common.Hash]*PendingMessage{
			txHash: {
				TxHash:      txHash,
				BlockNumber: 19_234_500,
				Slot:        918_100,
				LogIndex:    3,
				State:       MessageStateDetected,
				Attempts:    1,
				LastError:   "transient rpc error",
			},
		},
	}
}

// Loading before any state has been stored should return an empty RelayerState, not an error.
func TestDatabaseStoreLoadMissingKey(t *testing.T) {
	store := newTestStore(t)
	state, err := store.Load()
	require.NoError(t, err)
	require.NotNil(t, state)
	require.NotNil(t, state.PendingMessages)
	require.Empty(t, state.PendingMessages)
	require.Zero(t, state.LastScannedBlock)
	require.Zero(t, state.LastAppliedSlot)
}

// Saves and loads a RelayerState, verifying that the loaded state matches the saved state.
func TestDatabaseStoreSaveLoadRoundTrip(t *testing.T) {
	store := newTestStore(t)
	txHash := common.HexToHash("0xabc123")
	targetState := sampleState(txHash)

	require.NoError(t, store.Save(targetState))
	loadedState, err := store.Load()

	require.NoError(t, err)
	require.Equal(t, targetState.LastScannedBlock, loadedState.LastScannedBlock)
	require.Equal(t, targetState.LastAppliedSlot, loadedState.LastAppliedSlot)
	require.Equal(t, targetState.ConfirmedAnchors, loadedState.ConfirmedAnchors)
	require.Contains(t, loadedState.PendingMessages, txHash)
	require.Equal(t, targetState.PendingMessages[txHash], loadedState.PendingMessages[txHash])
}

// Helper function to run a load test case with given stored contents and expected error behavior.
func runLoadTestCase(t *testing.T, contents string, targetErr bool) {
	t.Helper()
	// Setup initial stored state
	db := newMemDatabase()
	relayerID := common.HexToHash("0x01")
	require.NoError(t, db.Put(relayerID, database.ZKRelayerStateKey, []byte(contents)))
	store := NewDatabaseStore(db, relayerID)

	state, err := store.Load()
	if targetErr {
		require.Error(t, err)
		return
	}
	require.NoError(t, err)

	// Validate the pending map is not nil and writable
	require.NotNil(t, state.PendingMessages)

	txHash := common.HexToHash("0xabc123")
	state.PendingMessages[txHash] = &PendingMessage{TxHash: txHash}
	require.Len(t, state.PendingMessages, 1)
}

// Loading a malformed state should return an error, not silently return a new state.
func TestDatabaseStoreLoadMalformedState(t *testing.T) {
	testCases := []struct {
		name      string
		contents  string
		targetErr bool
	}{
		{name: "invalid json", contents: `{not valid json`, targetErr: true},
		{name: "empty object", contents: `{}`, targetErr: false},
		{name: "null pending map", contents: `{"pending": null}`, targetErr: false},
		{name: "missing pending key", contents: `{"lastScannedBlock": 42}`, targetErr: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			runLoadTestCase(t, tc.contents, tc.targetErr)
		})
	}
}

// Loading a valid state twice should yield the same result both times
func TestDatabaseStoreLoadIsIdempotent(t *testing.T) {
	store := newTestStore(t)
	require.NoError(t, store.Save(&RelayerState{
		LastScannedBlock: 12345678,
		PendingMessages:  map[common.Hash]*PendingMessage{},
	}))

	first, err := store.Load()
	require.NoError(t, err)
	second, err := store.Load()
	require.NoError(t, err)
	require.Equal(t, first, second)
}
