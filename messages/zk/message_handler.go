// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package zk implements the message handler for Teleporter messages
// originating on Ethereum and verified on an Avalanche destination via the
// ZKAdapter (an IAdapter implementation), using Boundless zero-knowledge proofs
// of Ethereum consensus.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package zk

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	zkadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKAdapter"
	teleporterUtils "github.com/ava-labs/icm-services/icm-contracts/utils/teleporter-utils"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/messages/teleporterv2"
	"github.com/ava-labs/icm-services/messages/zk/proofs"
	"github.com/ava-labs/icm-services/messages/zk/proofs/beacon"
	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator"
	"github.com/ava-labs/icm-services/vms"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/ethclient"
	"go.uber.org/zap"
)

const (
	// processTimeout bounds one delivery attempt end to end: the anchor scan, two beacon state
	// fetches of ~100-200MB each, and the destination transaction.
	processTimeout = 15 * time.Minute

	// anchorScanTimeout bounds the backward ConfirmedBeaconBlock scan on its own, so a scan
	// that finds no anchor and walks toward genesis cannot use up the proof and send budget.
	anchorScanTimeout = 2 * time.Minute

	// anchorPollInterval is how often the handler re-scans for a confirmed anchor when none
	// covers the message's slot yet. New anchors land at the consensus proof cadence, so the
	// handler waits for the next transition instead of failing the attempt.
	anchorPollInterval = 30 * time.Second
)

type factory struct {
	messageConfig   *Config
	protocolAddress common.Address
	sourceClient    *ethclient.Client
	// proofBuilder is shared by all handlers so parsed beacon states are cached across messages.
	proofBuilder *proofs.ProofBuilder
}

type messageHandler struct {
	logger            logging.Logger
	sourceClient      *ethclient.Client
	destinationClient vms.DestinationClient
	metrics           messages.Metrics
	messageConfig     *Config
	proofBuilder      *proofs.ProofBuilder
	// sourceMessage contains the protocol address (the adapter).
	sourceMessage     messages.SourceMessage
	teleporterMessage *teleportermessengerv2.TeleporterMessageV2
	sourceBlockNumber uint64
	sourceLogIndex    uint
	// sourceReceiptsRoot is the receipts root of the message's source block, which the execution
	// proof must resolve to.
	sourceReceiptsRoot common.Hash
	slot               uint64
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
		proofBuilder:    proofs.NewProofBuilder(beacon.NewBeaconClient(messageConfig.BeaconAPIURL)),
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

	teleporterMessage, err := parseMessage(message)
	if err != nil {
		return nil, fmt.Errorf("failed to parse teleporter v2 message: %w", err)
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
		sourceClient:       f.sourceClient,
		destinationClient:  destinationClient,
		metrics:            metrics,
		messageConfig:      f.messageConfig,
		proofBuilder:       f.proofBuilder,
		sourceMessage:      *message,
		teleporterMessage:  teleporterMessage,
		sourceBlockNumber:  blockNumber,
		sourceLogIndex:     logIndex,
		sourceReceiptsRoot: header.ReceiptHash,
		slot:               slot,
	}, nil
}

