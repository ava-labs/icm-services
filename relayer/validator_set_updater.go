// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package relayer

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/icm-services/peers"
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator"
	"github.com/ava-labs/icm-services/vms/evm"
	"github.com/ava-labs/libevm/common"
	"go.uber.org/zap"
)

const (
	gasLimit = 1000000
)

// ValidatorSetUpdater monitors P-chain validator sets and posts updates
// to external EVM chains' AvalancheValidatorSetRegistry contracts.
type ValidatorSetUpdater struct {
	logger              logging.Logger
	validatorManager    *peers.ValidatorManager
	signatureAggregator *aggregator.SignatureAggregator

	// External EVM clients keyed by chain ID string (e.g., "1" for Ethereum)
	externalEVMClients map[string]*evm.ExternalEVMDestinationClient

	// Source L1 subnet IDs to monitor
	sourceSubnetIDs []ids.ID

	// Polling configuration
	pollIntervalSeconds uint64

	// Per-source-subnet tracking to avoid redundant updates
	lastPostedHeights map[ids.ID]uint64   // sourceSubnetID -> P-chain height
	lastPostedHashes  map[ids.ID][32]byte // sourceSubnetID -> hash of validator set

	registryAddress common.Address
}

// NewValidatorSetUpdater creates a new validator set updater.
func NewValidatorSetUpdater(
	logger logging.Logger,
	validatorManager *peers.ValidatorManager,
	signatureAggregator *aggregator.SignatureAggregator,
	externalEVMClients map[string]*evm.ExternalEVMDestinationClient,
	sourceSubnetIDs []ids.ID,
	pollIntervalSeconds uint64,
) (*ValidatorSetUpdater, error) {
	if len(externalEVMClients) == 0 {
		return nil, fmt.Errorf("no external EVM clients configured")
	}
	if len(sourceSubnetIDs) == 0 {
		return nil, fmt.Errorf("no source subnet IDs configured")
	}

	return &ValidatorSetUpdater{
		logger:              logger,
		validatorManager:    validatorManager,
		signatureAggregator: signatureAggregator,
		externalEVMClients:  externalEVMClients,
		sourceSubnetIDs:     sourceSubnetIDs,
		pollIntervalSeconds: pollIntervalSeconds,
		lastPostedHeights:   make(map[ids.ID]uint64),
		lastPostedHashes:    make(map[ids.ID][32]byte),
	}, nil
}

// Start begins the polling loop to check for validator set updates.
func (u *ValidatorSetUpdater) Start(ctx context.Context) error {
	ticker := time.NewTicker(time.Duration(u.pollIntervalSeconds) * time.Second)
	defer ticker.Stop()

	u.logger.Info("Validator set updater started",
		zap.Uint64("pollIntervalSeconds", u.pollIntervalSeconds),
		zap.Int("numExternalEVMClients", len(u.externalEVMClients)),
		zap.Int("numSourceSubnets", len(u.sourceSubnetIDs)),
	)

	// Initial check
	if err := u.checkAndUpdate(ctx); err != nil {
		u.logger.Error("Initial validator set update check failed", zap.Error(err))
	}

	for {
		select {
		case <-ticker.C:
			if err := u.checkAndUpdate(ctx); err != nil {
				u.logger.Error("Periodic validator set update check failed", zap.Error(err))
			}
		case <-ctx.Done():
			u.logger.Info("Validator set updater stopped")
			return nil
		}
	}
}

