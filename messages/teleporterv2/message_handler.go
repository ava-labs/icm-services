// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporterv2

import (
	"context"
	"fmt"
	"slices"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	warpPayload "github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	merkleregistry "github.com/ava-labs/icm-services/abi-bindings/go/MerkleValidatorSetRegistry"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	teleporterUtils "github.com/ava-labs/icm-services/icm-contracts/utils/teleporter-utils"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/peers/clients"
	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/icm-services/relayer/validatorupdater"
	"github.com/ava-labs/icm-services/vms"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"go.uber.org/zap"
)

const (
	// gasLimitBufferNumerator/Denominator adds a safety margin on top of the estimated gas, since
	// Merkle verification gas (software BLS, multi-proof) is sensitive to the signer set.
	gasLimitBufferNumerator   = 5
	gasLimitBufferDenominator = 4
)

type factory struct {
	messageConfig   *Config
	protocolAddress common.Address
	pChainClient    clients.CanonicalValidatorState
	sourceSubnetID  ids.ID
}

type messageHandler struct {
	logger            logging.Logger
	teleporterMessage *teleportermessengerv2.TeleporterMessageV2
	unsignedMessage   *warp.UnsignedMessage
	destinationClient vms.DestinationClient
	pChainClient      clients.CanonicalValidatorState
	sourceSubnetID    ids.ID
	// teleporterMessageID identifies the message for delivery/replay checks.
	teleporterMessageID ids.ID
	messageConfig       *Config
	// teleporterAddress is the TeleporterMessengerV2 contract (receive target + message ID + status).
	teleporterAddress common.Address
	logFields         []zap.Field
}

// NewMessageHandlerFactory creates a factory for the TeleporterV2 Merkle verification path.
func NewMessageHandlerFactory(
	messageProtocolAddress common.Address,
	messageProtocolConfig config.MessageProtocolConfig,
	pChainClient clients.CanonicalValidatorState,
	sourceSubnetID ids.ID,
) (messages.MessageHandlerFactory, error) {
	messageConfig, err := ConfigFromMap(messageProtocolConfig.Settings)
	if err != nil {
		return nil, fmt.Errorf("invalid teleporter v2 config: %w", err)
	}
	if pChainClient == nil {
		return nil, fmt.Errorf("teleporter v2 merkle handler requires a P-Chain client")
	}

	return &factory{
		messageConfig:   messageConfig,
		protocolAddress: messageProtocolAddress,
		pChainClient:    pChainClient,
		sourceSubnetID:  sourceSubnetID,
	}, nil
}

func (f *factory) NewMessageHandler(
	logger logging.Logger,
	unsignedMessage *warp.UnsignedMessage,
	destinationClient vms.DestinationClient,
) (messages.MessageHandler, error) {
	teleporterMessage, err := parseTeleporterMessage(unsignedMessage)
	if err != nil {
		logger.Error(
			"Failed to parse teleporter v2 message.",
			zap.Stringer("warpMessageID", unsignedMessage.ID()),
			zap.Error(err),
		)
		return nil, err
	}

	teleporterMessageID, err := teleporterUtils.CalculateMessageID(
		f.messageConfig.teleporterAddress(),
		unsignedMessage.SourceChainID,
		teleporterMessage.DestinationBlockchainID,
		teleporterMessage.MessageNonce,
	)
	if err != nil {
		logger.Error(
			"Failed to calculate Teleporter v2 message ID.",
			zap.Stringer("warpMessageID", unsignedMessage.ID()),
			zap.Error(err),
		)
		return nil, err
	}

	logFields := []zap.Field{
		zap.Stringer("warpMessageID", unsignedMessage.ID()),
		zap.Stringer("teleporterMessageID", teleporterMessageID),
		zap.Stringer("destinationBlockchainID", ids.ID(teleporterMessage.DestinationBlockchainID)),
		zap.Stringer("adapterAddress", f.protocolAddress),
		zap.Stringer("teleporterAddress", f.messageConfig.teleporterAddress()),
	}
	return &messageHandler{
		logger:              logger.With(logFields...),
		teleporterMessage:   teleporterMessage,
		unsignedMessage:     unsignedMessage,
		destinationClient:   destinationClient,
		pChainClient:        f.pChainClient,
		sourceSubnetID:      f.sourceSubnetID,
		teleporterMessageID: teleporterMessageID,
		messageConfig:       f.messageConfig,
		teleporterAddress:   f.messageConfig.teleporterAddress(),
		logFields:           logFields,
	}, nil
}

