// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporterv2

import (
	"context"
	"fmt"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/evm/predicate"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	warpadapter "github.com/ava-labs/icm-services/abi-bindings/go/teleporterV2/WarpAdapter"
	gasUtils "github.com/ava-labs/icm-services/icm-contracts/utils/gas-utils"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/icm-services/vms"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"go.uber.org/zap"
)

// warpMessageHandler relays TeleporterV2 messages whose destination messenger uses the
// WarpAdapter for verification. The signed warp message is included as a transaction predicate
// and the attestation passed to receiveCrossChainMessage is the index of that message in the
// predicate, which the WarpAdapter reads back through the warp precompile.
type warpMessageHandler struct {
	logger              logging.Logger
	teleporterMessage   *teleportermessengerv2.TeleporterMessageV2
	unsignedMessage     *warp.UnsignedMessage
	destinationClient   vms.DestinationClient
	signatureAggregator *aggregator.SignatureAggregator
	metrics             messages.Metrics
	signingSubnetID     ids.ID
	quorumNumerator     uint64
	// teleporterMessageID identifies the message for delivery/replay checks.
	teleporterMessageID ids.ID
	messageConfig       *Config
	// teleporterAddress is the TeleporterMessengerV2 contract (receive target + message ID + status).
	teleporterAddress common.Address
}

// ShouldSendMessage returns true if the message should be relayed to the destination chain.
func (m *warpMessageHandler) ShouldSendMessage() (bool, error) {
	// RequiredGasLimit is a Solidity uint256 (*big.Int). Calling Uint64() on a value that does not
	// fit in 64 bits is undefined, so treat any non-uint64 value as exceeding the block gas limit.
	destBlockGasLimit := m.destinationClient.BlockGasLimit()
	if !m.teleporterMessage.RequiredGasLimit.IsUint64() ||
		m.teleporterMessage.RequiredGasLimit.Uint64() > destBlockGasLimit {
		m.logger.Info(
			"Gas limit exceeds maximum threshold",
			zap.Stringer("requiredGasLimit", m.teleporterMessage.RequiredGasLimit),
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

// ProcessMessage relays the message to the destination chain by aggregating a signature for it
// and sending it via SendMessage. It does not retry on failure or checkpoint the height.
// Returns the transaction hash if the message is successfully relayed.
func (m *warpMessageHandler) ProcessMessage() (common.Hash, error) {
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

	ctx, cancel := context.WithTimeout(context.Background(), utils.DefaultCreateSignedMessageTimeout)
	defer cancel()

	// Determine the appropriate P-Chain height for validator set selection
	pchainHeight, err := m.destinationClient.GetPChainHeightForDestination(ctx)
	if err != nil {
		m.metrics.IncFailedRelayMessageCount("failed to determine P-Chain height")
		return common.Hash{}, fmt.Errorf("failed to determine P-Chain height for validator set: %w", err)
	}

	startCreateSignedMessageTime := time.Now()
	signedMessage, err := m.signatureAggregator.CreateSignedMessage(
		ctx,
		m.logger,
		m.unsignedMessage,
		nil,
		m.signingSubnetID,
		m.quorumNumerator,
		pchainHeight,
	)
	m.metrics.IncFetchSignatureAppRequestCount()
	if err != nil {
		m.metrics.IncFailedRelayMessageCount("failed to create signed warp message via AppRequest network")
		return common.Hash{}, fmt.Errorf("failed to create signed warp message via AppRequest network: %w", err)
	}
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

// SendMessage delivers the signed message to the destination TeleporterMessengerV2, attaching it
// as a transaction predicate for the WarpAdapter to verify through the warp precompile.
func (m *warpMessageHandler) SendMessage(signedMessage *warp.Message) (common.Hash, error) {
	m.logger.Info("Sending message to destination chain")
	numSigners, err := signedMessage.Signature.NumSigners()
	if err != nil {
		m.logger.Error("Failed to get number of signers")
		return common.Hash{}, err
	}

	gasLimit, err := gasUtils.CalculateReceiveMessageGasLimit(
		numSigners,
		m.teleporterMessage.RequiredGasLimit,
		len(predicate.New(signedMessage.Bytes())),
		len(signedMessage.Payload),
		len(m.teleporterMessage.Receipts),
	)
	if err != nil {
		m.logger.Error("Failed to calculate gas limit for receiveCrossChainMessage call")
		return common.Hash{}, err
	}

	// The signed warp message is the only message attached to the transaction predicate, so the
	// attestation the WarpAdapter verifies against is index 0.
	callData, err := warpadapter.PackReceiveCrossChainMessage(
		*m.teleporterMessage,
		signedMessage.SourceChainID,
		0,
		m.messageConfig.rewardAddress(),
	)
	if err != nil {
		m.logger.Error("Failed packing receiveCrossChainMessage call data")
		return common.Hash{}, err
	}

	accessList := utils.SignedWarpMessageToAccessList(signedMessage)

	receipt, err := m.destinationClient.SendTx(
		m.logger,
		accessList,
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
	if receipt.Status == types.ReceiptStatusSuccessful {
		log.Info("Delivered message to destination chain")
		return txHash, nil
	}

	// The transaction reverted. A common benign cause is that the message was already delivered
	// by another relayer, so check delivery status before treating the revert as a failure.
	teleporterMessenger, err := m.getTeleporterMessenger()
	if err != nil {
		log.Error("Transaction failed", zap.Error(err))
		return common.Hash{}, fmt.Errorf("transaction failed with status: %d", receipt.Status)
	}

	delivered, err := teleporterMessenger.MessageReceived(&bind.CallOpts{}, m.teleporterMessageID)
	if err != nil {
		log.Error("Transaction failed", zap.Error(err))
		return common.Hash{}, fmt.Errorf("transaction failed with status: %d", receipt.Status)
	}
	if delivered {
		log.Info("Execution reverted: message already delivered to destination.")
		return txHash, nil
	}

	log.Error("Transaction failed")
	return common.Hash{}, fmt.Errorf("transaction failed with status: %d", receipt.Status)
}

func (m *warpMessageHandler) getTeleporterMessenger() (*teleportermessengerv2.TeleporterMessengerV2, error) {
	messenger, err := teleportermessengerv2.NewTeleporterMessengerV2(m.teleporterAddress, m.destinationClient.Client())
	if err != nil {
		return nil, fmt.Errorf("failed to bind teleporter v2 messenger: %w", err)
	}
	return messenger, nil
}
