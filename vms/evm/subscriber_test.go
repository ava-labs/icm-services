// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"math/big"
	"testing"
	"time"

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
	// logs served by FilterLogs, filtered by the query's block range or block hash. Block hashes
	// are those produced by HeadByNumber, i.e. the block number as a hash.
	logs []types.Log
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
	matching := []types.Log{}
	for _, log := range c.logs {
		if q.BlockHash != nil {
			if log.BlockHash == *q.BlockHash {
				matching = append(matching, log)
			}
			continue
		}
		if log.BlockNumber >= q.FromBlock.Uint64() && log.BlockNumber <= q.ToBlock.Uint64() {
			matching = append(matching, log)
		}
	}
	return matching, nil
}

func (c *subscriberClientStub) SubscribeFilterLogs(
	ctx context.Context,
	q ethereum.FilterQuery,
	ch chan<- types.Log,
) (ethereum.Subscription, error) {
	c.numSubscribeFilterLogsCalls++
	return nil, nil
}

// stubLog returns a log in block [blockNumber] with the block hash produced by HeadByNumber.
func stubLog(blockNumber uint64, index uint) types.Log {
	return types.Log{
		BlockNumber: blockNumber,
		BlockHash:   common.BigToHash(new(big.Int).SetUint64(blockNumber)),
		Index:       index,
	}
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

			// The reported ranges must tile [input, latest] contiguously and in order.
			nextBlock := tc.input
			for nextBlock <= tc.latest {
				block := <-subscriberUnderTest.ICMBlocks()
				require.Equal(t, nextBlock, block.FromBlock)
				require.GreaterOrEqual(t, block.ToBlock, block.FromBlock)
				require.LessOrEqual(t, block.ToBlock, tc.latest)
				require.Empty(t, block.Logs)
				require.True(t, block.IsCatchup)
				nextBlock = block.ToBlock + 1
			}
			require.Zero(t, len(subscriberUnderTest.ICMBlocks()))
			require.EqualValues(t, expectedFilterLogCalls, stubRPCClient.numFilterLogCalls)
			require.EqualValues(t, expectedHeadCalls, stubRPCClient.numHeadByNumberCalls)
		})
	}
}

// Blocks with logs are reported individually, folding in the empty blocks before them; the empty
// blocks after the last block with logs are reported as one range.
func TestProcessBlockRangeGroupsLogsByBlock(t *testing.T) {
	errChan := make(chan error, 1)
	subscriberUnderTest, stubRPCClient := makeSubscriberWithMockEthClient(t, errChan)
	stubRPCClient.logs = []types.Log{
		stubLog(103, 0),
		stubLog(103, 1),
		stubLog(107, 0),
		// Outside the processed range, must not be reported.
		stubLog(150, 0),
	}

	require.NoError(t, subscriberUnderTest.processBlockRange(100, 120))
	require.Equal(t, 1, stubRPCClient.numFilterLogCalls)

	expected := []*ICMBlockInfo{
		{FromBlock: 100, ToBlock: 103, Logs: []types.Log{stubLog(103, 0), stubLog(103, 1)}, IsCatchup: true},
		{FromBlock: 104, ToBlock: 107, Logs: []types.Log{stubLog(107, 0)}, IsCatchup: true},
		{FromBlock: 108, ToBlock: 120, IsCatchup: true},
	}
	for _, want := range expected {
		got := <-subscriberUnderTest.ICMBlocks()
		require.Equal(t, want.FromBlock, got.FromBlock)
		require.Equal(t, want.ToBlock, got.ToBlock)
		require.Equal(t, want.Logs, got.Logs)
		require.True(t, got.IsCatchup)
	}
	require.Zero(t, len(subscriberUnderTest.ICMBlocks()))
}

// A range whose last block contains logs has no trailing empty range.
func TestProcessBlockRangeLogsInLastBlock(t *testing.T) {
	errChan := make(chan error, 1)
	subscriberUnderTest, stubRPCClient := makeSubscriberWithMockEthClient(t, errChan)
	stubRPCClient.logs = []types.Log{stubLog(120, 0)}

	require.NoError(t, subscriberUnderTest.processBlockRange(100, 120))

	got := <-subscriberUnderTest.ICMBlocks()
	require.Equal(t, uint64(100), got.FromBlock)
	require.Equal(t, uint64(120), got.ToBlock)
	require.Equal(t, []types.Log{stubLog(120, 0)}, got.Logs)
	require.Zero(t, len(subscriberUnderTest.ICMBlocks()))
}

// Each block that the subscription notifies about is reported once, with all of its matching
// logs fetched by block hash on the first notification.
func TestBlocksInfoFromLogs(t *testing.T) {
	errChan := make(chan error, 1)
	subscriberUnderTest, stubRPCClient := makeSubscriberWithMockEthClient(t, errChan)
	stubRPCClient.logs = []types.Log{
		stubLog(10, 0),
		stubLog(10, 1),
		stubLog(10, 2),
		stubLog(15, 0),
	}

	// The subscription delivers the block's logs one at a time.
	for _, log := range stubRPCClient.logs {
		subscriberUnderTest.logs <- log
	}
	// A removed log is ignored altogether.
	removed := stubLog(16, 0)
	removed.Removed = true
	subscriberUnderTest.logs <- removed

	block := <-subscriberUnderTest.ICMBlocks()
	require.Equal(t, uint64(10), block.FromBlock)
	require.Equal(t, uint64(10), block.ToBlock)
	require.Equal(t, stubRPCClient.logs[:3], block.Logs)
	require.False(t, block.IsCatchup)

	block = <-subscriberUnderTest.ICMBlocks()
	require.Equal(t, uint64(15), block.FromBlock)
	require.Equal(t, uint64(15), block.ToBlock)
	require.Equal(t, stubRPCClient.logs[3:], block.Logs)
	require.False(t, block.IsCatchup)

	select {
	case block := <-subscriberUnderTest.ICMBlocks():
		require.Fail(t, "unexpected block", "block %d", block.ToBlock)
	case <-time.After(50 * time.Millisecond):
	}
	require.Empty(t, errChan)
	// One fetch per block, regardless of the number of notifications for it.
	require.Equal(t, 2, stubRPCClient.numFilterLogCalls)
}
