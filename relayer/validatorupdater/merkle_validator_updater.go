// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validatorupdater

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	warpvdrs "github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/logging"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	warppayload "github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	merklevalidatorsetregistry "github.com/ava-labs/icm-services/abi-bindings/go/MerkleValidatorSetRegistry"
	"github.com/ava-labs/icm-services/peers/clients"
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/ethclient"
	"go.uber.org/zap"
)

type MerkleSetUpdater struct {
	logger              logging.Logger
	pChainClient        clients.CanonicalValidatorState
	signatureAggregator *aggregator.SignatureAggregator
	ethClient           *ethclient.Client
	contract            *merklevalidatorsetregistry.MerkleValidatorSetRegistry
	contractAddress     common.Address
	txOpts              *bind.TransactOpts

	networkID    uint32
	blockchainID ids.ID
	subnetID     ids.ID
	pollInterval time.Duration

	maxUpdateInterval time.Duration

	localValidatorSet []*Validator
	localPChainHeight uint64
	localTotalWeight  uint64
	lastUpdateTime    time.Time
	initialized       bool
}

func NewMerkleSetUpdater(
	logger logging.Logger,
	pChainClient clients.CanonicalValidatorState,
	signatureAggregator *aggregator.SignatureAggregator,
	ethClient *ethclient.Client,
	contract *merklevalidatorsetregistry.MerkleValidatorSetRegistry,
	contractAddress common.Address,
	txOpts *bind.TransactOpts,
	networkID uint32,
	blockchainID ids.ID,
	subnetID ids.ID,
	pollInterval time.Duration,
	maxUpdateInterval time.Duration,
) *MerkleSetUpdater {
	return &MerkleSetUpdater{
		logger:              logger,
		pChainClient:        pChainClient,
		signatureAggregator: signatureAggregator,
		ethClient:           ethClient,
		contract:            contract,
		contractAddress:     contractAddress,
		txOpts:              txOpts,
		networkID:           networkID,
		blockchainID:        blockchainID,
		subnetID:            subnetID,
		pollInterval:        pollInterval,
		maxUpdateInterval:   maxUpdateInterval,
	}
}

// Start runs the polling loop that detects validator set changes and updates the contract.
func (s *MerkleSetUpdater) Start(ctx context.Context) error {
	s.logger.Info("Starting MerkleSetUpdater",
		zap.Stringer("blockchainID", s.blockchainID),
		zap.Duration("maxUpdateInterval", s.maxUpdateInterval),
	)

	if err := s.checkAndUpdate(ctx); err != nil {
		s.logger.Error("Initial update failed", zap.Error(err))
	}

	ticker := time.NewTicker(s.pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			s.logger.Info("MerkleSetUpdater stopping")
			return ctx.Err()
		case <-ticker.C:
			if err := s.checkAndUpdate(ctx); err != nil {
				s.logger.Error("Update check failed", zap.Error(err))
			}
		}
	}
}

