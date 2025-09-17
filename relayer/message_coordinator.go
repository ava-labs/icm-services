// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package relayer

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/icm-services/database"
	"github.com/ava-labs/icm-services/messages"
	relayerTypes "github.com/ava-labs/icm-services/types"
	"github.com/ava-labs/icm-services/utils"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
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
	sourceClients           map[ids.ID]ethclient.Client
}

func NewMessageCoordinator(
	logger logging.Logger,
	messageHandlerFactories map[ids.ID]map[common.Address]messages.MessageHandlerFactory,
	applicationRelayers map[common.Hash]*ApplicationRelayer,
	sourceClients map[ids.ID]ethclient.Client,
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
	warpMessageInfo *relayerTypes.WarpMessageInfo,
) (
	*ApplicationRelayer,
	messages.MessageHandler,
	error,
) {
	// Check that the warp message is from a supported message protocol contract address.
	//nolint:lll
	messageHandlerFactory, supportedMessageProtocol := mc.messageHandlerFactories[warpMessageInfo.UnsignedMessage.SourceChainID][warpMessageInfo.SourceAddress]
	if !supportedMessageProtocol {
		// Do not return an error here because it is expected for there to be messages from other contracts
		// than just the ones supported by a single listener instance.
		mc.logger.Debug(
			"Warp message from unsupported message protocol address. Not relaying.",
			zap.Stringer("protocolAddress", warpMessageInfo.SourceAddress),
		)
		return nil, nil, nil
	}
	routeInfo, err := messageHandlerFactory.GetMessageRoutingInfo(warpMessageInfo.UnsignedMessage)
	if err != nil {
		mc.logger.Error("Failed to create message handler", zap.Error(err))
		return nil, nil, err
	}

	appRelayer := mc.getApplicationRelayer(
		routeInfo.SourceChainID,
		routeInfo.SenderAddress,
		routeInfo.DestinationChainID,
		routeInfo.DestinationAddress,
	)
	mc.logger.Info(
		"Unpacked warp message",
		zap.Stringer("sourceBlockchainID", routeInfo.SourceChainID),
		zap.Stringer("originSenderAddress", routeInfo.SenderAddress),
		zap.Stringer("destinationBlockchainID", routeInfo.DestinationChainID),
		zap.Stringer("destinationAddress", routeInfo.DestinationAddress),
		zap.Stringer("warpMessageID", warpMessageInfo.UnsignedMessage.ID()),
		zap.Bool("foundAppRelayer", appRelayer != nil),
	)
	if appRelayer == nil {
		return nil, nil, nil
	}

	messageHandler, err := messageHandlerFactory.NewMessageHandler(
		warpMessageInfo.UnsignedMessage,
		appRelayer.destinationClient,
	)
	if err != nil {
		mc.logger.Error("Failed to create message handler", zap.Error(err))
		return nil, nil, err
	}
	return appRelayer, messageHandler, nil
}

// Unpacks the Warp message and fetches the appropriate application relayer
// Checks for the following registered keys. At most one of these keys should be registered.
// 1. An exact match on sourceBlockchainID, destinationBlockchainID, originSenderAddress, and destinationAddress
// 2. A match on sourceBlockchainID and destinationBlockchainID, with a specific originSenderAddress and
// any destinationAddress
// 3. A match on sourceBlockchainID and destinationBlockchainID, with any originSenderAddress and a
// specific destinationAddress
// 4. A match on sourceBlockchainID and destinationBlockchainID, with any originSenderAddress and any
// destinationAddress
func (mc *MessageCoordinator) getApplicationRelayer(
	sourceBlockchainID ids.ID,
	originSenderAddress common.Address,
	destinationBlockchainID ids.ID,
	destinationAddress common.Address,
) *ApplicationRelayer {
	// Check for an exact match
	applicationRelayerID := database.CalculateRelayerID(
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

func (mc *MessageCoordinator) ProcessWarpMessage(warpMessage *relayerTypes.WarpMessageInfo) (common.Hash, error) {
	appRelayer, handler, err := mc.getAppRelayerMessageHandler(warpMessage)
	if err != nil {
		mc.logger.Error(
			"Failed to parse Warp message.",
			zap.Error(err),
			zap.Stringer("warpMessageID", warpMessage.UnsignedMessage.ID()),
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

	warpMessage, err := FetchWarpMessage(ethClient, messageID, blockNum)
	if err != nil {
		mc.logger.Error(
			"Failed to fetch warp from blockchain",
			zap.Stringer("blockchainID", blockchainID),
			zap.Error(err),
		)
		return common.Hash{}, fmt.Errorf("could not fetch warp message from ID: %w", err)
	}

	return mc.ProcessWarpMessage(warpMessage)
}

// Meant to be ran asynchronously. Errors should be sent to errChan.
func (mc *MessageCoordinator) ProcessBlock(
	icmBlockInfo *relayerTypes.WarpBlockInfo,
	blockchainID ids.ID,
	errChan chan error,
) {
	mc.logger.Debug(
		"Processing block",
		zap.Uint64("blockNumber", icmBlockInfo.BlockNumber),
		zap.Stringer("blockchainID", blockchainID),
	)

	// Register each message in the block with the appropriate application relayer
	messageHandlers := make(map[common.Hash][]messages.MessageHandler)
	for _, warpLogInfo := range icmBlockInfo.Messages {
		appRelayer, handler, err := mc.getAppRelayerMessageHandler(warpLogInfo)
		if err != nil {
			mc.logger.Error(
				"Failed to parse message",
				zap.Stringer("blockchainID", warpLogInfo.UnsignedMessage.SourceChainID),
				zap.Stringer("protocolAddress", warpLogInfo.SourceAddress),
				zap.Error(err),
			)
			continue
		}
		if appRelayer == nil {
			mc.logger.Debug(
				"Application relayer not found. Skipping message relay",
				zap.Stringer("warpMessageID", warpLogInfo.UnsignedMessage.ID()),
				zap.Stringer("sourceBlockchainID", warpLogInfo.UnsignedMessage.SourceChainID),
				zap.Stringer("originSenderAddress", warpLogInfo.SourceAddress),
				zap.Stringer("originTxID", warpLogInfo.SourceTxID),
			)
			continue
		}
		mc.logger.Info(
			"Registering message handler",
			zap.Stringer("relayerID", appRelayer.relayerID.ID),
			zap.Stringer("warpMessageID", warpLogInfo.UnsignedMessage.ID()),
			zap.Stringer("sourceBlockchainID", warpLogInfo.UnsignedMessage.SourceChainID),
			zap.Stringer("originSenderAddress", warpLogInfo.SourceAddress),
			zap.Stringer("originTxID", warpLogInfo.SourceTxID),
		)
		messageHandlers[appRelayer.relayerID.ID] = append(messageHandlers[appRelayer.relayerID.ID], handler)
	}
	// Initiate message relay of all registered messages
	for _, appRelayer := range mc.applicationRelayers {
		if appRelayer.sourceBlockchain.GetBlockchainID() != blockchainID {
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

func FetchWarpMessage(
	ethClient ethclient.Client,
	warpID ids.ID,
	blockNum *big.Int,
) (*relayerTypes.WarpMessageInfo, error) {
	fetchLogsCtx, fetchLogsCtxCancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
	defer fetchLogsCtxCancel()
	logs, err := ethClient.FilterLogs(fetchLogsCtx, ethereum.FilterQuery{
		Topics:    [][]common.Hash{{relayerTypes.WarpPrecompileLogFilter}, nil, {common.Hash(warpID)}},
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

	return relayerTypes.NewWarpMessageInfo(logs[0])
}
