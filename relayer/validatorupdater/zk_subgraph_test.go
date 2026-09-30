// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validatorupdater

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// newSubgraphServer returns a test server that responds to every request with
// the given body and status, capturing the last request body for inspection.
func newSubgraphServer(t *testing.T, status int, body string) (*httptest.Server, *string) {
	t.Helper()
	var lastRequest string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		lastRequest = string(reqBody)
		w.WriteHeader(status)
		_, err = w.Write([]byte(body))
		require.NoError(t, err)
	}))
	t.Cleanup(server.Close)
	return server, &lastRequest
}

// Happy path. Proofs must come back decoded (hex to bytes, string slot to uint64)
// and in the ascending order the subgraph returned them, since transitions apply
// sequentially against the contract's current state.
func TestConsensusProofsAfterSlot(t *testing.T) {
	response := `{
		"data": {
			"signalEthereumProofs": [
				{"finalizedSlot": "1000", "journalData": "0xdead", "seal": "0xbeef"},
				{"finalizedSlot": "1032", "journalData": "0xcafe", "seal": "0xf00d"}
			]
		}
	}`
	server, lastRequest := newSubgraphServer(t, http.StatusOK, response)
	client := NewSubgraphClient(server.URL)

	proofs, err := client.ConsensusProofsAfterSlot(context.Background(), 968, 4)
	require.NoError(t, err)
	require.Len(t, proofs, 2)

	require.Equal(t, uint64(1000), proofs[0].FinalizedSlot)
	require.Equal(t, []byte{0xde, 0xad}, proofs[0].JournalData)
	require.Equal(t, []byte{0xbe, 0xef}, proofs[0].Seal)
	require.Equal(t, uint64(1032), proofs[1].FinalizedSlot)
	require.Equal(t, []byte{0xca, 0xfe}, proofs[1].JournalData)
	require.Equal(t, []byte{0xf0, 0x0d}, proofs[1].Seal)

	// The query must carry the cursor and limit so the subgraph does the
	// filtering; fetching everything and filtering locally would not scale.
	require.Contains(t, *lastRequest, `finalizedSlot_gt: \"968\"`)
	require.Contains(t, *lastRequest, "first: 4")

	// The request body must be valid JSON with a query field.
	var envelope struct {
		Query string `json:"query"`
	}
	require.NoError(t, json.Unmarshal([]byte(*lastRequest), &envelope))
	require.True(t, strings.Contains(envelope.Query, "signalEthereumProofs"))
}

// No new proofs is a normal condition (the tick simply does nothing), so it
// must decode as an empty slice, not an error.
func TestConsensusProofsAfterSlotEmpty(t *testing.T) {
	server, _ := newSubgraphServer(t, http.StatusOK, `{"data": {"signalEthereumProofs": []}}`)
	client := NewSubgraphClient(server.URL)

	proofs, err := client.ConsensusProofsAfterSlot(context.Background(), 968, 4)
	require.NoError(t, err)
	require.Empty(t, proofs)
}

// GraphQL reports query failures in the response's errors field with HTTP 200;
// without surfacing them, a malformed query would look identical to "no new
// proofs" and silently stall consensus updates.
func TestConsensusProofsAfterSlotGraphQLError(t *testing.T) {
	response := `{"errors": [{"message": "Type SignalEthereumProof has no field finalizedSlots"}]}`
	server, _ := newSubgraphServer(t, http.StatusOK, response)
	client := NewSubgraphClient(server.URL)

	_, err := client.ConsensusProofsAfterSlot(context.Background(), 968, 4)
	require.ErrorContains(t, err, "no field finalizedSlots")
}

// Non-200 responses (gateway errors, rate limiting) must error rather than be
// parsed as empty results.
func TestConsensusProofsAfterSlotHTTPError(t *testing.T) {
	server, _ := newSubgraphServer(t, http.StatusBadGateway, "bad gateway")
	client := NewSubgraphClient(server.URL)

	_, err := client.ConsensusProofsAfterSlot(context.Background(), 968, 4)
	require.ErrorContains(t, err, "status 502")
}

// Malformed JSON (a proxy error page, a truncated response) must surface as a
// decode error.
func TestConsensusProofsAfterSlotBadJSON(t *testing.T) {
	server, _ := newSubgraphServer(t, http.StatusOK, `<html>not json</html>`)
	client := NewSubgraphClient(server.URL)

	_, err := client.ConsensusProofsAfterSlot(context.Background(), 968, 4)
	require.ErrorContains(t, err, "failed to decode")
}
