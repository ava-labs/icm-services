// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package valiatorupdater

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"math/big"
	"sort"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/logging"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/message"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	diffupdater "github.com/ava-labs/icm-services/abi-bindings/go/DiffUpdater"
	"github.com/ava-labs/icm-services/peers/clients"
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/ethclient"
	"go.uber.org/zap"
)

type DiffSetUpdater struct {
	logger              logging.Logger
	pChainClient        clients.CanonicalValidatorState
	signatureAggregator *aggregator.SignatureAggregator
	ethClient           *ethclient.Client
	contract            *diffupdater.DiffUpdater
	contractAddress     common.Address
	txOpts              *bind.TransactOpts

	networkID    uint32
	blockchainID ids.ID
	subnetID     ids.ID
	shardSize    uint32
	pollInterval time.Duration
}

type DiffSetUpdaterConfig struct {
	Logger              logging.Logger
	PChainClient        clients.CanonicalValidatorState
	SignatureAggregator *aggregator.SignatureAggregator
	EthClient           *ethclient.Client
	Contract            *diffupdater.DiffUpdater
	ContractAddress     common.Address
	TxOpts              *bind.TransactOpts

	NetworkID    uint32
	BlockchainID ids.ID
	SubnetID     ids.ID
	ShardSize    uint32
	PollInterval time.Duration
}

func NewDiffSetUpdater(cfg DiffSetUpdaterConfig) *DiffSetUpdater {
	shardSize := cfg.ShardSize
	if shardSize == 0 {
		shardSize = defaultShardSize
	}
	pollInterval := cfg.PollInterval
	if pollInterval == 0 {
		pollInterval = defaultPollInterval
	}
	return &DiffSetUpdater{
		logger:              cfg.Logger,
		pChainClient:        cfg.PChainClient,
		signatureAggregator: cfg.SignatureAggregator,
		ethClient:           cfg.EthClient,
		contract:            cfg.Contract,
		contractAddress:     cfg.ContractAddress,
		txOpts:              cfg.TxOpts,
		networkID:           cfg.NetworkID,
		blockchainID:        cfg.BlockchainID,
		subnetID:            cfg.SubnetID,
		shardSize:           shardSize,
		pollInterval:        pollInterval,
	}
}

// Start runs the polling loop that detects validator set changes and updates the contract.
func (d *DiffSetUpdater) Start(ctx context.Context) error {
	d.logger.Info("Starting DiffSetUpdater",
		zap.Stringer("blockchainID", d.blockchainID),
		zap.Uint32("shardSize", d.shardSize),
	)

	if err := d.checkAndUpdate(ctx); err != nil {
		d.logger.Error("Initial update failed", zap.Error(err))
	}

	ticker := time.NewTicker(d.pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			d.logger.Info("DiffSetUpdater stopping")
			return ctx.Err()
		case <-ticker.C:
			if err := d.checkAndUpdate(ctx); err != nil {
				d.logger.Error("Update check failed", zap.Error(err))
			}
		}
	}
}

func (d *DiffSetUpdater) checkAndUpdate(ctx context.Context) error {
	pChainHeight, err := d.pChainClient.GetLatestHeight(ctx)
	if err != nil {
		return fmt.Errorf("failed to get P-chain height: %w", err)
	}

	onChainVS, err := d.contract.GetValidatorSet(&bind.CallOpts{Context: ctx}, d.blockchainID)
	if err != nil {
		return fmt.Errorf("failed to get on-chain validator set: %w", err)
	}

	if onChainVS.PChainHeight >= pChainHeight {
		d.logger.Debug("On-chain validator set is up to date",
			zap.Uint64("pChainHeight", pChainHeight),
			zap.Uint64("onChainHeight", onChainVS.PChainHeight),
		)
		return nil
	}

	isFirstRegistration := onChainVS.TotalWeight == 0

	d.logger.Info("Validator set update needed",
		zap.Uint64("onChainHeight", onChainVS.PChainHeight),
		zap.Uint64("pChainHeight", pChainHeight),
		zap.Bool("isFirstRegistration", isFirstRegistration),
		zap.Int("onChainValidatorCount", len(onChainVS.Validators)),
		zap.Uint64("onChainTotalWeight", onChainVS.TotalWeight),
	)

	if isFirstRegistration {
		return d.performFullSetUpdate(ctx, pChainHeight, onChainVS)
	}
	return d.performDiffUpdate(ctx, pChainHeight, onChainVS)
}

