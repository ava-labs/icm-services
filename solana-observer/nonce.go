// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"sync"
)

// NonceStore persists the next OracleMessage nonce per (sourceType, sourceAddress)
// pair. OracleAdapter's replay protection is keyed on that tuple + nonce, so
// two observer runs must not hand out the same value even across restarts.
type NonceStore struct {
	mu    sync.Mutex
	path  string
	State map[string]uint64 `json:"state"`
}

// LoadNonces loads an existing store or creates an empty one if the file does
// not exist yet.
func LoadNonces(path string) (*NonceStore, error) {
	s := &NonceStore{path: path, State: map[string]uint64{}}
	b, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return s, nil
	}
	if err != nil {
		return nil, fmt.Errorf("read nonce file: %w", err)
	}
	if err := json.Unmarshal(b, s); err != nil {
		return nil, fmt.Errorf("parse nonce file: %w", err)
	}
	if s.State == nil {
		s.State = map[string]uint64{}
	}
	return s, nil
}

// Next returns the next unused nonce for the given source and advances the
// in-memory counter. Callers must invoke Save to persist across restarts.
func (s *NonceStore) Next(sourceType, sourceAddress string) uint64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	key := sourceType + "/" + sourceAddress
	n := s.State[key] + 1
	s.State[key] = n
	return n
}

// Save atomically writes the store to disk. Uses a temp-file + rename to avoid
// leaving a truncated file on crash.
func (s *NonceStore) Save() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	tmp := s.path + ".tmp"
	if err := os.WriteFile(tmp, b, 0o600); err != nil {
		return fmt.Errorf("write nonce tmp file: %w", err)
	}
	if err := os.Rename(tmp, s.path); err != nil {
		return fmt.Errorf("rename nonce file: %w", err)
	}
	return nil
}
