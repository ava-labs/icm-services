// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporterv2

import (
	"context"
	"fmt"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/icm-services/vms"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"go.uber.org/zap"
)

// handlerBase holds the state and behavior shared by the TeleporterV2 handlers. Both verification
// paths relay the same TeleporterMessageV2 to the same destination TeleporterMessengerV2 and make
// the same send/skip decisions; they differ only in which validator set the signature is
// aggregated over and how the attestation is presented to the destination adapter.
type handlerBase struct {
	logger              logging.Logger
	teleporterMessage   *teleportermessengerv2.TeleporterMessageV2
	unsignedMessage     *warp.UnsignedMessage
	destinationClient   vms.DestinationClient
	signatureAggregator *aggregator.SignatureAggregator
	metrics             messages.Metrics
	quorumNumerator     uint64
	// teleporterMessageID identifies the message for delivery/replay checks.
	teleporterMessageID ids.ID
	messageConfig       *Config
	// teleporterAddress is the TeleporterMessengerV2 contract (receive target + message ID + status).
	teleporterAddress common.Address
}

// ShouldSendMessage returns true if the message should be relayed to the destination chain.
func (m *handlerBase) ShouldSendMessage() (bool, error) {
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

	if !messages.ContainsAllowedRelayer(
		m.teleporterMessage.AllowedRelayerAddresses,
		m.destinationClient.SenderAddresses(),
	) {
		m.logger.Info("Relayer EOA not allowed to deliver this message.")
		return false, nil
	}

	delivered, err := m.messageDelivered()
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

// relay runs the send/skip decision and the surrounding metrics and logging that are common to
// both verification paths, delegating the signing and delivery of the message to [deliver].
// It does not retry on failure or checkpoint the height. Returns the transaction hash if the
// message is successfully relayed, or the zero hash if it should not be sent or [deliver] found
// nothing left to send.
func (m *handlerBase) relay(deliver func(ctx context.Context) (common.Hash, error)) (common.Hash, error) {
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

	txHash, err := deliver(ctx)
	if err != nil {
		return common.Hash{}, err
	}
	if txHash == (common.Hash{}) {
		// Nothing was sent: the message turned out to be delivered already while sending it.
		return common.Hash{}, nil
	}
	m.logger.Info(
		"Finished relaying message to destination chain",
		zap.Stringer("txID", txHash),
	)
	m.metrics.IncSuccessfulRelayMessageCount()
	return txHash, nil
}

// signMessage aggregates a signature over the message from the [signingSubnetID] validator set at
// [pChainHeight], recording the request count and latency metrics.
func (m *handlerBase) signMessage(
	ctx context.Context,
	signingSubnetID ids.ID,
	pChainHeight uint64,
) (*warp.Message, error) {
	startCreateSignedMessageTime := time.Now()
	signedMessage, err := m.signatureAggregator.CreateSignedMessage(
		ctx,
		m.logger,
		m.unsignedMessage,
		nil,
		signingSubnetID,
		m.quorumNumerator,
		pChainHeight,
	)
	m.metrics.IncFetchSignatureAppRequestCount()
	if err != nil {
		m.metrics.IncFailedRelayMessageCount("failed to create signed warp message via AppRequest network")
		return nil, fmt.Errorf("failed to create signed warp message via AppRequest network: %w", err)
	}
	m.metrics.SetCreateSignedMessageLatencyMS(float64(time.Since(startCreateSignedMessageTime).Milliseconds()))
	return signedMessage, nil
}

// sender delivers the message to the destination TeleporterMessengerV2.
func (m *handlerBase) sender() *messages.Sender {
	return &messages.Sender{
		// Receipts drive the gas a delivery needs, so they belong in every delivery log line.
		Logger:            m.logger.With(zap.Int("numReceipts", len(m.teleporterMessage.Receipts))),
		DestinationClient: m.destinationClient,
		To:                m.teleporterAddress,
		AllowedRelayers:   m.teleporterMessage.AllowedRelayerAddresses,
		Delivered:         m.messageDelivered,
	}
}

// messageDelivered reports whether the destination TeleporterMessengerV2 has received the message.
func (m *handlerBase) messageDelivered() (bool, error) {
	teleporterMessenger, err := m.getTeleporterMessenger()
	if err != nil {
		return false, err
	}
	return teleporterMessenger.MessageReceived(&bind.CallOpts{}, m.teleporterMessageID)
}

func (m *handlerBase) getTeleporterMessenger() (*teleportermessengerv2.TeleporterMessengerV2, error) {
	messenger, err := teleportermessengerv2.NewTeleporterMessengerV2(m.teleporterAddress, m.destinationClient.Client())
	if err != nil {
		return nil, fmt.Errorf("failed to bind teleporter v2 messenger: %w", err)
	}
	return messenger, nil
}
