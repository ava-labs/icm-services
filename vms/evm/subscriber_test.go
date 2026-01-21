// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	basecfg "github.com/ava-labs/icm-services/config"
	"github.com/ava-labs/icm-services/relayer/config"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/core/types"
	"github.com/stretchr/testify/require"
)

var _ SubscriberRPCClient = (*subscriberClientStub)(nil)
var _ SubscriberWSClient = (*subscriberClientStub)(nil)

type subscriberClientStub struct {
	blockNumber       uint64
	numFilterLogCalls int
}

func (c *subscriberClientStub) BlockNumber(ctx context.Context) (uint64, error) {
	return c.blockNumber, nil
}

func (c *subscriberClientStub) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	c.numFilterLogCalls++
	return []types.Log{}, nil
}

func (c *subscriberClientStub) SubscribeNewHead(
	ctx context.Context,
	ch chan<- *types.Header,
) (ethereum.Subscription, error) {
	return nil, nil
}

func makeSubscriberWithMockEthClient(t *testing.T, errChan chan error) (*Subscriber, *subscriberClientStub) {
	sourceSubnet := config.SourceBlockchain{
		SubnetID:     "2TGBXcnwx5PqiXWiqxAKUaNSqDguXNh1mxnp82jui68hxJSZAx",
		BlockchainID: "S4mMqUXe7vHsGiRAma6bv3CKnyaLssyAxmQ2KvFpX1KEvfFCD",
		RPCEndpoint: basecfg.APIConfig{
			BaseURL: "https://subnets.avax.network/mysubnet/rpc",
		},
	}

	stubRPCClient := &subscriberClientStub{}
	blockchainID, err := ids.FromString(sourceSubnet.BlockchainID)
	require.NoError(t, err)
	subscriber := NewSubscriber(logging.NoLog{}, blockchainID, stubRPCClient, stubRPCClient, errChan)

	return subscriber, stubRPCClient
}

func TestProcessFromHeight(t *testing.T) {
	testCases := []struct {
		name   string
		latest uint64
		input  uint64
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
			errChan := make(chan error, 1)
			subscriberUnderTest, stubRPCClient := makeSubscriberWithMockEthClient(t, errChan)

			stubRPCClient.blockNumber = tc.latest
			var expectedFilterLogCalls uint64
			if tc.latest > tc.input {
				expectedFilterLogCalls = (tc.latest-tc.input+1)/MaxBlocksPerRequest + 1
			}
			subscriberUnderTest.ProcessFromHeight(tc.input, tc.latest)
			require.Empty(t, errChan)

			if tc.latest > tc.input {
				for i := tc.input; i <= tc.latest; i++ {
					block := <-subscriberUnderTest.ICMBlocks()
					require.Equal(t, i, block.BlockNumber)
					require.Empty(t, block.Messages)
				}
			}
			require.Zero(t, len(subscriberUnderTest.ICMBlocks()))
			require.EqualValues(t, expectedFilterLogCalls, stubRPCClient.numFilterLogCalls)
		})
	}
}
