// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	basecfg "github.com/ava-labs/icm-services/config"
	"github.com/ava-labs/icm-services/relayer/config"
	mock_ethclient "github.com/ava-labs/icm-services/vms/evm/mocks"
	"github.com/ava-labs/icm-services/vms/evm/signer"
	"github.com/ava-labs/libevm/core/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

var destinationSubnet = config.DestinationBlockchain{
	SubnetID:     "2TGBXcnwx5PqiXWiqxAKUaNSqDguXNh1mxnp82jui68hxJSZAx",
	BlockchainID: "S4mMqUXe7vHsGiRAma6bv3CKnyaLssyAxmQ2KvFpX1KEvfFCD",
	RPCEndpoint: basecfg.APIConfig{
		BaseURL: "https://subnets.avax.network/mysubnet/rpc",
	},
	AccountPrivateKeys: []string{"56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"},
}

func TestGetFeePerGas(t *testing.T) {
	testCases := []struct {
		name                       string
		maxBaseFee                 *big.Int
		suggestedPriorityFeeBuffer *big.Int
		maxPriorityFeePerGas       *big.Int
		estimateBaseFee            *big.Int
		estimateBaseFeeErr         error
		suggestGasTipCap           *big.Int
		suggestGasTipCapErr        error
		expectedGasFeeCap          *big.Int
		expectedGasTipCap          *big.Int
		expectedError              error
	}{
		{
			name:                       "use configured max base fee",
			maxBaseFee:                 big.NewInt(1),
			suggestedPriorityFeeBuffer: big.NewInt(0),
			maxPriorityFeePerGas:       big.NewInt(0),
			suggestGasTipCap:           big.NewInt(0),
			expectedGasFeeCap:          big.NewInt(1),
			expectedGasTipCap:          big.NewInt(0),
		},
		{
			name:                       "error estimate base fee",
			maxBaseFee:                 big.NewInt(0),
			suggestedPriorityFeeBuffer: big.NewInt(0),
			maxPriorityFeePerGas:       big.NewInt(0),
			estimateBaseFee:            nil,
			estimateBaseFeeErr:         context.DeadlineExceeded,
			expectedError:              context.DeadlineExceeded,
		},
		{
			name:                       "use base fee estimate multiple",
			maxBaseFee:                 big.NewInt(0),
			suggestedPriorityFeeBuffer: big.NewInt(0),
			maxPriorityFeePerGas:       big.NewInt(0),
			estimateBaseFee:            big.NewInt(1),
			suggestGasTipCap:           big.NewInt(0),
			expectedGasFeeCap:          big.NewInt(1 * defaultBaseFeeFactor),
			expectedGasTipCap:          big.NewInt(0),
		},
		{
			name:                "error suggest gas tip cap",
			maxBaseFee:          big.NewInt(1),
			suggestGasTipCapErr: context.DeadlineExceeded,
			expectedError:       context.DeadlineExceeded,
		},
		{
			name:                       "suggest gas tip cap + buffer > max priority fee per gas",
			maxBaseFee:                 big.NewInt(1),
			suggestedPriorityFeeBuffer: big.NewInt(2),
			maxPriorityFeePerGas:       big.NewInt(3),
			suggestGasTipCap:           big.NewInt(2),
			expectedGasFeeCap:          big.NewInt(4),
			expectedGasTipCap:          big.NewInt(3),
		},
		{
			name:                       "suggest gas tip cap + buffer < max priority fee per gas",
			maxBaseFee:                 big.NewInt(1),
			suggestedPriorityFeeBuffer: big.NewInt(2),
			maxPriorityFeePerGas:       big.NewInt(10),
			suggestGasTipCap:           big.NewInt(2),
			expectedGasFeeCap:          big.NewInt(5),
			expectedGasTipCap:          big.NewInt(4),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockClient := mock_ethclient.NewMockDestinationRPCClient(ctrl)
			gasFeeConfig := GasFeeConfig{
				maxBaseFee:                 test.maxBaseFee,
				suggestedPriorityFeeBuffer: test.suggestedPriorityFeeBuffer,
				maxPriorityFeePerGas:       test.maxPriorityFeePerGas,
			}
			destClient := destinationClient{
				logger:       logging.NoLog{},
				avaRPCClient: mockClient,
				gasFeeConfig: &gasFeeConfig,
			}

			estimatedBaseFeeTimes := 0
			if test.maxBaseFee.Cmp(big.NewInt(0)) == 0 {
				estimatedBaseFeeTimes = 1
			}
			suggestGasTipCapTimes := 0
			if test.estimateBaseFeeErr == nil {
				suggestGasTipCapTimes = 1
			}

			gomock.InOrder(
				mockClient.EXPECT().EstimateBaseFee(gomock.Any()).Return(
					test.estimateBaseFee,
					test.estimateBaseFeeErr,
				).Times(estimatedBaseFeeTimes),
				mockClient.EXPECT().SuggestGasTipCap(gomock.Any()).Return(
					test.suggestGasTipCap,
					test.suggestGasTipCapErr,
				).Times(suggestGasTipCapTimes),
			)

			gasFeeCap, gasTipCap, err := destClient.getFeePerGas()
			require.ErrorIs(t, test.expectedError, err)
			require.Equal(t, test.expectedGasFeeCap, gasFeeCap)
			require.Equal(t, test.expectedGasTipCap, gasTipCap)
		})
	}
}

