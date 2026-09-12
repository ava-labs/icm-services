// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporterv2

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/icm-services/vms"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
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

// relay runs the send/skip decision and the surrounding metrics and logging that are common to
// both verification paths, delegating the signing and delivery of the message to [deliver].
// It does not retry on failure or checkpoint the height. Returns the transaction hash if the
// message is successfully relayed, or the zero hash if it should not be sent.
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

// sendTxAndConfirm delivers [callData] to the destination TeleporterMessengerV2 and resolves the
// receipt. A reverted transaction is not necessarily a failure: another relayer may have delivered
// the message first, so delivery status is checked before the revert is reported as an error.
func (m *handlerBase) sendTxAndConfirm(
	accessList types.AccessList,
	gasLimit uint64,
	callData []byte,
) (common.Hash, error) {
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

func (m *handlerBase) getTeleporterMessenger() (*teleportermessengerv2.TeleporterMessengerV2, error) {
	messenger, err := teleportermessengerv2.NewTeleporterMessengerV2(m.teleporterAddress, m.destinationClient.Client())
	if err != nil {
		return nil, fmt.Errorf("failed to bind teleporter v2 messenger: %w", err)
	}
	return messenger, nil
}

func isAllowedRelayer(allowedRelayers []common.Address, eoa common.Address) bool {
	if len(allowedRelayers) == 0 {
		return true
	}
	return slices.Contains(allowedRelayers, eoa)
}

func containsAllowedRelayer(allowedRelayers []common.Address, eoas []common.Address) bool {
	if len(allowedRelayers) == 0 {
		return true
	}
	return slices.ContainsFunc(eoas, func(eoa common.Address) bool {
		return slices.Contains(allowedRelayers, eoa)
	})
}