// checkAndUpdate checks for validator set changes and posts updates if needed.
func (u *ValidatorSetUpdater) checkAndUpdate(ctx context.Context) error {
	// Get latest P-chain height (no error return)
	pChainHeight := u.validatorManager.GetLatestSyncedPChainHeight()

	u.logger.Debug("Checking for validator set updates",
		zap.Uint64("pChainHeight", pChainHeight),
	)

	// Get all validator sets at this height
	allValidatorSets, err := u.validatorManager.GetAllValidatorSets(ctx, pChainHeight)
	if err != nil {
		return fmt.Errorf("failed to get validator sets at height %d: %w", pChainHeight, err)
	}

	// For each external EVM destination
	for chainID, client := range u.externalEVMClients {
		if err := u.updateDestination(ctx, chainID, client, pChainHeight, allValidatorSets); err != nil {
			u.logger.Error("Failed to update destination",
				zap.String("chainID", chainID),
				zap.Error(err),
			)
			// Continue with other destinations
		}
	}

	return nil
}

// updateDestination checks and updates a single external EVM destination.
func (u *ValidatorSetUpdater) updateDestination(
	ctx context.Context,
	chainID string,
	client *evm.ExternalEVMDestinationClient,
	currentPChainHeight uint64,
	allValidatorSets map[ids.ID]validators.WarpSet,
) error {
	// Get the P-chain height that the destination registry knows about
	registryPChainHeight, err := client.GetPChainHeightForDestination(ctx)
	if err != nil {
		return fmt.Errorf("failed to get P-chain height from registry: %w", err)
	}

	u.logger.Debug("Comparing P-chain heights",
		zap.String("chainID", chainID),
		zap.Uint64("currentPChainHeight", currentPChainHeight),
		zap.Uint64("registryPChainHeight", registryPChainHeight),
	)

	// If registry is up to date, no need to update
	if registryPChainHeight >= currentPChainHeight {
		return nil
	}

	// For each source subnet we're monitoring, check if update is needed
	for _, subnetID := range u.sourceSubnetIDs {
		validatorSet, ok := allValidatorSets[subnetID]
		if !ok {
			u.logger.Warn("Validator set not found for subnet",
				zap.String("chainID", chainID),
				zap.Stringer("subnetID", subnetID),
			)
			continue
		}

		if err := u.updateSubnetValidators(
			ctx,
			chainID,
			client,
			subnetID,
			currentPChainHeight,
			registryPChainHeight,
			validatorSet,
		); err != nil {
			u.logger.Error("Failed to update subnet validators",
				zap.String("chainID", chainID),
				zap.Stringer("subnetID", subnetID),
				zap.Error(err),
			)
			// Continue with other subnets
		}
	}

	return nil
}

// updateSubnetValidators updates the validator set for a specific subnet on a destination.
func (u *ValidatorSetUpdater) updateSubnetValidators(
	ctx context.Context,
	chainID string,
	client *evm.ExternalEVMDestinationClient,
	subnetID ids.ID,
	currentPChainHeight uint64,
	registryPChainHeight uint64,
	validatorSet validators.WarpSet,
) error {
	// Check if validator set has changed from last posted
	validatorHash := u.computeValidatorSetHash(validatorSet)
	if lastHash, ok := u.lastPostedHashes[subnetID]; ok {
		if lastHash == validatorHash {
			u.logger.Debug("Validator set unchanged, skipping update",
				zap.String("chainID", chainID),
				zap.Stringer("subnetID", subnetID),
			)
			return nil
		}
	}

	u.logger.Info("Validator set changed, preparing signed update",
		zap.String("chainID", chainID),
		zap.Stringer("subnetID", subnetID),
		zap.Uint64("currentPChainHeight", currentPChainHeight),
		zap.Uint64("registryPChainHeight", registryPChainHeight),
		zap.Int("numValidators", len(validatorSet.Validators)),
		zap.Uint64("totalWeight", validatorSet.TotalWeight),
	)

	// Step 1: Create unsigned Warp message with validator set update payload
	unsignedMessage, err := u.createValidatorSetUpdateMessage(subnetID, currentPChainHeight, validatorSet)
	if err != nil {
		return fmt.Errorf("failed to create validator set update message: %w", err)
	}

	// Step 2: Sign message using validators at the registry's known P-chain height
	// This ensures the message can be verified by the registry
	signedMessage, err := u.signMessage(ctx, unsignedMessage, subnetID, registryPChainHeight)
	if err != nil {
		return fmt.Errorf("failed to sign message: %w", err)
	}

	// Step 3: Send to external EVM via SendTx
	err = u.sendValidatorSetUpdate(ctx, client, signedMessage)
	if err != nil {
		return fmt.Errorf("failed to send validator set update: %w", err)
	}

	// Update tracking state after successful send
	u.lastPostedHeights[subnetID] = currentPChainHeight
	u.lastPostedHashes[subnetID] = validatorHash

	u.logger.Info("Successfully sent validator set update",
		zap.String("chainID", chainID),
		zap.Stringer("subnetID", subnetID),
		zap.Uint64("pChainHeight", currentPChainHeight),
	)

	return nil
}

