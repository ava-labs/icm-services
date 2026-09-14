// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"math/big"
	"sync/atomic"
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
	numFilterLogCalls           atomic.Int64
	numSubscribeFilterLogsCalls atomic.Int64
	numBlockHeaderByNumberCalls atomic.Int64
	// logs served by FilterLogs, filtered by the query's block range or block hash. Block hashes
	// are those produced by BlockHeaderByNumber, i.e. the block number as a hash.
	logs []types.Log
	// blockNumberDelay is how long BlockNumber takes to answer.
	blockNumberDelay time.Duration
	// onSubscribe, if set, is called by SubscribeFilterLogs with the subscription's channel, so
	// that a test can deliver logs as soon as the subscription opens.
	onSubscribe func(ch chan<- types.Log)
}

func (c *subscriberClientStub) BlockNumber(ctx context.Context) (uint64, error) {
	time.Sleep(c.blockNumberDelay)
	return c.blockNumber, nil
}

func (c *subscriberClientStub) BlockHeaderByNumber(ctx context.Context, number *big.Int) (*BlockHeader, error) {
	c.numBlockHeaderByNumberCalls.Add(1)
	return &BlockHeader{
		Hash:   common.BigToHash(number),
		Number: (*hexutil.Big)(new(big.Int).Set(number)),
	}, nil
}

func (c *subscriberClientStub) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	c.numFilterLogCalls.Add(1)
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
	c.numSubscribeFilterLogsCalls.Add(1)
	if c.onSubscribe != nil {
		c.onSubscribe(ch)
	}
	return nil, nil
}

// stubLog returns a log in block [blockNumber] with the block hash produced by BlockHeaderByNumber.
func stubLog(blockNumber uint64, index uint) types.Log {
	return types.Log{
		BlockNumber: blockNumber,
		BlockHash:   common.BigToHash(new(big.Int).SetUint64(blockNumber)),
		Index:       index,
	}
}

// testStartingHeight is the first block the subscribers under test account for.
const testStartingHeight = 1

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
		stubRPCClient,
		stubRPCClient,
		errChan,
		EventFilter{Topics: [][]common.Hash{{warp.WarpABI.Events["SendWarpMessage"].ID}}},
		testStartingHeight,
	)

	return subscriber, stubRPCClient
}

