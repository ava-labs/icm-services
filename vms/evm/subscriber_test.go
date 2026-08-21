// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"math/big"
	"testing"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/contracts/warp"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	basecfg "github.com/ava-labs/icm-services/config"
	"github.com/ava-labs/icm-services/relayer/config"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/core/types"
	"github.com/stretchr/testify/require"
)

var _ SubscriberRPCClient = (*subscriberClientStub)(nil)
var _ SubscriberWSClient = (*subscriberClientStub)(nil)

type subscriberClientStub struct {
	blockNumber                 uint64
	numFilterLogCalls           int
	numSubscribeFilterLogsCalls int
	numHeadByNumberCalls        int
}

func (c *subscriberClientStub) BlockNumber(ctx context.Context) (uint64, error) {
	return c.blockNumber, nil
}

func (c *subscriberClientStub) HeadByNumber(ctx context.Context, number *big.Int) (*BlockHead, error) {
	c.numHeadByNumberCalls++
	return &BlockHead{
		Hash:   common.BigToHash(number),
		Number: (*hexutil.Big)(new(big.Int).Set(number)),
	}, nil
}

func (c *subscriberClientStub) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	c.numFilterLogCalls++
	return []types.Log{}, nil
}

func (c *subscriberClientStub) SubscribeFilterLogs(
	ctx context.Context,
	q ethereum.FilterQuery,
	ch chan<- types.Log,
) (ethereum.Subscription, error) {
	c.numSubscribeFilterLogsCalls++
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
	subscriber := NewSubscriber(
		logging.NoLog{},
		blockchainID,
		false,
		stubRPCClient,
		stubRPCClient,
		errChan,
		EventFilter{Topics: [][]common.Hash{{warp.WarpABI.Events["SendWarpMessage"].ID}}},
	)

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
			name:   "span smaller than strict tail",
			latest: 50,
			input:  45,
		},
		{
			name:   "span exactly the strict tail",
			latest: 50,
			input:  41,
		},
		{
			name:   "span one greater than strict tail",
			latest: 50,
			input:  40,
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

			// The last strictTailBlocks blocks are processed one by one via
			// HeadByNumber (no FilterLogs here: the stub's bloom is empty and
			// the test chain is not the primary network, so the bloom gate
			// skips the log fetch); everything older is served by chunked
			// range queries.
			var expectedFilterLogCalls, expectedHeadCalls uint64
			if tc.latest >= tc.input {
				span := tc.latest - tc.input + 1
				if span > strictTailBlocks {
					expectedHeadCalls = strictTailBlocks
					rangeSpan := span - strictTailBlocks
					expectedFilterLogCalls = (rangeSpan + MaxBlocksPerRequest - 1) / MaxBlocksPerRequest
				} else {
					expectedHeadCalls = span
				}
			}

			subscriberUnderTest.ProcessFromHeight(tc.input, tc.latest)
			require.Empty(t, errChan)

			if tc.latest >= tc.input {
				for i := tc.input; i <= tc.latest; i++ {
					block := <-subscriberUnderTest.ICMBlocks()
					require.Equal(t, i, block.BlockNumber)
					require.Empty(t, block.Logs)
					require.True(t, block.IsCatchup)
				}
			}
			require.Zero(t, len(subscriberUnderTest.ICMBlocks()))
			require.EqualValues(t, expectedFilterLogCalls, stubRPCClient.numFilterLogCalls)
			require.EqualValues(t, expectedHeadCalls, stubRPCClient.numHeadByNumberCalls)
		})
	}
}
