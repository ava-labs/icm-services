// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"
	"math/big"
	"sync"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/awm-relayer/config"
	mock_ethclient "github.com/ava-labs/awm-relayer/vms/evm/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var destinationSubnet = config.DestinationSubnet{
	SubnetID:          "2TGBXcnwx5PqiXWiqxAKUaNSqDguXNh1mxnp82jui68hxJSZAx",
	ChainID:           "S4mMqUXe7vHsGiRAma6bv3CKnyaLssyAxmQ2KvFpX1KEvfFCD",
	VM:                config.EVM.String(),
	APINodeHost:       "127.0.0.1",
	APINodePort:       9650,
	EncryptConnection: false,
	RPCEndpoint:       "https://subnets.avax.network/mysubnet/rpc",
	AccountPrivateKey: "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027",
}

func TestSendTx(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := mock_ethclient.NewMockClient(ctrl)
	pk, eoa, err := destinationSubnet.GetRelayerAccountInfo()
	require.NoError(t, err)

	destinationClient := &destinationClient{
		lock:   &sync.Mutex{},
		logger: logging.NoLog{},
		client: mockClient,
		pk:     pk,
		eoa:    eoa,
	}

	testError := fmt.Errorf("call errored")
	testCases := []struct {
		name                  string
		chainIDErr            error
		chainIDTimes          int
		estimateBaseFeeErr    error
		estimateBaseFeeTimes  int
		suggestGasTipCapErr   error
		suggestGasTipCapTimes int
		sendTransactionErr    error
		sendTransactionTimes  int
		expectError           bool
	}{
		{
			name:                  "valid",
			chainIDTimes:          1,
			estimateBaseFeeTimes:  1,
			suggestGasTipCapTimes: 1,
			sendTransactionTimes:  1,
		},
		{
			name:         "invalid chainID",
			chainIDErr:   testError,
			chainIDTimes: 1,
			expectError:  true,
		},
		{
			name:                 "invalid estimateBaseFee",
			chainIDTimes:         1,
			estimateBaseFeeErr:   testError,
			estimateBaseFeeTimes: 1,
			expectError:          true,
		},
		{
			name:                  "invalid suggestGasTipCap",
			chainIDTimes:          1,
			estimateBaseFeeTimes:  1,
			suggestGasTipCapErr:   testError,
			suggestGasTipCapTimes: 1,
			expectError:           true,
		},
		{
			name:                  "invalid sendTransaction",
			chainIDTimes:          1,
			estimateBaseFeeTimes:  1,
			suggestGasTipCapTimes: 1,
			sendTransactionErr:    testError,
			sendTransactionTimes:  1,
			expectError:           true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			warpMsg := &avalancheWarp.Message{}
			toAddress := "0x27aE10273D17Cd7e80de8580A51f476960626e5f"

			gomock.InOrder(
				mockClient.EXPECT().ChainID(gomock.Any()).Return(new(big.Int), test.chainIDErr).Times(test.chainIDTimes),
				mockClient.EXPECT().EstimateBaseFee(gomock.Any()).Return(new(big.Int), test.estimateBaseFeeErr).Times(test.estimateBaseFeeTimes),
				mockClient.EXPECT().SuggestGasTipCap(gomock.Any()).Return(new(big.Int), test.suggestGasTipCapErr).Times(test.suggestGasTipCapTimes),
				mockClient.EXPECT().SendTransaction(gomock.Any(), gomock.Any()).Return(test.sendTransactionErr).Times(test.sendTransactionTimes),
			)

			err := destinationClient.SendTx(warpMsg, toAddress, 0, []byte{})
			if test.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestIsAllowedRelayer(t *testing.T) {
	_, eoa, err := destinationSubnet.GetRelayerAccountInfo()
	require.NoError(t, err)

	tdc := &destinationClient{
		logger: logging.NoLog{},
		eoa:    eoa,
	}

	invalidRelayer := common.Address{1}
	testCases := []struct {
		name            string
		allowedRelayers []common.Address
		expected        bool
	}{
		{
			name:            "empty allowed relayers",
			allowedRelayers: []common.Address{},
			expected:        true,
		},
		{
			name:            "invalid relayer",
			allowedRelayers: []common.Address{invalidRelayer},
			expected:        false,
		},
		{
			name:            "valid relayer",
			allowedRelayers: []common.Address{eoa},
			expected:        true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			output := tdc.isAllowedRelayer(test.allowedRelayers)
			require.Equal(t, test.expected, output)
		})
	}
}

func TestIsDestination(t *testing.T) {
	validChainID := ids.ID{1, 2, 3, 4, 5, 6, 7, 8}
	invalidChainID := ids.Empty
	logger := logging.NoLog{}
	tdc := &destinationClient{
		logger:             logger,
		destinationChainID: validChainID,
	}

	testCases := []struct {
		name     string
		chainID  ids.ID
		expected bool
	}{
		{
			name:     "valid",
			chainID:  validChainID,
			expected: true,
		},
		{
			name:     "invalid",
			chainID:  invalidChainID,
			expected: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			output := tdc.isDestination(test.chainID)
			require.Equal(t, test.expected, output)
		})
	}
}
