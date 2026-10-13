// Copyright (C) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peers

import (
	"context"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	avalancheCommon "github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
)

// recordingTimeoutManager is a timer.AdaptiveTimeoutManager that records Put and Remove calls
// instead of scheduling anything, so tests can assert exactly which timeouts a handler
// operation registers or cancels.
type recordingTimeoutManager struct {
	puts    []ids.RequestID
	removes []ids.RequestID
}

func (*recordingTimeoutManager) Dispatch()                      {}
func (*recordingTimeoutManager) Stop()                          {}
func (*recordingTimeoutManager) TimeoutDuration() time.Duration { return time.Second }
func (*recordingTimeoutManager) ObserveLatency(time.Duration)   {}

func (m *recordingTimeoutManager) Put(id ids.RequestID, _ bool, _ func()) {
	m.puts = append(m.puts, id)
}

func (m *recordingTimeoutManager) Remove(id ids.RequestID) {
	m.removes = append(m.removes, id)
}

func newTestHandler(t *testing.T) (*RelayerExternalHandler, *recordingTimeoutManager) {
	handler, err := NewRelayerExternalHandler(logging.NoLog{}, metrics, prometheus.NewRegistry())
	require.NoError(t, err)
	timeoutManager := &recordingTimeoutManager{}
	handler.timeoutManager = timeoutManager
	return handler, timeoutManager
}

// timeoutID is the key under which the aggregator registers a node's response timeout.
func timeoutID(nodeID ids.NodeID, chainID ids.ID, requestID uint32) ids.RequestID {
	return ids.RequestID{
		NodeID:    nodeID,
		ChainID:   chainID,
		RequestID: requestID,
		Op:        byte(message.AppResponseOp),
	}
}

func appResponse(nodeID ids.NodeID, chainID ids.ID, requestID uint32) *message.InboundMessage {
	return message.InboundAppResponse(chainID, requestID, nil, nodeID)
}

func requireNoTrackedRequests(t *testing.T, h *RelayerExternalHandler) {
	h.lock.Lock()
	defer h.lock.Unlock()
	require.Empty(t, h.requestedNodes)
	require.Empty(t, h.responseChans)
	require.Empty(t, h.responsesCount)
}

func requireTrackedRequests(t *testing.T, h *RelayerExternalHandler, n int) {
	h.lock.Lock()
	defer h.lock.Unlock()
	require.Len(t, h.requestedNodes, n)
	require.Len(t, h.responseChans, n)
	require.Len(t, h.responsesCount, n)
}

func requireClosed(t *testing.T, ch chan message.InboundMessage) {
	select {
	case _, ok := <-ch:
		require.False(t, ok, "expected the response channel to be closed, got a message")
	default:
		require.Fail(t, "expected the response channel to be closed")
	}
}

func TestExternalHandlerUnregisterRequestedNodes(t *testing.T) {
	handler, timeoutManager := newTestHandler(t)
	chainID := ids.GenerateTestID()
	const requestID uint32 = 7
	reached1 := ids.GenerateTestNodeID()
	reached2 := ids.GenerateTestNodeID()
	unreached := ids.GenerateTestNodeID()

	responseChan := handler.RegisterRequestID(requestID, set.Of(reached1, reached2, unreached))
	for _, nodeID := range []ids.NodeID{reached1, reached2, unreached} {
		handler.RegisterAppRequest(timeoutID(nodeID, chainID, requestID))
	}
	require.Len(t, timeoutManager.puts, 3)

	// The send did not reach one node: drop it, and only it, along with its timeout.
	handler.UnregisterRequestedNodes(requestID, chainID, []ids.NodeID{unreached})
	require.Equal(t, []ids.RequestID{timeoutID(unreached, chainID, requestID)}, timeoutManager.removes)
	requireTrackedRequests(t, handler, 1)

	// Unregistering it again, or a node that was never requested, is a no-op.
	handler.UnregisterRequestedNodes(requestID, chainID, []ids.NodeID{unreached, ids.GenerateTestNodeID()})
	require.Len(t, timeoutManager.removes, 1)
	requireTrackedRequests(t, handler, 1)

	// The two reached nodes respond; the channel closes after exactly those two responses.
	for _, nodeID := range []ids.NodeID{reached1, reached2} {
		handler.HandleInbound(context.Background(), appResponse(nodeID, chainID, requestID))
	}
	for _, nodeID := range []ids.NodeID{reached1, reached2} {
		response := <-responseChan
		require.Equal(t, nodeID, response.NodeID)
	}
	requireClosed(t, responseChan)
	requireNoTrackedRequests(t, handler)
}

