package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/stretchr/testify/require"
)

// The relayer API endpoints must be registered on the mux they are given, and that mux must not
// expose anything registered elsewhere in the process, in particular the /metrics endpoint.
func TestRelayerAPIRegistrationIsScopedToItsMux(t *testing.T) {
	mux := http.NewServeMux()
	HandleRelay(mux, logging.NoLog{}, nil)
	HandleRelayMessage(mux, logging.NoLog{}, nil)

	// Registering on the default mux must have no effect on the API mux.
	http.Handle("/relayer-api-test-metrics", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	for _, path := range []string{RelayAPIPath, RelayMessageAPIPath} {
		// An undecodable body short-circuits before the message coordinator is used,
		// so a 400 confirms the endpoint is routed on this mux.
		req := httptest.NewRequest(http.MethodPost, path, strings.NewReader("not json"))
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)
		require.Equal(t, http.StatusBadRequest, rec.Code, "expected %s to be served by the API mux", path)
	}

	for _, path := range []string{"/metrics", "/relayer-api-test-metrics"} {
		req := httptest.NewRequest(http.MethodGet, path, nil)
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)
		require.Equal(t, http.StatusNotFound, rec.Code, "expected %s to be absent from the API mux", path)
	}
}