// ---------------------------------------------------------------------------
// First registration: treat all validators as additions in a diff
// ---------------------------------------------------------------------------

func (d *DiffSetUpdater) performFullSetUpdate(
	ctx context.Context,
	pChainHeight uint64,
	onChainVS diffupdater.ValidatorSet,
) error {
	newValidators, err := d.fetchSortedValidators(ctx, pChainHeight)
	if err != nil {
		return fmt.Errorf("failed to fetch validators: %w", err)
	}

	changes := make([]message.ValidatorChange, len(newValidators))
	for i, v := range newValidators {
		changes[i] = message.ValidatorChange{
			UncompressedPublicKeyBytes: v.UncompressedPublicKeyBytes,
			Weight:                     v.Weight,
		}
	}

	pChainTimestamp := uint64(time.Now().Unix())

	shardBytesList, shardHashes, err := d.shardDiff(
		d.blockchainID,
		onChainVS.PChainHeight,
		onChainVS.PChainTimestamp,
		pChainHeight,
		pChainTimestamp,
		nil, // empty starting set for first registration
		changes,
	)
	if err != nil {
		return fmt.Errorf("failed to shard diff: %w", err)
	}

	metadataMsg, err := message.NewValidatorSetMetadata(
		d.blockchainID,
		pChainHeight,
		pChainTimestamp,
		shardHashes,
	)
	if err != nil {
		return fmt.Errorf("failed to create ValidatorSetMetadata: %w", err)
	}

	addressedCall, err := payload.NewAddressedCall(nil, metadataMsg.Bytes())
	if err != nil {
		return fmt.Errorf("failed to create addressed call: %w", err)
	}

	unsignedMsg, err := avalancheWarp.NewUnsignedMessage(
		d.networkID,
		constants.PlatformChainID,
		addressedCall.Bytes(),
	)
	if err != nil {
		return fmt.Errorf("failed to create unsigned warp message: %w", err)
	}

	// First L1 registration: contract verifies signatures against the P-chain set.
	signingSubnet := constants.PrimaryNetworkID
	d.logger.Info("Signing diff (first L1 registration)",
		zap.Stringer("signingSubnet", signingSubnet),
	)

	signedMsg, err := d.signatureAggregator.CreateSignedMessage(
		ctx,
		d.logger,
		unsignedMsg,
		nil,
		signingSubnet,
		defaultQuorumPercentage,
		defaultQuorumPercentageBuf,
		onChainVS.PChainHeight,
	)
	if err != nil {
		return fmt.Errorf("failed to sign message: %w", err)
	}

	return d.sendDiffUpdate(ctx, signedMsg, shardBytesList)
}

// ---------------------------------------------------------------------------
// Diff update: compute changes between on-chain and P-chain validator sets
// ---------------------------------------------------------------------------

