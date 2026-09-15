package metrics

import (
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/stretchr/testify/require"
)

// freePort reserves an ephemeral port and releases it so that the metrics server can bind to it.
func freePort(t *testing.T) uint16 {
	t.Helper()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	port := listener.Addr().(*net.TCPAddr).Port
	require.NoError(t, listener.Close())
	return uint16(port)
}

// The metrics listener must serve only /metrics. Anything registered on http.DefaultServeMux
// elsewhere in the process (e.g. the service API endpoints) must not be reachable on it.
func TestMetricsServerDoesNotServeDefaultServeMux(t *testing.T) {
	const sentinelPath = "/sentinel-api-endpoint"
	http.Handle(sentinelPath, http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	port := freePort(t)
	_, err := StartMetricsServer(logging.NoLog{}, port, []string{"test"})
	require.NoError(t, err)

	baseURL := fmt.Sprintf("http://127.0.0.1:%d", port)
	client := &http.Client{Timeout: 5 * time.Second}

	require.Eventually(t, func() bool {
		resp, err := client.Get(baseURL + MetricsPath)
		if err != nil {
			return false
		}
		defer resp.Body.Close()
		return resp.StatusCode == http.StatusOK
	}, 10*time.Second, 50*time.Millisecond, "metrics endpoint never became available")

	resp, err := client.Get(baseURL + sentinelPath)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