func (s *MerkleSetUpdater) checkAndUpdate(ctx context.Context) error {
	if !s.initialized {
		return s.initializeLocalState(ctx)
	}

	pChainHeight, err := s.pChainClient.GetLatestHeight(ctx)
	if err != nil {
		return fmt.Errorf("failed to get P-chain height: %w", err)
	}

	if pChainHeight <= s.localPChainHeight {
		s.logger.Debug("No new P-chain blocks",
			zap.Uint64("pChainHeight", pChainHeight),
			zap.Uint64("localPChainHeight", s.localPChainHeight),
		)
		return nil
	}
	pChainTimestamp, err := s.pChainClient.GetBlockTimestampAtHeight(ctx, pChainHeight)
	if err != nil {
		return fmt.Errorf("failed to get P-chain timestamp: %w", err)
	}
	newValidators, err := s.fetchSortedValidators(ctx, pChainHeight)
	if err != nil {
		return fmt.Errorf("failed to fetch validators at height %d: %w", pChainHeight, err)
	}

	validatorSetUpdate, err := s.nextUpdate(newValidators, pChainHeight, pChainTimestamp)
	if err != nil {
		return fmt.Errorf("failed to build validator set commitment: %w", err)
	}
	if validatorSetUpdate == nil {
		s.logger.Debug("No validator changes and not stale, skipping")
		return nil
	}

	s.logger.Info("Validator set update triggered",
		zap.Uint64("localPChainHeight", s.localPChainHeight),
		zap.Uint64("pChainHeight", pChainHeight),
		zap.String("newCommitmentRoot", fmt.Sprintf("0x%x", validatorSetUpdate.RootHash[:])),
		zap.Bool("stale", s.isStale()),
	)

	if err := s.performUpdate(ctx, s.localPChainHeight, validatorSetUpdate, false); err != nil {
		s.logger.Warn("Merkle root update failed, re-syncing from contract on next tick", zap.Error(err))
		s.initialized = false
		return nil
	}

	s.localValidatorSet = newValidators
	s.localPChainHeight = pChainHeight
	s.localTotalWeight = sumWeights(newValidators)
	s.lastUpdateTime = time.Now()

	return nil
}

func (s *MerkleSetUpdater) initializeLocalState(ctx context.Context) error {
	onChainVS, err := s.contract.GetValidatorSetCommitment(&bind.CallOpts{Context: ctx}, s.blockchainID)
	if err != nil {
		return fmt.Errorf("failed to get on-chain validator set: %w", err)
	}

	isFirstRegistration := onChainVS.TotalWeight == 0
	if isFirstRegistration {
		pChainHeight, err := s.pChainClient.GetLatestHeight(ctx)
		if err != nil {
			return fmt.Errorf("failed to get P-chain height: %w", err)
		}
		pChainTimestamp, err := s.pChainClient.GetBlockTimestampAtHeight(ctx, pChainHeight)
		if err != nil {
			return fmt.Errorf("failed to get P-chain timestamp: %w", err)
		}
		newValidators, err := s.fetchSortedValidators(ctx, pChainHeight)
		if err != nil {
			return fmt.Errorf("failed to fetch validators after first registration: %w", err)
		}
		cmt, err := NewValidatorSetMerkleCommitment(s.blockchainID, newValidators, pChainHeight, pChainTimestamp)
		if err != nil {
			return fmt.Errorf("failed to build validator set commitment: %w", err)
		}

		s.logger.Info("First registration detected, performing update",
			zap.Uint64("pChainHeight", pChainHeight),
		)
		if err := s.performUpdate(ctx, onChainVS.PChainHeight, cmt, true); err != nil {
			return err
		}
		s.localValidatorSet = newValidators
		s.localPChainHeight = pChainHeight
		s.localTotalWeight = sumWeights(newValidators)
		s.lastUpdateTime = time.Now()
		s.initialized = true
		return nil
	}

	validators, err := s.fetchSortedValidators(ctx, onChainVS.PChainHeight)
	if err != nil {
		return fmt.Errorf("failed to fetch P-chain validators at on-chain height %d: %w",
			onChainVS.PChainHeight, err)
	}

	s.localValidatorSet = validators
	s.localPChainHeight = onChainVS.PChainHeight
	s.localTotalWeight = sumWeights(validators)
	s.lastUpdateTime = time.Now()
	s.initialized = true

	s.logger.Info("Initialized local validator set from contract",
		zap.Uint64("pChainHeight", onChainVS.PChainHeight),
		zap.Int("numValidators", len(validators)),
		zap.Uint64("totalWeight", s.localTotalWeight),
	)

	return nil
}

func (s *MerkleSetUpdater) isStale() bool {
	if s.maxUpdateInterval == 0 {
		return false
	}
	return time.Since(s.lastUpdateTime) >= s.maxUpdateInterval
}