func (f *factory) GetMessageRoutingInfo(
	unsignedMessage *warp.UnsignedMessage,
) (messages.MessageRoutingInfo, error) {
	teleporterMessage, err := parseTeleporterMessage(unsignedMessage)
	if err != nil {
		return messages.MessageRoutingInfo{}, fmt.Errorf("failed to parse teleporter v2 message: %w", err)
	}
	return messages.MessageRoutingInfo{
		SourceChainID:      unsignedMessage.SourceChainID,
		SenderAddress:      teleporterMessage.OriginSenderAddress,
		DestinationChainID: ids.ID(teleporterMessage.DestinationBlockchainID),
		DestinationAddress: teleporterMessage.DestinationAddress,
	}, nil
}

func (m *messageHandler) GetUnsignedMessage() *warp.UnsignedMessage {
	return m.unsignedMessage
}

func (m *messageHandler) LoggerWithContext(logger logging.Logger) logging.Logger {
	return logger.With(m.logFields...)
}

// ShouldSendMessage returns true if the message should be relayed to the destination chain.
func (m *messageHandler) ShouldSendMessage() (bool, error) {
	requiredGasLimit := m.teleporterMessage.RequiredGasLimit.Uint64()
	destBlockGasLimit := m.destinationClient.BlockGasLimit()
	if requiredGasLimit > destBlockGasLimit {
		m.logger.Info(
			"Gas limit exceeds maximum threshold",
			zap.Uint64("requiredGasLimit", requiredGasLimit),
			zap.Uint64("blockGasLimit", destBlockGasLimit),
		)
		return false, nil
	}

	if !containsAllowedRelayer(m.teleporterMessage.AllowedRelayerAddresses, m.destinationClient.SenderAddresses()) {
		m.logger.Info("Relayer EOA not allowed to deliver this message.")
		return false, nil
	}

	teleporterMessenger, err := m.getTeleporterMessenger()
	if err != nil {
		return false, err
	}
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

	return true, nil
}

// SendMessage builds a Merkle attestation for the signed message and delivers it to the
// destination TeleporterMessengerV2, whose verifier is a MerkleValidatorSetRegistry.
func (m *messageHandler) SendMessage(signedMessage *warp.Message) (common.Hash, error) {
	ctx := context.Background()
	sourceChainID := m.unsignedMessage.SourceChainID

	bitSetSig, ok := signedMessage.Signature.(*avalancheWarp.BitSetSignature)
	if !ok {
		return common.Hash{}, fmt.Errorf("expected BitSetSignature, got %T", signedMessage.Signature)
	}

	// Fetch the validator set committed under the registry's stored Merkle root. The relayer must
	// build the multi-proof against the exact set (and P-chain height) the root was built from, so
	// the leaves resolve against the stored root and weights match the committed total.
	validators, err := m.fetchCommittedValidators(ctx, sourceChainID)
	if err != nil {
		m.logger.Error("Failed to fetch committed validator set", zap.Error(err))
		return common.Hash{}, err
	}

	attestation, err := validatorupdater.NewValidatorSetMerkleAttestation(validators, bitSetSig)
	if err != nil {
		m.logger.Error("Failed to build Merkle attestation", zap.Error(err))
		return common.Hash{}, fmt.Errorf("failed to build merkle attestation: %w", err)
	}

	callData, err := teleportermessengerv2.PackReceiveCrossChainMessageMerkle(
		*m.teleporterMessage,
		m.unsignedMessage.NetworkID,
		sourceChainID,
		attestation.Bytes(),
		m.messageConfig.rewardAddress(),
	)
	if err != nil {
		m.logger.Error("Failed to pack receiveCrossChainMessage call data", zap.Error(err))
		return common.Hash{}, err
	}

	gasLimit, err := m.estimateGasLimit(ctx, callData)
	if err != nil {
		m.logger.Error("Failed to estimate gas limit", zap.Error(err))
		return common.Hash{}, err
	}

	receipt, err := m.destinationClient.SendTx(
		m.logger,
		nil, // No access list: verification reads the attestation from calldata, not a predicate.
		set.Of(m.teleporterMessage.AllowedRelayerAddresses...),
		m.teleporterAddress,
		gasLimit,
		callData,
	)
	if err != nil {
		m.logger.Error("Failed to send tx.", zap.Error(err))
		return common.Hash{}, err
	}

	txHash := receipt.TxHash
	log := m.logger.With(zap.Stringer("txID", txHash))
	if receipt.Status != types.ReceiptStatusSuccessful {
		teleporterMessenger, msgErr := m.getTeleporterMessenger()
		if msgErr == nil {
			delivered, derr := teleporterMessenger.MessageReceived(&bind.CallOpts{}, m.teleporterMessageID)
			if derr == nil && delivered {
				log.Info("Execution reverted: message already delivered to destination.")
				return txHash, nil
			}
		}
		log.Error("Transaction failed")
		return common.Hash{}, fmt.Errorf("transaction failed with status: %d", receipt.Status)
	}

	log.Info("Delivered message to destination chain")
	return txHash, nil
}