func TestSendTx(t *testing.T) {
	var destClient destinationClient
	txSigners, err := signer.NewTxSigners(destinationSubnet.AccountPrivateKeys)
	require.NoError(t, err)

	signer := &concurrentSigner{
		logger:            logging.NoLog{},
		signer:            txSigners[0],
		currentNonce:      0,
		messageChan:       make(chan txData),
		queuedTxSemaphore: make(chan struct{}, poolTxsPerAccount),
		destinationClient: &destClient,
	}
	go signer.processIncomingTransactions()

	testError := fmt.Errorf("call errored")
	testCases := []struct {
		name                  string
		chainIDErr            error
		chainIDTimes          int
		maxBaseFee            *big.Int
		estimateBaseFeeErr    error
		estimateBaseFeeTimes  int
		suggestGasTipCapErr   error
		suggestGasTipCapTimes int
		sendTransactionErr    error
		sendTransactionTimes  int
		txReceiptTimes        int
		expectError           bool
	}{
		{
			name:                  "valid - use base fee estimate",
			chainIDTimes:          1,
			maxBaseFee:            big.NewInt(0),
			estimateBaseFeeTimes:  1,
			suggestGasTipCapTimes: 1,
			sendTransactionTimes:  1,
			txReceiptTimes:        1,
		},
		{
			name:                  "valid - max base fee",
			chainIDTimes:          1,
			maxBaseFee:            big.NewInt(100),
			estimateBaseFeeTimes:  0,
			suggestGasTipCapTimes: 1,
			sendTransactionTimes:  1,
			txReceiptTimes:        1,
		},
		{
			name:                 "invalid estimateBaseFee",
			maxBaseFee:           big.NewInt(0),
			estimateBaseFeeErr:   testError,
			estimateBaseFeeTimes: 1,
			expectError:          true,
		},
		{
			name:                  "invalid suggestGasTipCap",
			maxBaseFee:            big.NewInt(0),
			estimateBaseFeeTimes:  1,
			suggestGasTipCapErr:   testError,
			suggestGasTipCapTimes: 1,
			expectError:           true,
		},
		{
			name:                  "invalid sendTransaction",
			chainIDTimes:          1,
			maxBaseFee:            big.NewInt(0),
			estimateBaseFeeTimes:  1,
			suggestGasTipCapTimes: 1,
			sendTransactionErr:    testError,
			sendTransactionTimes:  1,
			expectError:           true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockClient := mock_ethclient.NewMockDestinationRPCClient(ctrl)
			gasFeeConfig := GasFeeConfig{
				maxBaseFee:                 test.maxBaseFee,
				suggestedPriorityFeeBuffer: big.NewInt(0),
				maxPriorityFeePerGas:       big.NewInt(0),
			}
			destClient = destinationClient{
				readonlyConcurrentSigners: []*readonlyConcurrentSigner{
					(*readonlyConcurrentSigner)(signer),
				},
				logger:             logging.NoLog{},
				avaRPCClient:       mockClient,
				evmChainID:         big.NewInt(5),
				gasFeeConfig:       &gasFeeConfig,
				blockGasLimit:      0,
				txInclusionTimeout: 30 * time.Second,
			}
			warpMsg := &avalancheWarp.Message{}
			toAddress := "0x27aE10273D17Cd7e80de8580A51f476960626e5f"

			gomock.InOrder(
				mockClient.EXPECT().EstimateBaseFee(gomock.Any()).Return(
					big.NewInt(100_000),
					test.estimateBaseFeeErr,
				).Times(test.estimateBaseFeeTimes),
				mockClient.EXPECT().SuggestGasTipCap(gomock.Any()).Return(
					big.NewInt(0),
					test.suggestGasTipCapErr,
				).Times(test.suggestGasTipCapTimes),
				mockClient.EXPECT().SendTransaction(gomock.Any(), gomock.Any()).Return(
					test.sendTransactionErr,
				).Times(test.sendTransactionTimes),
				mockClient.EXPECT().
					TransactionReceipt(gomock.Any(), gomock.Any()).
					Return(
						&types.Receipt{
							Status: types.ReceiptStatusSuccessful,
						},
						nil,
					).Times(test.txReceiptTimes),
			)

			_, err := destClient.SendTx(warpMsg, nil, toAddress, 0, []byte{})
			if test.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// TestDestinationClient_QueryParamsForwarding verifies that query parameters are forwarded correctly
func TestDestinationClient_QueryParamsForwarding(t *testing.T) {
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
			// Track received query params
			receivedParams := make(map[string]string)

			// Create test server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Capture query params
				for key := range tt.queryParams {
					receivedParams[key] = r.URL.Query().Get(key)
				}

				// Return a valid JSON-RPC response for ChainID call
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x1"}`))
			}))
			defer server.Close()

			// Create destination blockchain config with query params
			destinationBlockchain := config.DestinationBlockchain{
				SubnetID:     "2TGBXcnwx5PqiXWiqxAKUaNSqDguXNh1mxnp82jui68hxJSZAx",
				BlockchainID: "S4mMqUXe7vHsGiRAma6bv3CKnyaLssyAxmQ2KvFpX1KEvfFCD",
				RPCEndpoint: basecfg.APIConfig{
					BaseURL:     server.URL,
					QueryParams: tt.queryParams,
				},
				AccountPrivateKeys: []string{"56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"},
			}

			// Create destination client (this will make ChainID call)
			logger := logging.NoLog{}
			_, err := NewDestinationClient(logger, &destinationBlockchain, time.Minute)
			require.NoError(t, err)

			// Verify all query params were received
			for key, expectedValue := range tt.queryParams {
				actualValue := receivedParams[key]
				require.Equal(t, expectedValue, actualValue,
					"Query param %s: expected %s, got %s", key, expectedValue, actualValue)
			}
		})
	}
}