// createValidatorSetUpdateMessage creates an unsigned Warp message for validator set update.
func (u *ValidatorSetUpdater) createValidatorSetUpdateMessage(
	subnetID ids.ID,
	pChainHeight uint64,
	validatorSet validators.WarpSet,
) (*warp.UnsignedMessage, error) {
	// TODO: Implement message creation
	// The payload should contain:
	// - P-chain height
	// - Subnet ID
	// - List of validators (nodeID, publicKey, weight)
	//
	// payload := encodeValidatorSetUpdatePayload(pChainHeight, subnetID, validatorSet)
	// unsignedMessage, err := warp.NewUnsignedMessage(networkID, sourceChainID, payload)

	_ = subnetID
	_ = pChainHeight
	_ = validatorSet

	return nil, fmt.Errorf("createValidatorSetUpdateMessage not implemented")
}

// signMessage signs the message using SignatureAggregator.
// Uses the P-chain height from the registry so the external chain can verify.
func (u *ValidatorSetUpdater) signMessage(
	ctx context.Context,
	unsignedMessage *warp.UnsignedMessage,
	subnetID ids.ID,
	pChainHeight uint64,
) (*warp.Message, error) {
	// Use the same signing method as normal message relay
	// The pChainHeight should be from the registry (registryPChainHeight)
	// so the external chain can verify against its known validator set

	// TODO: sign the message
	signedMessage := &warp.Message{
		UnsignedMessage: *unsignedMessage,
		Signature:       &warp.BitSetSignature{},
	}

	return signedMessage, nil
}

// sendValidatorSetUpdate sends the signed message to the external EVM.
func (u *ValidatorSetUpdater) sendValidatorSetUpdate(
	ctx context.Context,
	client *evm.ExternalEVMDestinationClient,
	signedMessage *warp.Message,
) error {
	// TODO: Construct callData for registry contract and send via client.SendTx

	// callData := registryABI.Pack("updateValidatorSet", signedMessage.Bytes())
	receipt, err := client.SendTx(
		signedMessage,
		nil,                     // Any sender can deliver
		u.registryAddress.Hex(), // Registry contract address
		gasLimit,
		[]byte{},
	)
	if err != nil {
		return fmt.Errorf("failed to send validator set update: %w", err)
	}
	if receipt == nil {
		return fmt.Errorf("failed to send validator set update: no receipt")
	}

	return fmt.Errorf("sendValidatorSetUpdate not implemented")
}

// computeValidatorSetHash computes a hash of the validator set for change detection.
func (u *ValidatorSetUpdater) computeValidatorSetHash(validatorSet validators.WarpSet) [32]byte {
	// Hash the validator public keys and weights
	h := sha256.New()
	for _, v := range validatorSet.Validators {
		h.Write(v.PublicKeyBytes)
		// Write weight as bytes
		weightBytes := make([]byte, 8)
		for i := 0; i < 8; i++ {
			weightBytes[i] = byte(v.Weight >> (56 - i*8))
		}
		h.Write(weightBytes)
	}
	var result [32]byte
	copy(result[:], h.Sum(nil))
	return result
}
