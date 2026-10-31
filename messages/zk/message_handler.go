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
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	zkadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKAdapter"
	teleporterUtils "github.com/ava-labs/icm-services/icm-contracts/utils/teleporter-utils"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/messages/zk/proofs"
	"github.com/ava-labs/icm-services/messages/zk/proofs/beacon"
	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator"
	"github.com/ava-labs/icm-services/vms"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/ethclient"
	"go.uber.org/zap"
)

const (
	// processTimeout bounds one delivery attempt end to end: the anchor scan, two beacon state
	// fetches of ~100-200MB each, and the destination transaction.
	processTimeout = 15 * time.Minute

	// gasLimitBufferNumerator/Denominator adds a safety margin on top of the estimated gas, since
	// on-chain SSZ and MPT verification cost depends on proof depth.
	gasLimitBufferNumerator   = 5
	gasLimitBufferDenominator = 4
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
//  1. Skip the message if the destination messenger has already received it.
//  2. Select a confirmed anchor on the ZKAdapter whose history covers the message's slot.
//  3. Build the SSZ execution proof (anchor block root -> receipts root) and the MPT receipt
//     proof (receipts root -> the message's log), and check they agree on the receipts root.
//  4. Encode both as the attestation and send receiveCrossChainMessage to the messenger, which
//     hands the attestation to the ZKAdapter for on-chain verification.
func (m *messageHandler) ProcessMessage() (common.Hash, error) {
	ctx, cancel := context.WithTimeout(context.Background(), processTimeout)
	defer cancel()

	messageID, err := m.messageID()
	if err != nil {
		return common.Hash{}, err
	}
	logger := m.logger.With(zap.Stringer("teleporterMessageID", messageID))

	// Skip messages the destination already received, so duplicates observed on the source
	// chain (e.g. after a relayer restart) cost no proof building.
	delivered, err := m.messageDelivered(ctx, messageID)
	if err != nil {
		m.metrics.IncFailedRelayMessageCount("failed to check delivery status")
		return common.Hash{}, err
	}
	if delivered {
		logger.Info("Message already delivered to destination")
		return common.Hash{}, nil
	}

	adapter, err := zkadapter.NewZKAdapter(m.messageConfig.adapterAddress(), m.destinationClient.Client())
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to bind ZKAdapter: %w", err)
	}
	anchor, err := findAnchor(ctx, m.destinationClient.Client(), adapter, m.slot)
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

	txHash, err := m.sendMessage(ctx, adapter, execProof, receiptProof)
	if err != nil {
		// A revert after the delivery check may mean another relayer delivered first.
		if delivered, checkErr := m.messageDelivered(ctx, messageID); checkErr == nil && delivered {
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

// sendMessage encodes the attestation, packs the receiveCrossChainMessage call, and sends it to
// the destination TeleporterMessengerV2.
func (m *messageHandler) sendMessage(
	ctx context.Context,
	adapter *zkadapter.ZKAdapter,
	execProof *zkadapter.ExecutionProof,
	receiptProof *proofs.ReceiptProof,
) (common.Hash, error) {
	attestation, err := encodeAttestation(execProof, receiptProof)
	if err != nil {
		return common.Hash{}, err
	}

	// The adapter checks the message's source network and chain against its own configuration,
	// so take both from the adapter rather than duplicating them in the relayer config.
	sourceNetworkID, err := adapter.SourceNetworkId(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to read adapter source network ID: %w", err)
	}
	sourceChainID, err := adapter.SourceChainId(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to read adapter source chain ID: %w", err)
	}

	callData, err := packReceiveCrossChainMessage(
		teleportermessengerv2.TeleporterICMMessage{
			Message:            *m.teleporterMessage,
			SourceNetworkID:    sourceNetworkID,
			SourceBlockchainID: common.BigToHash(sourceChainID),
			Attestation:        attestation,
		},
		m.messageConfig.rewardAddress(),
	)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack receiveCrossChainMessage: %w", err)
	}

	teleporterAddress := m.messageConfig.teleporterAddress()
	gasLimit, err := m.estimateGasLimit(ctx, teleporterAddress, callData)
	if err != nil {
		return common.Hash{}, err
	}

	receipt, err := m.destinationClient.SendTx(
		m.logger, nil, set.Of(m.destinationClient.SenderAddresses()...), teleporterAddress, gasLimit, callData,
	)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to send transaction: %w", err)
	}

	if err := checkReceiptStatus(receipt, gasLimit); err != nil {
		return common.Hash{}, err
	}
	return receipt.TxHash, nil
}

// estimateGasLimit estimates the gas for the receiveCrossChainMessage call and applies a safety
// buffer, capped at the destination's block gas limit. Falls back to the block gas limit if
// estimation fails.
func (m *messageHandler) estimateGasLimit(
	ctx context.Context,
	teleporterAddress common.Address,
	callData []byte,
) (uint64, error) {
	blockGasLimit := m.destinationClient.BlockGasLimit()
	estimated, err := m.destinationClient.Client().EstimateGas(ctx, ethereum.CallMsg{
		From: m.selectSenderAddress(),
		To:   &teleporterAddress,
		Data: callData,
	})
	if err != nil {
		m.logger.Warn("Gas estimation failed, falling back to block gas limit",
			zap.Error(err),
			zap.Uint64("blockGasLimit", blockGasLimit),
		)
		if blockGasLimit == 0 {
			return 0, fmt.Errorf("failed to estimate gas and no block gas limit configured: %w", err)
		}
		return blockGasLimit, nil
	}
	buffered := estimated * gasLimitBufferNumerator / gasLimitBufferDenominator
	if blockGasLimit != 0 && buffered > blockGasLimit {
		buffered = blockGasLimit
	}
	return buffered, nil
}

// selectSenderAddress picks a relayer EOA eligible to deliver the message for gas estimation.
func (m *messageHandler) selectSenderAddress() common.Address {
	senders := m.destinationClient.SenderAddresses()
	for _, sender := range senders {
		for _, allowed := range m.teleporterMessage.AllowedRelayerAddresses {
			if sender == allowed {
				return sender
			}
		}
	}
	if len(senders) > 0 {
		return senders[0]
	}
	return common.Address{}
}

// messageID computes the Teleporter message ID the destination messenger tracks delivery by.
func (m *messageHandler) messageID() (ids.ID, error) {
	id, err := teleporterUtils.CalculateMessageID(
		m.messageConfig.teleporterAddress(),
		m.sourceMessage.SourceBlockchainID,
		ids.ID(m.teleporterMessage.DestinationBlockchainID),
		m.teleporterMessage.MessageNonce,
	)
	if err != nil {
		return ids.Empty, fmt.Errorf("failed to calculate teleporter message ID: %w", err)
	}
	return id, nil
}

// messageDelivered reports whether the destination messenger has already received the message.
func (m *messageHandler) messageDelivered(ctx context.Context, messageID ids.ID) (bool, error) {
	messenger, err := teleportermessengerv2.NewTeleporterMessengerV2(
		m.messageConfig.teleporterAddress(), m.destinationClient.Client(),
	)
	if err != nil {
		return false, fmt.Errorf("failed to bind TeleporterMessengerV2: %w", err)
	}
	delivered, err := messenger.MessageReceived(&bind.CallOpts{Context: ctx}, messageID)
	if err != nil {
		return false, fmt.Errorf("failed to check message delivery status: %w", err)
	}
	return delivered, nil
}

// packReceiveCrossChainMessage packs the calldata for the messenger's receiveCrossChainMessage.
func packReceiveCrossChainMessage(
	message teleportermessengerv2.TeleporterICMMessage,
	rewardAddress common.Address,
) ([]byte, error) {
	messengerABI, err := teleportermessengerv2.TeleporterMessengerV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return messengerABI.Pack("receiveCrossChainMessage", message, rewardAddress)
}

// checkReceiptStatus maps a failed delivery receipt to the relayer's error taxonomy: running
// out of gas is non-retryable since a resend would burn the same gas again.
func checkReceiptStatus(receipt *types.Receipt, gasLimit uint64) error {
	if receipt.Status == types.ReceiptStatusSuccessful {
		return nil
	}
	if receipt.GasUsed >= gasLimit {
		return fmt.Errorf("%w: used %d of %d", messages.ErrDeliveryOutOfGas, receipt.GasUsed, gasLimit)
	}
	return fmt.Errorf("transaction %s reverted", receipt.TxHash.Hex())
}
