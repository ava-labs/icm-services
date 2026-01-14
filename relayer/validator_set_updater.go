// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package relayer

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	"github.com/ava-labs/avalanchego/vms/platformvm/block"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/message"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	avalanchevalidatorsetregistry "github.com/ava-labs/icm-services/abi-bindings/go/AvalancheValidatorSetRegistry"
	"github.com/ava-labs/icm-services/peers"
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator"
	"github.com/ava-labs/icm-services/vms/evm"
	"go.uber.org/zap"
)

const (
	// Gas limit for validator set update transactions
	// BLS12-381 precompile operations (signature verification, hash to curve) are gas-intensive
	// and require more gas than typical EVM operations
	gasLimit = 5_000_000

	// Default quorum percentage for signing
	defaultQuorumPercentage = 67
)

// ValidatorSetUpdater monitors P-chain validator sets and posts updates
// to external EVM chains' AvalancheValidatorSetRegistry contracts.
type ValidatorSetUpdater struct {
	logger              logging.Logger
	validatorManager    *peers.ValidatorManager
	signatureAggregator *aggregator.SignatureAggregator
	networkID           uint32

	// P-chain client for fetching block timestamps
	pChainClient *platformvm.Client

	// External EVM clients keyed by chain ID string (e.g., "1" for Ethereum)
	externalEVMClients map[string]*evm.ExternalEVMDestinationClient

	// Source L1 subnet IDs to monitor
	sourceSubnetIDs []ids.ID

	// Source blockchain IDs (needed for message creation)
	// Maps subnetID -> blockchainID
	sourceBlockchainIDs map[ids.ID]ids.ID

	// Polling configuration
	pollIntervalSeconds uint64

	// Per-source-subnet tracking to avoid redundant updates
	lastPostedHeights map[ids.ID]uint64   // sourceSubnetID -> P-chain height
	lastPostedHashes  map[ids.ID][32]byte // sourceSubnetID -> hash of validator set
}

