// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package offchainregistry

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	warpPayload "github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	teleporterregistry "github.com/ava-labs/icm-services/abi-bindings/go/teleporter/registry/TeleporterRegistry"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/icm-services/vms"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"go.uber.org/zap"
)

var OffChainRegistrySourceAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")

const (
	addProtocolVersionGasLimit  uint64 = 500_000
	revertVersionNotFoundString        = "TeleporterRegistry: version not found"
)

type factory struct {
	registryAddress common.Address
}

type messageHandler struct {
	logger              logging.Logger
	unsignedMessage     *warp.UnsignedMessage
	destinationClient   vms.DestinationClient
	signatureAggregator *aggregator.SignatureAggregator
	metrics             messages.Metrics
	registryAddress     common.Address
	signingSubnetID     ids.ID
	quorumNumerator     uint64
}

func NewMessageHandlerFactory(
	messageProtocolConfig config.MessageProtocolConfig,
) (messages.MessageHandlerFactory, error) {
	// Marshal the map and unmarshal into the off-chain registry config
	data, err := json.Marshal(messageProtocolConfig.Settings)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal off-chain registry config: %w", err)
	}
	var messageConfig Config
	if err := json.Unmarshal(data, &messageConfig); err != nil {
		return nil, fmt.Errorf("failed to unmarshal off-chain registry config: %w", err)
	}

	if err := messageConfig.Validate(); err != nil {
		return nil, fmt.Errorf("invalid off-chain registry config: %w", err)
	}
	return &factory{
		registryAddress: common.HexToAddress(messageConfig.TeleporterRegistryAddress),
	}, nil
}

func (f *factory) NewMessageHandler(
	logger logging.Logger,
	unsignedMessage *warp.UnsignedMessage,
	destinationClient vms.DestinationClient,
	signatureAggregator *aggregator.SignatureAggregator,
	metrics messages.Metrics,
	signingSubnetID ids.ID,
	quorumNumerator uint64,
) (messages.MessageHandler, error) {
	logFields := []zap.Field{
		zap.Stringer("warpMessageID", unsignedMessage.ID()),
		zap.Stringer("destinationBlockchainID", destinationClient.DestinationBlockchainID()),
	}
	return &messageHandler{
		logger:              logger.With(logFields...),
		unsignedMessage:     unsignedMessage,
		destinationClient:   destinationClient,
		signatureAggregator: signatureAggregator,
		metrics:             metrics,
		registryAddress:     f.registryAddress,
		signingSubnetID:     signingSubnetID,
		quorumNumerator:     quorumNumerator,
	}, nil
}

func (f *factory) GetMessageRoutingInfo(
	unsignedMessage *warp.UnsignedMessage,
) (messages.MessageRoutingInfo, error) {
	return messages.MessageRoutingInfo{
			SourceChainID:      unsignedMessage.SourceChainID,
			SenderAddress:      OffChainRegistrySourceAddress, // Off-chain registry messages have a zero address as the sender
			DestinationChainID: unsignedMessage.SourceChainID,
			DestinationAddress: f.registryAddress,
		},
		nil
}

// ShouldSendMessage returns false if any contract is already registered as the specified version
// in the TeleporterRegistry contract. This is because a single contract address can be registered
// to multiple versions, but each version may only map to a single contract address.
func (m *messageHandler) ShouldSendMessage() (bool, error) {
	addressedPayload, err := warpPayload.ParseAddressedCall(m.unsignedMessage.Payload)
	if err != nil {
		m.logger.Error(
			"Failed parsing addressed payload",
			zap.Error(err),
		)
		return false, err
	}
	entry, destination, err := teleporterregistry.UnpackTeleporterRegistryWarpPayload(
		addressedPayload.Payload,
	)
	if err != nil {
		m.logger.Error(
			"Failed unpacking teleporter registry warp payload",
			zap.Error(err),
		)
		return false, err
	}
	if destination != m.registryAddress {
		m.logger.Info(
			"Message is not intended for the configured registry",
			zap.Stringer("destination", destination),
			zap.Stringer("configuredRegistry", m.registryAddress),
		)
		return false, nil
	}

	// Check if the version is already registered in the TeleporterRegistry contract.
	registry, err := teleporterregistry.NewTeleporterRegistryCaller(m.registryAddress, m.destinationClient.Client())
	if err != nil {
		m.logger.Error(
			"Failed to create TeleporterRegistry caller",
			zap.Error(err),
		)
		return false, err
	}

	address, err := registry.GetAddressFromVersion(&bind.CallOpts{}, entry.Version)
	if err != nil {
		if strings.Contains(err.Error(), revertVersionNotFoundString) {
			return true, nil
		}
		m.logger.Error(
			"Failed to get address from version",
			zap.Error(err),
		)
		return false, err
	}

	m.logger.Info(
		"Version is already registered in the TeleporterRegistry contract",
		zap.Stringer("version", entry.Version),
		zap.Stringer("registeredAddress", address),
	)
	return false, nil
}