// A starting height of 0 must not wrap the highest dispatched block around to
// MaxUint64, which would make every live block look already dispatched.
func TestNewSubscriberStartingHeightZero(t *testing.T) {
	subscriber := NewSubscriber(
		logging.NoLog{},
		ids.Empty,
		&subscriberClientStub{},
		&subscriberClientStub{},
		make(chan error, 1),
		EventFilter{},
		0,
	)
	require.Equal(t, uint64(0), subscriber.highestDispatchedBlock)
	// Genesis counts as dispatched, so the first live block's range starts at 1.
	require.Equal(t, uint64(1), subscriber.dispatchLiveBlock(5))
	require.Equal(t, uint64(5), subscriber.highestDispatchedBlock)
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

			// The last strictTailBlocks blocks are processed one by one, each
			// with a BlockHeaderByNumber call and a FilterLogs call by hash;
			// everything older is served by chunked range queries.
			var expectedFilterLogCalls, expectedHeaderCalls uint64
			if tc.latest >= tc.input {
				span := tc.latest - tc.input + 1
				if span > strictTailBlocks {
					expectedHeaderCalls = strictTailBlocks
					rangeSpan := span - strictTailBlocks
					expectedFilterLogCalls = (rangeSpan + MaxBlocksPerRequest - 1) / MaxBlocksPerRequest
				} else {
					expectedHeaderCalls = span
				}
				expectedFilterLogCalls += expectedHeaderCalls
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
				nextBlock = block.ToBlock + 1
			}
			require.Zero(t, len(subscriberUnderTest.ICMBlocks()))
			require.EqualValues(t, expectedFilterLogCalls, stubRPCClient.numFilterLogCalls.Load())
			require.EqualValues(t, expectedHeaderCalls, stubRPCClient.numBlockHeaderByNumberCalls.Load())
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
	require.EqualValues(t, 1, stubRPCClient.numFilterLogCalls.Load())

	expected := []*ICMBlockInfo{
		{FromBlock: 100, ToBlock: 103, Logs: []types.Log{stubLog(103, 0), stubLog(103, 1)}},
		{FromBlock: 104, ToBlock: 107, Logs: []types.Log{stubLog(107, 0)}},
		{FromBlock: 108, ToBlock: 120},
	}
	for _, want := range expected {
		got := <-subscriberUnderTest.ICMBlocks()
		require.Equal(t, want.FromBlock, got.FromBlock)
		require.Equal(t, want.ToBlock, got.ToBlock)
		require.Equal(t, want.Logs, got.Logs)
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
// logs fetched by block hash on the first notification. Its range starts right after the highest
// block dispatched so far, so that the empty blocks in between are accounted for.
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
	require.Equal(t, uint64(testStartingHeight), block.FromBlock)
	require.Equal(t, uint64(10), block.ToBlock)
	require.Equal(t, stubRPCClient.logs[:3], block.Logs)

	block = <-subscriberUnderTest.ICMBlocks()
	require.Equal(t, uint64(11), block.FromBlock)
	require.Equal(t, uint64(15), block.ToBlock)
	require.Equal(t, stubRPCClient.logs[3:], block.Logs)

	select {
	case block := <-subscriberUnderTest.ICMBlocks():
		require.Fail(t, "unexpected block", "block %d", block.ToBlock)
	case <-time.After(50 * time.Millisecond):
	}
	require.Empty(t, errChan)
	// One fetch per block, regardless of the number of notifications for it.
	require.EqualValues(t, 2, stubRPCClient.numFilterLogCalls.Load())
}

// Subscribing dispatches catch-up of every block up to the subscribed node's head, and blocks
// reported by the subscription afterwards continue from the head. Blocks the subscription reports
// that catch-up already covers are reported as is.
func TestSubscribeDispatchesCatchup(t *testing.T) {
	errChan := make(chan error, 1)
	subscriberUnderTest, stubRPCClient := makeSubscriberWithMockEthClient(t, errChan)
	const head = 30
	stubRPCClient.blockNumber = head
	// Block 5 is served by catch-up's range query, block 25 by its strict tail.
	stubRPCClient.logs = []types.Log{stubLog(5, 0), stubLog(25, 0), stubLog(40, 0)}

	require.NoError(t, subscriberUnderTest.Subscribe(time.Second))
	require.EqualValues(t, 1, stubRPCClient.numSubscribeFilterLogsCalls.Load())

	// Catch-up tiles [testStartingHeight, head], reporting the logs of blocks 5 and 25 on the way.
	nextBlock := uint64(testStartingHeight)
	var caughtUpLogs []types.Log
	for nextBlock <= head {
		block := <-subscriberUnderTest.ICMBlocks()
		require.Equal(t, nextBlock, block.FromBlock)
		require.LessOrEqual(t, block.ToBlock, uint64(head))
		caughtUpLogs = append(caughtUpLogs, block.Logs...)
		nextBlock = block.ToBlock + 1
	}
	require.Equal(t, []types.Log{stubLog(5, 0), stubLog(25, 0)}, caughtUpLogs)
	require.Empty(t, errChan)

	// A block the subscription reports that catch-up already covered is reported as is.
	subscriberUnderTest.logs <- stubLog(25, 0)
	block := <-subscriberUnderTest.ICMBlocks()
	require.Equal(t, uint64(25), block.FromBlock)
	require.Equal(t, uint64(25), block.ToBlock)

	// The next block above the head is reported from right after the head.
	subscriberUnderTest.logs <- stubLog(40, 0)
	block = <-subscriberUnderTest.ICMBlocks()
	require.Equal(t, uint64(head+1), block.FromBlock)
	require.Equal(t, uint64(40), block.ToBlock)
	require.Equal(t, []types.Log{stubLog(40, 0)}, block.Logs)

	// Resubscribing to a node that is behind has nothing to catch up on, and the next block
	// continues from the highest block dispatched so far.
	stubRPCClient.blockNumber = 35
	require.NoError(t, subscriberUnderTest.Subscribe(time.Second))
	subscriberUnderTest.logs <- stubLog(42, 0)
	block = <-subscriberUnderTest.ICMBlocks()
	require.Equal(t, uint64(41), block.FromBlock)
	require.Equal(t, uint64(42), block.ToBlock)

	select {
	case block := <-subscriberUnderTest.ICMBlocks():
		require.Fail(t, "unexpected block", "block %d", block.ToBlock)
	case <-time.After(50 * time.Millisecond):
	}
	require.Empty(t, errChan)
}

// drainCatchup reads the catch-up ranges tiling [fromBlock, toBlock] off the blocks channel,
// returning the logs they carried.
func drainCatchup(t *testing.T, subscriber *Subscriber, fromBlock, toBlock uint64) []types.Log {
	var logs []types.Log
	nextBlock := fromBlock
	for nextBlock <= toBlock {
		block := <-subscriber.ICMBlocks()
		require.Equal(t, nextBlock, block.FromBlock)
		require.LessOrEqual(t, block.ToBlock, toBlock)
		logs = append(logs, block.Logs...)
		nextBlock = block.ToBlock + 1
	}
	return logs
}

// ProcessIdleBlocks catches up to the chain head only if the subscription has not reported a new
// block since the previous call.
// A log the subscription delivers while Subscribe is still reading the chain head must wait until
// the catch-up range is fixed. Dispatched first, it would fold the whole catch-up range into its
// own as if those blocks held no matching logs, and catch-up would then skip them.
func TestSubscribeLiveLogDuringCatchupSetup(t *testing.T) {
	errChan := make(chan error, 1)
	subscriberUnderTest, stubRPCClient := makeSubscriberWithMockEthClient(t, errChan)
	const head = 30
	const liveBlock = head + 1
	stubRPCClient.blockNumber = head
	stubRPCClient.logs = []types.Log{stubLog(5, 0), stubLog(liveBlock, 0)}
	// Deliver the live block the moment the subscription opens, and make the head read slow
	// enough that the live path would otherwise reach the dispatch lock first.
	stubRPCClient.onSubscribe = func(ch chan<- types.Log) { ch <- stubLog(liveBlock, 0) }
	stubRPCClient.blockNumberDelay = 100 * time.Millisecond

	require.NoError(t, subscriberUnderTest.Subscribe(time.Second))

	// Catch-up tiles [testStartingHeight, head] and the live block is reported on its own;
	// the two are dispatched by different goroutines, so their order is not fixed.
	var logs []types.Log
	covered := uint64(0)
	for covered < liveBlock-testStartingHeight+1 {
		block := <-subscriberUnderTest.ICMBlocks()
		if block.ToBlock == liveBlock {
			require.Equal(t, uint64(liveBlock), block.FromBlock, "live block folded over catch-up range")
		} else {
			require.LessOrEqual(t, block.ToBlock, uint64(head))
		}
		covered += block.ToBlock - block.FromBlock + 1
		logs = append(logs, block.Logs...)
	}
	require.ElementsMatch(t, []types.Log{stubLog(5, 0), stubLog(liveBlock, 0)}, logs)
	require.Empty(t, errChan)
}

func TestProcessIdleBlocks(t *testing.T) {
	errChan := make(chan error, 1)
	subscriberUnderTest, stubRPCClient := makeSubscriberWithMockEthClient(t, errChan)
	stubRPCClient.blockNumber = 30
	require.NoError(t, subscriberUnderTest.Subscribe(time.Second))
	require.Empty(t, drainCatchup(t, subscriberUnderTest, testStartingHeight, 30))

	// Nothing reported since subscribing: catch up to the new head, finding a log on the way.
	stubRPCClient.blockNumber = 60
	stubRPCClient.logs = []types.Log{stubLog(45, 0)}
	require.NoError(t, subscriberUnderTest.ProcessIdleBlocks())
	require.Equal(t, []types.Log{stubLog(45, 0)}, drainCatchup(t, subscriberUnderTest, 31, 60))

	// A block reported by the subscription continues from the caught-up head...
	stubRPCClient.logs = append(stubRPCClient.logs, stubLog(70, 0))
	subscriberUnderTest.logs <- stubLog(70, 0)
	block := <-subscriberUnderTest.ICMBlocks()
	require.Equal(t, uint64(61), block.FromBlock)
	require.Equal(t, uint64(70), block.ToBlock)

	// ...and suppresses the next idle check, since the checkpoint advanced with it.
	stubRPCClient.blockNumber = 80
	require.NoError(t, subscriberUnderTest.ProcessIdleBlocks())
	select {
	case block := <-subscriberUnderTest.ICMBlocks():
		require.Fail(t, "unexpected block", "block %d", block.ToBlock)
	case <-time.After(50 * time.Millisecond):
	}

	// The check after that is idle again and catches up to the head.
	require.NoError(t, subscriberUnderTest.ProcessIdleBlocks())
	require.Empty(t, drainCatchup(t, subscriberUnderTest, 71, 80))

	// A head that has not moved has nothing to catch up on.
	require.NoError(t, subscriberUnderTest.ProcessIdleBlocks())
	select {
	case block := <-subscriberUnderTest.ICMBlocks():
		require.Fail(t, "unexpected block", "block %d", block.ToBlock)
	case <-time.After(50 * time.Millisecond):
	}
	require.Empty(t, errChan)
}
