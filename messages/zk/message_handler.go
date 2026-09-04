// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package zk implements the message handler for Teleporter messages
// originating on Ethereum and verified on an Avalanche destination via the
// ZKAdapter (an IAdapter implementation), using Boundless zero-knowledge proofs
// of Ethereum consensus.

package zk

import (
	"context"
	"fmt"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator"
	"github.com/ava-labs/icm-services/vms"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/ethclient"
	"go.uber.org/zap"
)

type factory struct {
	messageConfig   *Config
	protocolAddress common.Address
	sourceClient    *ethclient.Client
}

type messageHandler struct {
	logger            logging.Logger
	sourceClient      *ethclient.Client
	destinationClient vms.DestinationClient
	metrics           messages.Metrics
	messageConfig     *Config
	// sourceMessage contains the protocol address (the adapter).
	sourceMessage     messages.SourceMessage
	sourceBlockNumber uint64
	sourceLogIndex    uint
	slot              uint64
}

func NewMessageHandlerFactory(
	messageProtocolAddress common.Address,
	messageProtocolConfig config.MessageProtocolConfig,
	sourceClient *ethclient.Client,
) (messages.MessageHandlerFactory, error) {
	messageConfig, err := ConfigFromMap(messageProtocolConfig.Settings)
	if err != nil {
		return nil, fmt.Errorf("invalid zk message protocol config: %w", err)
	}
	return &factory{
		messageConfig:   messageConfig,
		protocolAddress: messageProtocolAddress,
		sourceClient:    sourceClient,
	}, nil
}

func (f *factory) GetMessageRoutingInfo(
	message *messages.SourceMessage,
) (messages.MessageRoutingInfo, error) {
	teleporterMessage, err := parseMessage(message)
	if err != nil {
		return messages.MessageRoutingInfo{}, fmt.Errorf("failed to parse teleporter v2 message: %w", err)
	}
	return messages.MessageRoutingInfo{
		SourceChainID:      message.SourceBlockchainID,
		SenderAddress:      message.ProtocolAddress,
		DestinationChainID: ids.ID(teleporterMessage.DestinationBlockchainID),
		DestinationAddress: message.ProtocolAddress,
	}, nil
}

// NewMessageHandler locates the message's log on the source chain and returns
// the handler that will prove and deliver it. The signature aggregator,
// signing subnet, and quorum numerator are unused as the attestation is a Boundless
// consensus proof verified on-chain. Issue: https://github.com/ava-labs/icm-services/issues/1462
func (f *factory) NewMessageHandler(
	logger logging.Logger,
	message *messages.SourceMessage,
	destinationClient vms.DestinationClient,
	_ *aggregator.SignatureAggregator,
	metrics messages.Metrics,
	_ ids.ID,
	_ uint64,
) (messages.MessageHandler, error) {
	if message.SourceTxID == (common.Hash{}) {
		return nil, fmt.Errorf("SourceTxID on zk messages must be non-empty")
	}
	if len(message.Payload) == 0 {
		return nil, fmt.Errorf("empty TeleporterV2MessageSent log data")
	}
	// From the source message transaction id, derive the transaction receipt and
	// match the TeleporterV2MessageSent event log by emitter, topic, and payload.
	receipt, err := f.sourceClient.TransactionReceipt(context.Background(), message.SourceTxID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch source receipt %s: %w", message.SourceTxID.Hex(), err)
	}
	logIndex, err := findMessageLog(receipt, message.ProtocolAddress, message.Payload)
	if err != nil {
		return nil, err
	}
	blockNumber := receipt.BlockNumber.Uint64()
	header, err := f.sourceClient.HeaderByNumber(context.Background(), receipt.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch source header %d: %w", blockNumber, err)
	}
	slot := (header.Time - f.messageConfig.BeaconGenesisTime) / f.messageConfig.SecondsPerSlot // the floor

	return &messageHandler{
		logger: logger.With(
			zap.String("sourceTxHash", message.SourceTxID.Hex()),
			zap.Uint64("blockNumber", blockNumber),
			zap.Uint64("slot", slot),
		),
		sourceClient:      f.sourceClient,
		destinationClient: destinationClient,
		metrics:           metrics,
		messageConfig:     f.messageConfig,
		sourceMessage:     *message,
		sourceBlockNumber: blockNumber,
		sourceLogIndex:    logIndex,
		slot:              slot,
	}, nil
}

// ProcessMessage delivers the message to the destination via the TeleporterV2 messenger,
// attaching a zk attestation.
// TODO: Stubbed out. Implement this in follow up work
func (m *messageHandler) ProcessMessage() (common.Hash, error) {
	return common.Hash{}, fmt.Errorf("zk message delivery not yet implemented")
}
