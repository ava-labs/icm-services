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
	"go.uber.org/zap"
)

//
// CheckpointManager commits keys to be written to the database in a thread safe manner.
//

type CheckpointManager struct {
	logger          logging.Logger
	metrics         *CheckpointManagerMetrics
	database        database.RelayerDatabase
	writeSignal     chan struct{}
	relayerID       database.RelayerID
	committedHeight uint64
	lock            *sync.RWMutex
	pendingCommits  *blockRangeHeap
	// Update the dirty flag when committedHeight is updated
	dirty bool
}

func NewCheckpointManager(
	logger logging.Logger,
	metrics *CheckpointManagerMetrics,
	db database.RelayerDatabase,
	writeSignal chan struct{},
	relayerID database.RelayerID,
	startingHeight uint64,
) (*CheckpointManager, error) {
	logger = logger.With(zap.Stringer("relayerID", relayerID.ID))

	h := &blockRangeHeap{}
	heap.Init(h)
	logger.Info(
		"Creating checkpoint manager",
		zap.Uint64("startingHeight", startingHeight),
	)

	storedHeight, err := database.GetLatestProcessedBlockHeight(db, relayerID)
	if err != nil && !database.IsKeyNotFoundError(err) {
		logger.Error("Failed to get latest processed block height", zap.Error(err))
		return nil, fmt.Errorf("failed to get the latest processed block height: %w", err)
	}

	committedHeight := max(storedHeight, startingHeight)

	metrics.UpdateCommittedHeight(relayerID, committedHeight)
	metrics.UpdatePendingCommitsHeapLength(relayerID, 0)

	return &CheckpointManager{
		logger:          logger,
		metrics:         metrics,
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

	cm.logger.Verbo("Writing height",
		zap.Uint64("height", cm.committedHeight),
	)
	err := cm.database.Put(
		cm.relayerID.ID,
		database.LatestProcessedBlockKey,
		[]byte(strconv.FormatUint(cm.committedHeight, 10)),
	)
	if err != nil {
		cm.logger.Error("Failed to write latest processed block height", zap.Error(err))
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

// StageCommittedHeights records that every block in [fromHeight, toHeight] has been processed
// and queues the heights to be written to the database. Heights are committed in sequence, so if
// [fromHeight] is not directly after the current committedHeight, the range is instead cached in
// memory until the ranges covering the heights before it have been staged.
// Ranges may overlap heights that were already committed or staged, e.g. when a block is
// processed both by catch-up and from the subscription; only their uncommitted heights count.
// TODO: We should only stage heights once all app relayers for a given source chain have staged
func (cm *CheckpointManager) StageCommittedHeights(fromHeight, toHeight uint64) {
	cm.lock.Lock()
	defer cm.lock.Unlock()

	log := cm.logger.With(
		zap.Uint64("stagingFromHeight", fromHeight),
		zap.Uint64("stagingToHeight", toHeight),
	)

	if fromHeight > toHeight {
		log.Error("Attempting to commit an inverted height range. Skipping.")
		return
	}
	if toHeight <= cm.committedHeight {
		log.Debug(
			"Attempting to commit heights less than or equal to the committed height. Skipping.",
			zap.Uint64("committedHeight", cm.committedHeight),
		)
		return
	}

	// First push the range onto the pending commits min heap
	// This will ensure that the heights are committed in order
	heap.Push(cm.pendingCommits, blockRange{from: fromHeight, to: toHeight})
	cm.metrics.UpdatePendingCommitsHeapLength(cm.relayerID, cm.pendingCommits.Len())
	log.Verbo(
		"Pending committed heights",
		zap.Uint64("maxCommittedHeight", cm.committedHeight),
	)

	// Commit every pending range that starts at or below the next uncommitted height.
	for cm.pendingCommits.Len() > 0 && cm.pendingCommits.Peek().from <= cm.committedHeight+1 {
		r := heap.Pop(cm.pendingCommits).(blockRange)
		if r.to <= cm.committedHeight {
			// Every height in the range was already committed by an overlapping range.
			continue
		}
		log.Verbo("Committing heights", zap.Uint64("toHeight", r.to))
		cm.committedHeight = r.to
		cm.dirty = true
		cm.metrics.UpdateCommittedHeight(cm.relayerID, cm.committedHeight)
	}
	cm.metrics.UpdatePendingCommitsHeapLength(cm.relayerID, cm.pendingCommits.Len())
}