// TestDestinationClient_HTTPHeadersForwarding verifies that HTTP headers are forwarded correctly
func TestDestinationClient_HTTPHeadersForwarding(t *testing.T) {
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
			// Track received headers
			receivedHeaders := make(map[string]string)

			// Create test server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Capture headers
				for key := range tt.httpHeaders {
					receivedHeaders[key] = r.Header.Get(key)
				}

				// Return a valid JSON-RPC response for ChainID call
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x1"}`))
			}))
			defer server.Close()

			// Create destination blockchain config with headers
			destinationBlockchain := config.DestinationBlockchain{
				SubnetID:     "2TGBXcnwx5PqiXWiqxAKUaNSqDguXNh1mxnp82jui68hxJSZAx",
				BlockchainID: "S4mMqUXe7vHsGiRAma6bv3CKnyaLssyAxmQ2KvFpX1KEvfFCD",
				RPCEndpoint: basecfg.APIConfig{
					BaseURL:     server.URL,
					HTTPHeaders: tt.httpHeaders,
				},
				AccountPrivateKeys: []string{"56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"},
			}

			// Create destination client (this will make ChainID call)
			logger := logging.NoLog{}
			_, err := NewDestinationClient(logger, &destinationBlockchain, time.Minute)
			require.NoError(t, err)

			// Verify all headers were received
			for key, expectedValue := range tt.httpHeaders {
				actualValue := receivedHeaders[key]
				require.Equal(t, expectedValue, actualValue,
					"Header %s: expected %s, got %s", key, expectedValue, actualValue)
			}
		})
	}
}

// TestDestinationClient_CombinedQueryParamsAndHeaders verifies both work together
func TestDestinationClient_CombinedQueryParamsAndHeaders(t *testing.T) {
	queryParams := map[string]string{
		"token":   "query-token",
		"api-key": "query-key",
	}
	httpHeaders := map[string]string{
		"Authorization": "Bearer header-token",
		"X-API-Key":     "header-key",
	}

	// Track what the server receives
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
		w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x1"}`))
	}))
	defer server.Close()

	destinationBlockchain := config.DestinationBlockchain{
		SubnetID:     "2TGBXcnwx5PqiXWiqxAKUaNSqDguXNh1mxnp82jui68hxJSZAx",
		BlockchainID: "S4mMqUXe7vHsGiRAma6bv3CKnyaLssyAxmQ2KvFpX1KEvfFCD",
		RPCEndpoint: basecfg.APIConfig{
			BaseURL:     server.URL,
			QueryParams: queryParams,
			HTTPHeaders: httpHeaders,
		},
		AccountPrivateKeys: []string{"56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"},
	}

	_, err := NewDestinationClient(logging.NoLog{}, &destinationBlockchain, time.Minute)
	require.NoError(t, err)

	for key, expectedValue := range queryParams {
		require.Equal(t, expectedValue, receivedParams[key],
			"Query param %s not forwarded correctly", key)
	}

	for key, expectedValue := range httpHeaders {
		require.Equal(t, expectedValue, receivedHeaders[key],
			"Header %s not forwarded correctly", key)
	}
}