// NewValidatorSetUpdater creates a new validator set updater.
func NewValidatorSetUpdater(
	logger logging.Logger,
	validatorManager *peers.ValidatorManager,
	signatureAggregator *aggregator.SignatureAggregator,
	networkID uint32,
	pChainClient *platformvm.Client,
	externalEVMClients map[string]*evm.ExternalEVMDestinationClient,
	sourceSubnetIDs []ids.ID,
	sourceBlockchainIDs map[ids.ID]ids.ID,
	pollIntervalSeconds uint64,
) (*ValidatorSetUpdater, error) {
	if len(externalEVMClients) == 0 {
		return nil, fmt.Errorf("no external EVM clients configured")
	}
	if len(sourceSubnetIDs) == 0 {
		return nil, fmt.Errorf("no source subnet IDs configured")
	}
	if pChainClient == nil {
		return nil, fmt.Errorf("P-chain client is required")
	}

	return &ValidatorSetUpdater{
		logger:              logger,
		validatorManager:    validatorManager,
		signatureAggregator: signatureAggregator,
		networkID:           networkID,
		pChainClient:        pChainClient,
		externalEVMClients:  externalEVMClients,
		sourceSubnetIDs:     sourceSubnetIDs,
		sourceBlockchainIDs: sourceBlockchainIDs,
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

	// Wait for P2P network to stabilize before making signature requests.
	// The signature aggregator needs fully established P2P connections to collect signatures.
	startupDelay := 10 * time.Second
	u.logger.Info("Waiting for P2P network to stabilize before initial update", zap.Duration("delay", startupDelay))
	select {
	case <-time.After(startupDelay):
	case <-ctx.Done():
		return nil
	}

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
	for chainID, clientWithRegistry := range u.externalEVMClients {
		if err := u.updateDestination(ctx, chainID, clientWithRegistry, pChainHeight, allValidatorSets); err != nil {
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
		// If registry doesn't have any validator sets yet, use height 0
		u.logger.Warn("Failed to get P-chain height from registry, assuming 0",
			zap.String("chainID", chainID),
			zap.Error(err),
		)
		registryPChainHeight = 0
	}

	u.logger.Debug("Comparing P-chain heights",
		zap.String("chainID", chainID),
		zap.Uint64("currentPChainHeight", currentPChainHeight),
		zap.Uint64("registryPChainHeight", registryPChainHeight),
	)

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

	// Get blockchainID for this subnet
	blockchainID, ok := u.sourceBlockchainIDs[subnetID]
	if !ok {
		return fmt.Errorf("no blockchain ID found for subnet %s", subnetID)
	}

	// Step 1: Create unsigned Warp message with validator set update payload
	unsignedMessage, validatorsBytes, validatorSetHash, err := u.createValidatorSetUpdateMessage(
		ctx,
		blockchainID,
		currentPChainHeight,
		validatorSet,
	)
	if err != nil {
		return fmt.Errorf("failed to create validator set update message: %w", err)
	}

	// Step 2: Sign message using P-chain validators via signature aggregation
	// For initial registration (registryPChainHeight == 0), use current P-chain height
	// For updates, use the registry's known P-chain height to maintain update chain continuity
	signingHeight := registryPChainHeight
	if signingHeight == 0 {
		signingHeight = currentPChainHeight
	}
	signedMessage, err := u.signMessage(ctx, unsignedMessage, subnetID, signingHeight)
	if err != nil {
		return fmt.Errorf("failed to sign message: %w", err)
	}

	// Step 3: Send to external EVM via SendTx
	err = u.sendValidatorSetUpdate(ctx, client, signedMessage, validatorsBytes)
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
		zap.Stringer("validatorSetHash", validatorSetHash),
	)

	return nil
}

// createValidatorSetUpdateMessage creates an unsigned Warp message for validator set update.
// This uses the ACP-118 ValidatorSetState message type from avalanchego.
// The message is signed by P-chain validators using ids.Empty as source chain.
func (u *ValidatorSetUpdater) createValidatorSetUpdateMessage(
	ctx context.Context,
	blockchainID ids.ID,
	pChainHeight uint64,
	validatorSet validators.WarpSet,
) (*warp.UnsignedMessage, []byte, ids.ID, error) {
	// Get the actual P-chain block timestamp at the given height
	pChainTimestamp, err := u.getPChainBlockTimestamp(ctx, pChainHeight)
	if err != nil {
		return nil, nil, ids.Empty, fmt.Errorf("failed to get P-chain block timestamp: %w", err)
	}

	// Compute validator set hash using the same algorithm as avalanchego
	validatorSetHash, err := u.computeValidatorSetHashForMessage(validatorSet)
	if err != nil {
		return nil, nil, ids.Empty, fmt.Errorf("failed to compute validator set hash: %w", err)
	}

	u.logger.Debug("Creating ValidatorSetState message",
		zap.Stringer("blockchainID", blockchainID),
		zap.Uint64("pChainHeight", pChainHeight),
		zap.Uint64("pChainTimestamp", pChainTimestamp),
		zap.Stringer("validatorSetHash", validatorSetHash),
	)

	// Create ValidatorSetState payload using avalanchego's message package
	validatorSetState, err := message.NewValidatorSetState(
		blockchainID,
		pChainHeight,
		pChainTimestamp,
		validatorSetHash,
	)
	if err != nil {
		return nil, nil, ids.Empty, fmt.Errorf("failed to create ValidatorSetState: %w", err)
	}

	// Create AddressedCall with empty source address
	addressedCall, err := payload.NewAddressedCall(nil, validatorSetState.Bytes())
	if err != nil {
		return nil, nil, ids.Empty, fmt.Errorf("failed to create addressed call: %w", err)
	}

	// Create the unsigned Warp message
	// Use ids.Empty as source chain for P-chain signed messages (ValidatorSetState)
	unsignedMessage, err := warp.NewUnsignedMessage(
		u.networkID,
		ids.Empty, // P-chain messages use empty source chain ID
		addressedCall.Bytes(),
	)
	if err != nil {
		return nil, nil, ids.Empty, fmt.Errorf("failed to create unsigned message: %w", err)
	}

	// Serialize validators for contract submission
	validatorsBytes := serializeValidatorsForContract(validatorSet)

	u.logger.Debug("Created validator set update message",
		zap.Stringer("messageID", unsignedMessage.ID()),
		zap.Stringer("blockchainID", blockchainID),
		zap.Int("payloadLen", len(unsignedMessage.Payload)),
	)

	return unsignedMessage, validatorsBytes, validatorSetHash, nil
}

// getPChainBlockTimestamp fetches the block at the given P-chain height and extracts its timestamp.
func (u *ValidatorSetUpdater) getPChainBlockTimestamp(ctx context.Context, height uint64) (uint64, error) {
	blockBytes, err := u.pChainClient.GetBlockByHeight(ctx, height)
	if err != nil {
		return 0, fmt.Errorf("failed to get block at height %d: %w", height, err)
	}

	// Parse the block using avalanchego's block codec
	blk, err := block.Parse(block.Codec, blockBytes)
	if err != nil {
		return 0, fmt.Errorf("failed to parse block: %w", err)
	}

	// Extract timestamp from BanffBlock
	banffBlock, ok := blk.(block.BanffBlock)
	if !ok {
		return 0, fmt.Errorf("block at height %d is not a BanffBlock", height)
	}

	return uint64(banffBlock.Timestamp().Unix()), nil
}

// computeValidatorSetHashForMessage computes the validator set hash using the same algorithm
// as avalanchego's verifyValidatorSetState. This uses uncompressed public keys (96 bytes)
// and the message.Codec for serialization.
func (u *ValidatorSetUpdater) computeValidatorSetHashForMessage(validatorSet validators.WarpSet) (ids.ID, error) {
	// Convert to message.Validator format with uncompressed public keys
	validators := make([]*message.Validator, len(validatorSet.Validators))
	for i, v := range validatorSet.Validators {
		// PublicKeyBytes should already be uncompressed (96 bytes) from the WarpSet
		if len(v.PublicKeyBytes) != bls.PublicKeyLen {
			// If it's compressed (48 bytes), we need to decompress it
			if v.PublicKey != nil {
				validators[i] = &message.Validator{
					UncompressedPublicKeyBytes: [96]byte(bls.PublicKeyToUncompressedBytes(v.PublicKey)),
					Weight:                     v.Weight,
				}
			} else {
				return ids.Empty, fmt.Errorf("validator %d has no public key", i)
			}
		} else {
			validators[i] = &message.Validator{
				UncompressedPublicKeyBytes: [96]byte(v.PublicKeyBytes),
				Weight:                     v.Weight,
			}
		}
	}

	// Marshal using message.Codec (same as avalanchego's verification)
	bytes, err := message.Codec.Marshal(message.CodecVersion, validators)
	if err != nil {
		return ids.Empty, fmt.Errorf("failed to marshal validators: %w", err)
	}

	// SHA256 hash
	hash := sha256.Sum256(bytes)
	return ids.ID(hash), nil
}

// serializeValidatorsForContract serializes validators for the registry contract.
// Uses uncompressed public keys (96 bytes) to match avalanchego's format.
func serializeValidatorsForContract(validatorSet validators.WarpSet) []byte {
	// Validators are already sorted in canonical order (by uncompressed public key bytes)
	// from validators.FlattenValidatorSet in avalanchego

	// Use message.Codec to serialize validators in the same format as the hash computation
	validatorSlice := make([]*message.Validator, len(validatorSet.Validators))
	for i, v := range validatorSet.Validators {
		if len(v.PublicKeyBytes) == bls.PublicKeyLen {
			// Already uncompressed (96 bytes)
			validatorSlice[i] = &message.Validator{
				UncompressedPublicKeyBytes: [96]byte(v.PublicKeyBytes),
				Weight:                     v.Weight,
			}
		} else if v.PublicKey != nil {
			// Convert to uncompressed
			validatorSlice[i] = &message.Validator{
				UncompressedPublicKeyBytes: [96]byte(bls.PublicKeyToUncompressedBytes(v.PublicKey)),
				Weight:                     v.Weight,
			}
		}
	}

	// Marshal using message.Codec for consistency
	data, err := message.Codec.Marshal(message.CodecVersion, validatorSlice)
	if err != nil {
		// This should not happen with valid validator data
		return nil
	}

	return data
}

// signMessage signs the message using SignatureAggregator.
// ValidatorSetState messages are signed by P-chain validators.
func (u *ValidatorSetUpdater) signMessage(
	ctx context.Context,
	unsignedMessage *warp.UnsignedMessage,
	subnetID ids.ID,
	pChainHeight uint64,
) (*warp.Message, error) {
	u.logger.Debug("Signing ValidatorSetState message via P-chain validators",
		zap.Stringer("messageID", unsignedMessage.ID()),
		zap.Stringer("subnetID", subnetID),
		zap.Uint64("pChainHeight", pChainHeight),
	)

	// Use signature aggregator to collect signatures from P-chain validators
	// The signing subnet for ValidatorSetState is the subnet of the blockchain
	// we're reporting the validator set for
	signedMessage, err := u.signatureAggregator.CreateSignedMessage(
		ctx,
		u.logger,
		unsignedMessage,
		nil, // No justification needed for ValidatorSetState
		subnetID,
		defaultQuorumPercentage,
		0, // No buffer
		pChainHeight,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to aggregate signatures: %w", err)
	}

	u.logger.Info("Successfully signed ValidatorSetState message",
		zap.Stringer("messageID", signedMessage.ID()),
		zap.Stringer("subnetID", subnetID),
		zap.Uint64("pChainHeight", pChainHeight),
	)

	return signedMessage, nil
}

// sendValidatorSetUpdate sends the signed message to the external EVM.
// It checks if a validator set is already registered and uses updateValidatorSet if so,
// otherwise it uses registerValidatorSet for the initial registration.
func (u *ValidatorSetUpdater) sendValidatorSetUpdate(
	ctx context.Context,
	client *evm.ExternalEVMDestinationClient,
	signedMessage *warp.Message,
	validatorsBytes []byte,
) error {
	// Get the ABI for packing the call data
	registryABI, err := avalanchevalidatorsetregistry.AvalancheValidatorSetRegistryMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get registry ABI: %w", err)
	}

	// Convert warp.Message to ICMMessage format
	bitSetSig, ok := signedMessage.Signature.(*warp.BitSetSignature)
	if !ok {
		return fmt.Errorf("unexpected signature type: %T", signedMessage.Signature)
	}

	// Convert the 96-byte compressed signature to 192-byte uncompressed format
	// The contract expects uncompressed BLS signatures (192 bytes)
	sig, err := bls.SignatureFromBytes(bitSetSig.Signature[:])
	if err != nil {
		return fmt.Errorf("failed to parse signature: %w", err)
	}
	// Serialize returns the uncompressed 192-byte format
	uncompressedSig := sig.Serialize()

	icmMessage := avalanchevalidatorsetregistry.ICMMessage{
		UnsignedMessage: avalanchevalidatorsetregistry.ICMUnsignedMessage{
			AvalancheNetworkID:          u.networkID,
			AvalancheSourceBlockchainID: signedMessage.UnsignedMessage.SourceChainID,
			Payload:                     signedMessage.UnsignedMessage.Payload,
		},
		UnsignedMessageBytes: signedMessage.UnsignedMessage.Bytes(),
		Signature: avalanchevalidatorsetregistry.ICMSignature{
			Signers:   bitSetSig.Signers,
			Signature: uncompressedSig,
		},
	}

	// Check if a validator set is already registered by querying nextValidatorSetID
	nextID, err := client.GetNextValidatorSetID(ctx)
	if err != nil {
		return fmt.Errorf("failed to get next validator set ID: %w", err)
	}

	var callData []byte
	if nextID == 0 {
		// No validator sets registered yet - use registerValidatorSet
		u.logger.Info("No existing validator set, using registerValidatorSet")
		callData, err = registryABI.Pack("registerValidatorSet", icmMessage, validatorsBytes)
		if err != nil {
			return fmt.Errorf("failed to pack registerValidatorSet call data: %w", err)
		}
	} else {
		// Validator set exists - use updateValidatorSet with current validator set ID
		currentID, err := client.GetCurrentValidatorSetID(ctx)
		if err != nil {
			return fmt.Errorf("failed to get current validator set ID: %w", err)
		}
		u.logger.Info("Existing validator set found, using updateValidatorSet",
			zap.Uint64("currentValidatorSetID", currentID),
		)
		// Convert to *big.Int as required by the contract
		validatorSetID := new(big.Int).SetUint64(currentID)
		callData, err = registryABI.Pack("updateValidatorSet", validatorSetID, icmMessage, validatorsBytes)
		if err != nil {
			return fmt.Errorf("failed to pack updateValidatorSet call data: %w", err)
		}
	}

	u.logger.Info("Sending validator set update transaction",
		zap.Stringer("registryAddress", client.RegistryAddress()),
		zap.Int("callDataLen", len(callData)),
		zap.Int("validatorsBytesLen", len(validatorsBytes)),
		zap.Int("signersLen", len(bitSetSig.Signers)),
		zap.Int("signatureLen", len(uncompressedSig)),
		zap.Stringer("messageID", signedMessage.ID()),
		zap.Uint32("networkID", u.networkID),
		zap.Stringer("sourceChainID", signedMessage.UnsignedMessage.SourceChainID),
		zap.Binary("unsignedMessageBytes", signedMessage.UnsignedMessage.Bytes()),
	)

	// Simulate the call first to check for potential errors
	_, simErr := client.SimulateCall(ctx, client.RegistryAddress().Hex(), callData)
	if simErr != nil {
		u.logger.Error("Transaction simulation failed",
			zap.Error(simErr),
			zap.Stringer("registryAddress", client.RegistryAddress()),
			zap.Binary("signers", bitSetSig.Signers),
		)
		return fmt.Errorf("transaction simulation failed: %w", simErr)
	}

	// Send the transaction
	receipt, err := client.SendTx(
		signedMessage,
		nil, // Any sender can deliver
		client.RegistryAddress().Hex(),
		gasLimit,
		callData,
	)
	if err != nil {
		return fmt.Errorf("failed to send transaction: %w", err)
	}
	if receipt == nil {
		return fmt.Errorf("no receipt returned")
	}

	if receipt.Status != 1 {
		// Try to get more error details by simulating the call at the block where it failed
		blockNum := new(big.Int).SetUint64(receipt.BlockNumber.Uint64() - 1)
		_, revertErr := client.SimulateCallAtBlock(ctx, client.RegistryAddress().Hex(), callData, blockNum)
		u.logger.Error("Transaction reverted",
			zap.Stringer("txHash", receipt.TxHash),
			zap.Uint64("blockNumber", receipt.BlockNumber.Uint64()),
			zap.Uint64("gasUsed", receipt.GasUsed),
			zap.Uint64("gasLimit", gasLimit),
			zap.Error(revertErr),
		)
		return fmt.Errorf("transaction failed with status %d: %v", receipt.Status, revertErr)
	}

	u.logger.Info("Validator set update transaction confirmed",
		zap.Stringer("txHash", receipt.TxHash),
		zap.Uint64("gasUsed", receipt.GasUsed),
	)

	return nil
}

