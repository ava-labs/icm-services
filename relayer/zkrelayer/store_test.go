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

// Loading before any state has been stored should return an empty RelayerState, not an error.
func TestDatabaseStoreLoadMissingKey(t *testing.T) {
	store := newTestStore(t)
	state, err := store.Load()
	require.NoError(t, err)
	require.NotNil(t, state)
	require.Zero(t, state.ScanFromBlock)
}

// Saves and loads a RelayerState, verifying that the loaded state matches the saved state.
func TestDatabaseStoreSaveLoadRoundTrip(t *testing.T) {
	store := newTestStore(t)
	targetState := &RelayerState{ScanFromBlock: 12_245_678}

	require.NoError(t, store.Save(targetState))
	loadedState, err := store.Load()

	require.NoError(t, err)
	require.Equal(t, targetState.ScanFromBlock, loadedState.ScanFromBlock)
}

// Helper function to run a load test case with given stored contents and expected error behavior.
func runLoadTestCase(t *testing.T, contents string, targetErr bool, targetBlock uint64) {
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
	require.Equal(t, targetBlock, state.ScanFromBlock)
}

// Loading a malformed state should return an error, not silently return a new state.
func TestDatabaseStoreLoadMalformedState(t *testing.T) {
	testCases := []struct {
		name        string
		contents    string
		targetErr   bool
		targetBlock uint64
	}{
		{name: "invalid json", contents: `{not valid json`, targetErr: true},
		{name: "empty object", contents: `{}`, targetErr: false, targetBlock: 0},
		{name: "valid cursor", contents: `{"scanFromBlock": 42}`, targetErr: false, targetBlock: 42},
		{name: "unknown fields ignored", contents: `{"scanFromBlock": 7, "someFutureField": true}`,
			targetErr: false, targetBlock: 7},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			runLoadTestCase(t, tc.contents, tc.targetErr, tc.targetBlock)
		})
	}
}

// Loading a valid state twice should yield the same result both times
func TestDatabaseStoreLoadIsIdempotent(t *testing.T) {
	store := newTestStore(t)
	require.NoError(t, store.Save(&RelayerState{ScanFromBlock: 12345678}))

	first, err := store.Load()
	require.NoError(t, err)
	second, err := store.Load()
	require.NoError(t, err)
	require.Equal(t, first, second)
}
