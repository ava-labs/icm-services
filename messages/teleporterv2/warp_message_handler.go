// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporterv2

import (
	"context"
	"errors"
	"fmt"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/evm/predicate"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	warpadapter "github.com/ava-labs/icm-services/abi-bindings/go/teleporterV2/WarpAdapter"
	gasUtils "github.com/ava-labs/icm-services/icm-contracts/utils/gas-utils"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/icm-services/vms/evm"
	"github.com/ava-labs/libevm/common"
	"go.uber.org/zap"
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

		txHash, err := m.SendMessage(ctx, signedMessage)
		if err != nil {
			m.metrics.IncFailedRelayMessageCount("failed to send warp message")
			return common.Hash{}, fmt.Errorf("failed to send warp message: %w", err)
		}
		return txHash, nil
	})
}

// SendMessage delivers the signed message to the destination TeleporterMessengerV2, attaching it
// as a transaction predicate for the WarpAdapter to verify through the warp precompile.
func (m *warpMessageHandler) SendMessage(ctx context.Context, signedMessage *warp.Message) (common.Hash, error) {
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

	// Simulate the delivery before paying for it. The call carries the signed message in its
	// access list like the transaction would, but a node does not verify predicates for calls:
	// getVerifiedWarpMessage reports the message as valid and the simulation runs every check the
	// messenger makes after verification against the message itself. A revert here is therefore
	// reproducible from the message content, independent of the attestation, and would be mined
	// identically on every retry. The verification failures a retry with a fresh signature can fix
	// are exactly the ones the simulation cannot produce, so those still go through to the chain.
	err = m.simulateDelivery(ctx, accessList, gasLimit, callData)
	var revert *evm.RevertError
	switch {
	case err == nil:
	case !errors.As(err, &revert):
		m.logger.Error("Failed to simulate delivery", zap.Error(err))
		return common.Hash{}, err
	case revert.OutOfGas:
		// Running out of gas in simulation is most likely deterministic too, but it is also what a
		// node's RPC gas cap looks like when it is below the gas limit. Let the chain decide: if
		// the transaction runs out of gas there as well, sendTxAndConfirm abandons the message
		// after this one attempt.
		m.logger.Warn(
			"Delivery ran out of gas in simulation, sending anyway",
			zap.Uint64("gasLimit", gasLimit),
			zap.String("reason", revert.Reason),
		)
	default:
		return m.abandonSimulatedRevert(revert)
	}

	return m.sendTxAndConfirm(accessList, gasLimit, callData)
}