// computeValidatorSetHash computes a hash of the validator set for change detection.
// This uses the same hash computation as computeValidatorSetHashForMessage for consistency.
func (u *ValidatorSetUpdater) computeValidatorSetHash(validatorSet validators.WarpSet) [32]byte {
	// Validators are already in canonical order from the WarpSet
	// Hash using uncompressed public keys and weights
	h := sha256.New()
	for _, v := range validatorSet.Validators {
		// Use uncompressed public key bytes
		if len(v.PublicKeyBytes) == bls.PublicKeyLen {
			h.Write(v.PublicKeyBytes)
		} else if v.PublicKey != nil {
			h.Write(bls.PublicKeyToUncompressedBytes(v.PublicKey))
		}
		// Write weight as 8 bytes big-endian
		weightBytes := make([]byte, 8)
		weightBytes[0] = byte(v.Weight >> 56)
		weightBytes[1] = byte(v.Weight >> 48)
		weightBytes[2] = byte(v.Weight >> 40)
		weightBytes[3] = byte(v.Weight >> 32)
		weightBytes[4] = byte(v.Weight >> 24)
		weightBytes[5] = byte(v.Weight >> 16)
		weightBytes[6] = byte(v.Weight >> 8)
		weightBytes[7] = byte(v.Weight)
		h.Write(weightBytes)
	}
	var result [32]byte
	copy(result[:], h.Sum(nil))
	return result
}
