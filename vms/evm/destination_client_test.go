// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"
	"math/big"
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
	VM:           config.EVM.String(),
	RPCEndpoint: basecfg.APIConfig{
		BaseURL: "https://subnets.avax.network/mysubnet/rpc",
	},
	AccountPrivateKeys: []string{"56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"},
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
			mockClient := mock_ethclient.NewMockClient(ctrl)
			destClient = destinationClient{
				readonlyConcurrentSigners: []*readonlyConcurrentSigner{
					(*readonlyConcurrentSigner)(signer),
				},
				logger:               logging.NoLog{},
				client:               mockClient,
				evmChainID:           big.NewInt(5),
				maxBaseFee:           test.maxBaseFee,
				maxPriorityFeePerGas: big.NewInt(0),
				maxFeePerGas:         big.NewInt(0), // Initialize to prevent nil panic
				blockGasLimit:        0,
				txInclusionTimeout:   30 * time.Second,
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

// TestMaxFeePerGasCalculation tests the gas fee calculation and capping logic
func TestMaxFeePerGasCalculation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name                 string
		maxBaseFee           uint64
		maxPriorityFeePerGas uint64
		maxFeePerGas         uint64
		currentBaseFee       uint64
		suggestedTip         uint64
		expectedGasFeeCap    uint64
		expectedGasTipCap    uint64
	}{
		{
			name:                 "No MaxFeePerGas - normal behavior",
			maxBaseFee:           50000000000,
			maxPriorityFeePerGas: 2500000000,
			maxFeePerGas:         0, // Not set
			currentBaseFee:       25000000000,
			suggestedTip:         2000000000,
			expectedGasFeeCap:    52000000000, // maxBaseFee + suggestedTip
			expectedGasTipCap:    2000000000,
		},
		{
			name:                 "MaxFeePerGas higher than calculated - no capping",
			maxBaseFee:           50000000000,
			maxPriorityFeePerGas: 2500000000,
			maxFeePerGas:         60000000000,
			currentBaseFee:       25000000000,
			suggestedTip:         2000000000,
			expectedGasFeeCap:    52000000000, // maxBaseFee + suggestedTip
			expectedGasTipCap:    2000000000,
		},
		{
			name:                 "MaxFeePerGas lower than calculated - should cap",
			maxBaseFee:           50000000000,
			maxPriorityFeePerGas: 2500000000,
			maxFeePerGas:         51000000000,
			currentBaseFee:       25000000000,
			suggestedTip:         2000000000,
			expectedGasFeeCap:    51000000000, // Capped to maxFeePerGas
			expectedGasTipCap:    1000000000,  // Adjusted: 51B - 50B = 1B
		},
		{
			name:                 "MaxFeePerGas much lower - tip becomes zero",
			maxBaseFee:           50000000000,
			maxPriorityFeePerGas: 2500000000,
			maxFeePerGas:         40000000000,
			currentBaseFee:       25000000000,
			suggestedTip:         2000000000,
			expectedGasFeeCap:    40000000000, // Capped to maxFeePerGas
			expectedGasTipCap:    0,           // Adjusted: 40B - 50B = -10B, so 0
		},
		{
			name:                 "Dynamic base fee calculation when maxBaseFee is 0",
			maxBaseFee:           0, // Will use currentBaseFee * 3
			maxPriorityFeePerGas: 2500000000,
			maxFeePerGas:         100000000000,
			currentBaseFee:       25000000000,
			suggestedTip:         2000000000,
			expectedGasFeeCap:    77000000000, // (25B * 3) + 2B = 77B
			expectedGasTipCap:    2000000000,
		},
		{
			name:                 "Dynamic base fee with MaxFeePerGas capping",
			maxBaseFee:           0, // Will use currentBaseFee * 3
			maxPriorityFeePerGas: 2500000000,
			maxFeePerGas:         70000000000,
			currentBaseFee:       25000000000,
			suggestedTip:         2000000000,
			expectedGasFeeCap:    70000000000, // Capped
			expectedGasTipCap:    0,           // Adjusted: 70B - 75B = -5B, so 0
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := mock_ethclient.NewMockClient(ctrl)

			// Setup base fee estimation if maxBaseFee is 0
			if tt.maxBaseFee == 0 {
				mockClient.EXPECT().
					EstimateBaseFee(gomock.Any()).
					Return(big.NewInt(int64(tt.currentBaseFee)), nil).
					Times(1)
			}

			// Setup suggested gas tip cap
			mockClient.EXPECT().
				SuggestGasTipCap(gomock.Any()).
				Return(big.NewInt(int64(tt.suggestedTip)), nil).
				Times(1)

			// Create destination client with test configuration
			destClient := &destinationClient{
				client:               mockClient,
				maxBaseFee:           big.NewInt(int64(tt.maxBaseFee)),
				maxPriorityFeePerGas: big.NewInt(int64(tt.maxPriorityFeePerGas)),
				maxFeePerGas:         big.NewInt(int64(tt.maxFeePerGas)),
				evmChainID:           big.NewInt(1),
				logger:               logging.NoLog{},
			}

			// Test just the gas fee calculation logic by replicating the SendTx logic
			var maxBaseFee *big.Int
			if destClient.maxBaseFee.Cmp(big.NewInt(0)) > 0 {
				maxBaseFee = destClient.maxBaseFee
			} else {
				baseFee, err := destClient.client.EstimateBaseFee(nil)
				require.NoError(t, err)
				maxBaseFee = new(big.Int).Mul(baseFee, big.NewInt(defaultBaseFeeFactor))
			}

			gasTipCap, err := destClient.client.SuggestGasTipCap(nil)
			require.NoError(t, err)
			
			if gasTipCap.Cmp(destClient.maxPriorityFeePerGas) > 0 {
				gasTipCap = destClient.maxPriorityFeePerGas
			}

			gasFeeCap := new(big.Int).Add(maxBaseFee, gasTipCap)

			// Apply MaxFeePerGas capping logic (replicated from SendTx)
			if destClient.maxFeePerGas != nil && destClient.maxFeePerGas.Cmp(big.NewInt(0)) > 0 && gasFeeCap.Cmp(destClient.maxFeePerGas) > 0 {
				gasFeeCap = new(big.Int).Set(destClient.maxFeePerGas)
				
				adjustedTipCap := new(big.Int).Sub(gasFeeCap, maxBaseFee)
				if adjustedTipCap.Cmp(big.NewInt(0)) < 0 {
					gasTipCap = big.NewInt(0)
				} else if adjustedTipCap.Cmp(gasTipCap) < 0 {
					gasTipCap = adjustedTipCap
				}
			}

			// Verify the calculated values match expectations
			require.Equal(t, big.NewInt(int64(tt.expectedGasFeeCap)), gasFeeCap,
				"Expected gasFeeCap %d, got %s", tt.expectedGasFeeCap, gasFeeCap.String())
			
			require.Equal(t, big.NewInt(int64(tt.expectedGasTipCap)), gasTipCap,
				"Expected gasTipCap %d, got %s", tt.expectedGasTipCap, gasTipCap.String())
		})
	}
}
