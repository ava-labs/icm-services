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

// RelayerState is the persisted state of the relayer used on startup. It is kept minimal as
// a single scan cursor. Everything else the relayer needs is reconstructed on startup, i.e.,
// pending messages by rescanning the source chain from the cursor (checking delivery
// status against the destination), and the consensus position and confirmed anchors by
// reading the destination contract.
type RelayerState struct {
	// ScanFromBlock is the source-chain block number to resume scanning from. It is the
	// block of the earliest pending (not yet confirmed delivered) message, or the last
	// scanned block if no messages are pending, so a restart rescans exactly the window
	// that may still contain undelivered messages.
	ScanFromBlock uint64 `json:"scanFromBlock"`
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
	state := &RelayerState{}
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
