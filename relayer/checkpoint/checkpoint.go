// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package checkpoint

import (
	"container/heap"
	"fmt"
	"strconv"
	"sync"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/icm-services/database"
	"github.com/ava-labs/icm-services/utils"
	"go.uber.org/zap"
)

//
// CheckpointManager commits keys to be written to the database in a thread safe manner.
//

type CheckpointManager struct {
	logger          logging.Logger
	database        database.RelayerDatabase
	writeSignal     chan struct{}
	relayerID       database.RelayerID
	committedHeight uint64
	lock            *sync.RWMutex
	pendingCommits  *utils.UInt64Heap
	// Update the dirty flag when committedHeight is updated
	dirty bool
}

func NewCheckpointManager(
	logger logging.Logger,
	db database.RelayerDatabase,
	writeSignal chan struct{},
	relayerID database.RelayerID,
	startingHeight uint64,
) (*CheckpointManager, error) {
	h := &utils.UInt64Heap{}
	heap.Init(h)
	logger.Info(
		"Creating checkpoint manager",
		zap.String("relayerID", relayerID.ID.String()),
		zap.Uint64("startingHeight", startingHeight),
	)

	storedHeight, err := database.GetLatestProcessedBlockHeight(db, relayerID)
	if err != nil && !database.IsKeyNotFoundError(err) {
		logger.Error(
			"Failed to get latest processed block height",
			zap.Error(err),
			zap.String("relayerID", relayerID.ID.String()),
		)
		return nil, fmt.Errorf("failed to get the latest processed block height: %w", err)
	}

	committedHeight := max(storedHeight, startingHeight)

	return &CheckpointManager{
		logger:          logger,
		database:        db,
		writeSignal:     writeSignal,
		relayerID:       relayerID,
		committedHeight: committedHeight,
		lock:            &sync.RWMutex{},
		pendingCommits:  h,
		dirty:           true,
	}, nil
}

func (cm *CheckpointManager) Run() {
	go cm.listenForWriteSignal()
}

func (cm *CheckpointManager) writeToDatabase() {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	// Defensively ensure we're not writing the default value
	// If committedHeight is not changed, we can skip the write
	if cm.committedHeight == 0 || !cm.dirty {
		return
	}

	cm.logger.Verbo(
		"Writing height",
		zap.Uint64("height", cm.committedHeight),
		zap.String("relayerID", cm.relayerID.ID.String()),
	)
	err := cm.database.Put(
		cm.relayerID.ID,
		database.LatestProcessedBlockKey,
		[]byte(strconv.FormatUint(cm.committedHeight, 10)),
	)
	if err != nil {
		cm.logger.Error(
			"Failed to write latest processed block height",
			zap.Error(err),
			zap.String("relayerID", cm.relayerID.ID.String()),
		)
		return
	}

	// Reset the dirty flag after successfully write to db
	cm.dirty = false
}

func (cm *CheckpointManager) listenForWriteSignal() {
	for range cm.writeSignal {
		cm.writeToDatabase()
	}
}

// StageCommittedHeight queues a height to be written to the database.
// Heights are committed in sequence, so if height is not exactly one
// greater than the current committedHeight, it is instead cached in memory
// to potentially be committed later.
// TODO: We should only stage heights once all app relayers for a given source chain have staged
func (cm *CheckpointManager) StageCommittedHeight(height uint64) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	if height <= cm.committedHeight {
		cm.logger.Debug(
			"Attempting to commit height less than or equal to the committed height. Skipping.",
			zap.Uint64("height", height),
			zap.Uint64("committedHeight", cm.committedHeight),
			zap.String("relayerID", cm.relayerID.ID.String()),
		)
		return
	}

	// First push the height onto the pending commits min heap
	// This will ensure that the heights are committed in order
	heap.Push(cm.pendingCommits, height)
	cm.logger.Verbo(
		"Pending committed heights",
		zap.Any("maxPendingHeight", height),
		zap.Uint64("maxCommittedHeight", cm.committedHeight),
		zap.String("relayerID", cm.relayerID.ID.String()),
	)

	for cm.pendingCommits.Peek() == cm.committedHeight+1 {
		h := heap.Pop(cm.pendingCommits).(uint64)
		cm.logger.Verbo(
			"Committing height",
			zap.Uint64("height", height),
			zap.String("relayerID", cm.relayerID.ID.String()),
		)
		cm.committedHeight = h
		cm.dirty = true
		if cm.pendingCommits.Len() == 0 {
			break
		}
	}
}