// fetchCommittedValidators reads the registry's committed P-chain height for the source chain and
// returns the source subnet's validator set at that height, sorted by BLS public key to match the
// canonical ordering used to build the committed Merkle root and the signer bitset.
func (m *messageHandler) fetchCommittedValidators(
	ctx context.Context,
	sourceChainID ids.ID,
) ([]*validatorupdater.Validator, error) {
	registry, err := merkleregistry.NewMerkleValidatorSetRegistry(
		m.messageConfig.registryAddress(),
		m.destinationClient.Client(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to bind merkle registry: %w", err)
	}

	commitment, err := registry.GetValidatorSetCommitment(&bind.CallOpts{Context: ctx}, sourceChainID)
	if err != nil {
		return nil, fmt.Errorf("failed to read committed validator set: %w", err)
	}
	if commitment.TotalWeight == 0 {
		return nil, fmt.Errorf("no validator set registered for source chain %s", sourceChainID)
	}

	subnetValidators, err := m.pChainClient.GetValidatorsAt(ctx, m.sourceSubnetID, commitment.PChainHeight)
	if err != nil {
		return nil, fmt.Errorf("failed to get validators at height %d: %w", commitment.PChainHeight, err)
	}

	validatorList := make([]*validatorupdater.Validator, 0, len(subnetValidators))
	for _, vdr := range subnetValidators {
		if vdr.PublicKey == nil {
			continue
		}
		validatorList = append(validatorList, &validatorupdater.Validator{
			UncompressedPublicKeyBytes: [96]byte(vdr.PublicKey.Serialize()),
			Weight:                     vdr.Weight,
		})
	}
	validatorupdater.SortValidators(validatorList)
	return validatorList, nil
}

// estimateGasLimit estimates the gas for the receiveCrossChainMessage call and applies a safety
// buffer. Falls back to the configured block gas limit if estimation fails.
func (m *messageHandler) estimateGasLimit(ctx context.Context, callData []byte) (uint64, error) {
	from := m.selectSenderAddress()
	estimated, err := m.destinationClient.Client().EstimateGas(ctx, ethereum.CallMsg{
		From: from,
		To:   &m.teleporterAddress,
		Data: callData,
	})
	if err != nil {
		blockGasLimit := m.destinationClient.BlockGasLimit()
		m.logger.Warn(
			"Gas estimation failed, falling back to block gas limit",
			zap.Error(err),
			zap.Uint64("blockGasLimit", blockGasLimit),
		)
		if blockGasLimit == 0 {
			return 0, fmt.Errorf("failed to estimate gas and no block gas limit configured: %w", err)
		}
		return blockGasLimit, nil
	}
	buffered := estimated * gasLimitBufferNumerator / gasLimitBufferDenominator
	if blockGasLimit := m.destinationClient.BlockGasLimit(); blockGasLimit != 0 && buffered > blockGasLimit {
		buffered = blockGasLimit
	}
	return buffered, nil
}

// selectSenderAddress picks a relayer EOA eligible to deliver the message for gas estimation.
func (m *messageHandler) selectSenderAddress() common.Address {
	senders := m.destinationClient.SenderAddresses()
	for _, sender := range senders {
		if isAllowedRelayer(m.teleporterMessage.AllowedRelayerAddresses, sender) {
			return sender
		}
	}
	if len(senders) > 0 {
		return senders[0]
	}
	return common.Address{}
}

func (m *messageHandler) getTeleporterMessenger() (*teleportermessengerv2.TeleporterMessengerV2, error) {
	messenger, err := teleportermessengerv2.NewTeleporterMessengerV2(m.teleporterAddress, m.destinationClient.Client())
	if err != nil {
		return nil, fmt.Errorf("failed to bind teleporter v2 messenger: %w", err)
	}
	return messenger, nil
}

func parseTeleporterMessage(
	unsignedMessage *warp.UnsignedMessage,
) (*teleportermessengerv2.TeleporterMessageV2, error) {
	addressedPayload, err := warpPayload.ParseAddressedCall(unsignedMessage.Payload)
	if err != nil {
		return nil, fmt.Errorf("failed parsing addressed payload: %w", err)
	}
	return ParseTeleporterMessageV2(addressedPayload.Payload)
}

func isAllowedRelayer(allowedRelayers []common.Address, eoa common.Address) bool {
	if len(allowedRelayers) == 0 {
		return true
	}
	return slices.Contains(allowedRelayers, eoa)
}

func containsAllowedRelayer(allowedRelayers []common.Address, eoas []common.Address) bool {
	for _, eoa := range eoas {
		if isAllowedRelayer(allowedRelayers, eoa) {
			return true
		}
	}
	return false
}
