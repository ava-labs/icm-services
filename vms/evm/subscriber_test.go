// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"math/big"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	basecfg "github.com/ava-labs/icm-services/config"
	"github.com/ava-labs/icm-services/relayer/config"
	mock_ethclient "github.com/ava-labs/icm-services/vms/evm/mocks"
	"github.com/ava-labs/libevm/core/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func makeSubscriberWithMockEthClient(t *testing.T) (*subscriber, *mock_ethclient.MockClient) {
	sourceSubnet := config.SourceBlockchain{
		SubnetID:     "2TGBXcnwx5PqiXWiqxAKUaNSqDguXNh1mxnp82jui68hxJSZAx",
		BlockchainID: "S4mMqUXe7vHsGiRAma6bv3CKnyaLssyAxmQ2KvFpX1KEvfFCD",
		VM:           config.EVM.String(),
		RPCEndpoint: basecfg.APIConfig{
			BaseURL: "https://subnets.avax.network/mysubnet/rpc",
		},
	}

	logger := logging.NoLog{}

	mockEthClient := mock_ethclient.NewMockClient(gomock.NewController(t))
	blockchainID, err := ids.FromString(sourceSubnet.BlockchainID)
	require.NoError(t, err)
	subscriber := NewSubscriber(logger, blockchainID, mockEthClient, mockEthClient)

	return subscriber, mockEthClient
}

func TestProcessFromHeight(t *testing.T) {
	testCases := []struct {
		name   string
		latest int64
		input  int64
	}{
		{
			name:   "zero to max blocks",
			latest: 200,
			input:  0,
		},
		{
			name:   "max blocks",
			latest: 1000,
			input:  800,
		},
		{
			name:   "greater than max blocks",
			latest: 1000,
			input:  700,
		},
		{
			name:   "many rounds greater than max blocks",
			latest: 19642,
			input:  751,
		},
		{
			name:   "latest is less than max blocks",
			latest: 96,
			input:  41,
		},
		{
			name:   "invalid starting block number",
			latest: 50,
			input:  51,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			subscriberUnderTest, mockEthClient := makeSubscriberWithMockEthClient(t)

			mockEthClient.
				EXPECT().
				BlockNumber(gomock.Any()).
				Return(uint64(tc.latest), nil).
				Times(1)
			if tc.latest > tc.input {
				expectedFilterLogCalls := (tc.latest-tc.input+1)/MaxBlocksPerRequest + 1
				mockEthClient.EXPECT().FilterLogs(
					gomock.Any(),
					gomock.Any(),
				).Return(
					[]types.Log{},
					nil,
				).Times(int(expectedFilterLogCalls))
			}
			done := make(chan bool, 1)
			subscriberUnderTest.ProcessFromHeight(big.NewInt(tc.input), done)
			result := <-done
			require.True(t, result)

			if tc.latest > tc.input {
				for i := tc.input; i <= tc.latest; i++ {
					block := <-subscriberUnderTest.ICMBlocks()
					require.Equal(t, uint64(i), block.BlockNumber)
					require.Empty(t, block.Messages)
				}
			}
			require.Zero(t, len(subscriberUnderTest.ICMBlocks()))
		})
	}
}
