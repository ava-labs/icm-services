// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package relayer

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/contracts/warp"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/icm-services/database"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/icm-services/vms/evm"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/ethclient"
	"go.uber.org/zap"
)

// MessageCoordinator contains all the logic required to process messages in the relayer.
// Other components such as the listeners or the API should pass messages to the MessageCoordinator
// so that it can parse the message(s) and pass them the proper ApplicationRelayer.
type MessageCoordinator struct {
	logger logging.Logger
	// Maps Source blockchain ID and protocol address to a Message Handler Factory
	messageHandlerFactories map[ids.ID]map[common.Address]messages.MessageHandlerFactory
	applicationRelayers     map[common.Hash]*ApplicationRelayer
	sourceClients           map[ids.ID]*ethclient.Client
}

func NewMessageCoordinator(
	logger logging.Logger,
	messageHandlerFactories map[ids.ID]map[common.Address]messages.MessageHandlerFactory,
	applicationRelayers map[common.Hash]*ApplicationRelayer,
	sourceClients map[ids.ID]*ethclient.Client,
) *MessageCoordinator {
	return &MessageCoordinator{
		logger:                  logger,
		messageHandlerFactories: messageHandlerFactories,
		applicationRelayers:     applicationRelayers,
		sourceClients:           sourceClients,
	}
}

// getAppRelayerMessageHandler returns the ApplicationRelayer that is configured to handle this message,
// as well as a one-time MessageHandler instance that the ApplicationRelayer uses to relay this specific message.
// The MessageHandler and ApplicationRelayer are decoupled to support batch workflows in which a single
// ApplicationRelayer processes multiple messages (using their corresponding MessageHandlers) in a single shot.
func (mc *MessageCoordinator) getAppRelayerMessageHandler(
	message *messages.SourceMessage,
) (
	*ApplicationRelayer,
	messages.MessageHandler,
	error,
) {
	// Check that the message is from a supported message protocol contract address.
	protocolAddress := message.ProtocolAddress
	messageHandlerFactory, supportedMessageProtocol :=
		mc.messageHandlerFactories[message.SourceBlockchainID][protocolAddress]
	if !supportedMessageProtocol {
		// Do not return an error here because it is expected for there to be messages from other contracts
		// than just the ones supported by a single listener instance.
		mc.logger.Debug(
			"Message from unsupported message protocol address. Not relaying.",
			zap.Stringer("protocolAddress", protocolAddress),
		)
		return nil, nil, nil
	}
	routeInfo, err := messageHandlerFactory.GetMessageRoutingInfo(message)
	if err != nil {
		mc.logger.Error("Failed to get message routing info", zap.Error(err))
		return nil, nil, err
	}

	appRelayer := mc.getApplicationRelayer(
		protocolAddress,
		routeInfo.SourceChainID,
		routeInfo.SenderAddress,
		routeInfo.DestinationChainID,
		routeInfo.DestinationAddress,
	)
	mc.logger.Info(
		"Unpacked message",
		zap.Stringer("sourceBlockchainID", routeInfo.SourceChainID),
		zap.Stringer("originSenderAddress", routeInfo.SenderAddress),
		zap.Stringer("destinationBlockchainID", routeInfo.DestinationChainID),
		zap.Stringer("destinationAddress", routeInfo.DestinationAddress),
		zap.Stringer("originTxID", message.SourceTxID),
		zap.Bool("foundAppRelayer", appRelayer != nil),
	)
	if appRelayer == nil {
		return nil, nil, nil
	}

	messageHandler, err := messageHandlerFactory.NewMessageHandler(
		appRelayer.logger,
		message,
		appRelayer.destinationClient,
		appRelayer.signatureAggregator,
		appRelayer.metrics,
		appRelayer.signingSubnetID,
		appRelayer.warpConfig.QuorumNumerator,
	)
	if err != nil {
		mc.logger.Error("Failed to create message handler", zap.Error(err))
		return nil, nil, err
	}
	return appRelayer, messageHandler, nil
}

