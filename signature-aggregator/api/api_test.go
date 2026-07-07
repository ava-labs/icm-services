// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package api

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	networkP2P "github.com/ava-labs/avalanchego/network/p2p"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/logging"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator"
	"github.com/ava-labs/icm-services/signature-aggregator/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
)

// stubAggregator fails every CreateSignedMessage call with a fixed error.
type stubAggregator struct {
	err   error
	calls int
}

func (s *stubAggregator) CreateSignedMessage(
	context.Context,
	logging.Logger,
	*avalancheWarp.UnsignedMessage,
	[]byte,
	ids.ID,
	uint64,
	uint64,
	uint64,
) (*avalancheWarp.Message, error) {
	s.calls++
	return nil, s.err
}

// postValidAggregationRequest sends a well-formed request to a handler backed by [agg] and
// returns the recorded response.
func postValidAggregationRequest(t *testing.T, agg signatureAggregator) *httptest.ResponseRecorder {
	unsignedMessage, err := avalancheWarp.NewUnsignedMessage(
		constants.UnitTestID, ids.GenerateTestID(), []byte("payload"),
	)
	require.NoError(t, err)
	body, err := json.Marshal(AggregateSignatureRequest{
		Message: hex.EncodeToString(unsignedMessage.Bytes()),
	})
	require.NoError(t, err)

	handler := signatureAggregationAPIHandler(
		logging.NoLog{},
		metrics.NewSignatureAggregatorMetrics(prometheus.NewRegistry()),
		agg,
		networkP2P.SignatureRequestHandlerID,
	)
	req := httptest.NewRequest(http.MethodPost, APIPath, strings.NewReader(string(body)))
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	return rec
}

func decodeErrorResponse(t *testing.T, rec *httptest.ResponseRecorder) string {
	var resp AggregateSignatureErrorResponse
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp), "body: %s", rec.Body.String())
	return resp.Error
}

// The endpoint is unauthenticated, so an aggregation failure must never echo the error chain:
// upstream RPC transport errors embed the full request URL, including any API key carried in
// the P-Chain endpoint's path or query string.
func TestAggregateSignaturesErrorResponseDoesNotLeakUpstreamError(t *testing.T) {
	const (
		secret   = "sk_live_SECRET_API_KEY"
		endpoint = "https://rpc.internal.example.com:9650/ext/bc/P?token=" + secret
		dialErr  = "dial tcp 10.1.2.3:9650: connect: connection refused"
	)
	transportErr := &url.Error{Op: "Post", URL: endpoint, Err: errors.New(dialErr)}
	// Mirrors how the aggregator wraps a failed P-Chain lookup before returning it to the API.
	agg := &stubAggregator{err: fmt.Errorf("failed to get subnet: %w", transportErr)}
	// Guard against the fixture itself being harmless.
	require.Contains(t, agg.err.Error(), secret)

	rec := postValidAggregationRequest(t, agg)
	require.Equal(t, 1, agg.calls, "request must reach the aggregator")
	require.Equal(t, http.StatusInternalServerError, rec.Code)
	require.Equal(t, "failed to aggregate signatures", decodeErrorResponse(t, rec))
	for _, leaked := range []string{secret, "rpc.internal.example.com", "9650", "10.1.2.3", "dial tcp", "token="} {
		require.NotContains(t, rec.Body.String(), leaked)
	}
}

// Sentinel errors defined by the aggregator carry fixed, caller-relevant text and are the only
// errors reported specifically; their wrapping context is still dropped.
func TestAggregateSignaturesErrorResponseClassification(t *testing.T) {
	tests := []struct {
		name       string
		err        error
		wantStatus int
		wantError  string
	}{
		{
			name:       "request too large",
			err:        fmt.Errorf("%w: 5000000 > 4000000 bytes", aggregator.ErrRequestTooLarge),
			wantStatus: http.StatusRequestEntityTooLarge,
			wantError:  aggregator.ErrRequestTooLarge.Error(),
		},
		{
			name:       "not enough connected stake",
			err:        fmt.Errorf("wrapped: %w", aggregator.ErrNotEnoughConnectedStake),
			wantStatus: http.StatusInternalServerError,
			wantError:  aggregator.ErrNotEnoughConnectedStake.Error(),
		},
		{
			name:       "not enough signatures",
			err:        aggregator.ErrNotEnoughSignatures,
			wantStatus: http.StatusInternalServerError,
			wantError:  aggregator.ErrNotEnoughSignatures.Error(),
		},
		{
			name:       "anything else is generic",
			err:        errors.New("source blockchain not found for chain ID abc"),
			wantStatus: http.StatusInternalServerError,
			wantError:  "failed to aggregate signatures",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := postValidAggregationRequest(t, &stubAggregator{err: tt.err})
			require.Equal(t, tt.wantStatus, rec.Code)
			require.Equal(t, tt.wantError, decodeErrorResponse(t, rec))
			// The response carries only the fixed text, never the wrapping context.
			require.NotContains(t, rec.Body.String(), "wrapped")
			require.NotContains(t, rec.Body.String(), "5000000")
			require.NotContains(t, rec.Body.String(), "abc")
		})
	}
}

// The handler must reject a body larger than MaxRequestBodySize before decoding it, and must not
// reject a body that fits. Both cases fail before the aggregator is used, so it can be nil.
func TestAggregateSignaturesBodySizeLimit(t *testing.T) {
	mux := http.NewServeMux()
	HandleAggregateSignaturesByRawMsgRequest(
		mux,
		logging.NoLog{},
		metrics.NewSignatureAggregatorMetrics(prometheus.NewRegistry()),
		nil,
	)

	// One hex string that alone exceeds the whole body limit, which is the shape of the
	// memory exhaustion attack: the decoder must stop reading at the limit.
	oversized := `{"message":"` + strings.Repeat("a", MaxRequestBodySize) + `"}`
	req := httptest.NewRequest(http.MethodPost, APIPath, strings.NewReader(oversized))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	require.Equal(t, http.StatusRequestEntityTooLarge, rec.Code)
	require.Contains(t, rec.Body.String(), "Request body too large")

	// A body exactly at the limit is read in full: JSON whitespace padding up to the limit
	// followed by an invalid token must produce a decode error, not a size error.
	atLimit := strings.Repeat(" ", MaxRequestBodySize-len("not json")) + "not json"
	require.Len(t, atLimit, MaxRequestBodySize)
	req = httptest.NewRequest(http.MethodPost, APIPath, strings.NewReader(atLimit))
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	require.Equal(t, http.StatusBadRequest, rec.Code)
	require.Contains(t, rec.Body.String(), "Could not decode request body")
}

func TestTruncateForLog(t *testing.T) {
	short := strings.Repeat("a", maxLoggedFieldLen)
	require.Equal(t, short, truncateForLog(short))

	long := strings.Repeat("a", maxLoggedFieldLen+1)
	require.Equal(t, short+"...", truncateForLog(long))
}
