// Copyright (C) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package clients

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/ava-labs/icm-services/config"
	"github.com/stretchr/testify/require"
)

// TestProposerVMAPI_AllMethodsForwardQueryParams uses reflection to verify that ALL methods
// on ProposerVMAPI correctly forward query parameters
func TestProposerVMAPI_AllMethodsForwardQueryParams(t *testing.T) {
	queryParams := map[string]string{
		"token":   "test-token-123",
		"api-key": "secret-key-789",
	}

	var lastReceivedParams map[string]string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Capture query params from this request
		params := make(map[string]string)
		for key := range queryParams {
			params[key] = r.URL.Query().Get(key)
		}
		lastReceivedParams = params

		// Return a generic valid JSON-RPC response
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

	// Use reflection to call all methods
	clientValue := reflect.ValueOf(client)
	clientType := clientValue.Type()

	for i := 0; i < clientType.NumMethod(); i++ {
		method := clientType.Method(i)
		methodName := method.Name

		// Skip unexported methods
		if method.PkgPath != "" {
			continue
		}

		t.Run(methodName, func(t *testing.T) {
			// Prepare arguments for the method
			args := []reflect.Value{clientValue}

			// Add context as first argument if the method takes one
			methodType := method.Type
			if methodType.NumIn() > 1 && methodType.In(1).String() == "context.Context" {
				args = append(args, reflect.ValueOf(t.Context()))
			}

			// Call the method (may fail due to mock response format, but we only care about HTTP request)
			method.Func.Call(args)

			// Verify query params were forwarded for this method
			require.NotNil(t, lastReceivedParams, "Method %s did not forward query parameters", methodName)
			for key, expectedValue := range queryParams {
				require.Equal(t, expectedValue, lastReceivedParams[key],
					"Method %s: query param %s not forwarded correctly", methodName, key)
			}
		})
	}
}

// TestProposerVMAPI_AllMethodsForwardHTTPHeaders uses reflection to verify that ALL methods
// on ProposerVMAPI correctly forward HTTP headers
func TestProposerVMAPI_AllMethodsForwardHTTPHeaders(t *testing.T) {
	httpHeaders := map[string]string{
		"Authorization": "Bearer test-token",
		"X-API-Key":     "secret-key",
	}

	var lastReceivedHeaders map[string]string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := make(map[string]string)
		for key := range httpHeaders {
			headers[key] = r.Header.Get(key)
		}
		lastReceivedHeaders = headers

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":{"Number":1,"PChainHeight":100,"StartTime":1234567890}}`))
	}))
	defer server.Close()

	apiConfig := &config.APIConfig{
		BaseURL:     server.URL,
		HTTPHeaders: httpHeaders,
	}
	client := NewProposerVMAPI(server.URL, "test-chain", apiConfig)

	clientValue := reflect.ValueOf(client)
	clientType := clientValue.Type()

	for i := 0; i < clientType.NumMethod(); i++ {
		method := clientType.Method(i)
		methodName := method.Name

		if method.PkgPath != "" {
			continue
		}

		t.Run(methodName, func(t *testing.T) {
			args := []reflect.Value{clientValue}

			methodType := method.Type
			if methodType.NumIn() > 1 && methodType.In(1).String() == "context.Context" {
				args = append(args, reflect.ValueOf(t.Context()))
			}

			method.Func.Call(args)

			require.NotNil(t, lastReceivedHeaders, "Method %s did not forward HTTP headers", methodName)
			for key, expectedValue := range httpHeaders {
				require.Equal(t, expectedValue, lastReceivedHeaders[key],
					"Method %s: header %s not forwarded correctly", methodName, key)
			}
		})
	}
}
