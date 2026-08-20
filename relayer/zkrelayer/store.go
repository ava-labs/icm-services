package zkrelayer

import (
	"encoding/json"
	"os"

	"github.com/ava-labs/libevm/common"
)

// Store is an interface for persisting the relayer state.
type Store interface {
	Load() (*RelayerState, error)
	Save(*RelayerState) error
}

// RelayerState is the persisted state of the relayer used on startup.
type RelayerState struct {
	// LastScannedBlock is the last source-chain block number that was scanned for message events.
	LastScannedBlock uint64 `json:"lastScannedBlock"`
	// LastAppliedSlot is the last beacon chain slot for which a Boundless proof was applied to
	// the destination chain.
	LastAppliedSlot uint64 `json:"lastAppliedSlot"`
	// ConfirmedAnchors is the list of beacon chain slots for which a Boundless proof has been
	// applied to the destination chain.
	ConfirmedAnchors []uint64 `json:"confirmedAnchors"`
	// PendingMessages is a map of source-chain transaction hashes to their corresponding
	// PendingMessage records.
	PendingMessages map[common.Hash]*PendingMessage `json:"pending"`
}

type FileStore struct{ Path string }

// Load reads the relayer state from the file at f.Path. If the file does not exist, it returns
// an empty RelayerState.
func (f *FileStore) Load() (*RelayerState, error) {
	// Initialize empty relayer state
	state := &RelayerState{PendingMessages: map[common.Hash]*PendingMessage{}}
	data, err := os.ReadFile(f.Path)
	if os.IsNotExist(err) {
		return state, nil
	}
	if err != nil {
		return nil, err
	}
	// Unmarshal the JSON data into the RelayerState struct
	if err := json.Unmarshal(data, state); err != nil {
		return nil, err
	}
	if state.PendingMessages == nil {
		state.PendingMessages = map[common.Hash]*PendingMessage{}
	}
	return state, nil
}

// Save `writes the relayer state to the file at f.Path.
func (f *FileStore) Save(state *RelayerState) error {
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}

	// Write to a temporary file first to ensure an atomic update.
	// This prevents file corruption if the application crashes mid-write
	// because it will not reach creating the final file.
	tmp := f.Path + ".tmp"
	if err := os.WriteFile(tmp, data, 0o600); err != nil {
		return err
	}

	// Atomically replace the old file with the new one.
	// The temp file is renamed to the final path, which is an atomic operation on
	// most filesystems.
	return os.Rename(tmp, f.Path)
}
