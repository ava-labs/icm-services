// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peers

import (
	"context"
	"fmt"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/network/p2p"
	protop2p "github.com/ava-labs/avalanchego/proto/pb/p2p"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	avagoCommon "github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/subnets"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

var _ common.AppSender = (*AppRequestSender)(nil)

type AppRequestSender struct {
	logger         logging.Logger
	network        AppRequestNetwork
	p2pNetwork     *p2p.Network
	messageCreator message.Creator
	SourceChainID  ids.ID
	subnetID       ids.ID
	allower        subnets.Allower
}

func NewP2PClient(
	logger logging.Logger,
	network AppRequestNetwork,
	messageCreator message.Creator,
	subnetID ids.ID,
	chainID ids.ID,
) (*p2p.Client, error) {
	sender := &AppRequestSender{
		logger:         logger,
		network:        network,
		messageCreator: messageCreator,
		subnetID:       subnetID,
		allower:        subnets.NoOpAllower,
		SourceChainID:  chainID,
	}
	p2pNetwork, err := p2p.NewNetwork(
		logger,
		sender,
		prometheus.DefaultRegisterer,
		fmt.Sprintf("sig_agg_%s", subnetID),
	)
	if err != nil {
		return nil, err
	}
	sender.p2pNetwork = p2pNetwork
	return p2pNetwork.NewClient(p2p.SignatureRequestHandlerID), nil
}

func (s *AppRequestSender) SendAppResponse(
	ctx context.Context,
	nodeID ids.NodeID,
	requestID uint32,
	appResponseBytes []byte,
) error {
	return nil
}

// Following methods implement the AppSender interface

func (s *AppRequestSender) SendAppRequest(
	ctx context.Context,
	nodeIDs set.Set[ids.NodeID],
	requestID uint32,
	appRequestBytes []byte,
) error {
	connectedValidators, err := s.network.GetConnectedCanonicalValidators(s.subnetID)
	if err != nil {
		return err
	}
	for nodeID := range nodeIDs {
		_, ok := connectedValidators.NodeValidatorIndexMap[nodeID]
		if !ok {
			s.logger.Warn(
				"Failed to make app request to node -- not connected",
				zap.String("nodeID", nodeID.String()),
				zap.Error(err),
			)
		}
		// Register a timeout response for each queried node
		reqID := ids.RequestID{
			NodeID:    nodeID,
			ChainID:   s.SourceChainID,
			RequestID: requestID,
			Op:        byte(message.AppResponseOp),
		}
		s.network.RegisterAppRequest(reqID)
	}
	outMsg, err := s.messageCreator.AppRequest(s.SourceChainID, requestID, DefaultAppRequestTimeout, appRequestBytes)
	if err != nil {
		return err
	}

	failedToSend := 0
	sentTo := s.network.Send(outMsg, avagoCommon.SendConfig{NodeIDs: nodeIDs}, s.subnetID, s.allower)
	s.logger.Warn("Sent app request", zap.Int("sentTo", sentTo.Len()))
	for nodeID := range nodeIDs {
		if !sentTo.Contains(nodeID) {
			// s.logger.Warn(
			// 	"Failed to make app request to node",
			// 	zap.String("nodeID", nodeID.String()),
			// 	zap.Error(err),
			// )
			failedToSend += 1
		}
		// s.network.RegisterAppRequest(reqID)
	}
	if failedToSend > 0 {
		s.logger.Warn("Failed to send app request to some nodes", zap.Int("failedToSend", failedToSend))
	}
	responseChan := s.network.RegisterRequestID(requestID, len(connectedValidators.ValidatorSet))
	go s.handleResponses(len(connectedValidators.ValidatorSet), responseChan)
	return nil
}

func (s *AppRequestSender) SendAppError(
	ctx context.Context,
	nodeID ids.NodeID,
	requestID uint32,
	errorCode int32,
	errorMessage string,
) error {
	return nil
}

func (s *AppRequestSender) SendAppGossip(
	ctx context.Context,
	config common.SendConfig,
	appGossipBytes []byte,
) error {
	return nil
}

func (s *AppRequestSender) handleResponses(
	responsesExpected int,
	responseChan chan message.InboundMessage,
) {
	for response := range responseChan {
		appResponse, ok := response.Message().(*protop2p.AppResponse)
		if ok {
			err := s.p2pNetwork.AppResponse(
				context.Background(),
				response.NodeID(),
				appResponse.GetRequestId(),
				appResponse.GetAppBytes())
			if err != nil {
				s.logger.Warn("Failed to handle AppResponse", zap.Error(err))
			}
		}
	}
}
