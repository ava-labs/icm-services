package zkrelayer

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ava-labs/libevm/common"
	"github.com/stretchr/testify/require"
)

// newTestStore returns a FileStore rooted in a per-test temp directory.
func newTestStore(t *testing.T) *FileStore {
	t.Helper()
	return &FileStore{Path: filepath.Join(t.TempDir(), "state.json")}
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

// Loading a missing file should return an empty RelayerState, not an error.
func TestFileStoreLoadMissingFile(t *testing.T) {
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
func TestFileStoreSaveLoadRoundTrip(t *testing.T) {
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

// Saving a RelayerState should remove any temporary file used during the save process.
func TestFileStoreSaveRemovesTempFile(t *testing.T) {
	store := newTestStore(t)

	require.NoError(t, store.Save(&RelayerState{
		PendingMessages: map[common.Hash]*PendingMessage{},
	}))

	// The .tmp file must be renamed away, leaving only the final state file.
	_, err := os.Stat(store.Path + ".tmp")
	require.True(t, os.IsNotExist(err))
	_, err = os.Stat(store.Path)
	require.NoError(t, err)
}

// Helper function to run a load test case with given file contents and expected error behavior.
func runLoadTestCase(t *testing.T, store *FileStore, contents string, targetErr bool) {
	t.Helper()
	// Setup initial file state
	require.NoError(t, os.WriteFile(store.Path, []byte(contents), 0o600))
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

// Loading a malformed state file should return an error, not silently return a new state.
func TestFileStoreLoadMalformedState(t *testing.T) {
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
			store := newTestStore(t)
			runLoadTestCase(t, store, tc.contents, tc.targetErr)
		})
	}
}

// Loading a valid state file twice should yield the same result both times
func TestFileStoreLoadIsIdempotent(t *testing.T) {
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