func (d *DiffSetUpdater) performDiffUpdate(
	ctx context.Context,
	pChainHeight uint64,
	onChainVS diffupdater.ValidatorSet,
) error {
	newValidators, err := d.fetchSortedValidators(ctx, pChainHeight)
	if err != nil {
		return fmt.Errorf("failed to fetch validators: %w", err)
	}

	oldValidators := onChainValidatorsToMessage(onChainVS.Validators)
	sort.Slice(oldValidators, func(i, j int) bool {
		return bytes.Compare(
			oldValidators[i].UncompressedPublicKeyBytes[:],
			oldValidators[j].UncompressedPublicKeyBytes[:],
		) < 0
	})

	changes, _ := computeValidatorDiff(oldValidators, newValidators)

	if len(changes) == 0 {
		d.logger.Info("No validator changes detected, skipping update")
		return nil
	}

	pChainTimestamp := uint64(time.Now().Unix())

	shardBytesList, shardHashes, err := d.shardDiff(
		d.blockchainID,
		onChainVS.PChainHeight,
		onChainVS.PChainTimestamp,
		pChainHeight,
		pChainTimestamp,
		oldValidators,
		changes,
	)
	if err != nil {
		return fmt.Errorf("failed to shard diff: %w", err)
	}

	metadataMsg, err := message.NewValidatorSetMetadata(
		d.blockchainID,
		pChainHeight,
		pChainTimestamp,
		shardHashes,
	)
	if err != nil {
		return fmt.Errorf("failed to create ValidatorSetMetadata: %w", err)
	}

	addressedCall, err := payload.NewAddressedCall(nil, metadataMsg.Bytes())
	if err != nil {
		return fmt.Errorf("failed to create addressed call: %w", err)
	}

	unsignedMsg, err := avalancheWarp.NewUnsignedMessage(
		d.networkID,
		constants.PlatformChainID,
		addressedCall.Bytes(),
	)
	if err != nil {
		return fmt.Errorf("failed to create unsigned warp message: %w", err)
	}

	// L1 already registered: contract verifyICMMessage uses that chain's registered
	// BLS keys (same as SubsetSetUpdater after first registration).
	signingSubnet := d.subnetID
	d.logger.Info("Signing diff update", zap.Stringer("signingSubnet", signingSubnet))

	signedMsg, err := d.signatureAggregator.CreateSignedMessage(
		ctx,
		d.logger,
		unsignedMsg,
		nil,
		signingSubnet,
		defaultQuorumPercentage,
		defaultQuorumPercentageBuf,
		onChainVS.PChainHeight,
	)
	if err != nil {
		return fmt.Errorf("failed to sign message: %w", err)
	}

	d.logger.Info("Sending diff update",
		zap.Int("numChanges", len(changes)),
		zap.Int("numShards", len(shardBytesList)),
	)

	return d.sendDiffUpdate(ctx, signedMsg, shardBytesList)
}

// ---------------------------------------------------------------------------
// Diff computation
// ---------------------------------------------------------------------------

// computeValidatorDiff performs an O(n+m) merge-walk over two sorted validator
// slices, producing a sorted list of changes and the count of additions.
func computeValidatorDiff(
	old, new_ []*message.Validator,
) ([]message.ValidatorChange, uint32) {
	var changes []message.ValidatorChange
	var numAdded uint32

	oi, ni := 0, 0
	for oi < len(old) || ni < len(new_) {
		var cmp int
		switch {
		case oi >= len(old):
			cmp = 1
		case ni >= len(new_):
			cmp = -1
		default:
			cmp = bytes.Compare(old[oi].UncompressedPublicKeyBytes[:], new_[ni].UncompressedPublicKeyBytes[:])
		}

		switch {
		case cmp < 0:
			// Removal: exists in old but not in new
			changes = append(changes, message.ValidatorChange{
				UncompressedPublicKeyBytes: old[oi].UncompressedPublicKeyBytes,
				Weight:                     0,
			})
			oi++
		case cmp > 0:
			// Addition: exists in new but not in old
			changes = append(changes, message.ValidatorChange{
				UncompressedPublicKeyBytes: new_[ni].UncompressedPublicKeyBytes,
				Weight:                     new_[ni].Weight,
			})
			numAdded++
			ni++
		default:
			// Same key: check if weight changed
			if old[oi].Weight != new_[ni].Weight {
				changes = append(changes, message.ValidatorChange{
					UncompressedPublicKeyBytes: new_[ni].UncompressedPublicKeyBytes,
					Weight:                     new_[ni].Weight,
				})
			}
			oi++
			ni++
		}
	}

	return changes, numAdded
}

// ---------------------------------------------------------------------------
// Merge-walk: mirrors Solidity's ValidatorSets.applyValidatorSetDiff
// ---------------------------------------------------------------------------