func (m *messageHandler) SendMessage(signedMessage *warp.Message) (common.Hash, error) {
	// Construct the transaction call data to call the TeleporterRegistry contract.
	// Only one off-chain registry Warp message is sent at a time, so we hardcode the index to 0 in the call.
	callData, err := teleporterregistry.PackAddProtocolVersion(0)
	if err != nil {
		m.logger.Error("Failed packing receiveCrossChainMessage call data")
		return common.Hash{}, err
	}

	accessList := utils.SignedWarpMessageToAccessList(signedMessage)

	receipt, err := m.destinationClient.SendTx(
		m.logger,
		accessList,
		nil,
		m.registryAddress,
		addProtocolVersionGasLimit,
		callData,
	)
	if err != nil {
		m.logger.Error(
			"Failed to send tx.",
			zap.Error(err),
		)
		return common.Hash{}, err
	}
	m.logger.Info("Sent message to destination chain")
	return receipt.TxHash, nil
}

// ProcessMessage relays the message to the destination chain by aggregating a signature for it
// and sending it via SendMessage. It does not retry on failure or checkpoint the height.
// Returns the transaction hash if the message is successfully relayed.
func (m *messageHandler) ProcessMessage() (common.Hash, error) {
	m.logger.Info("Relaying message")
	shouldSend, err := m.ShouldSendMessage()
	if err != nil {
		m.metrics.IncFailedRelayMessageCount("failed to check if message should be sent")
		return common.Hash{}, fmt.Errorf("failed to check if message should be sent: %w", err)
	}
	if !shouldSend {
		m.logger.Info("Message should not be sent")
		return common.Hash{}, nil
	}
	unsignedMessage := m.unsignedMessage

	startCreateSignedMessageTime := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), utils.DefaultCreateSignedMessageTimeout)
	defer cancel()

	quorumPercentageBuffer := utils.DefaultQuorumPercentageBuffer(m.quorumNumerator)
	// Determine the appropriate P-Chain height for validator set selection
	pchainHeight, err := m.destinationClient.GetPChainHeightForDestination(ctx)
	if err != nil {
		m.metrics.IncFailedRelayMessageCount("failed to determine P-Chain height")
		return common.Hash{}, fmt.Errorf("failed to determine P-Chain height for validator set: %w", err)
	}

	signedMessage, err := m.signatureAggregator.CreateSignedMessage(
		ctx,
		m.logger,
		unsignedMessage,
		nil,
		m.signingSubnetID,
		m.quorumNumerator,
		quorumPercentageBuffer,
		pchainHeight,
	)
	m.metrics.IncFetchSignatureAppRequestCount()
	if err != nil {
		m.metrics.IncFailedRelayMessageCount("failed to create signed warp message via AppRequest network")
		return common.Hash{}, fmt.Errorf("failed to create signed warp message via AppRequest network: %w", err)
	}

	// create signed message latency (ms)
	m.metrics.SetCreateSignedMessageLatencyMS(float64(time.Since(startCreateSignedMessageTime).Milliseconds()))

	txHash, err := m.SendMessage(signedMessage)
	if err != nil {
		m.metrics.IncFailedRelayMessageCount("failed to send warp message")
		return common.Hash{}, fmt.Errorf("failed to send warp message: %w", err)
	}
	m.logger.Info(
		"Finished relaying message to destination chain",
		zap.Stringer("txID", txHash),
	)
	m.metrics.IncSuccessfulRelayMessageCount()

	return txHash, nil
}
