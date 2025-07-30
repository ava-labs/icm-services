package peers

import (
	"github.com/prometheus/client_golang/prometheus"
)

type AppRequestNetworkMetrics struct {
	pChainAPICallLatencyMS prometheus.Histogram
	connects               prometheus.Counter
	disconnects            prometheus.Counter
}

func newAppRequestNetworkMetrics(registerer prometheus.Registerer) *AppRequestNetworkMetrics {
	m := AppRequestNetworkMetrics{
		pChainAPICallLatencyMS: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Name:    "p_chain_api_call_latency_ms",
				Help:    "Latency of calling p-chain rpc in milliseconds",
				Buckets: prometheus.ExponentialBucketsRange(100, 10000, 10),
			},
		),
		connects: prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: "connects",
				Help: "Number of connected events",
			},
		),
		disconnects: prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: "disconnects",
				Help: "Number of disconnected events",
			},
		),
	}
	registerer.MustRegister(m.pChainAPICallLatencyMS)
	registerer.MustRegister(m.connects)
	registerer.MustRegister(m.disconnects)

	return &m
}
