// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//go:generate go run go.uber.org/mock/mockgen -source=$GOFILE -destination=./mocks/mock_database.go -package=mocks

package database

import (
	"fmt"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/libevm/common"
	"github.com/pkg/errors"
)

var (
	ErrKeyNotFound              = errors.New("key not found")
	ErrRelayerIDNotFound        = errors.New("no database entry for relayer id")
	ErrDatabaseMisconfiguration = errors.New("database misconfiguration")
)

const (
	LatestProcessedBlockKey DataKey = iota
)

type DataKey int

func (k DataKey) String() string {
	switch k {
	case LatestProcessedBlockKey:
		return "latestProcessedBlock"
	}
	return "unknown"
}

// RelayerDatabase is a key-value store for relayer state, with each relayerID maintaining its own state.
// Implementations should be thread-safe.
type RelayerDatabase interface {
	Get(relayerID common.Hash, key DataKey) ([]byte, error)
	Put(relayerID common.Hash, key DataKey, value []byte) error
	Close() error
}

func NewDatabase(logger logging.Logger, cfg *config.Config) (RelayerDatabase, error) {
	if cfg.RedisURL != "" {
		db, err := NewRedisDatabase(logger, cfg.RedisURL, GetConfigRelayerIDs(cfg))
		if err != nil {
			return nil, fmt.Errorf("failed to create redis database: %w", err)
		}
		return db, nil
	} else {
		db, err := NewJSONFileStorage(logger, cfg.StorageLocation, GetConfigRelayerIDs(cfg))
		if err != nil {
			return nil, fmt.Errorf("failed to create json database: %w", err)
		}
		return db, nil
	}
}
