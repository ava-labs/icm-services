// Copyright (C) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ava-labs/icm-services/config"
	"github.com/stretchr/testify/require"
)

// TestProposerVMAPI_QueryParamsForwarding verifies that query parameters are forwarded correctly
func TestProposerVMAPI_QueryParamsForwarding(t *testing.T) {
	tests := []struct {
		name        string
		queryParams map[string]string
	}{
		{
			name: "single query param",
			queryParams: map[string]string{
				"token": "test-token-123",
			},
		},
		{
			name: "multiple query params",
			queryParams: map[string]string{
				"token":   "test-token-456",
				"api-key": "secret-key-789",
			},
		},
		{
			name: "query params with special characters",
			queryParams: map[string]string{
				"token": "token-with-dashes_and_underscores",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receivedParams := make(map[string]string)

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				for key := range tt.queryParams {
					receivedParams[key] = r.URL.Query().Get(key)
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":{"height":100}}`))
			}))
			defer server.Close()

			apiConfig := &config.APIConfig{
				BaseURL:     server.URL,
				QueryParams: tt.queryParams,
			}
			client := NewProposerVMAPI(server.URL, "test-chain", apiConfig)

			ctx := context.Background()
			client.GetProposedHeight(ctx)

			for key, expectedValue := range tt.queryParams {
				actualValue := receivedParams[key]
				require.Equal(t, expectedValue, actualValue,
					"Query param %s: expected %s, got %s", key, expectedValue, actualValue)
			}
		})
	}
}

// TestProposerVMAPI_HTTPHeadersForwarding verifies that HTTP headers are forwarded correctly
func TestProposerVMAPI_HTTPHeadersForwarding(t *testing.T) {
	tests := []struct {
		name        string
		httpHeaders map[string]string
	}{
		{
			name: "authorization header",
			httpHeaders: map[string]string{
				"Authorization": "Bearer test-token",
			},
		},
		{
			name: "multiple headers",
			httpHeaders: map[string]string{
				"Authorization": "Bearer test-token",
				"X-API-Key":     "secret-key",
				"X-Custom":      "custom-value",
			},
		},
		{
			name: "headers with special values",
			httpHeaders: map[string]string{
				"X-Token": "token-with-dashes-123",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receivedHeaders := make(map[string]string)

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				for key := range tt.httpHeaders {
					receivedHeaders[key] = r.Header.Get(key)
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":{"height":100}}`))
			}))
			defer server.Close()

			apiConfig := &config.APIConfig{
				BaseURL:     server.URL,
				HTTPHeaders: tt.httpHeaders,
			}
			client := NewProposerVMAPI(server.URL, "test-chain", apiConfig)

			ctx := context.Background()
			client.GetProposedHeight(ctx)

			// Verify all headers were received.
			for key, expectedValue := range tt.httpHeaders {
				actualValue := receivedHeaders[key]
				require.Equal(t, expectedValue, actualValue,
					"Header %s: expected %s, got %s", key, expectedValue, actualValue)
			}
		})
	}
}

// TestProposerVMAPI_CombinedQueryParamsAndHeaders verifies both work together
func TestProposerVMAPI_CombinedQueryParamsAndHeaders(t *testing.T) {
	queryParams := map[string]string{
		"token":   "query-token",
		"api-key": "query-key",
	}
	httpHeaders := map[string]string{
		"Authorization": "Bearer header-token",
		"X-API-Key":     "header-key",
	}

	receivedParams := make(map[string]string)
	receivedHeaders := make(map[string]string)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for key := range queryParams {
			receivedParams[key] = r.URL.Query().Get(key)
		}

		for key := range httpHeaders {
			receivedHeaders[key] = r.Header.Get(key)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":{"height":100}}`))
	}))
	defer server.Close()

	apiConfig := &config.APIConfig{
		BaseURL:     server.URL,
		QueryParams: queryParams,
		HTTPHeaders: httpHeaders,
	}
	client := NewProposerVMAPI(server.URL, "test-chain", apiConfig)

	ctx := context.Background()
	client.GetProposedHeight(ctx)

	for key, expectedValue := range queryParams {
		require.Equal(t, expectedValue, receivedParams[key],
			"Query param %s not forwarded correctly", key)
	}

	for key, expectedValue := range httpHeaders {
		require.Equal(t, expectedValue, receivedHeaders[key],
			"Header %s not forwarded correctly", key)
	}
}

// TestProposerVMAPI_GetCurrentEpochWithQueryParams tests GetCurrentEpoch with query params
func TestProposerVMAPI_GetCurrentEpochWithQueryParams(t *testing.T) {
	queryParams := map[string]string{
		"token": "epoch-test-token",
	}

	receivedParams := make(map[string]string)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for key := range queryParams {
			receivedParams[key] = r.URL.Query().Get(key)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":{"Number":1,"PChainHeight":100,"StartTime":1234567890}}`))
	}))
	defer server.Close()

	apiConfig := &config.APIConfig{
		BaseURL:     server.URL,
		QueryParams: queryParams,
	}
	client := NewProposerVMAPI(server.URL, "test-chain", apiConfig)

	ctx := context.Background()
	client.GetCurrentEpoch(ctx)

	for key, expectedValue := range queryParams {
		require.Equal(t, expectedValue, receivedParams[key],
			"Query param %s not forwarded in GetCurrentEpoch", key)
	}
}