func (s *MerkleSetUpdater) nextUpdate(
	newValidators []*Validator,
	pChainHeight uint64,
	pChainTimestamp uint64,
) (*ValidatorSetMerkleCommitment, error) {
	if BuildMerkleRoot(s.localValidatorSet) != BuildMerkleRoot(newValidators) || s.isStale() {
		return NewValidatorSetMerkleCommitment(s.blockchainID, newValidators, pChainHeight, pChainTimestamp)
	}
	return nil, nil
}

func (s *MerkleSetUpdater) performUpdate(
	ctx context.Context,
	onChainPChainHeight uint64,
	validatorSetUpdate *ValidatorSetMerkleCommitment,
	isFirstRegistration bool,
) error {
	var signingSubnet ids.ID
	if isFirstRegistration {
		// Contract verifies with the P-chain validator set; warp source is the P-chain.
		signingSubnet = constants.PrimaryNetworkID
	} else {
		// Contract verifies with the L1's registered set; preimage must use this chain ID.
		signingSubnet = s.subnetID
	}

	addressedCall, err := warppayload.NewAddressedCall(nil, validatorSetUpdate.Bytes())
	if err != nil {
		return fmt.Errorf("failed to create addressed call: %w", err)
	}

	unsignedMsg, err := avalancheWarp.NewUnsignedMessage(
		s.networkID,
		constants.PlatformChainID,
		addressedCall.Bytes(),
	)
	if err != nil {
		return fmt.Errorf("failed to create unsigned warp message: %w", err)
	}

	s.logger.Info("Signing new merkle root",
		zap.Bool("isFirstRegistration", isFirstRegistration),
		zap.Stringer("signingSubnet", signingSubnet),
	)

	signedMsg, err := s.signatureAggregator.CreateSignedMessage(
		ctx,
		s.logger,
		unsignedMsg,
		nil,
		signingSubnet,
		defaultQuorumPercentage,
		defaultQuorumPercentageBuf,
		onChainPChainHeight,
	)
	if err != nil {
		return fmt.Errorf("failed to sign message: %w", err)
	}

	return s.sendUpdate(ctx, signedMsg, onChainPChainHeight, isFirstRegistration)
}

func (s *MerkleSetUpdater) sendUpdate(
	ctx context.Context,
	signedMsg *avalancheWarp.Message,
	onChainPChainHeight uint64,
	isFirstRegistration bool,
) error {
	var attestationValidators []*Validator
	if isFirstRegistration {
		// First registration is verified against the P-chain Merkle root stored in the
		// contract constructor. Fetch the primary network validators used to build that
		// root so the attestation proof is computed against the correct ordered set.
		allValidatorSets, err := s.pChainClient.GetAllValidatorSets(ctx, onChainPChainHeight)
		if err != nil {
			return fmt.Errorf("failed to get P-chain validator sets at height %d for attestation: %w",
				onChainPChainHeight, err)
		}
		pChainWarpSet, ok := allValidatorSets[ids.Empty]
		if !ok {
			return fmt.Errorf("primary network not found in validator sets at height %d", onChainPChainHeight)
		}
		attestationValidators = warpSetToLocalValidators(pChainWarpSet)
	} else {
		attestationValidators = s.localValidatorSet
	}

	icmMessage, err := s.buildICMMessage(signedMsg, attestationValidators)
	if err != nil {
		return err
	}

	var tx *types.Transaction
	if isFirstRegistration {
		s.logger.Info("Sending registerValidatorSet (initial)")
		tx, err = s.contract.RegisterValidatorSet(s.txOpts, icmMessage, [32]byte(ids.Empty))
		if err != nil {
			return fmt.Errorf("registerValidatorSet failed: %w", err)
		}
	} else {
		s.logger.Info("Sending registerValidatorSet (update)")
		tx, err = s.contract.RegisterValidatorSet(s.txOpts, icmMessage, [32]byte(s.blockchainID))
		if err != nil {
			return fmt.Errorf("registerValidatorSet (update) failed: %w", err)
		}
	}

	receipt, err := bind.WaitMined(ctx, s.ethClient, tx)
	if err != nil {
		return fmt.Errorf("waiting for validator set tx: %w", err)
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return fmt.Errorf("validator set tx reverted: %s", tx.Hash().Hex())
	}
	s.logger.Info("validator set tx confirmed",
		zap.String("txHash", tx.Hash().Hex()),
		zap.Uint64("blockNumber", receipt.BlockNumber.Uint64()),
		zap.Bool("isFirstRegistration", isFirstRegistration),
	)
	return nil
}

