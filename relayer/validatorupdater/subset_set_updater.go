// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validatorupdater

import (
	"context"
	"crypto/sha256"
	"fmt"
	"sort"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/logging"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/message"
	warppayload "github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	subsetupdater "github.com/ava-labs/icm-services/abi-bindings/go/SubsetUpdater"
	"github.com/ava-labs/icm-services/peers/clients"
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/ethclient"
	"go.uber.org/zap"
)

const (
	defaultPollInterval               = 10 * time.Second
	defaultQuorumPercentage           = 67
	defaultQuorumPercentageBuf        = 5
	defaultShardSize           uint32 = 10
)

type SubsetSetUpdater struct {
	logger              logging.Logger
	pChainClient        clients.CanonicalValidatorState
	signatureAggregator *aggregator.SignatureAggregator
	ethClient           *ethclient.Client
	contract            *subsetupdater.SubsetUpdater
	contractAddress     common.Address
	txOpts              *bind.TransactOpts

	networkID    uint32
	blockchainID ids.ID
	subnetID     ids.ID
	shardSize    uint32
	pollInterval time.Duration
}

func NewSubsetSetUpdater(
	logger logging.Logger,
	pChainClient clients.CanonicalValidatorState,
	signatureAggregator *aggregator.SignatureAggregator,
	ethClient *ethclient.Client,
	contract *subsetupdater.SubsetUpdater,
	contractAddress common.Address,
	txOpts *bind.TransactOpts,
	networkID uint32,
	blockchainID ids.ID,
	subnetID ids.ID,
	shardSize uint32,
	pollInterval time.Duration,
) *SubsetSetUpdater {
	if shardSize == 0 {
		shardSize = defaultShardSize
	}
	if pollInterval == 0 {
		pollInterval = defaultPollInterval
	}
	return &SubsetSetUpdater{
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
		shardSize:           shardSize,
		pollInterval:        pollInterval,
	}
}

// Start runs the polling loop that detects validator set changes and updates the contract.
func (s *SubsetSetUpdater) Start(ctx context.Context) error {
	s.logger.Info("Starting SubsetSetUpdater",
		zap.Stringer("blockchainID", s.blockchainID),
		zap.Uint32("shardSize", s.shardSize),
	)

	if err := s.checkAndUpdate(ctx); err != nil {
		s.logger.Error("Initial update failed", zap.Error(err))
	}

	ticker := time.NewTicker(s.pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			s.logger.Info("SubsetSetUpdater stopping")
			return ctx.Err()
		case <-ticker.C:
			if err := s.checkAndUpdate(ctx); err != nil {
				s.logger.Error("Update check failed", zap.Error(err))
			}
		}
	}
}

func (s *SubsetSetUpdater) checkAndUpdate(ctx context.Context) error {
	pChainHeight, err := s.pChainClient.GetLatestHeight(ctx)
	if err != nil {
		return fmt.Errorf("failed to get P-chain height: %w", err)
	}

	onChainVS, err := s.contract.GetValidatorSet(&bind.CallOpts{Context: ctx}, s.blockchainID)
	if err != nil {
		return fmt.Errorf("failed to get on-chain validator set: %w", err)
	}

	if onChainVS.PChainHeight >= pChainHeight {
		s.logger.Debug("On-chain validator set is up to date",
			zap.Uint64("pChainHeight", pChainHeight),
			zap.Uint64("onChainHeight", onChainVS.PChainHeight),
		)
		return nil
	}

	isFirstRegistration := onChainVS.TotalWeight == 0

	s.logger.Info("Validator set update needed",
		zap.Uint64("onChainHeight", onChainVS.PChainHeight),
		zap.Uint64("pChainHeight", pChainHeight),
		zap.Bool("isFirstRegistration", isFirstRegistration),
	)

	return s.performFullSetUpdate(ctx, pChainHeight, onChainVS.PChainHeight, isFirstRegistration)
}

// ---------------------------------------------------------------------------
// Full-set (subset/shard) update
// ---------------------------------------------------------------------------

