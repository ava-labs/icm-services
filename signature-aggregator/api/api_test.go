// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

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
