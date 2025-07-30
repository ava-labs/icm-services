// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package relayer

import (
	"github.com/prometheus/client_golang/prometheus"
)

type ApplicationRelayerMetrics struct {
	successfulRelayMessageCount   *prometheus.CounterVec
	createSignedMessageLatencyMS  *prometheus.GaugeVec
	failedRelayMessageCount       *prometheus.CounterVec
	fetchSignatureAppRequestCount *prometheus.CounterVec
	fetchSignatureRPCCount        *prometheus.CounterVec
}

func NewApplicationRelayerMetrics(registerer prometheus.Registerer) *ApplicationRelayerMetrics {
	m := ApplicationRelayerMetrics{
		successfulRelayMessageCount: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "successful_relay_message_count",
				Help: "Number of messages that relayed successfully",
			},
			[]string{"destination_chain_id", "source_chain_id", "source_subnet_id"},
		),
		createSignedMessageLatencyMS: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "create_signed_message_latency_ms",
				Help: "Latency of creating a signed message in milliseconds",
			},
			[]string{"destination_chain_id", "source_chain_id", "source_subnet_id"},
		),
		failedRelayMessageCount: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "failed_relay_message_count",
				Help: "Number of messages that failed to relay",
			},
			[]string{"destination_chain_id", "source_chain_id", "source_subnet_id", "failure_reason"},
		),
		fetchSignatureAppRequestCount: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "fetch_signature_app_request_count",
				Help: "Number of aggregate signatures constructed via AppRequest",
			},
			[]string{"destination_chain_id", "source_chain_id", "source_subnet_id"},
		),
		fetchSignatureRPCCount: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "fetch_signature_rpc_count",
				Help: "Number of aggregate signatures fetched via Warp API",
			},
			[]string{"destination_chain_id", "source_chain_id", "source_subnet_id"},
		),
	}

	registerer.MustRegister(m.successfulRelayMessageCount)
	registerer.MustRegister(m.createSignedMessageLatencyMS)
	registerer.MustRegister(m.failedRelayMessageCount)
	registerer.MustRegister(m.fetchSignatureAppRequestCount)
	registerer.MustRegister(m.fetchSignatureRPCCount)

	return &m
}