// Fetches the application relayer registered for the message's routing information.
// Checks for the following registered keys. At most one of these keys should be registered.
// 1. An exact match on sourceBlockchainID, destinationBlockchainID, originSenderAddress, and destinationAddress
// 2. A match on sourceBlockchainID and destinationBlockchainID, with a specific originSenderAddress and
// any destinationAddress
// 3. A match on sourceBlockchainID and destinationBlockchainID, with any originSenderAddress and a
// specific destinationAddress
// 4. A match on sourceBlockchainID and destinationBlockchainID, with any originSenderAddress and any
// destinationAddress
func (mc *MessageCoordinator) getApplicationRelayer(
	protocolAddress common.Address,
	sourceBlockchainID ids.ID,
	originSenderAddress common.Address,
	destinationBlockchainID ids.ID,
	destinationAddress common.Address,
) *ApplicationRelayer {
	// Check for an exact match
	applicationRelayerID := database.CalculateRelayerID(
		protocolAddress,
		sourceBlockchainID,
		destinationBlockchainID,
		originSenderAddress,
		destinationAddress,
	)
	if applicationRelayer, ok := mc.applicationRelayers[applicationRelayerID]; ok {
		return applicationRelayer
	}

	// Check for a match on sourceBlockchainID and destinationBlockchainID, with a specific
	// originSenderAddress and any destinationAddress.
	applicationRelayerID = database.CalculateRelayerID(
		protocolAddress,
		sourceBlockchainID,
		destinationBlockchainID,
		originSenderAddress,
		database.AllAllowedAddress,
	)
	if applicationRelayer, ok := mc.applicationRelayers[applicationRelayerID]; ok {
		return applicationRelayer
	}

	// Check for a match on sourceBlockchainID and destinationBlockchainID, with any originSenderAddress
	// and a specific destinationAddress.
	applicationRelayerID = database.CalculateRelayerID(
		protocolAddress,
		sourceBlockchainID,
		destinationBlockchainID,
		database.AllAllowedAddress,
		destinationAddress,
	)
	if applicationRelayer, ok := mc.applicationRelayers[applicationRelayerID]; ok {
		return applicationRelayer
	}

	// Check for a match on sourceBlockchainID and destinationBlockchainID, with any originSenderAddress
	// and any destinationAddress.
	applicationRelayerID = database.CalculateRelayerID(
		protocolAddress,
		sourceBlockchainID,
		destinationBlockchainID,
		database.AllAllowedAddress,
		database.AllAllowedAddress,
	)
	if applicationRelayer, ok := mc.applicationRelayers[applicationRelayerID]; ok {
		return applicationRelayer
	}
	mc.logger.Debug(
		"Application relayer not found. Skipping message relay.",
		zap.Stringer("blockchainID", sourceBlockchainID),
		zap.Stringer("destinationBlockchainID", destinationBlockchainID),
		zap.Stringer("originSenderAddress", originSenderAddress),
		zap.Stringer("destinationAddress", destinationAddress),
	)
	return nil
}

func (mc *MessageCoordinator) ProcessMessage(message *messages.SourceMessage) (common.Hash, error) {
	appRelayer, handler, err := mc.getAppRelayerMessageHandler(message)
	if err != nil {
		mc.logger.Error(
			"Failed to parse message.",
			zap.Stringer("sourceBlockchainID", message.SourceBlockchainID),
			zap.Stringer("protocolAddress", message.ProtocolAddress),
			zap.Error(err),
		)
		return common.Hash{}, err
	}
	if appRelayer == nil {
		mc.logger.Error("Application relayer not found")
		return common.Hash{}, errors.New("application relayer not found")
	}

	return appRelayer.ProcessMessage(handler)
}

func (mc *MessageCoordinator) ProcessMessageID(
	blockchainID ids.ID,
	messageID ids.ID,
	blockNum *big.Int,
) (common.Hash, error) {
	ethClient, ok := mc.sourceClients[blockchainID]
	if !ok {
		mc.logger.Error(
			"Source client not found",
			zap.Stringer("blockchainID", blockchainID),
		)
		return common.Hash{}, fmt.Errorf("source client not set for blockchain: %s", blockchainID.String())
	}

	warpMessage, err := FetchWarpMessage(ethClient, blockchainID, messageID, blockNum)
	if err != nil {
		mc.logger.Error(
			"Failed to fetch warp from blockchain",
			zap.Stringer("blockchainID", blockchainID),
			zap.Error(err),
		)
		return common.Hash{}, fmt.Errorf("could not fetch warp message from ID: %w", err)
	}

	return mc.ProcessMessage(warpMessage)
}

