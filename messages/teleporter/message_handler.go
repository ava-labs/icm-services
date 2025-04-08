// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporter

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	warpPayload "github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	teleportermessenger "github.com/ava-labs/icm-contracts/abi-bindings/go/teleporter/TeleporterMessenger"
	gasUtils "github.com/ava-labs/icm-contracts/utils/gas-utils"
	teleporterUtils "github.com/ava-labs/icm-contracts/utils/teleporter-utils"
	"github.com/ava-labs/icm-services/messages"
	pbDecider "github.com/ava-labs/icm-services/proto/pb/decider"
	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/icm-services/vms"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	defaultBlockAcceptanceTimeout = 30 * time.Second
)

type factory struct {
	messageConfig   *Config
	protocolAddress common.Address
	logger          logging.Logger
	deciderClient   pbDecider.DeciderServiceClient
}

type messageHandler struct {
	logger              logging.Logger
	teleporterMessage   *teleportermessenger.TeleporterMessage
	unsignedMessage     *warp.UnsignedMessage
	factory             *factory
	deciderClient       pbDecider.DeciderServiceClient
	destinationClient   vms.DestinationClient
	teleporterMessageID ids.ID
	logFields           []zap.Field
}

// define an "empty" decider client to use when a connection isn't provided:
type emptyDeciderClient struct{}

func (s *emptyDeciderClient) ShouldSendMessage(
	_ context.Context,
	_ *pbDecider.ShouldSendMessageRequest,
	_ ...grpc.CallOption,
) (*pbDecider.ShouldSendMessageResponse, error) {
	return &pbDecider.ShouldSendMessageResponse{ShouldSendMessage: true}, nil
}

func NewMessageHandlerFactory(
	logger logging.Logger,
	messageProtocolAddress common.Address,
	messageProtocolConfig config.MessageProtocolConfig,
	deciderClientConn *grpc.ClientConn,
) (messages.MessageHandlerFactory, error) {
	messageConfig, err := ConfigFromMap(messageProtocolConfig.Settings)
	if err != nil {
		logger.Error(
			"Invalid Teleporter config.",
			zap.Error(err),
		)
		return nil, err
	}

	var deciderClient pbDecider.DeciderServiceClient
	if deciderClientConn == nil {
		deciderClient = &emptyDeciderClient{}
	} else {
		deciderClient = pbDecider.NewDeciderServiceClient(deciderClientConn)
	}

	return &factory{
		messageConfig:   messageConfig,
		protocolAddress: messageProtocolAddress,
		logger:          logger,
		deciderClient:   deciderClient,
	}, nil
}

func (f *factory) NewMessageHandler(
	unsignedMessage *warp.UnsignedMessage,
	destinationClient vms.DestinationClient,
) (messages.MessageHandler, error) {
	teleporterMessage, err := f.parseTeleporterMessage(unsignedMessage)
	if err != nil {
		f.logger.Error(
			"Failed to parse teleporter message.",
			zap.String("warpMessageID", unsignedMessage.ID().String()),
		)
		return nil, err
	}
	destinationBlockChainID := destinationClient.DestinationBlockchainID()
	teleporterMessageID, err := teleporterUtils.CalculateMessageID(
		f.protocolAddress,
		unsignedMessage.SourceChainID,
		destinationBlockChainID,
		teleporterMessage.MessageNonce,
	)
	if err != nil {
		f.logger.Error(
			"Failed to calculate Teleporter message ID.",
			zap.Stringer("warpMessageID", unsignedMessage.ID()),
			zap.Error(err),
		)
		return &messageHandler{}, err
	}

	logFields := []zap.Field{
		zap.Stringer("warpMessageID", unsignedMessage.ID()),
		zap.Stringer("teleporterMessageID", teleporterMessageID),
		zap.Stringer("destinationBlockchainID", destinationBlockChainID),
	}
	return &messageHandler{
		logger:            f.logger.With(logFields...),
		teleporterMessage: teleporterMessage,

		unsignedMessage:     unsignedMessage,
		factory:             f,
		deciderClient:       f.deciderClient,
		destinationClient:   destinationClient,
		teleporterMessageID: teleporterMessageID,
		logFields:           logFields,
	}, nil
}

