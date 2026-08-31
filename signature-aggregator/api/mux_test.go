package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/icm-services/signature-aggregator/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
)

// The aggregation endpoint must be registered on the mux it is given, and that mux must not
// expose anything registered elsewhere in the process, in particular the /metrics endpoint.
func TestAggregateSignaturesRegistrationIsScopedToItsMux(t *testing.T) {
	mux := http.NewServeMux()
	HandleAggregateSignaturesByRawMsgRequest(
		mux,
		logging.NoLog{},
		metrics.NewSignatureAggregatorMetrics(prometheus.NewRegistry()),
		nil,
	)

	// Registering on the default mux must have no effect on the API mux.
	http.Handle("/sig-agg-api-test-metrics", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// An undecodable body short-circuits before the aggregator is used,
	// so a 400 confirms the endpoint is routed on this mux.
	req := httptest.NewRequest(http.MethodPost, APIPath, strings.NewReader("not json"))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	require.Equal(t, http.StatusBadRequest, rec.Code)

	for _, path := range []string{"/metrics", "/sig-agg-api-test-metrics"} {
		req := httptest.NewRequest(http.MethodGet, path, nil)
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)
		require.Equal(t, http.StatusNotFound, rec.Code, "expected %s to be absent from the API mux", path)
	}
}