// ---------------------------------------------------------------------------
// Shared helpers
// ---------------------------------------------------------------------------

func (s *MerkleSetUpdater) fetchSortedValidators(
	ctx context.Context,
	pChainHeight uint64,
) ([]*Validator, error) {
	subnetValidatorSet, err := s.pChainClient.GetValidatorsAt(ctx, s.subnetID, pChainHeight)
	if err != nil {
		return nil, fmt.Errorf("failed to get validator sets: %w", err)
	}

	validators := make([]*Validator, len(subnetValidatorSet))
	ix := 0
	for _, vdr := range subnetValidatorSet {
		validators[ix] = &Validator{
			UncompressedPublicKeyBytes: [96]byte(vdr.PublicKey.Serialize()),
			Weight:                     vdr.Weight,
		}
		ix++
	}
	sort.Slice(validators, func(i, j int) bool {
		return string(validators[i].UncompressedPublicKeyBytes[:]) < string(validators[j].UncompressedPublicKeyBytes[:])
	})

	return validators, nil
}

func (s *MerkleSetUpdater) buildICMMessage(
	signedMsg *avalancheWarp.Message,
	validators []*Validator,
) (merklevalidatorsetregistry.ICMMessage, error) {
	addressedCall, err := warppayload.ParseAddressedCall(signedMsg.UnsignedMessage.Payload)
	if err != nil {
		return merklevalidatorsetregistry.ICMMessage{},
			fmt.Errorf("failed to parse addressed call from signed message: %w", err)
	}

	bitSetSig, ok := signedMsg.Signature.(*avalancheWarp.BitSetSignature)
	if !ok {
		return merklevalidatorsetregistry.ICMMessage{}, fmt.Errorf("expected BitSetSignature, got %T", signedMsg.Signature)
	}

	attestation, err := NewValidatorSetMerkleAttestation(validators, bitSetSig)
	if err != nil {
		return merklevalidatorsetregistry.ICMMessage{}, fmt.Errorf("failed to build validator set attestation: %w", err)
	}

	return merklevalidatorsetregistry.ICMMessage{
		RawMessage:         addressedCall.Payload,
		SourceNetworkID:    signedMsg.UnsignedMessage.NetworkID,
		SourceBlockchainID: signedMsg.UnsignedMessage.SourceChainID,
		Attestation:        attestation.Bytes(),
	}, nil
}

// warpSetToLocalValidators converts a canonical WarpSet to the []*Validator form
// expected by the merkle attestation builder. The returned slice is sorted by
// uncompressed public key bytes to match the canonical ordering used to build
// the on-chain merkle root.
func warpSetToLocalValidators(warpSet warpvdrs.WarpSet) []*Validator {
	out := make([]*Validator, len(warpSet.Validators))
	for i, vdr := range warpSet.Validators {
		out[i] = &Validator{
			UncompressedPublicKeyBytes: [96]byte(vdr.PublicKey.Serialize()),
			Weight:                     vdr.Weight,
		}
	}
	sort.Slice(out, func(i, j int) bool {
		return string(out[i].UncompressedPublicKeyBytes[:]) < string(out[j].UncompressedPublicKeyBytes[:])
	})
	return out
}