func TestExternalHandlerUnregisterAllRequestedNodesReleasesRequest(t *testing.T) {
	handler, timeoutManager := newTestHandler(t)
	chainID := ids.GenerateTestID()
	const requestID uint32 = 9
	nodeA := ids.GenerateTestNodeID()
	nodeB := ids.GenerateTestNodeID()

	responseChan := handler.RegisterRequestID(requestID, set.Of(nodeA, nodeB))
	handler.RegisterAppRequest(timeoutID(nodeA, chainID, requestID))
	handler.RegisterAppRequest(timeoutID(nodeB, chainID, requestID))

	// Nothing was reached: the request must be released immediately rather than waiting on
	// responses that can never come.
	handler.UnregisterRequestedNodes(requestID, chainID, []ids.NodeID{nodeA, nodeB})
	require.Len(t, timeoutManager.removes, 2)
	requireClosed(t, responseChan)
	requireNoTrackedRequests(t, handler)
}

func TestExternalHandlerIgnoresUnexpectedAndDuplicateResponses(t *testing.T) {
	handler, timeoutManager := newTestHandler(t)
	ctx := context.Background()
	chainID := ids.GenerateTestID()
	const requestID uint32 = 11
	nodeA := ids.GenerateTestNodeID()
	nodeB := ids.GenerateTestNodeID()
	stranger := ids.GenerateTestNodeID()

	responseChan := handler.RegisterRequestID(requestID, set.Of(nodeA, nodeB))
	handler.RegisterAppRequest(timeoutID(nodeA, chainID, requestID))
	handler.RegisterAppRequest(timeoutID(nodeB, chainID, requestID))

	// A response from a node that was never queried is dropped and must not cancel any timeout.
	handler.HandleInbound(ctx, appResponse(stranger, chainID, requestID))
	require.Empty(t, timeoutManager.removes)
	require.Empty(t, responseChan)

	// A response for a request ID that was never registered is dropped the same way.
	handler.HandleInbound(ctx, appResponse(nodeA, chainID, requestID+2))
	require.Empty(t, timeoutManager.removes)
	require.Empty(t, responseChan)

	// The first response from a queried node is forwarded and cancels only that node's timeout.
	handler.HandleInbound(ctx, appResponse(nodeA, chainID, requestID))
	require.Equal(t, []ids.RequestID{timeoutID(nodeA, chainID, requestID)}, timeoutManager.removes)
	require.Equal(t, nodeA, (<-responseChan).NodeID)

	// A duplicate from the same node is dropped: not forwarded, not counted, no further Remove.
	handler.HandleInbound(ctx, appResponse(nodeA, chainID, requestID))
	require.Len(t, timeoutManager.removes, 1)
	require.Empty(t, responseChan)
	requireTrackedRequests(t, handler, 1)

	// The remaining node's response completes the request.
	handler.HandleInbound(ctx, appResponse(nodeB, chainID, requestID))
	require.Equal(t, nodeB, (<-responseChan).NodeID)
	requireClosed(t, responseChan)
	requireNoTrackedRequests(t, handler)
}

func TestExternalHandlerAppErrorCancelsTimeoutOfExpectedResponse(t *testing.T) {
	handler, timeoutManager := newTestHandler(t)
	ctx := context.Background()
	chainID := ids.GenerateTestID()
	const requestID uint32 = 13
	nodeA := ids.GenerateTestNodeID()
	nodeB := ids.GenerateTestNodeID()

	responseChan := handler.RegisterRequestID(requestID, set.Of(nodeA, nodeB))
	handler.RegisterAppRequest(timeoutID(nodeA, chainID, requestID))
	handler.RegisterAppRequest(timeoutID(nodeB, chainID, requestID))

	// Timeouts are keyed by the expected AppResponse op. An AppError reply must cancel that
	// same timeout, otherwise it fires later and the node is counted a second time. It must
	// cancel only the erroring node's timeout: the other node's response is still outstanding,
	// so the request must remain tracked and its channel open.
	handler.HandleInbound(ctx, message.InboundAppError(
		nodeA,
		chainID,
		requestID,
		avalancheCommon.ErrTimeout.Code,
		avalancheCommon.ErrTimeout.Message,
	))
	require.Equal(t, []ids.RequestID{timeoutID(nodeA, chainID, requestID)}, timeoutManager.removes)
	response := <-responseChan
	require.Equal(t, nodeA, response.NodeID)
	require.Equal(t, message.AppErrorOp, response.Op)
	require.Empty(t, responseChan)
	requireTrackedRequests(t, handler, 1)

	// The other node's response is still expected; it cancels its own timeout and completes
	// the request.
	handler.HandleInbound(ctx, appResponse(nodeB, chainID, requestID))
	require.Equal(t, []ids.RequestID{
		timeoutID(nodeA, chainID, requestID),
		timeoutID(nodeB, chainID, requestID),
	}, timeoutManager.removes)
	response = <-responseChan
	require.Equal(t, nodeB, response.NodeID)
	require.Equal(t, message.AppResponseOp, response.Op)
	requireClosed(t, responseChan)
	requireNoTrackedRequests(t, handler)
}
