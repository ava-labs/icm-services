// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package relayer

import (
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/icm-services/database"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/messages/mocks"
	"github.com/ava-labs/icm-services/relayer/checkpoint"
	"github.com/ava-labs/icm-services/vms/evm"
	"github.com/ava-labs/libevm/common"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

// Each on-chain protocol has its own listener, and every listener's blocks reach the
// coordinator. A block range must only be committed for the relayers of the protocol whose
// listener produced it: committing it for another protocol's relayers would advance their
// checkpoints before their own listener has delivered (or even seen) their messages in that
// range, and after a restart those messages would be skipped for good.
func TestProcessBlockOnlyCommitsForListenerProtocol(t *testing.T) {
	ctrl := gomock.NewController(t)
	sourceBlockchainID := ids.GenerateTestID()
	destinationBlockchainID := ids.GenerateTestID()
	protocolA := common.HexToAddress("0xaa")
	protocolB := common.HexToAddress("0xbb")

	registry := prometheus.NewRegistry()
	checkpointMetrics := checkpoint.NewCheckpointManagerMetrics(registry)
	newRelayer := func(protocol common.Address) *ApplicationRelayer {
		id := database.NewRelayerID(
			protocol,
			sourceBlockchainID,
			destinationBlockchainID,
			database.AllAllowedAddress,
			database.AllAllowedAddress,
		)
		return &ApplicationRelayer{
			logger:                  logging.NoLog{},
			metrics:                 mocks.NewMockMetrics(ctrl),
			relayerID:               id,
			checkpointManager:       newTestCheckpointManagerForRelayer(t, ctrl, id, checkpointMetrics),
			processMessageSemaphore: make(chan struct{}, 1),
		}
	}
	relayerA := newRelayer(protocolA)
	relayerB := newRelayer(protocolB)

	coordinator := NewMessageCoordinator(
		logging.NoLog{},
		map[ids.ID]map[common.Address]messages.MessageHandlerFactory{
			sourceBlockchainID: {
				protocolA: mocks.NewMockMessageHandlerFactory(ctrl),
				protocolB: mocks.NewMockMessageHandlerFactory(ctrl),
			},
		},
		map[common.Hash]*ApplicationRelayer{
			relayerA.relayerID.ID: relayerA,
			relayerB.relayerID.ID: relayerB,
		},
		nil,
	)

	// Protocol A's listener saw blocks 1-5 with no messages, the common case for a foreign
	// protocol's listener: it never carries messages for protocol B's relayers.
	errChan := make(chan error, 1)
	coordinator.ProcessBlock(
		&evm.ICMBlockInfo{FromBlock: 1, ToBlock: 5},
		sourceBlockchainID,
		protocolA,
		errChan,
	)

	require.Eventually(t, func() bool {
		return committedHeight(t, registry, relayerA.relayerID) == 5
	}, time.Second, 10*time.Millisecond, "the listener's own relayer should commit the range")
	require.Never(t, func() bool {
		return committedHeight(t, registry, relayerB.relayerID) != 0
	}, 200*time.Millisecond, 10*time.Millisecond,
		"another protocol's relayer must not have the range committed for it")
	select {
	case err := <-errChan:
		t.Fatalf("unexpected error: %v", err)
	default:
	}
}

// committedHeight reads a relayer's committed height back from its checkpoint metrics.
func committedHeight(t *testing.T, registry *prometheus.Registry, id database.RelayerID) uint64 {
	t.Helper()
	families, err := registry.Gather()
	require.NoError(t, err)
	for _, family := range families {
		if family.GetName() != "checkpoint_committed_height" {
			continue
		}
		for _, metric := range family.GetMetric() {
			if hasLabel(metric, "relayer_id", id.ID.String()) {
				return uint64(metric.GetGauge().GetValue())
			}
		}
	}
	t.Fatalf("no committed height metric for relayer %s", id.ID.String())
	return 0
}

func hasLabel(metric *dto.Metric, name, value string) bool {
	for _, label := range metric.GetLabel() {
		if label.GetName() == name && label.GetValue() == value {
			return true
		}
	}
	return false
}