// applyChangesToValidators performs the same sorted merge-walk as the Solidity
// applyValidatorSetDiff function. Both currentSet and changes must be sorted
// by UncompressedPublicKeyBytes. Returns the new validator set.
func applyChangesToValidators(
	currentSet []*message.Validator,
	changes []message.ValidatorChange,
) []*message.Validator {
	var result []*message.Validator
	vi, ci := 0, 0
	for vi < len(currentSet) || ci < len(changes) {
		var cmp int
		switch {
		case ci >= len(changes):
			cmp = -1
		case vi >= len(currentSet):
			cmp = 1
		default:
			cmp = bytes.Compare(
				currentSet[vi].UncompressedPublicKeyBytes[:],
				changes[ci].UncompressedPublicKeyBytes[:],
			)
		}

		switch {
		case cmp < 0:
			result = append(result, currentSet[vi])
			vi++
		case cmp > 0:
			result = append(result, &message.Validator{
				UncompressedPublicKeyBytes: changes[ci].UncompressedPublicKeyBytes,
				Weight:                     changes[ci].Weight,
			})
			ci++
		default:
			if changes[ci].Weight != 0 {
				result = append(result, &message.Validator{
					UncompressedPublicKeyBytes: changes[ci].UncompressedPublicKeyBytes,
					Weight:                     changes[ci].Weight,
				})
			}
			vi++
			ci++
		}
	}
	return result
}

// ---------------------------------------------------------------------------
// Sharding
// ---------------------------------------------------------------------------

// shardDiff splits the changes into shard-sized chunks. Each shard carries a
// running currentValidatorSetHash that matches the result of applying that
// shard's changes to the running set (starting from oldValidators). This
// mirrors the per-shard hash check in the DiffUpdater contract.
func (d *DiffSetUpdater) shardDiff(
	blockchainID ids.ID,
	prevHeight, prevTimestamp uint64,
	currHeight, currTimestamp uint64,
	oldValidators []*message.Validator,
	changes []message.ValidatorChange,
) ([][]byte, []ids.ID, error) {
	ss := int(d.shardSize)
	numChanges := len(changes)
	numShards := (numChanges + ss - 1) / ss
	if numShards == 0 {
		numShards = 1
	}

	shardBytesList := make([][]byte, numShards)
	shardHashes := make([]ids.ID, numShards)
	runningSet := oldValidators

	for i := 0; i < numShards; i++ {
		start := i * ss
		end := start + ss
		if end > numChanges {
			end = numChanges
		}
		shardChanges := changes[start:end]

		existingKeys := make(map[[96]byte]struct{}, len(runningSet))
		for _, v := range runningSet {
			existingKeys[v.UncompressedPublicKeyBytes] = struct{}{}
		}
		var shardNumAdded uint32
		for _, c := range shardChanges {
			if c.Weight > 0 {
				if _, exists := existingKeys[c.UncompressedPublicKeyBytes]; !exists {
					shardNumAdded++
				}
			}
		}

		runningSet = applyChangesToValidators(runningSet, shardChanges)

		diff, err := message.NewValidatorSetDiff(
			blockchainID,
			prevHeight,
			prevTimestamp,
			currHeight,
			currTimestamp,
			shardChanges,
			shardNumAdded,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create ValidatorSetDiff for shard %d: %w", i, err)
		}

		shardBytes := diff.Bytes()
		hash := sha256.Sum256(shardBytes)
		shardBytesList[i] = shardBytes
		shardHashes[i] = ids.ID(hash)
	}

	return shardBytesList, shardHashes, nil
}

// ---------------------------------------------------------------------------
// Sending transactions
// ---------------------------------------------------------------------------

