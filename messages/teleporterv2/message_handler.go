// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporterv2

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
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
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/icm-services/vms"
	"github.com/ava-labs/icm-services/vms/evm"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
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

// messageHandler relays TeleporterV2 messages whose destination messenger verifies them against a
// MerkleValidatorSetRegistry. The attestation passed to receiveCrossChainMessage is a Merkle
// attestation over the validator set committed under the registry's stored root, so the signature
// must be aggregated over the source subnet's set at the committed P-chain height.
type messageHandler struct {
	handlerBase
	pChainClient   clients.CanonicalValidatorState
	sourceSubnetID ids.ID
}

// NewMessageHandlerFactory creates a factory for the TeleporterV2 message protocol. The
// verification path (Merkle registry or WarpAdapter) is selected by the config's verifier-type.
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

	return &factory{
		messageConfig:   messageConfig,
		protocolAddress: messageProtocolAddress,
		pChainClient:    pChainClient,
		sourceSubnetID:  sourceSubnetID,
	}, nil
}

// EventFilter returns the filter matching the Warp messages sent by the TeleporterV2 Warp adapter.
func (f *factory) EventFilter() evm.EventFilter {
	return messages.WarpEventFilter(f.protocolAddress)
}

func (f *factory) NewMessageHandler(
	logger logging.Logger,
	message *messages.SourceMessage,
	destinationClient vms.DestinationClient,
	signatureAggregator *aggregator.SignatureAggregator,
	metrics messages.Metrics,
	signingSubnetID ids.ID,
	quorumNumerator uint64,
) (messages.MessageHandler, error) {
	unsignedMessage, teleporterMessage, err := f.parseMessage(message)
	if err != nil {
		logger.Error(
			"Failed to parse teleporter v2 message.",
			zap.Stringer("sourceTxID", message.SourceTxID),
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

	base := handlerBase{
		logger:              logger.With(logFields...),
		teleporterMessage:   teleporterMessage,
		unsignedMessage:     unsignedMessage,
		destinationClient:   destinationClient,
		signatureAggregator: signatureAggregator,
		metrics:             metrics,
		quorumNumerator:     quorumNumerator,
		teleporterMessageID: teleporterMessageID,
		messageConfig:       f.messageConfig,
		teleporterAddress:   f.messageConfig.teleporterAddress(),
	}

	if f.messageConfig.VerifierType == VerifierTypeWarp {
		return &warpMessageHandler{
			handlerBase:     base,
			signingSubnetID: signingSubnetID,
		}, nil
	}

	return &messageHandler{
		handlerBase:    base,
		pChainClient:   f.pChainClient,
		sourceSubnetID: f.sourceSubnetID,
	}, nil
}

func (f *factory) GetMessageRoutingInfo(
	message *messages.SourceMessage,
) (messages.MessageRoutingInfo, error) {
	unsignedMessage, teleporterMessage, err := f.parseMessage(message)
	if err != nil {
		return messages.MessageRoutingInfo{}, fmt.Errorf("failed to parse teleporter v2 message: %w", err)
	}
	return messages.MessageRoutingInfo{
		SourceChainID:      unsignedMessage.SourceChainID,
		SenderAddress:      teleporterMessage.OriginSenderAddress,
		DestinationChainID: teleporterMessage.DestinationBlockchainID,
		DestinationAddress: teleporterMessage.DestinationAddress,
	}, nil
}

// ProcessMessage relays the message to the destination chain by aggregating a signature over the
// committed validator set and delivering it via SendMessage. It does not retry on failure or
// checkpoint the height. Returns the transaction hash if the message is successfully relayed.
func (m *messageHandler) ProcessMessage() (common.Hash, error) {
	return m.relay(func(ctx context.Context) (common.Hash, error) {
		sourceChainID := m.unsignedMessage.SourceChainID

		// Fetch the validator set commitment stored under the registry's Merkle root. The signature
		// must be aggregated over the exact set (and P-chain height) the root was built from, so the
		// signer bitset and weights match the committed total and the leaves resolve against the
		// stored root.
		commitment, err := m.fetchCommitment(ctx, sourceChainID)
		if err != nil {
			m.metrics.IncFailedRelayMessageCount("failed to read committed validator set")
			m.logger.Error("Failed to read committed validator set", zap.Error(err))
			return common.Hash{}, err
		}

		// Fetch the canonical warp validator set at the committed P-chain height. This is the same
		// representation the committed Merkle root was built from and the signature aggregator's
		// signer bitset indexes into: validators without a BLS key are omitted and validators
		// sharing a BLS key are merged into a single entry.
		validators, err := validatorupdater.FetchCanonicalValidators(
			ctx, m.pChainClient, m.sourceSubnetID, commitment.PChainHeight,
		)
		if err != nil {
			m.metrics.IncFailedRelayMessageCount("failed to fetch committed validators")
			m.logger.Error("Failed to fetch committed validator set",
				zap.Uint64("pChainHeight", commitment.PChainHeight),
				zap.Error(err),
			)
			return common.Hash{}, err
		}

		signedMessage, err := m.signMessage(ctx, m.sourceSubnetID, commitment.PChainHeight)
		if err != nil {
			return common.Hash{}, err
		}

		txHash, err := m.SendMessage(ctx, signedMessage, validators)
		if err != nil {
			m.metrics.IncFailedRelayMessageCount("failed to send warp message")
			return common.Hash{}, fmt.Errorf("failed to send warp message: %w", err)
		}
		return txHash, nil
	})
}

// SendMessage builds a Merkle attestation for the signed message and delivers it to the
// destination TeleporterMessengerV2, whose verifier is a MerkleValidatorSetRegistry. The
// [validators] must be the committed set the signature was aggregated over.
func (m *messageHandler) SendMessage(
	ctx context.Context,
	signedMessage *warp.Message,
	validators []*validatorupdater.Validator,
) (common.Hash, error) {
	bitSetSig, ok := signedMessage.Signature.(*avalancheWarp.BitSetSignature)
	if !ok {
		return common.Hash{}, fmt.Errorf("expected BitSetSignature, got %T", signedMessage.Signature)
	}

	attestation, err := validatorupdater.NewValidatorSetMerkleAttestation(validators, bitSetSig)
	if err != nil {
		m.logger.Error("Failed to build Merkle attestation", zap.Error(err))
		return common.Hash{}, fmt.Errorf("failed to build merkle attestation: %w", err)
	}

	callData, err := teleportermessengerv2.PackReceiveCrossChainMessageMerkle(
		*m.teleporterMessage,
		m.unsignedMessage.NetworkID,
		m.unsignedMessage.SourceChainID,
		attestation.Bytes(),
		m.messageConfig.rewardAddress(),
	)
	if err != nil {
		m.logger.Error("Failed to pack receiveCrossChainMessage call data", zap.Error(err))
		return common.Hash{}, err
	}

	gasLimit, err := m.estimateGasLimit(ctx, callData)
	if err != nil {
		// A non-retryable failure is already described by its error and is logged again when the
		// relayer abandons the message, so it is not reported as an error here.
		if !errors.Is(err, messages.ErrNonRetryable) {
			m.logger.Error("Failed to estimate gas limit", zap.Error(err))
		}
		return common.Hash{}, err
	}

	// No access list: verification reads the attestation from calldata, not a predicate.
	return m.sendTxAndConfirm(nil, gasLimit, callData)
}

// fetchCommitment reads the registry's stored validator set commitment for the source chain,
// which pins the P-chain height the committed Merkle root was built from.
func (m *messageHandler) fetchCommitment(
	ctx context.Context,
	sourceChainID ids.ID,
) (merkleregistry.ValidatorSetMerkleCommitment, error) {
	registry, err := merkleregistry.NewMerkleValidatorSetRegistry(
		m.messageConfig.registryAddress(),
		m.destinationClient.Client(),
	)
	if err != nil {
		return merkleregistry.ValidatorSetMerkleCommitment{}, fmt.Errorf("failed to bind merkle registry: %w", err)
	}

	commitment, err := registry.GetValidatorSetCommitment(&bind.CallOpts{Context: ctx}, sourceChainID)
	if err != nil {
		return merkleregistry.ValidatorSetMerkleCommitment{}, fmt.Errorf("failed to read committed validator set: %w", err)
	}
	if commitment.TotalWeight == 0 {
		return merkleregistry.ValidatorSetMerkleCommitment{},
			fmt.Errorf("no validator set registered for source chain %s", sourceChainID)
	}
	return commitment, nil
}

// estimateGasLimit estimates the gas for the receiveCrossChainMessage call and applies a safety
// buffer, bounded by the configured block gas limit. Estimation is the only simulation performed
// before the delivery is signed and broadcast, so it never resolves to a gas limit the estimate has
// not shown to be sufficient.
func (m *messageHandler) estimateGasLimit(ctx context.Context, callData []byte) (uint64, error) {
	from := m.selectSenderAddress()
	blockGasLimit := m.destinationClient.BlockGasLimit()
	// Bound the node's search at the gas the delivery transaction can actually carry, so a
	// successful estimate is one we can honor.
	estimated, err := m.destinationClient.Client().EstimateGas(ctx, ethereum.CallMsg{
		From: from,
		To:   &m.teleporterAddress,
		Gas:  blockGasLimit,
		Data: callData,
	})
	if err != nil {
		if isUndeliverableEstimateError(err) {
			return 0, fmt.Errorf("%w: %w", messages.ErrDeliveryExceedsBlockGasLimit, err)
		}
		return 0, fmt.Errorf("failed to estimate gas: %w", err)
	}
	if blockGasLimit != 0 && estimated > blockGasLimit {
		return 0, fmt.Errorf(
			"%w: estimated gas %d exceeds block gas limit %d",
			messages.ErrDeliveryExceedsBlockGasLimit, estimated, blockGasLimit,
		)
	}
	buffered := estimated * gasLimitBufferNumerator / gasLimitBufferDenominator
	// Capping is safe only because the unbuffered estimate already fits under the limit, so the
	// capped transaction still carries enough gas to execute.
	if blockGasLimit != 0 && buffered > blockGasLimit {
		buffered = blockGasLimit
	}
	return buffered, nil
}

// undeliverableEstimateErrors are the eth_estimateGas failures that mean the delivery cannot
// succeed at any gas value within the block gas limit.
var undeliverableEstimateErrors = []string{
	// The node exhausted its search range without the call succeeding.
	"gas required exceeds allowance",
	"always failing transaction",
	// TeleporterMessengerV2's revert reason when gasleft() is below requiredGasLimit.
	"insufficient gas",
}

// isUndeliverableEstimateError distinguishes conclusive gas-envelope failures from errors that may
// clear on a retry. Reverts that are not gas related, such as a verification failure caused by the
// destination registry's committed validator set changing mid-delivery, are treated as retryable.
func isUndeliverableEstimateError(err error) bool {
	msg := strings.ToLower(err.Error())
	for _, undeliverable := range undeliverableEstimateErrors {
		if strings.Contains(msg, undeliverable) {
			return true
		}
	}
	return false
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

// parseMessage decodes the TeleporterV2 message that [message] was sent as. Both adapters send
// their messages through the Warp precompile, so the source message payload is the encoded
// unsigned Warp message that carries the Teleporter message as its addressed payload. The two
// adapters encode that payload differently: the Merkle registry adapter uses the packed
// serializeTeleporterMessageV2 layout, while the WarpAdapter uses abi.encode.
func (f *factory) parseMessage(
	message *messages.SourceMessage,
) (*warp.UnsignedMessage, *teleportermessengerv2.TeleporterMessageV2, error) {
	unsignedMessage, err := utils.UnpackWarpMessage(message.Payload)
	if err != nil {
		return nil, nil, fmt.Errorf("failed parsing unsigned warp message: %w", err)
	}
	addressedPayload, err := warpPayload.ParseAddressedCall(unsignedMessage.Payload)
	if err != nil {
		return nil, nil, fmt.Errorf("failed parsing addressed payload: %w", err)
	}

	var teleporterMessage *teleportermessengerv2.TeleporterMessageV2
	if f.messageConfig.VerifierType == VerifierTypeWarp {
		teleporterMessage = &teleportermessengerv2.TeleporterMessageV2{}
		if err := teleporterMessage.Unpack(addressedPayload.Payload); err != nil {
			return nil, nil, fmt.Errorf("failed unpacking teleporter v2 message: %w", err)
		}
	} else {
		teleporterMessage, err = ParseTeleporterMessageV2(addressedPayload.Payload)
		if err != nil {
			return nil, nil, err
		}
	}
	return unsignedMessage, teleporterMessage, nil
}
