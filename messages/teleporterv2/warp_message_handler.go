// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporterv2

import (
	"context"
	"fmt"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/evm/predicate"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	warpadapter "github.com/ava-labs/icm-services/abi-bindings/go/teleporterV2/WarpAdapter"
	gasUtils "github.com/ava-labs/icm-services/icm-contracts/utils/gas-utils"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/libevm/common"
)

// warpMessageHandler relays TeleporterV2 messages whose destination messenger uses the
// WarpAdapter for verification. The signed warp message is included as a transaction predicate
// and the attestation passed to receiveCrossChainMessage is the index of that message in the
// predicate, which the WarpAdapter reads back through the warp precompile.
type warpMessageHandler struct {
	handlerBase
	signingSubnetID ids.ID
}

// ProcessMessage relays the message to the destination chain by aggregating a signature for it
// and sending it via SendMessage. It does not retry on failure or checkpoint the height.
// Returns the transaction hash if the message is successfully relayed.
func (m *warpMessageHandler) ProcessMessage() (common.Hash, error) {
	return m.relay(func(ctx context.Context) (common.Hash, error) {
		// Determine the appropriate P-Chain height for validator set selection
		pchainHeight, err := m.destinationClient.GetPChainHeightForDestination(ctx)
		if err != nil {
			m.metrics.IncFailedRelayMessageCount("failed to determine P-Chain height")
			return common.Hash{}, fmt.Errorf("failed to determine P-Chain height for validator set: %w", err)
		}

		signedMessage, err := m.signMessage(ctx, m.signingSubnetID, pchainHeight)
		if err != nil {
			return common.Hash{}, err
		}

		txHash, err := m.SendMessage(signedMessage)
		if err != nil {
			m.metrics.IncFailedRelayMessageCount("failed to send warp message")
			return common.Hash{}, fmt.Errorf("failed to send warp message: %w", err)
		}
		return txHash, nil
	})
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

	return m.sendTxAndConfirm(utils.SignedWarpMessageToAccessList(signedMessage), gasLimit, callData)
}