func (d *DiffSetUpdater) sendDiffUpdate(
	ctx context.Context,
	signedMsg *avalancheWarp.Message,
	shardBytesList [][]byte,
) error {
	icmMessage, err := buildDiffICMMessage(signedMsg)
	if err != nil {
		return err
	}

	d.logger.Info("Sending registerValidatorSet",
		zap.Int("totalShards", len(shardBytesList)),
		zap.Int("firstShardSize", len(shardBytesList[0])),
	)
	tx, err := d.contract.RegisterValidatorSet(d.txOpts, icmMessage, shardBytesList[0])
	if err != nil {
		return fmt.Errorf("registerValidatorSet failed: %w", err)
	}
	receipt, err := bind.WaitMined(ctx, d.ethClient, tx)
	if err != nil {
		return fmt.Errorf("waiting for registerValidatorSet tx: %w", err)
	}
	if receipt.Status == 0 {
		return fmt.Errorf("registerValidatorSet tx reverted: %s", tx.Hash().Hex())
	}
	d.logger.Info("registerValidatorSet confirmed",
		zap.String("txHash", tx.Hash().Hex()),
		zap.Uint64("blockNumber", receipt.BlockNumber.Uint64()),
	)

	for i := 1; i < len(shardBytesList); i++ {
		shard := diffupdater.ValidatorSetShard{
			ShardNumber:           uint64(i + 1),
			AvalancheBlockchainID: d.blockchainID,
		}
		d.logger.Info("Sending updateValidatorSet",
			zap.Int("shardNumber", i+1),
		)
		tx, err := d.contract.UpdateValidatorSet(d.txOpts, shard, shardBytesList[i])
		if err != nil {
			return fmt.Errorf("updateValidatorSet (shard %d) failed: %w", i+1, err)
		}
		receipt, err := bind.WaitMined(ctx, d.ethClient, tx)
		if err != nil {
			return fmt.Errorf("waiting for updateValidatorSet tx (shard %d): %w", i+1, err)
		}
		if receipt.Status == 0 {
			return fmt.Errorf("updateValidatorSet tx (shard %d) reverted: %s", i+1, tx.Hash().Hex())
		}
		d.logger.Info("updateValidatorSet confirmed",
			zap.Int("shardNumber", i+1),
			zap.String("txHash", tx.Hash().Hex()),
		)
	}

	d.logger.Info("All shards submitted successfully",
		zap.Int("numShards", len(shardBytesList)),
	)
	return nil
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func (d *DiffSetUpdater) fetchSortedValidators(
	ctx context.Context,
	pChainHeight uint64,
) ([]*message.Validator, error) {
	allValidatorSets, err := d.pChainClient.GetAllValidatorSets(ctx, pChainHeight)
	if err != nil {
		return nil, fmt.Errorf("failed to get validator sets: %w", err)
	}

	vdrSet, ok := allValidatorSets[d.subnetID]
	if !ok {
		return nil, fmt.Errorf("subnet %s not found in validator sets at height %d", d.subnetID, pChainHeight)
	}

	validators := make([]*message.Validator, len(vdrSet.Validators))
	for i, vdr := range vdrSet.Validators {
		validators[i] = &message.Validator{
			UncompressedPublicKeyBytes: [96]byte(vdr.PublicKey.Serialize()),
			Weight:                     vdr.Weight,
		}
	}
	sort.Slice(validators, func(i, j int) bool {
		return string(validators[i].UncompressedPublicKeyBytes[:]) < string(validators[j].UncompressedPublicKeyBytes[:])
	})

	return validators, nil
}

func buildDiffICMMessage(signedMsg *avalancheWarp.Message) (diffupdater.ICMMessage, error) {
	addressedCall, err := payload.ParseAddressedCall(signedMsg.UnsignedMessage.Payload)
	if err != nil {
		return diffupdater.ICMMessage{}, fmt.Errorf("failed to parse addressed call from signed message: %w", err)
	}

	bitSetSig, ok := signedMsg.Signature.(*avalancheWarp.BitSetSignature)
	if !ok {
		return diffupdater.ICMMessage{}, fmt.Errorf("expected BitSetSignature, got %T", signedMsg.Signature)
	}

	sig, err := bls.SignatureFromBytes(bitSetSig.Signature[:])
	if err != nil {
		return diffupdater.ICMMessage{}, fmt.Errorf("failed to decompress BLS signature: %w", err)
	}
	uncompressedSig := sig.Serialize()

	attestation := make([]byte, 0, len(bitSetSig.Signers)+len(uncompressedSig))
	attestation = append(attestation, bitSetSig.Signers...)
	attestation = append(attestation, uncompressedSig...)

	return diffupdater.ICMMessage{
		RawMessage:         addressedCall.Payload,
		SourceNetworkID:    signedMsg.UnsignedMessage.NetworkID,
		SourceBlockchainID: signedMsg.UnsignedMessage.SourceChainID,
		Attestation:        attestation,
	}, nil
}

// unPadOnChainBlsPublicKey converts a padded on-chain BLS public key to the 96-byte
// uncompressed form used in warp messages. When padded is 128 bytes, the layout matches
// Solidity BLST.unPadUncompressedBlsPublicKey (see BLST.sol assembly). Otherwise the
// first 96 bytes are copied as a best-effort fallback for unexpected formats.
func unPadOnChainBlsPublicKey(padded []byte) [96]byte {
	var pk [96]byte
	if len(padded) != 128 {
		copy(pk[:], padded)
		return pk
	}
	// X: mstore(res+0x20, mload(pk+0x30)); mstore(res+0x30, mload(pk+0x40))
	copy(pk[0:32], padded[16:48])
	copy(pk[16:48], padded[32:64])
	// Y: mstore(res+0x50, mload(pk+0x70)); mstore(res+0x60, mload(pk+0x80))
	copy(pk[48:80], padded[80:112])
	copy(pk[64:96], padded[96:128])
	return pk
}

// onChainValidatorsToMessage converts on-chain Validator structs (with padded
// 128-byte BLS keys) to message.Validator structs (with uncompressed 96-byte keys).
func onChainValidatorsToMessage(validators []diffupdater.Validator) []*message.Validator {
	result := make([]*message.Validator, len(validators))
	for i, v := range validators {
		result[i] = &message.Validator{
			UncompressedPublicKeyBytes: unPadOnChainBlsPublicKey(v.BlsPublicKey),
			Weight:                     v.Weight,
		}
	}
	return result
}

// ShardValidatorsAsDiff creates ValidatorSetDiff (type ID 5) shards suitable
// for bootstrapping a DiffUpdater contract from an empty validator set.
// All validators are treated as additions; each shard carries a subset of changes.
func ShardValidatorsAsDiff(
	validators []*message.Validator,
	shardSize uint32,
	blockchainID ids.ID,
	prevHeight, prevTimestamp uint64,
	currHeight, currTimestamp uint64,
) ([][]byte, []ids.ID, error) {
	ss := int(shardSize)
	numValidators := len(validators)
	numShards := (numValidators + ss - 1) / ss
	if numShards == 0 {
		numShards = 1
	}

	shardBytesList := make([][]byte, numShards)
	shardHashes := make([]ids.ID, numShards)

	for i := 0; i < numShards; i++ {
		start := i * ss
		end := start + ss
		if end > numValidators {
			end = numValidators
		}
		shardValidators := validators[start:end]

		changes := make([]message.ValidatorChange, len(shardValidators))
		for j, v := range shardValidators {
			changes[j] = message.ValidatorChange{
				UncompressedPublicKeyBytes: v.UncompressedPublicKeyBytes,
				Weight:                     v.Weight,
			}
		}
		numAdded := uint32(len(changes))

		diff, err := message.NewValidatorSetDiff(
			blockchainID,
			prevHeight,
			prevTimestamp,
			currHeight,
			currTimestamp,
			changes,
			numAdded,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create ValidatorSetDiff shard %d: %w", i, err)
		}

		shardBytes := diff.Bytes()
		hash := sha256.Sum256(shardBytes)
		shardBytesList[i] = shardBytes
		shardHashes[i] = ids.ID(hash)
	}

	return shardBytesList, shardHashes, nil
}

// ---------------------------------------------------------------------------
// Test / convenience helpers
// ---------------------------------------------------------------------------

// PerformSingleUpdate is a convenience method for tests: builds, signs, and sends
// a single diff update at the given P-chain height.
func (d *DiffSetUpdater) PerformSingleUpdate(ctx context.Context, pChainHeight uint64) error {
	onChainVS, err := d.contract.GetValidatorSet(&bind.CallOpts{Context: ctx}, d.blockchainID)
	if err != nil {
		return fmt.Errorf("failed to get on-chain validator set: %w", err)
	}

	if onChainVS.TotalWeight == 0 {
		return d.performFullSetUpdate(ctx, pChainHeight, onChainVS)
	}
	return d.performDiffUpdate(ctx, pChainHeight, onChainVS)
}

// GetOnChainValidatorSet returns the validator set stored on-chain for the updater's blockchain ID.
func (d *DiffSetUpdater) GetOnChainValidatorSet(
	ctx context.Context,
) (diffupdater.ValidatorSet, error) {
	return d.contract.GetValidatorSet(
		&bind.CallOpts{Context: ctx},
		d.blockchainID,
	)
}

// SetTxValue sets the value on the tx opts (useful for initial deployment gas).
func (d *DiffSetUpdater) SetTxValue(val *big.Int) {
	d.txOpts.Value = val
}
