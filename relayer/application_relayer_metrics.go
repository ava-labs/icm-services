// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package relayer

import (
	"github.com/ava-labs/icm-services/database"
	"github.com/ava-labs/icm-services/messages"
	"github.com/prometheus/client_golang/prometheus"
)

type ApplicationRelayerMetrics struct {
	successfulRelayMessageCount   *prometheus.CounterVec
	createSignedMessageLatencyMS  *prometheus.GaugeVec
	failedRelayMessageCount       *prometheus.CounterVec
	fetchSignatureAppRequestCount *prometheus.CounterVec
}

func NewApplicationRelayerMetrics(registerer prometheus.Registerer) *ApplicationRelayerMetrics {
	m := ApplicationRelayerMetrics{
		successfulRelayMessageCount: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "successful_relay_message_count",
				Help: "Number of messages that relayed successfully",
			},
			[]string{"destination_chain_id", "source_chain_id"},
		),
		createSignedMessageLatencyMS: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "create_signed_message_latency_ms",
				Help: "Latency of creating a signed message in milliseconds",
			},
			[]string{"destination_chain_id", "source_chain_id"},
		),
		failedRelayMessageCount: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "failed_relay_message_count",
				Help: "Number of messages that failed to relay",
			},
			[]string{"destination_chain_id", "source_chain_id", "failure_reason"},
		),
		fetchSignatureAppRequestCount: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "fetch_signature_app_request_count",
				Help: "Number of aggregate signatures constructed via AppRequest",
			},
			[]string{"destination_chain_id", "source_chain_id"},
		),
	}

	registerer.MustRegister(m.successfulRelayMessageCount)
	registerer.MustRegister(m.createSignedMessageLatencyMS)
	registerer.MustRegister(m.failedRelayMessageCount)
	registerer.MustRegister(m.fetchSignatureAppRequestCount)

	return &m
}

// relayerMetrics binds ApplicationRelayerMetrics to a single relayer's label set so that it can be
// handed to a message handler as messages.Metrics, without the handler depending on the relayer.
type relayerMetrics struct {
	metrics            *ApplicationRelayerMetrics
	destinationChainID string
	sourceChainID      string
}

var _ messages.Metrics = (*relayerMetrics)(nil)

// forRelayer returns a messages.Metrics that records metrics for the given relayer.
func (m *ApplicationRelayerMetrics) forRelayer(relayerID database.RelayerID) *relayerMetrics {
	return &relayerMetrics{
		metrics:            m,
		destinationChainID: relayerID.DestinationBlockchainID.String(),
		sourceChainID:      relayerID.SourceBlockchainID.String(),
	}
}

func (m *relayerMetrics) IncSuccessfulRelayMessageCount() {
	m.metrics.successfulRelayMessageCount.
		WithLabelValues(m.destinationChainID, m.sourceChainID).Inc()
}

func (m *relayerMetrics) IncFailedRelayMessageCount(failureReason string) {
	m.metrics.failedRelayMessageCount.
		WithLabelValues(m.destinationChainID, m.sourceChainID, failureReason).Inc()
}

func (m *relayerMetrics) SetCreateSignedMessageLatencyMS(latency float64) {
	m.metrics.createSignedMessageLatencyMS.
		WithLabelValues(m.destinationChainID, m.sourceChainID).Set(latency)
}

func (m *relayerMetrics) IncFetchSignatureAppRequestCount() {
	m.metrics.fetchSignatureAppRequestCount.
		WithLabelValues(m.destinationChainID, m.sourceChainID).Inc()
}