// TestDestinationClient_AllRPCCallsForwardQueryParams verifies that ALL RPC calls
// made by the EVM client correctly forward query parameters
func TestDestinationClient_AllRPCCallsForwardQueryParams(t *testing.T) {
	queryParams := map[string]string{
		"token":   "test-token-123",
		"api-key": "secret-key-789",
	}

	requestCount := 0
	receivedParams := make([]map[string]string, 0)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := make(map[string]string)
		for key := range queryParams {
			params[key] = r.URL.Query().Get(key)
		}
		receivedParams = append(receivedParams, params)
		requestCount++

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if requestCount == 1 {
			w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x1"}`))
		} else {
			w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x0"}`))
		}
	}))
	defer server.Close()

	destinationBlockchain := config.DestinationBlockchain{
		SubnetID:     "2TGBXcnwx5PqiXWiqxAKUaNSqDguXNh1mxnp82jui68hxJSZAx",
		BlockchainID: "S4mMqUXe7vHsGiRAma6bv3CKnyaLssyAxmQ2KvFpX1KEvfFCD",
		RPCEndpoint: basecfg.APIConfig{
			BaseURL:     server.URL,
			QueryParams: queryParams,
		},
		AccountPrivateKeys: []string{"56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"},
	}

	client, err := NewDestinationClient(logging.NoLog{}, &destinationBlockchain, time.Minute)
	require.NoError(t, err)

	ctx := t.Context()
	client.avaRPCClient.BlockNumber(ctx)
	ethClient := client.Client()
	ethClient.SuggestGasPrice(ctx)
	ethClient.SuggestGasTipCap(ctx)

	require.Greater(t, len(receivedParams), 0, "No requests were made")
	for i, params := range receivedParams {
		for key, expectedValue := range queryParams {
			require.Equal(t, expectedValue, params[key],
				"Request %d: query param %s not forwarded correctly", i, key)
		}
	}
}

// TestDestinationClient_AllRPCCallsForwardHTTPHeaders verifies that ALL RPC calls
// made by the EVM client correctly forward HTTP headers
func TestDestinationClient_AllRPCCallsForwardHTTPHeaders(t *testing.T) {
	httpHeaders := map[string]string{
		"Authorization": "Bearer test-token",
		"X-API-Key":     "secret-key",
	}

	requestCount := 0
	receivedHeaders := make([]map[string]string, 0)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := make(map[string]string)
		for key := range httpHeaders {
			headers[key] = r.Header.Get(key)
		}
		receivedHeaders = append(receivedHeaders, headers)
		requestCount++

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if requestCount == 1 {
			w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x1"}`))
		} else {
			w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x0"}`))
		}
	}))
	defer server.Close()

	destinationBlockchain := config.DestinationBlockchain{
		SubnetID:     "2TGBXcnwx5PqiXWiqxAKUaNSqDguXNh1mxnp82jui68hxJSZAx",
		BlockchainID: "S4mMqUXe7vHsGiRAma6bv3CKnyaLssyAxmQ2KvFpX1KEvfFCD",
		RPCEndpoint: basecfg.APIConfig{
			BaseURL:     server.URL,
			HTTPHeaders: httpHeaders,
		},
		AccountPrivateKeys: []string{"56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"},
	}

	client, err := NewDestinationClient(logging.NoLog{}, &destinationBlockchain, time.Minute)
	require.NoError(t, err)

	client.avaRPCClient.BlockNumber(t.Context())
	ethClient := client.Client()
	ethClient.SuggestGasPrice(t.Context())
	ethClient.SuggestGasTipCap(t.Context())

	require.Greater(t, len(receivedHeaders), 0, "No requests were made")
	for i, headers := range receivedHeaders {
		for key, expectedValue := range httpHeaders {
			require.Equal(t, expectedValue, headers[key],
				"Request %d: header %s not forwarded correctly", i, key)
		}
	}
}
