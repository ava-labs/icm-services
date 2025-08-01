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
func (h *RelayerExternalHandler) HandleInbound(_ context.Context, inboundMessage message.InboundMessage) {
	h.log.Debug(
		"Handling app response",
		zap.Stringer("op", inboundMessage.Op()),
		zap.Stringer("from", inboundMessage.NodeID()),
	)
	if inboundMessage.Op() == message.AppResponseOp || inboundMessage.Op() == message.AppErrorOp {
		if inboundMessage.Op() == message.AppErrorOp {
			h.log.Debug("Received AppError message", zap.Stringer("message", inboundMessage.Message()))
		}
		h.registerAppResponse(inboundMessage)
	} else {
		h.log.Debug("Ignoring message", zap.Stringer("op", inboundMessage.Op()))
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
	responseChan := make(chan message.InboundMessage, numExpectedResponses)
	h.responseChans[requestID] = responseChan
	h.responsesCount[requestID] = expectedResponses{
		expected: numExpectedResponses,
	}
	return responseChan
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

// registerAppResponse registers an AppResponse with the timeout manager
func (h *RelayerExternalHandler) registerAppResponse(inboundMessage message.InboundMessage) {
	h.lock.Lock()
	defer h.lock.Unlock()

	// Extract the message fields
	m := inboundMessage.Message()

	chainID, err := message.GetChainID(m)
	if err != nil {
		h.log.Error("Could not get chainID from message")
		inboundMessage.OnFinishedHandling()
		return
	}
	requestID, ok := message.GetRequestID(m)
	if !ok {
		h.log.Error("Could not get requestID from message")
		inboundMessage.OnFinishedHandling()
		return
	}

	// Remove the timeout on the request
	reqID := ids.RequestID{
		NodeID:    inboundMessage.NodeID(),
		ChainID:   chainID,
		RequestID: requestID,
		Op:        byte(inboundMessage.Op()),
	}
	h.timeoutManager.Remove(reqID)

	// If the message is from an unexpected node, we ignore it
	if !h.isRequestedNode(requestID, reqID.NodeID) {
		h.log.Debug(
			"Received response from unexpected node",
			zap.Stringer("nodeID", reqID.NodeID),
			zap.Uint32("requestID", requestID),
		)
		return
	}

	// Dispatch to the appropriate response channel
	if responseChan, ok := h.responseChans[requestID]; ok {
		responseChan <- inboundMessage
	} else {
		h.log.Debug("Could not find response channel for request", zap.Uint32("requestID", requestID))
		return
	}

	// Check for the expected number of responses, and clear from the map if all expected responses have been received
	// TODO: we can improve performance here by independently locking the response channel and response count maps
	responses, ok := h.responsesCount[requestID]
	if !ok {
		h.log.Error("Could not find expected responses for request", zap.Uint32("requestID", requestID))
		return
	}
	received := responses.received + 1
	if received == responses.expected {
		close(h.responseChans[requestID])
		delete(h.responseChans, requestID)
		delete(h.responsesCount, requestID)
		delete(h.requestedNodes, requestID)
	} else {
		h.responsesCount[requestID] = expectedResponses{
			expected: responses.expected,
			received: received,
		}
	}
}

func (h *RelayerExternalHandler) isRequestedNode(
	requestID uint32,
	nodeID ids.NodeID,
) bool {
	if requestedNodes, ok := h.requestedNodes[requestID]; ok {
		return requestedNodes.Contains(nodeID)
	}
	return false
}
