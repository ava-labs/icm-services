// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
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

// TestInfoAPI_QueryParamsForwarding verifies that query parameters are forwarded correctly
func TestInfoAPI_QueryParamsForwarding(t *testing.T) {
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
				w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":12345}`))
			}))
			defer server.Close()

			apiConfig := &config.APIConfig{
				BaseURL:     server.URL,
				QueryParams: tt.queryParams,
			}
			client, err := NewInfoAPI(apiConfig)
			require.NoError(t, err)

			ctx := context.Background()
			client.GetNetworkID(ctx)

			for key, expectedValue := range tt.queryParams {
				actualValue := receivedParams[key]
				require.Equal(t, expectedValue, actualValue,
					"Query param %s: expected %s, got %s", key, expectedValue, actualValue)
			}
		})
	}
}

// TestInfoAPI_HTTPHeadersForwarding verifies that HTTP headers are forwarded correctly
func TestInfoAPI_HTTPHeadersForwarding(t *testing.T) {
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
				w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":12345}`))
			}))
			defer server.Close()

			apiConfig := &config.APIConfig{
				BaseURL:     server.URL,
				HTTPHeaders: tt.httpHeaders,
			}
			client, err := NewInfoAPI(apiConfig)
			require.NoError(t, err)

			ctx := context.Background()
			client.GetNetworkID(ctx)

			for key, expectedValue := range tt.httpHeaders {
				actualValue := receivedHeaders[key]
				require.Equal(t, expectedValue, actualValue,
					"Header %s: expected %s, got %s", key, expectedValue, actualValue)
			}
		})
	}
}

// TestInfoAPI_CombinedQueryParamsAndHeaders verifies both work together
func TestInfoAPI_CombinedQueryParamsAndHeaders(t *testing.T) {
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
		w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"mainnet"}`))
	}))
	defer server.Close()

	apiConfig := &config.APIConfig{
		BaseURL:     server.URL,
		QueryParams: queryParams,
		HTTPHeaders: httpHeaders,
	}
	client, err := NewInfoAPI(apiConfig)
	require.NoError(t, err)

	ctx := context.Background()
	client.GetNetworkName(ctx)

	for key, expectedValue := range queryParams {
		require.Equal(t, expectedValue, receivedParams[key],
			"Query param %s not forwarded correctly", key)
	}

	for key, expectedValue := range httpHeaders {
		require.Equal(t, expectedValue, receivedHeaders[key],
			"Header %s not forwarded correctly", key)
	}
}

// TestInfoAPI_MultipleMethodsWithParams tests different Info API methods with params
func TestInfoAPI_MultipleMethodsWithParams(t *testing.T) {
	queryParams := map[string]string{
		"token": "multi-method-token",
	}

	methodsCalled := make(map[string]bool)
	receivedParams := make(map[string]map[string]string)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		methodsCalled[r.URL.Path] = true
		params := make(map[string]string)
		for key := range queryParams {
			params[key] = r.URL.Query().Get(key)
		}
		receivedParams[r.URL.Path] = params

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"test-result"}`))
	}))
	defer server.Close()

	apiConfig := &config.APIConfig{
		BaseURL:     server.URL,
		QueryParams: queryParams,
	}
	client, err := NewInfoAPI(apiConfig)
	require.NoError(t, err)

	ctx := context.Background()

	client.GetNetworkName(ctx)
	client.GetNetworkID(ctx)

	for _, params := range receivedParams {
		for key, expectedValue := range queryParams {
			require.Equal(t, expectedValue, params[key],
				"Query param %s not forwarded correctly", key)
		}
	}
}

// TestInfoAPI_NoQueryParamsOrHeaders verifies client works without additional params
func TestInfoAPI_NoQueryParamsOrHeaders(t *testing.T) {
	requestReceived := false

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestReceived = true
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":12345}`))
	}))
	defer server.Close()

	apiConfig := &config.APIConfig{
		BaseURL: server.URL,
	}
	client, err := NewInfoAPI(apiConfig)
	require.NoError(t, err)

	ctx := context.Background()
	client.GetNetworkID(ctx)

	require.True(t, requestReceived, "Request should have been sent to server")
}