// Meant to be ran asynchronously. Errors should be sent to errChan.
// The logs of [icmBlockInfo] are expected to match the event filter of the message protocol at
// [protocolAddress], since that is what the subscriber that produced them filtered on.
func (mc *MessageCoordinator) ProcessBlock(
	icmBlockInfo *evm.ICMBlockInfo,
	blockchainID ids.ID,
	protocolAddress common.Address,
	errChan chan error,
) {
	mc.logger.Debug(
		"Processing block",
		zap.Uint64("blockNumber", icmBlockInfo.BlockNumber),
		zap.Stringer("blockchainID", blockchainID),
		zap.Stringer("protocolAddress", protocolAddress),
	)

	// Register each message in the block with the appropriate application relayer
	messageHandlers := make(map[common.Hash][]messages.MessageHandler)
	for _, log := range icmBlockInfo.Logs {
		message := messages.NewSourceMessage(blockchainID, protocolAddress, log)
		appRelayer, handler, err := mc.getAppRelayerMessageHandler(message)
		if err != nil {
			mc.logger.Error(
				"Failed to parse message",
				zap.Stringer("blockchainID", blockchainID),
				zap.Stringer("protocolAddress", protocolAddress),
				zap.Stringer("originTxID", message.SourceTxID),
				zap.Error(err),
			)
			continue
		}
		if appRelayer == nil {
			mc.logger.Debug(
				"Application relayer not found. Skipping message relay",
				zap.Stringer("sourceBlockchainID", blockchainID),
				zap.Stringer("protocolAddress", protocolAddress),
				zap.Stringer("originTxID", message.SourceTxID),
			)
			continue
		}
		mc.logger.Info(
			"Registering message handler",
			zap.Stringer("relayerID", appRelayer.relayerID.ID),
			zap.Stringer("sourceBlockchainID", blockchainID),
			zap.Stringer("protocolAddress", protocolAddress),
			zap.Stringer("originTxID", message.SourceTxID),
		)
		messageHandlers[appRelayer.relayerID.ID] = append(messageHandlers[appRelayer.relayerID.ID], handler)
	}
	// Initiate message relay of all registered messages
	for _, appRelayer := range mc.applicationRelayers {
		if appRelayer.relayerID.SourceBlockchainID != blockchainID {
			continue
		}
		// Dispatch all messages in the block to the appropriate application relayer.
		// An empty slice is still a valid argument to ProcessHeight; in this case the height is immediately committed.
		handlers := messageHandlers[appRelayer.relayerID.ID]
		mc.logger.Verbo(
			"Dispatching to app relayer",
			zap.Stringer("relayerID", appRelayer.relayerID.ID),
			zap.Int("numMessages", len(handlers)),
		)
		go appRelayer.ProcessHeight(icmBlockInfo.BlockNumber, handlers, errChan)
	}
}

// FetchWarpMessage fetches the Warp message with ID [warpID] that was sent in block [blockNum] on
// [blockchainID]. Only message protocols that send their messages through the Warp precompile can
// be looked up by Warp message ID.
func FetchWarpMessage(
	ethClient *ethclient.Client,
	blockchainID ids.ID,
	warpID ids.ID,
	blockNum *big.Int,
) (*messages.SourceMessage, error) {
	fetchLogsCtx, fetchLogsCtxCancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
	defer fetchLogsCtxCancel()
	logs, err := ethClient.FilterLogs(fetchLogsCtx, ethereum.FilterQuery{
		Topics:    [][]common.Hash{{messages.WarpPrecompileLogFilter}, nil, {common.Hash(warpID)}},
		Addresses: []common.Address{warp.ContractAddress},
		FromBlock: blockNum,
		ToBlock:   blockNum,
	})
	if err != nil {
		return nil, fmt.Errorf("could not fetch logs: %w", err)
	}
	if len(logs) != 1 {
		return nil, fmt.Errorf("found more than 1 log: %d", len(logs))
	}

	return messages.NewSourceMessageFromWarpLog(blockchainID, logs[0])
}