func (s *SubsetSetUpdater) performFullSetUpdate(
	ctx context.Context,
	pChainHeight uint64,
	onChainPChainHeight uint64,
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

	shardBytesList, subsetUpdateMsg, err := s.buildSubsetUpdate(ctx, pChainHeight)
	if err != nil {
		return fmt.Errorf("failed to build subset update: %w", err)
	}

	s.logger.Info("Signing subset update",
		zap.Bool("isFirstRegistration", isFirstRegistration),
		zap.Stringer("signingSubnet", signingSubnet),
	)

	signedMsg, err := s.signatureAggregator.CreateSignedMessage(
		ctx,
		s.logger,
		subsetUpdateMsg,
		nil,
		signingSubnet,
		defaultQuorumPercentage,
		defaultQuorumPercentageBuf,
		onChainPChainHeight,
	)
	if err != nil {
		return fmt.Errorf("failed to sign message: %w", err)
	}

	return s.sendSubsetUpdate(ctx, signedMsg, shardBytesList)
}

func (s *SubsetSetUpdater) buildSubsetUpdate(
	ctx context.Context,
	pChainHeight uint64,
) ([][]byte, *avalancheWarp.UnsignedMessage, error) {
	validators, err := s.fetchSortedValidators(ctx, pChainHeight)
	if err != nil {
		return nil, nil, err
	}

	shardBytesList, shardHashes, err := ShardValidators(validators, int(s.shardSize))
	if err != nil {
		return nil, nil, err
	}

	pChainTimestamp, err := s.pChainClient.GetBlockTimestampAtHeight(ctx, pChainHeight)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get P-chain block timestamp at height %d: %w", pChainHeight, err)
	}

	subsetUpdatePayload, err := NewValidatorSetMetadata(
		s.blockchainID,
		pChainHeight,
		pChainTimestamp,
		shardHashes,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create SubsetUpdate: %w", err)
	}

	addressedCall, err := warppayload.NewAddressedCall(nil, subsetUpdatePayload.Bytes())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create addressed call: %w", err)
	}

	unsignedMsg, err := avalancheWarp.NewUnsignedMessage(
		s.networkID,
		constants.PlatformChainID,
		addressedCall.Bytes(),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create unsigned warp message: %w", err)
	}

	s.logger.Info("Built SubsetUpdate message",
		zap.Uint64("pChainHeight", pChainHeight),
		zap.Uint64("pChainTimestamp", pChainTimestamp),
		zap.Int("numValidators", len(validators)),
		zap.Int("numShards", len(shardBytesList)),
	)

	return shardBytesList, unsignedMsg, nil
}

func (s *SubsetSetUpdater) sendSubsetUpdate(
	ctx context.Context,
	signedMsg *avalancheWarp.Message,
	shardBytesList [][]byte,
) error {
	icmMessage, err := buildICMMessage(signedMsg)
	if err != nil {
		return err
	}

	s.logger.Info("Sending registerValidatorSet",
		zap.Int("totalShards", len(shardBytesList)),
	)
	tx, err := s.contract.RegisterValidatorSet(s.txOpts, icmMessage, shardBytesList[0])
	if err != nil {
		return fmt.Errorf("registerValidatorSet failed: %w", err)
	}
	receipt, err := bind.WaitMined(ctx, s.ethClient, tx)
	if err != nil {
		return fmt.Errorf("waiting for registerValidatorSet tx: %w", err)
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return fmt.Errorf("registerValidatorSet tx reverted: %s", tx.Hash().Hex())
	}
	s.logger.Info("registerValidatorSet confirmed",
		zap.String("txHash", tx.Hash().Hex()),
		zap.Uint64("blockNumber", receipt.BlockNumber.Uint64()),
	)

	for i := 1; i < len(shardBytesList); i++ {
		shard := subsetupdater.ValidatorSetShard{
			ShardNumber:           uint64(i + 1),
			AvalancheBlockchainID: s.blockchainID,
		}
		s.logger.Info("Sending updateValidatorSet",
			zap.Int("shardNumber", i+1),
		)
		tx, err := s.contract.UpdateValidatorSet(s.txOpts, shard, shardBytesList[i])
		if err != nil {
			return fmt.Errorf("updateValidatorSet (shard %d) failed: %w", i+1, err)
		}
		receipt, err := bind.WaitMined(ctx, s.ethClient, tx)
		if err != nil {
			return fmt.Errorf("waiting for updateValidatorSet tx (shard %d): %w", i+1, err)
		}
		if receipt.Status == types.ReceiptStatusFailed {
			return fmt.Errorf("updateValidatorSet tx (shard %d) reverted: %s", i+1, tx.Hash().Hex())
		}
		s.logger.Info("updateValidatorSet confirmed",
			zap.Int("shardNumber", i+1),
			zap.String("txHash", tx.Hash().Hex()),
		)
	}

	s.logger.Info("All shards submitted successfully",
		zap.Int("numShards", len(shardBytesList)),
	)
	return nil
}