// ProcessMessage delivers the message to the destination via the TeleporterV2 messenger,
// attaching a zk attestation. It does not retry on failure or checkpoint the height. Returns
// the transaction hash if the message is successfully relayed.
//
// The steps are:
//  1. Skip messages the destination messenger is certain to reject from content alone, and
//     messages it has already received.
//  2. Select a confirmed anchor on the ZKAdapter whose history covers the message's slot.
//  3. Build the SSZ execution proof (anchor block root -> receipts root) and the MPT receipt
//     proof (receipts root -> the message's log), and check they agree on the receipts root.
//  4. Encode both as the attestation and send receiveCrossChainMessage to the messenger, which
//     hands the attestation to the ZKAdapter for on-chain verification.
func (m *messageHandler) ProcessMessage() (common.Hash, error) {
	ctx, cancel := context.WithTimeout(context.Background(), processTimeout)
	defer cancel()

	// Reject undeliverable messages before any network calls, proof building, or gas. Anyone
	// can call the adapter's sendMessage naming an arbitrary origin, so logs reaching the
	// handler are untrusted.
	if ok, reason := m.shouldDeliverMessage(); !ok {
		m.logger.Warn("Skipping undeliverable message", zap.String("reason", reason))
		m.metrics.IncFailedRelayMessageCount(reason)
		return common.Hash{}, nil
	}

	adapter, err := zkadapter.NewZKAdapter(m.messageConfig.adapterAddress(), m.destinationClient.Client())
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to bind ZKAdapter: %w", err)
	}
	// The messenger keys delivery by the adapter-attested source blockchain ID, so the
	// delivered-checks must use the same ID sendMessage packs, not the relayer config's.
	sourceNetworkID, err := adapter.SourceNetworkId(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to read adapter source network ID: %w", err)
	}
	sourceChainID, err := adapter.SourceChainId(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to read adapter source chain ID: %w", err)
	}

	messageID, err := teleporterUtils.CalculateMessageID(
		m.messageConfig.teleporterAddress(),
		ids.ID(common.BigToHash(sourceChainID)),
		ids.ID(m.teleporterMessage.DestinationBlockchainID),
		m.teleporterMessage.MessageNonce,
	)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to calculate teleporter message ID: %w", err)
	}
	logger := m.logger.With(zap.Stringer("teleporterMessageID", messageID))

	// Skip messages the destination already received, so duplicates observed on the source
	// chain (e.g. after a relayer restart) cost no proof building.
	delivered, err := teleporterv2.MessageDelivered(
		ctx, m.destinationClient.Client(), m.messageConfig.teleporterAddress(), messageID,
	)
	if err != nil {
		m.metrics.IncFailedRelayMessageCount("failed to check delivery status")
		return common.Hash{}, err
	}
	if delivered {
		logger.Info("Message already delivered to destination")
		return common.Hash{}, nil
	}

	anchor, err := m.awaitAnchor(ctx, adapter)
	if err != nil {
		m.metrics.IncFailedRelayMessageCount("failed to select anchor")
		return common.Hash{}, err
	}
	logger.Info("Selected anchor", zap.Uint64("anchorSlot", anchor.slot))

	execProof, err := m.proofBuilder.BuildExecutionProofForSlots(ctx, anchor.slot, m.slot, anchor.root)
	if err != nil {
		m.metrics.IncFailedRelayMessageCount("failed to build execution proof")
		return common.Hash{}, fmt.Errorf("failed to build execution proof: %w", err)
	}
	receiptProof, err := proofs.BuildReceiptProof(
		ctx, m.sourceClient, m.sourceBlockNumber, m.sourceMessage.SourceTxID, m.sourceLogIndex,
	)
	if err != nil {
		m.metrics.IncFailedRelayMessageCount("failed to build receipt proof")
		return common.Hash{}, fmt.Errorf("failed to build receipt proof: %w", err)
	}
	// The execution proof resolves to the receipts root of the beacon state's execution payload,
	// and the receipt proof is built against the source block header's receipts root. They must
	// be the same block, or the contract would verify the log against the wrong receipts trie.
	if common.Hash(execProof.TargetReceiptsRoot) != m.sourceReceiptsRoot {
		m.metrics.IncFailedRelayMessageCount("receipts root mismatch")
		return common.Hash{}, fmt.Errorf(
			"execution proof receipts root %s does not match source block receipts root %s",
			common.Hash(execProof.TargetReceiptsRoot), m.sourceReceiptsRoot,
		)
	}

	txHash, err := m.sendMessage(ctx, sourceNetworkID, sourceChainID, execProof, receiptProof)
	if err != nil {
		// A revert (simulated or mined) after the delivery check may mean another relayer
		// delivered first.
		delivered, checkErr := teleporterv2.MessageDelivered(
			ctx, m.destinationClient.Client(), m.messageConfig.teleporterAddress(), messageID,
		)
		if checkErr == nil && delivered {
			logger.Info("Message delivered by another relayer")
			return common.Hash{}, nil
		}
		m.metrics.IncFailedRelayMessageCount("failed to send message")
		return common.Hash{}, err
	}
	m.metrics.IncSuccessfulRelayMessageCount()
	logger.Info("Delivered message", zap.Stringer("txHash", txHash))
	return txHash, nil
}

// shouldDeliverMessage rejects messages the destination messenger is certain to refuse from
// content alone, so spoofed or undeliverable messages cost no proof building or gas. Returning
// false skips the message without an error, so the framework checkpoints past it instead of
// retrying a message that can never land.
func (m *messageHandler) shouldDeliverMessage() (bool, string) {
	if m.teleporterMessage.OriginTeleporterAddress != m.messageConfig.teleporterAddress() {
		return false, "origin is not the configured teleporter"
	}

	if !messages.ContainsAllowedRelayer(
		m.teleporterMessage.AllowedRelayerAddresses, m.destinationClient.SenderAddresses(),
	) {
		return false, "no relayer key in the message's allowed relayers"
	}
	return true, ""
}

// sendMessage encodes the attestation, packs the receiveCrossChainMessage call, and sends it to
// the destination TeleporterMessengerV2.
func (m *messageHandler) sendMessage(
	ctx context.Context,
	sourceNetworkID uint32,
	sourceChainID *big.Int,
	execProof *zkadapter.ExecutionProof,
	receiptProof *proofs.ReceiptProof,
) (common.Hash, error) {
	attestation, err := encodeAttestation(execProof, receiptProof)
	if err != nil {
		return common.Hash{}, err
	}

	callData, err := teleportermessengerv2.PackReceiveCrossChainMessageMerkle(
		*m.teleporterMessage,
		sourceNetworkID,
		ids.ID(common.BigToHash(sourceChainID)),
		attestation,
		m.messageConfig.rewardAddress(),
	)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack receiveCrossChainMessage: %w", err)
	}

	teleporterAddress := m.messageConfig.teleporterAddress()
	sender := messages.SelectSenderAddress(
		m.destinationClient.SenderAddresses(), m.teleporterMessage.AllowedRelayerAddresses,
	)
	gasLimit, err := messages.EstimateGasLimit(
		ctx, m.logger, m.destinationClient, sender, teleporterAddress, callData,
	)
	if err != nil {
		return common.Hash{}, err
	}

	// Restrict the signing key to the message's allowed relayers.
	// shouldDeliverMessage guarantees an eligible key exists when the set is non-empty.
	receipt, err := m.destinationClient.SendTx(
		m.logger,
		nil,
		set.Of(m.teleporterMessage.AllowedRelayerAddresses...),
		teleporterAddress,
		gasLimit,
		callData,
	)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to send transaction: %w", err)
	}

	if err := messages.CheckReceiptStatus(receipt, gasLimit); err != nil {
		return common.Hash{}, err
	}
	return receipt.TxHash, nil
}
