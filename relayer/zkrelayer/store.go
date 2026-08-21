package zkrelayer

import (
	"encoding/json"
	"errors"

	"github.com/ava-labs/icm-services/database"
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

// DatabaseStore persists the relayer state as a single JSON value in the shared
// RelayerDatabase, keyed by this relayer's ID.
type DatabaseStore struct {
	db        database.RelayerDatabase
	relayerID common.Hash
}

func NewDatabaseStore(db database.RelayerDatabase, relayerID common.Hash) *DatabaseStore {
	return &DatabaseStore{db: db, relayerID: relayerID}
}

// Load reads the relayer state from the database. If no state has been stored yet, it returns
// an empty RelayerState.
func (s *DatabaseStore) Load() (*RelayerState, error) {
	// Initialize empty relayer state
	state := &RelayerState{PendingMessages: map[common.Hash]*PendingMessage{}}
	data, err := s.db.Get(s.relayerID, database.ZKRelayerStateKey)
	if errors.Is(err, database.ErrKeyNotFound) || errors.Is(err, database.ErrRelayerIDNotFound) {
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

// Save writes the relayer state to the database.
func (s *DatabaseStore) Save(state *RelayerState) error {
	data, err := json.Marshal(state)
	if err != nil {
		return err
	}
	return s.db.Put(s.relayerID, database.ZKRelayerStateKey, data)
}