// ---------------------------------------------------------------------------
// Shared helpers
// ---------------------------------------------------------------------------

func (s *SubsetSetUpdater) fetchSortedValidators(
	ctx context.Context,
	pChainHeight uint64,
) ([]*Validator, error) {
	allValidatorSets, err := s.pChainClient.GetAllValidatorSets(ctx, pChainHeight)
	if err != nil {
		return nil, fmt.Errorf("failed to get validator sets: %w", err)
	}

	vdrSet, ok := allValidatorSets[s.subnetID]
	if !ok {
		return nil, fmt.Errorf("subnet %s not found in validator sets at height %d", s.subnetID, pChainHeight)
	}

	validators := make([]*Validator, len(vdrSet.Validators))
	for i, vdr := range vdrSet.Validators {
		validators[i] = &Validator{
			UncompressedPublicKeyBytes: [96]byte(vdr.PublicKey.Serialize()),
			Weight:                     vdr.Weight,
		}
	}
	sort.Slice(validators, func(i, j int) bool {
		return string(validators[i].UncompressedPublicKeyBytes[:]) < string(validators[j].UncompressedPublicKeyBytes[:])
	})

	return validators, nil
}

// ShardValidators splits a sorted validator slice into shards, marshaling
// each shard and computing its sha256 hash.
func ShardValidators(
	validators []*Validator,
	shardSize int,
) ([][]byte, []ids.ID, error) {
	numValidators := len(validators)
	numShards := (numValidators + shardSize - 1) / shardSize
	if numShards == 0 {
		numShards = 1
	}

	shardHashes := make([]ids.ID, numShards)
	shardBytesList := make([][]byte, numShards)
	for i := 0; i < numShards; i++ {
		start := i * shardSize
		end := start + shardSize
		if end > numValidators {
			end = numValidators
		}
		shard := validators[start:end]
		shardBytes, err := message.Codec.Marshal(message.CodecVersion, shard)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to marshal shard %d: %w", i, err)
		}
		hash := sha256.Sum256(shardBytes)
		shardHashes[i] = ids.ID(hash)
		shardBytesList[i] = shardBytes
	}
	return shardBytesList, shardHashes, nil
}

func buildICMMessage(signedMsg *avalancheWarp.Message) (subsetupdater.ICMMessage, error) {
	addressedCall, err := warppayload.ParseAddressedCall(signedMsg.UnsignedMessage.Payload)
	if err != nil {
		return subsetupdater.ICMMessage{}, fmt.Errorf("failed to parse addressed call from signed message: %w", err)
	}

	bitSetSig, ok := signedMsg.Signature.(*avalancheWarp.BitSetSignature)
	if !ok {
		return subsetupdater.ICMMessage{}, fmt.Errorf("expected BitSetSignature, got %T", signedMsg.Signature)
	}

	// The contract expects uncompressed BLS signatures (192 bytes) whereas the
	// warp BitSetSignature stores compressed signatures (96 bytes).
	sig, err := bls.SignatureFromBytes(bitSetSig.Signature[:])
	if err != nil {
		return subsetupdater.ICMMessage{}, fmt.Errorf("failed to decompress BLS signature: %w", err)
	}
	uncompressedSig := sig.Serialize()

	// Contract-expected attestation format: raw signers bitset || uncompressed BLS signature (192 bytes)
	attestation := make([]byte, 0, len(bitSetSig.Signers)+len(uncompressedSig))
	attestation = append(attestation, bitSetSig.Signers...)
	attestation = append(attestation, uncompressedSig...)

	return subsetupdater.ICMMessage{
		RawMessage:         addressedCall.Payload,
		SourceNetworkID:    signedMsg.UnsignedMessage.NetworkID,
		SourceBlockchainID: signedMsg.UnsignedMessage.SourceChainID,
		Attestation:        attestation,
	}, nil
}
