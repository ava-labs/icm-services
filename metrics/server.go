package metrics

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ava-labs/avalanchego/api/metrics"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

const (
	MetricsPath = "/metrics"

	metricsReadHeaderTimeout = 10 * time.Second
)

// Starts a metrics server on the given port and registers the provided names with the metrics gatherer.
// Returns a map of registries, keyed by the provided names.
func StartMetricsServer(logger logging.Logger, port uint16, names []string) (map[string]*prometheus.Registry, error) {
	gatherer := metrics.NewPrefixGatherer()

	registries := make(map[string]*prometheus.Registry, len(names))
	for _, name := range names {
		registry := prometheus.NewRegistry()
		if err := gatherer.Register(name, registry); err != nil {
			return nil, err
		}
		registries[name] = registry
	}

	// Serve a dedicated mux rather than http.DefaultServeMux so that the metrics listener
	// exposes only /metrics, and never the service APIs registered elsewhere in the process.
	mux := http.NewServeMux()
	mux.Handle(
		MetricsPath,
		promhttp.HandlerFor(gatherer, promhttp.HandlerOpts{}),
	)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           mux,
		ReadHeaderTimeout: metricsReadHeaderTimeout,
	}

	go func() {
		logger.Info(
			"Starting metrics server...",
			zap.Uint16("port", port),
		)
		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			logger.Info("Metrics check server closed")
		} else if err != nil {
			logger.Fatal("Metrics check server exited with error", zap.Error(err))
			os.Exit(1)
		}
	}()

	return registries, nil
}