func (f *factory) GetMessageRoutingInfo(unsignedMessage *warp.UnsignedMessage) (messages.MessageRoutingInfo, error) {
	teleporterMessage, err := f.parseTeleporterMessage(unsignedMessage)
	if err != nil {
		f.logger.Error(
			"Failed to parse teleporter message.",
			zap.String("warpMessageID", unsignedMessage.ID().String()),
		)
		return messages.MessageRoutingInfo{}, err
	}
	return messages.MessageRoutingInfo{
		SourceChainID:      unsignedMessage.SourceChainID,
		SenderAddress:      teleporterMessage.OriginSenderAddress,
		DestinationChainID: teleporterMessage.DestinationBlockchainID,
		DestinationAddress: teleporterMessage.DestinationAddress,
	}, nil
}

func isAllowedRelayer(allowedRelayers []common.Address, eoa common.Address) bool {
	// If no allowed relayer addresses were set, then anyone can relay it.
	if len(allowedRelayers) == 0 {
		return true
	}

	return slices.Contains(allowedRelayers, eoa)
}

func (m *messageHandler) GetUnsignedMessage() *warp.UnsignedMessage {
	return m.unsignedMessage
}

func (m *messageHandler) GetMessageRoutingInfo() (
	ids.ID,
	common.Address,
	ids.ID,
	common.Address,
	error,
) {
	return m.unsignedMessage.SourceChainID,
		m.teleporterMessage.OriginSenderAddress,
		m.teleporterMessage.DestinationBlockchainID,
		m.teleporterMessage.DestinationAddress,
		nil
}

// ShouldSendMessage returns true if the message should be sent to the destination chain
func (m *messageHandler) ShouldSendMessage() (bool, error) {
	requiredGasLimit := m.teleporterMessage.RequiredGasLimit.Uint64()
	destBlockGasLimit := m.destinationClient.BlockGasLimit()
	// Check if the specified gas limit is below the maximum threshold
	if requiredGasLimit > destBlockGasLimit {
		m.logger.Info(
			"Gas limit exceeds maximum threshold",
			zap.Uint64("requiredGasLimit", m.teleporterMessage.RequiredGasLimit.Uint64()),
			zap.Uint64("blockGasLimit", destBlockGasLimit),
		)
		return false, nil
	}

	// Check if the relayer is allowed to deliver this message
	senderAddress := m.destinationClient.SenderAddress()
	if !isAllowedRelayer(m.teleporterMessage.AllowedRelayerAddresses, senderAddress) {
		m.logger.Info("Relayer EOA not allowed to deliver this message.")
		return false, nil
	}

	// Check if the message has already been delivered to the destination chain
	teleporterMessenger := m.factory.getTeleporterMessenger(m.destinationClient)
	delivered, err := teleporterMessenger.MessageReceived(&bind.CallOpts{}, m.teleporterMessageID)
	if err != nil {
		m.logger.Error(
			"Failed to check if message has been delivered to destination chain.",
			zap.Error(err),
		)
		return false, err
	}
	if delivered {
		m.logger.Info("Message already delivered to destination.")
		return false, nil
	}

	// Dispatch to the external decider service. If the service is unavailable or returns
	// an error, then use the decision that has already been made, i.e. return true
	decision, err := m.getShouldSendMessageFromDecider()
	if err != nil {
		m.logger.Warn("Error delegating to decider")
		return true, nil
	}
	if !decision {
		m.logger.Info("Decider rejected message")
	}
	return decision, nil
}

// Queries the decider service to determine whether this message should be
// sent. If the decider client is nil, returns true.
func (m *messageHandler) getShouldSendMessageFromDecider() (bool, error) {
	warpMsgID := m.unsignedMessage.ID()

	ctx, cancelCtx := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelCtx()
	response, err := m.deciderClient.ShouldSendMessage(
		ctx,
		&pbDecider.ShouldSendMessageRequest{
			NetworkId:           m.unsignedMessage.NetworkID,
			SourceChainId:       m.unsignedMessage.SourceChainID[:],
			Payload:             m.unsignedMessage.Payload,
			BytesRepresentation: m.unsignedMessage.Bytes(),
			Id:                  warpMsgID[:],
		},
	)
	if err != nil {
		m.logger.Error("Error response from decider.", zap.Error(err))
		return false, err
	}

	return response.ShouldSendMessage, nil
}

