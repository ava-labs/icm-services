// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sync"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/icm-services/config"
	"github.com/stretchr/testify/require"
)

// TestInfoAPI_AllMethodsForwardQueryParams uses reflection to verify that ALL methods
// on InfoAPI correctly forward query parameters
func TestInfoAPI_AllMethodsForwardQueryParams(t *testing.T) {
	queryParams := map[string]string{
		"token":   "test-token-123",
		"api-key": "secret-key-789",
	}

	var mu sync.Mutex
	methodsCalled := make(map[string]bool)
	receivedParams := make(map[string]map[string]string)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()

		methodsCalled[r.URL.Path] = true
		params := make(map[string]string)
		for key := range queryParams {
			params[key] = r.URL.Query().Get(key)
		}
		receivedParams[r.URL.Path] = params

		// Return a generic valid JSON-RPC response that works for most Info API methods
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

	clientValue := reflect.ValueOf(client)
	clientType := clientValue.Type()

	ctx := context.Background()

	for i := 0; i < clientType.NumMethod(); i++ {
		method := clientType.Method(i)
		methodName := method.Name

		if method.PkgPath != "" {
			continue
		}

		t.Run(methodName, func(t *testing.T) {
			args := []reflect.Value{clientValue}

			methodType := method.Type
			// Build arguments for the method based on its signature
			for argIdx := 1; argIdx < methodType.NumIn(); argIdx++ {
				argType := methodType.In(argIdx)

				switch argType.String() {
				case "context.Context":
					args = append(args, reflect.ValueOf(ctx))
				case "string":
					args = append(args, reflect.ValueOf("test-string"))
				case "ids.ID":
					testID := ids.GenerateTestID()
					args = append(args, reflect.ValueOf(testID))
				case "[]ids.NodeID":
					args = append(args, reflect.ValueOf([]ids.NodeID{}))
				default:
					// For other types, use zero value
					args = append(args, reflect.Zero(argType))
				}
			}

			// Call the method
			method.Func.Call(args)

			mu.Lock()
			defer mu.Unlock()

			var foundParams map[string]string
			for _, params := range receivedParams {
				allPresent := true
				for key, expectedValue := range queryParams {
					if params[key] != expectedValue {
						allPresent = false
						break
					}
				}
				if allPresent {
					foundParams = params
					break
				}
			}

			require.NotNil(t, foundParams, "Method %s did not forward query parameters", methodName)
			for key, expectedValue := range queryParams {
				require.Equal(t, expectedValue, foundParams[key],
					"Method %s: query param %s not forwarded correctly", methodName, key)
			}
		})
	}
}

// TestInfoAPI_AllMethodsForwardHTTPHeaders uses reflection to verify that ALL methods
// on InfoAPI correctly forward HTTP headers
func TestInfoAPI_AllMethodsForwardHTTPHeaders(t *testing.T) {
	httpHeaders := map[string]string{
		"Authorization": "Bearer test-token",
		"X-API-Key":     "secret-key",
	}

	var mu sync.Mutex
	methodsCalled := make(map[string]bool)
	receivedHeaders := make(map[string]map[string]string)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()

		methodsCalled[r.URL.Path] = true
		headers := make(map[string]string)
		for key := range httpHeaders {
			headers[key] = r.Header.Get(key)
		}
		receivedHeaders[r.URL.Path] = headers

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"test-result"}`))
	}))
	defer server.Close()

	apiConfig := &config.APIConfig{
		BaseURL:     server.URL,
		HTTPHeaders: httpHeaders,
	}
	client, err := NewInfoAPI(apiConfig)
	require.NoError(t, err)

	clientValue := reflect.ValueOf(client)
	clientType := clientValue.Type()

	ctx := context.Background()

	for i := 0; i < clientType.NumMethod(); i++ {
		method := clientType.Method(i)
		methodName := method.Name

		if method.PkgPath != "" {
			continue
		}

		t.Run(methodName, func(t *testing.T) {
			args := []reflect.Value{clientValue}

			methodType := method.Type
			for argIdx := 1; argIdx < methodType.NumIn(); argIdx++ {
				argType := methodType.In(argIdx)

				switch argType.String() {
				case "context.Context":
					args = append(args, reflect.ValueOf(ctx))
				case "string":
					args = append(args, reflect.ValueOf("test-string"))
				case "ids.ID":
					testID := ids.GenerateTestID()
					args = append(args, reflect.ValueOf(testID))
				case "[]ids.NodeID":
					args = append(args, reflect.ValueOf([]ids.NodeID{}))
				default:
					args = append(args, reflect.Zero(argType))
				}
			}

			method.Func.Call(args)

			mu.Lock()
			defer mu.Unlock()

			var foundHeaders map[string]string
			for _, headers := range receivedHeaders {
				allPresent := true
				for key, expectedValue := range httpHeaders {
					if headers[key] != expectedValue {
						allPresent = false
						break
					}
				}
				if allPresent {
					foundHeaders = headers
					break
				}
			}

			require.NotNil(t, foundHeaders, "Method %s did not forward HTTP headers", methodName)
			for key, expectedValue := range httpHeaders {
				require.Equal(t, expectedValue, foundHeaders[key],
					"Method %s: header %s not forwarded correctly", methodName, key)
			}
		})
	}
}
