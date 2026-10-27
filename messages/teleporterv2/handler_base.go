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
	"github.com/ava-labs/icm-services/vms/evm"
	ethereum "github.com/ava-labs/libevm"
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

// simulateDelivery executes the delivery that sendTxAndConfirm would broadcast as an eth_call
// against the destination chain's latest state, without broadcasting anything. It returns nil if
// the call would succeed, an [*evm.RevertError] if the node executed the call and it failed, and
// any other error if the call could not be simulated.
//
// The call carries the same sender, calldata, access list and gas limit as the transaction. Note
// that nodes silently cap the gas of a call at their RPC gas cap, so the destination node must
// allow calls at least the configured block gas limit for the simulation to be faithful
// (subnet-evm and geth default to 50M, against a 12M default block gas limit here).
func (m *handlerBase) simulateDelivery(
	ctx context.Context,
	accessList types.AccessList,
	gasLimit uint64,
	callData []byte,
) error {
	ctx, cancel := context.WithTimeout(ctx, utils.DefaultRPCTimeout)
	defer cancel()
	_, err := m.destinationClient.Client().CallContract(ctx, ethereum.CallMsg{
		From:       m.selectSenderAddress(),
		To:         &m.teleporterAddress,
		Gas:        gasLimit,
		Data:       callData,
		AccessList: accessList,
	}, nil)
	if err == nil {
		return nil
	}
	if revert := evm.AsRevertError(err); revert != nil {
		return revert
	}
	return fmt.Errorf("failed to simulate delivery: %w", err)
}

// abandonSimulatedRevert reports a delivery that reverts in simulation. The one benign cause is
// that another relayer delivered the message in the meantime, in which case there is nothing left
// to do and the zero hash is returned. Any other revert is reproducible from the message, so the
// returned error is non-retryable: broadcasting the delivery would only mine the same revert, and
// retrying it with a fresh signature would mine it again. Abandoning the message lets the relayer
// checkpoint past it instead of exiting and replaying it on every restart.
func (m *handlerBase) abandonSimulatedRevert(revert *evm.RevertError) (common.Hash, error) {
	delivered, err := m.messageDelivered()
	if err != nil {
		m.logger.Error(
			"Failed to check if message has been delivered to destination chain.",
			zap.Error(err),
		)
		return common.Hash{}, fmt.Errorf("failed to check if message has been delivered: %w", err)
	}
	if delivered {
		m.logger.Info("Delivery reverts in simulation: message already delivered to destination.")
		return common.Hash{}, nil
	}

	m.logger.Error(
		"Delivery reverts in simulation, not sending",
		zap.String("revertReason", revert.Reason),
		zap.Int("numReceipts", len(m.teleporterMessage.Receipts)),
	)
	return common.Hash{}, fmt.Errorf("%w: %s", messages.ErrDeliveryReverts, revert.Reason)
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

	delivered, err := m.messageDelivered()
	if err != nil {
		log.Error("Transaction failed", zap.Error(err))
		return common.Hash{}, fmt.Errorf("transaction failed with status: %d", receipt.Status)
	}
	if delivered {
		log.Info("Execution reverted: message already delivered to destination.")
		return txHash, nil
	}

	// A revert that consumed the entire gas limit is an out-of-gas failure. The gas limit is
	// computed from the message, or estimated against the same state, so a retry would fail
	// identically and only bill the relayer for a second reverted transaction. Reverts that left
	// gas unspent are logic failures, such as verification against a validator set that has
	// churned since the signature was aggregated, which a fresh attempt may resolve.
	if receipt.GasUsed >= gasLimit {
		log.Error(
			"Transaction ran out of gas",
			zap.Uint64("gasUsed", receipt.GasUsed),
			zap.Uint64("gasLimit", gasLimit),
			zap.Int("numReceipts", len(m.teleporterMessage.Receipts)),
		)
		return common.Hash{}, fmt.Errorf(
			"%w: used %d of %d",
			messages.ErrDeliveryOutOfGas, receipt.GasUsed, gasLimit,
		)
	}

	log.Error("Transaction failed")
	return common.Hash{}, fmt.Errorf("transaction failed with status: %d", receipt.Status)
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

// selectSenderAddress picks a relayer EOA eligible to deliver the message, as the sender of
// simulated deliveries. Any eligible EOA gives the same result, since the only sender-dependent
// check the messenger makes is the allowed-relayer one.
func (m *handlerBase) selectSenderAddress() common.Address {
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