// SendMessage extracts the gasLimit and packs the call data to call the receiveCrossChainMessage
// method of the Teleporter contract, and dispatches transaction construction and broadcast to the
// destination client.
func (m *messageHandler) SendMessage(
	signedMessage *warp.Message,
) (common.Hash, error) {
	destinationBlockchainID := m.destinationClient.DestinationBlockchainID()
	teleporterMessageID, err := teleporterUtils.CalculateMessageID(
		m.factory.protocolAddress,
		signedMessage.SourceChainID,
		destinationBlockchainID,
		m.teleporterMessage.MessageNonce,
	)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to calculate Teleporter message ID: %w", err)
	}

	m.logger.Info("Sending message to destination chain")
	numSigners, err := signedMessage.Signature.NumSigners()
	if err != nil {
		m.logger.Error("Failed to get number of signers")
		return common.Hash{}, err
	}

	gasLimit, err := gasUtils.CalculateReceiveMessageGasLimit(
		numSigners,
		m.teleporterMessage.RequiredGasLimit,
		len(signedMessage.Bytes()),
		len(signedMessage.Payload),
		len(m.teleporterMessage.Receipts),
	)
	if err != nil {
		m.logger.Error("Failed to calculate gas limit for receiveCrossChainMessage call")
		return common.Hash{}, err
	}
	// Construct the transaction call data to call the receive cross chain message method of the receiver precompile.
	callData, err := teleportermessenger.PackReceiveCrossChainMessage(
		0,
		common.HexToAddress(m.factory.messageConfig.RewardAddress),
	)
	if err != nil {
		m.logger.Error("Failed packing receiveCrossChainMessage call data")
		return common.Hash{}, err
	}

	txHash, err := m.destinationClient.SendTx(
		signedMessage,
		m.factory.protocolAddress.Hex(),
		gasLimit,
		callData,
	)
	if err != nil {
		m.logger.Error("Failed to send tx.", zap.Error(err))
		return common.Hash{}, err
	}

	// Wait for the message to be included in a block before returning
	err = m.waitForReceipt(signedMessage, m.destinationClient, txHash, teleporterMessageID)
	if err != nil {
		return common.Hash{}, err
	}

	m.logger.Info(
		"Delivered message to destination chain",
		zap.String("txHash", txHash.String()),
	)
	return txHash, nil
}

func (m *messageHandler) GetLogContext() []zap.Field {
	return m.logFields
}

func (m *messageHandler) waitForReceipt(
	signedMessage *warp.Message,
	destinationClient vms.DestinationClient,
	txHash common.Hash,
	teleporterMessageID ids.ID,
) error {
	callCtx, callCtxCancel := context.WithTimeout(context.Background(), defaultBlockAcceptanceTimeout)
	defer callCtxCancel()
	var receipt *types.Receipt
	operation := func() (err error) {
		receipt, err = destinationClient.Client().(ethclient.Client).TransactionReceipt(callCtx, txHash)
		return err
	}
	err := utils.WithRetriesTimeout(m.logger, operation, defaultBlockAcceptanceTimeout)
	if err != nil {
		m.logger.Error(
			"Failed to get transaction receipt",
			zap.Error(err),
		)
		return err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		m.logger.Error(
			"Transaction failed",
			zap.String("txHash", txHash.String()),
		)
		return fmt.Errorf("transaction failed with status: %d", receipt.Status)
	}
	return nil
}

// parseTeleporterMessage returns the Warp message's corresponding Teleporter message from the cache if it exists.
// Otherwise parses the Warp message payload.
func (f *factory) parseTeleporterMessage(
	unsignedMessage *warp.UnsignedMessage,
) (*teleportermessenger.TeleporterMessage, error) {
	addressedPayload, err := warpPayload.ParseAddressedCall(unsignedMessage.Payload)
	if err != nil {
		f.logger.Error(
			"Failed parsing addressed payload",
			zap.Error(err),
		)
		return nil, err
	}
	var teleporterMessage teleportermessenger.TeleporterMessage
	err = teleporterMessage.Unpack(addressedPayload.Payload)
	if err != nil {
		f.logger.Error(
			"Failed unpacking teleporter message.",
			zap.String("warpMessageID", unsignedMessage.ID().String()),
		)
		return nil, err
	}

	return &teleporterMessage, nil
}

// getTeleporterMessenger returns the Teleporter messenger instance for the destination chain.
// Panic instead of returning errors because this should never happen, and if it does, we do not
// want to log and swallow the error, since operations after this will fail too.
func (f *factory) getTeleporterMessenger(
	destinationClient vms.DestinationClient,
) *teleportermessenger.TeleporterMessenger {
	client, ok := destinationClient.Client().(ethclient.Client)
	if !ok {
		panic(fmt.Sprintf(
			"Destination client for chain %s is not an Ethereum client",
			destinationClient.DestinationBlockchainID().String()),
		)
	}

	// Get the teleporter messenger contract
	teleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(f.protocolAddress, client)
	if err != nil {
		panic(fmt.Sprintf("Failed to get teleporter messenger contract: %s", err.Error()))
	}
	return teleporterMessenger
}
