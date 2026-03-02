// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package checkpoint

import (
	"github.com/ava-labs/icm-services/database"
	"github.com/prometheus/client_golang/prometheus"
)

type CheckpointManagerMetrics struct {
	pendingCommitsHeapLength *prometheus.GaugeVec
	committedHeight          *prometheus.GaugeVec
}

func NewCheckpointManagerMetrics(registerer prometheus.Registerer) *CheckpointManagerMetrics {
	m := CheckpointManagerMetrics{
		pendingCommitsHeapLength: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "checkpoint_pending_commits_heap_length",
				Help: "Number of pending commits in the heap",
			},
			[]string{"relayer_id", "destination_blockchain_id", "source_blockchain_id"},
		),
		committedHeight: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "checkpoint_committed_height",
				Help: "Current committed block height",
			},
			[]string{"relayer_id", "destination_blockchain_id", "source_blockchain_id"},
		),
	}

	registerer.MustRegister(m.pendingCommitsHeapLength)
	registerer.MustRegister(m.committedHeight)

	return &m
}

func (m *CheckpointManagerMetrics) UpdatePendingCommitsHeapLength(relayerID database.RelayerID, length int) {
	m.pendingCommitsHeapLength.WithLabelValues(
		relayerID.ID.String(),
		relayerID.DestinationBlockchainID.String(),
		relayerID.SourceBlockchainID.String(),
	).Set(float64(length))
}

func (m *CheckpointManagerMetrics) UpdateCommittedHeight(relayerID database.RelayerID, height uint64) {
	m.committedHeight.WithLabelValues(
		relayerID.ID.String(),
		relayerID.DestinationBlockchainID.String(),
		relayerID.SourceBlockchainID.String(),
	).Set(float64(height))
}
