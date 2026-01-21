// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package clients

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/icm-services/config"
	"github.com/stretchr/testify/require"
)

const contextContextType = "context.Context"

// TestInfoAPI_AllMethodsForwardQueryParams uses reflection to verify that ALL methods
// on InfoAPI correctly forward query parameters
func TestInfoAPI_AllMethodsForwardQueryParams(t *testing.T) {
	queryParams := map[string]string{
		"token":   "test-token-123",
		"api-key": "secret-key-789",
	}

	var lastReceivedParams map[string]string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := make(map[string]string)
		for key := range queryParams {
			params[key] = r.URL.Query().Get(key)
		}
		lastReceivedParams = params

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
				case contextContextType:
					args = append(args, reflect.ValueOf(t.Context()))
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

			require.NotNil(t, lastReceivedParams, "Method %s did not forward query parameters", methodName)
			for key, expectedValue := range queryParams {
				require.Equal(t, expectedValue, lastReceivedParams[key],
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

	var lastReceivedHeaders map[string]string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := make(map[string]string)
		for key := range httpHeaders {
			headers[key] = r.Header.Get(key)
		}
		lastReceivedHeaders = headers

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
				case contextContextType:
					args = append(args, reflect.ValueOf(t.Context()))
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

			require.NotNil(t, lastReceivedHeaders, "Method %s did not forward HTTP headers", methodName)
			for key, expectedValue := range httpHeaders {
				require.Equal(t, expectedValue, lastReceivedHeaders[key],
					"Method %s: header %s not forwarded correctly", methodName, key)
			}
		})
	}
}
