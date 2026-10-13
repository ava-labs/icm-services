// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peers

import (
	"context"
	"sync"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	avalancheCommon "github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/snow/networking/router"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/utils/timer"
	"github.com/ava-labs/avalanchego/version"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

var _ router.ExternalHandler = &RelayerExternalHandler{}

// Note: all of the external handler's methods are called on peer goroutines. It
// is possible for multiple concurrent calls to happen with different NodeIDs.
// However, a given NodeID will only be performing one call at a time.
type RelayerExternalHandler struct {
	log            logging.Logger
	lock           *sync.Mutex
	requestedNodes map[uint32]*set.Set[ids.NodeID]
	responseChans  map[uint32]chan message.InboundMessage
	responsesCount map[uint32]expectedResponses
	timeoutManager timer.AdaptiveTimeoutManager
	metrics        *AppRequestNetworkMetrics
}

// expectedResponses counts the number of responses and compares against the expected number of responses
type expectedResponses struct {
	expected, received int
}

// Create a new RelayerExternalHandler to forward relevant inbound app messages to the respective
// Teleporter application relayer, as well as handle timeouts.
func NewRelayerExternalHandler(
	logger logging.Logger,
	metrics *AppRequestNetworkMetrics,
	timeoutManagerRegistry prometheus.Registerer,
) (*RelayerExternalHandler, error) {
	// TODO: Leaving this static for now, but we may want to have this as a config option
	cfg := timer.AdaptiveTimeoutConfig{
		InitialTimeout:     constants.DefaultNetworkInitialTimeout,
		MinimumTimeout:     constants.DefaultNetworkMinimumTimeout,
		MaximumTimeout:     constants.DefaultNetworkMaximumTimeout,
		TimeoutCoefficient: constants.DefaultNetworkTimeoutCoefficient,
		TimeoutHalflife:    constants.DefaultNetworkTimeoutHalflife,
	}

	timeoutManager, err := timer.NewAdaptiveTimeoutManager(&cfg, timeoutManagerRegistry)
	if err != nil {
		logger.Error(
			"Failed to create timeout manager",
			zap.Error(err),
		)
		return nil, err
	}

	go timeoutManager.Dispatch()

	return &RelayerExternalHandler{
		log:            logger,
		lock:           &sync.Mutex{},
		requestedNodes: make(map[uint32]*set.Set[ids.NodeID]),
		responseChans:  make(map[uint32]chan message.InboundMessage),
		responsesCount: make(map[uint32]expectedResponses),
		timeoutManager: timeoutManager,
		metrics:        metrics,
	}, nil
}

// HandleInbound handles all inbound app message traffic. For the relayer, we only care about App Responses to
// signature request App Requests, and App Request Fail messages sent by the timeout manager.
// For each inboundMessage, OnFinishedHandling must be called exactly once. However, since we handle relayer messages
// async, we must call OnFinishedHandling manually across all code paths.
//
// This diagram illustrates how HandleInbound forwards relevant AppResponses to the corresponding
// Teleporter application relayer. On startup, one Relayer goroutine is created per source subnet,
// which listens to the subscriber for cross-chain messages. When a cross-chain message is picked
// up by a Relayer, HandleInbound routes AppResponses traffic to the appropriate Relayer.
func (h *RelayerExternalHandler) HandleInbound(_ context.Context, inboundMessage *message.InboundMessage) {
	h.log.Debug(
		"Handling app response",
		zap.Stringer("op", inboundMessage.Op),
		zap.Stringer("from", inboundMessage.NodeID),
	)
	if inboundMessage.Op == message.AppResponseOp || inboundMessage.Op == message.AppErrorOp {
		if inboundMessage.Op == message.AppErrorOp {
			h.log.Debug("Received AppError message", zap.Stringer("message", inboundMessage.Message))
		}
		h.registerAppResponse(*inboundMessage)
	} else {
		h.log.Debug("Ignoring message", zap.Stringer("op", inboundMessage.Op))
		inboundMessage.OnFinishedHandling()
	}
}

func (h *RelayerExternalHandler) Connected(nodeID ids.NodeID, version *version.Application, subnetID ids.ID) {
	h.log.Debug(
		"Connected",
		zap.Stringer("nodeID", nodeID),
		zap.Stringer("version", version),
		zap.Stringer("subnetID", subnetID),
	)
	h.metrics.connects.Inc()
}

func (h *RelayerExternalHandler) Disconnected(nodeID ids.NodeID) {
	h.log.Debug(
		"Disconnected",
		zap.Stringer("nodeID", nodeID),
	)
	h.metrics.disconnects.Inc()
}

// RegisterRequestID registers an AppRequest by requestID, and marks the number of
// expected responses, equivalent to the number of nodes requested. requestID should
// be globally unique for the lifetime of the AppRequest. This is upper bounded by the timeout duration.
// It must be called BEFORE the AppRequest is sent, so that no response can reach the handler
// before it knows about the request. Nodes the send subsequently fails to reach should be
// dropped from the expected set with UnregisterRequestedNodes.
// NOTE: This function must be called at most once per requestID. Multiple calls with the same requestID
// will result in a fatal log and process termination.
func (h *RelayerExternalHandler) RegisterRequestID(
	requestID uint32,
	requestedNodes set.Set[ids.NodeID],
) chan message.InboundMessage {
	// Create a channel to receive the response
	h.lock.Lock()
	defer h.lock.Unlock()

	h.log.Debug("Registering request ID", zap.Uint32("requestID", requestID))

	if _, exist := h.responseChans[requestID]; exist {
		h.log.Fatal(
			"RegisterRequestID called more than once for the same requestID",
			zap.Uint32("requestID", requestID),
		)
		return nil
	}

	setWithNode := set.NewSet[ids.NodeID](requestedNodes.Len())
	h.requestedNodes[requestID] = &setWithNode

	// Add the requested nodes to the map
	for nodeID := range requestedNodes {
		h.requestedNodes[requestID].Add(nodeID)
	}

	numExpectedResponses := requestedNodes.Len()
	// Each requested node is forwarded at most once, so the buffer can never fill up and
	// the forwarding send under h.lock never blocks.
	responseChan := make(chan message.InboundMessage, numExpectedResponses)
	h.responseChans[requestID] = responseChan
	h.responsesCount[requestID] = expectedResponses{
		expected: numExpectedResponses,
	}
	return responseChan
}

// UnregisterRequestedNodes removes [nodeIDs] from the set of nodes expected to respond to
// [requestID] (typically because the AppRequest could not be sent to them) and cancels the
// timeouts registered for them via RegisterAppRequest. Nodes that are not (or no longer)
// expected to respond are ignored. If no responses remain outstanding afterwards, the
// response channel is closed and the request's tracking state is released.
func (h *RelayerExternalHandler) UnregisterRequestedNodes(
	requestID uint32,
	chainID ids.ID,
	nodeIDs []ids.NodeID,
) {
	h.lock.Lock()
	defer h.lock.Unlock()

	requestedNodes, ok := h.requestedNodes[requestID]
	if !ok {
		return
	}
	removed := 0
	for _, nodeID := range nodeIDs {
		if !requestedNodes.Contains(nodeID) {
			continue
		}
		requestedNodes.Remove(nodeID)
		h.timeoutManager.Remove(ids.RequestID{
			NodeID:    nodeID,
			ChainID:   chainID,
			RequestID: requestID,
			Op:        byte(message.AppResponseOp),
		})
		removed++
	}
	if removed == 0 {
		return
	}
	h.log.Debug(
		"Unregistered requested nodes",
		zap.Uint32("requestID", requestID),
		zap.Int("numUnregistered", removed),
	)
	responses := h.responsesCount[requestID]
	responses.expected -= removed
	h.responsesCount[requestID] = responses
	h.finishRequestIfComplete(requestID)
}

// RegisterAppRequest registers an AppRequest with the timeout manager.
// If RegisterResponse is not called before the timeout, HandleInbound is called with
// an internally created AppRequestFailed message.
func (h *RelayerExternalHandler) RegisterAppRequest(reqID ids.RequestID) {
	inMsg := message.InboundAppError(
		reqID.NodeID,
		reqID.ChainID,
		reqID.RequestID,
		avalancheCommon.ErrTimeout.Code,
		avalancheCommon.ErrTimeout.Message,
	)
	h.timeoutManager.Put(reqID, false, func() {
		h.HandleInbound(context.Background(), inMsg)
	})
}

// registerAppResponse attributes an AppResponse (or AppError) to the request and node it
// answers, cancels that node's timeout, and forwards it to the request's response channel.
func (h *RelayerExternalHandler) registerAppResponse(inboundMessage message.InboundMessage) {
	h.lock.Lock()
	defer h.lock.Unlock()

	// Ownership of the message, and therefore the responsibility to call
	// OnFinishedHandling, is transferred to the consumer of the response channel once the
	// message is forwarded. On every other path this handler must finish the message
	// itself, otherwise the inbound message throttler's byte and processing-message
	// allocations for the sending node are held forever.
	forwarded := false
	defer func() {
		if !forwarded {
			inboundMessage.OnFinishedHandling()
		}
	}()

	// Extract the message fields
	m := inboundMessage.Message

	chainID, err := message.GetChainID(m)
	if err != nil {
		h.log.Error("Could not get chainID from message")
		return
	}
	requestID, ok := message.GetRequestID(m)
	if !ok {
		h.log.Error("Could not get requestID from message")
		return
	}

	log := h.log.With(
		zap.Stringer("nodeID", inboundMessage.NodeID),
		zap.Uint32("requestID", requestID),
	)

	// Attribute the response BEFORE touching any timeout: a response from a node that is not
	// (or no longer) expected to answer must never cancel a timeout that still guards a live
	// expected-response slot, otherwise that slot can never be resolved and the request's
	// channel and tracking state leak.
	requestedNodes, ok := h.requestedNodes[requestID]
	if !ok || !requestedNodes.Contains(inboundMessage.NodeID) {
		log.Debug("Received response from unexpected node")
		return
	}
	// Each requested node is expected to answer exactly once. Consume its slot so a duplicate
	// response, or the timeout-generated AppError for a node that already answered, is ignored.
	requestedNodes.Remove(inboundMessage.NodeID)

	// Timeouts are registered under the op of the expected successful response, so map
	// failure ops (AppError, including this handler's own timeout notifications) back to it.
	timeoutOp := inboundMessage.Op
	if responseOp, isFailed := message.FailedToResponseOps[inboundMessage.Op]; isFailed {
		timeoutOp = responseOp
	}
	h.timeoutManager.Remove(ids.RequestID{
		NodeID:    inboundMessage.NodeID,
		ChainID:   chainID,
		RequestID: requestID,
		Op:        byte(timeoutOp),
	})

	// Dispatch to the appropriate response channel
	responseChan, ok := h.responseChans[requestID]
	if !ok {
		log.Error("Could not find response channel for request")
		return
	}
	responseChan <- inboundMessage
	forwarded = true

	// TODO: we can improve performance here by independently locking the response channel and response count maps
	responses, ok := h.responsesCount[requestID]
	if !ok {
		log.Error("Could not find expected responses for request")
		return
	}
	responses.received++
	h.responsesCount[requestID] = responses
	h.finishRequestIfComplete(requestID)
}

// finishRequestIfComplete closes the response channel and releases all tracking state for
// [requestID] once every expected response has been received (or unregistered).
// Must be called with h.lock held.
func (h *RelayerExternalHandler) finishRequestIfComplete(requestID uint32) {
	responses, ok := h.responsesCount[requestID]
	if !ok || responses.received < responses.expected {
		return
	}
	if responseChan, ok := h.responseChans[requestID]; ok {
		close(responseChan)
	}
	delete(h.responseChans, requestID)
	delete(h.responsesCount, requestID)
	delete(h.requestedNodes, requestID)
}
