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
