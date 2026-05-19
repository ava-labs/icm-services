// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validatorupdater

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"sort"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/logging"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	warppayload "github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	diffupdater "github.com/ava-labs/icm-services/abi-bindings/go/DiffUpdater"
	"github.com/ava-labs/icm-services/peers/clients"
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
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

	weightChangeThresholdPct float64
	maxUpdateInterval        time.Duration

	localValidatorSet []*Validator
	localPChainHeight uint64
	localTotalWeight  uint64
	lastUpdateTime    time.Time
	initialized       bool
}

func NewDiffSetUpdater(
	logger logging.Logger,
	pChainClient clients.CanonicalValidatorState,
	signatureAggregator *aggregator.SignatureAggregator,
	ethClient *ethclient.Client,
	contract *diffupdater.DiffUpdater,
	contractAddress common.Address,
	txOpts *bind.TransactOpts,
	networkID uint32,
	blockchainID ids.ID,
	subnetID ids.ID,
	shardSize uint32,
	pollInterval time.Duration,
	weightChangeThresholdPct float64,
	maxUpdateInterval time.Duration,
) *DiffSetUpdater {
	if shardSize == 0 {
		shardSize = defaultShardSize
	}
	if pollInterval == 0 {
		pollInterval = defaultPollInterval
	}
	return &DiffSetUpdater{
		logger:                   logger,
		pChainClient:             pChainClient,
		signatureAggregator:      signatureAggregator,
		ethClient:                ethClient,
		contract:                 contract,
		contractAddress:          contractAddress,
		txOpts:                   txOpts,
		networkID:                networkID,
		blockchainID:             blockchainID,
		subnetID:                 subnetID,
		shardSize:                shardSize,
		pollInterval:             pollInterval,
		weightChangeThresholdPct: weightChangeThresholdPct,
		maxUpdateInterval:        maxUpdateInterval,
	}
}

// Start runs the polling loop that detects validator set changes and updates the contract.
func (d *DiffSetUpdater) Start(ctx context.Context) error {
	d.logger.Info("Starting DiffSetUpdater",
		zap.Stringer("blockchainID", d.blockchainID),
		zap.Uint32("shardSize", d.shardSize),
		zap.Float64("weightChangeThresholdPct", d.weightChangeThresholdPct),
		zap.Duration("maxUpdateInterval", d.maxUpdateInterval),
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
	if !d.initialized {
		return d.initializeLocalState(ctx)
	}

	pChainHeight, err := d.pChainClient.GetLatestHeight(ctx)
	if err != nil {
		return fmt.Errorf("failed to get P-chain height: %w", err)
	}

	if pChainHeight <= d.localPChainHeight {
		d.logger.Debug("No new P-chain blocks",
			zap.Uint64("pChainHeight", pChainHeight),
			zap.Uint64("localPChainHeight", d.localPChainHeight),
		)
		return nil
	}

	newValidators, err := d.fetchSortedValidators(ctx, pChainHeight)
	if err != nil {
		return fmt.Errorf("failed to fetch validators at height %d: %w", pChainHeight, err)
	}

	changes, _ := computeValidatorDiff(d.localValidatorSet, newValidators)

	stale := d.isStale()
	if len(changes) == 0 && !stale {
		d.logger.Debug("No validator changes and not stale, skipping")
		return nil
	}

	weightDelta := computeWeightDelta(d.localValidatorSet, changes)
	thresholdCrossed := d.localTotalWeight > 0 &&
		float64(weightDelta)/float64(d.localTotalWeight) >= d.weightChangeThresholdPct

	if !thresholdCrossed && !stale {
		d.logger.Debug("Below threshold, skipping update",
			zap.Uint64("weightDelta", weightDelta),
			zap.Uint64("localTotalWeight", d.localTotalWeight),
			zap.Float64("thresholdPct", d.weightChangeThresholdPct),
		)
		return nil
	}

	d.logger.Info("Validator set update triggered",
		zap.Uint64("localPChainHeight", d.localPChainHeight),
		zap.Uint64("pChainHeight", pChainHeight),
		zap.Int("numChanges", len(changes)),
		zap.Uint64("weightDelta", weightDelta),
		zap.Bool("thresholdCrossed", thresholdCrossed),
		zap.Bool("stale", stale),
	)

	if err := d.performDiffUpdate(ctx, pChainHeight, newValidators, changes); err != nil {
		d.logger.Warn("Diff update failed, re-syncing from contract on next tick", zap.Error(err))
		d.initialized = false
		return err
	}

	d.localValidatorSet = newValidators
	d.localPChainHeight = pChainHeight
	d.localTotalWeight = sumWeights(newValidators)
	d.lastUpdateTime = time.Now()

	return nil
}

func (d *DiffSetUpdater) initializeLocalState(ctx context.Context) error {
	onChainVS, err := d.contract.GetValidatorSet(&bind.CallOpts{Context: ctx}, d.blockchainID)
	if err != nil {
		return fmt.Errorf("failed to get on-chain validator set: %w", err)
	}

	isFirstRegistration := onChainVS.TotalWeight == 0
	if isFirstRegistration {
		pChainHeight, err := d.pChainClient.GetLatestHeight(ctx)
		if err != nil {
			return fmt.Errorf("failed to get P-chain height: %w", err)
		}
		d.logger.Info("First registration detected, performing full set update",
			zap.Uint64("pChainHeight", pChainHeight),
		)
		if err := d.performFullSetUpdate(ctx, pChainHeight, onChainVS); err != nil {
			return err
		}
		newValidators, err := d.fetchSortedValidators(ctx, pChainHeight)
		if err != nil {
			return fmt.Errorf("failed to fetch validators after first registration: %w", err)
		}
		d.localValidatorSet = newValidators
		d.localPChainHeight = pChainHeight
		d.localTotalWeight = sumWeights(newValidators)
		d.lastUpdateTime = time.Now()
		d.initialized = true
		return nil
	}

	validators, err := d.fetchSortedValidators(ctx, onChainVS.PChainHeight)
	if err != nil {
		return fmt.Errorf("failed to fetch P-chain validators at on-chain height %d: %w",
			onChainVS.PChainHeight, err)
	}

	d.localValidatorSet = validators
	d.localPChainHeight = onChainVS.PChainHeight
	d.localTotalWeight = sumWeights(validators)
	d.lastUpdateTime = time.Now()
	d.initialized = true

	d.logger.Info("Initialized local validator set from contract",
		zap.Uint64("pChainHeight", onChainVS.PChainHeight),
		zap.Int("numValidators", len(validators)),
		zap.Uint64("totalWeight", d.localTotalWeight),
	)

	return nil
}

func (d *DiffSetUpdater) isStale() bool {
	if d.maxUpdateInterval == 0 {
		return false
	}
	return time.Since(d.lastUpdateTime) >= d.maxUpdateInterval
}

func sumWeights(validators []*Validator) uint64 {
	var total uint64
	for _, v := range validators {
		total += v.Weight
	}
	return total
}

// computeWeightDelta computes the sum of absolute weight deltas across all
// changes. For additions the delta is the new weight; for removals it is the
// old weight; for modifications it is |new - old|.
func computeWeightDelta(oldSet []*Validator, changes []ValidatorChange) uint64 {
	oldWeights := make(map[[96]byte]uint64, len(oldSet))
	for _, v := range oldSet {
		oldWeights[v.UncompressedPublicKeyBytes] = v.Weight
	}

	var delta uint64
	for _, c := range changes {
		oldW := oldWeights[c.UncompressedPublicKeyBytes]
		newW := c.Weight
		if newW >= oldW {
			delta += newW - oldW
		} else {
			delta += oldW - newW
		}
	}
	return delta
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

	changes := make([]ValidatorChange, len(newValidators))
	for i, v := range newValidators {
		changes[i] = ValidatorChange{
			UncompressedPublicKeyBytes: v.UncompressedPublicKeyBytes,
			Weight:                     v.Weight,
		}
	}

	pChainTimestamp, err := d.pChainClient.GetBlockTimestampAtHeight(ctx, pChainHeight)
	if err != nil {
		return fmt.Errorf("failed to get P-chain block timestamp at height %d: %w", pChainHeight, err)
	}

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

	metadataMsg, err := NewValidatorSetMetadata(
		d.blockchainID,
		pChainHeight,
		pChainTimestamp,
		shardHashes,
	)
	if err != nil {
		return fmt.Errorf("failed to create ValidatorSetMetadata: %w", err)
	}

	addressedCall, err := warppayload.NewAddressedCall(nil, metadataMsg.Bytes())
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

	// 24-byte diff justification: shard_size(8) + prev_height(8) + prev_timestamp(8)
	var justification [24]byte
	binary.BigEndian.PutUint64(justification[:8], uint64(d.shardSize))
	binary.BigEndian.PutUint64(justification[8:16], onChainVS.PChainHeight)
	binary.BigEndian.PutUint64(justification[16:24], onChainVS.PChainTimestamp)

	signedMsg, err := d.signatureAggregator.CreateSignedMessage(
		ctx,
		d.logger,
		unsignedMsg,
		justification[:],
		signingSubnet,
		defaultQuorumPercentage,
		defaultQuorumPercentageBuf,
		pChainHeight,
	)
	if err != nil {
		return fmt.Errorf("failed to sign message: %w", err)
	}

	return d.sendDiffUpdate(ctx, signedMsg, shardBytesList)
}

// ---------------------------------------------------------------------------
// Diff update using local validator set tracking
// ---------------------------------------------------------------------------

func (d *DiffSetUpdater) performDiffUpdate(
	ctx context.Context,
	pChainHeight uint64,
	newValidators []*Validator,
	changes []ValidatorChange,
) error {
	localPChainTimestamp, err := d.pChainClient.GetBlockTimestampAtHeight(ctx, d.localPChainHeight)
	if err != nil {
		return fmt.Errorf("failed to get P-chain block timestamp at height %d: %w", d.localPChainHeight, err)
	}

	pChainTimestamp, err := d.pChainClient.GetBlockTimestampAtHeight(ctx, pChainHeight)
	if err != nil {
		return fmt.Errorf("failed to get P-chain block timestamp at height %d: %w", pChainHeight, err)
	}

	// If stale but no actual changes, send the diff with the full current set
	// as "no-op" changes to advance the on-chain height.
	if len(changes) == 0 {
		changes = make([]ValidatorChange, len(newValidators))
		for i, v := range newValidators {
			changes[i] = ValidatorChange{
				UncompressedPublicKeyBytes: v.UncompressedPublicKeyBytes,
				Weight:                     v.Weight,
			}
		}
	}

	shardBytesList, shardHashes, err := d.shardDiff(
		d.blockchainID,
		d.localPChainHeight,
		localPChainTimestamp,
		pChainHeight,
		pChainTimestamp,
		d.localValidatorSet,
		changes,
	)
	if err != nil {
		return fmt.Errorf("failed to shard diff: %w", err)
	}

	metadataMsg, err := NewValidatorSetMetadata(
		d.blockchainID,
		pChainHeight,
		pChainTimestamp,
		shardHashes,
	)
	if err != nil {
		return fmt.Errorf("failed to create ValidatorSetMetadata: %w", err)
	}

	addressedCall, err := warppayload.NewAddressedCall(nil, metadataMsg.Bytes())
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

	signingSubnet := d.subnetID
	d.logger.Info("Signing diff update", zap.Stringer("signingSubnet", signingSubnet))

	// 24-byte diff justification: shard_size(8) + prev_height(8) + prev_timestamp(8)
	var justification [24]byte
	binary.BigEndian.PutUint64(justification[:8], uint64(d.shardSize))
	binary.BigEndian.PutUint64(justification[8:16], d.localPChainHeight)
	binary.BigEndian.PutUint64(justification[16:24], localPChainTimestamp)

	signedMsg, err := d.signatureAggregator.CreateSignedMessage(
		ctx,
		d.logger,
		unsignedMsg,
		justification[:],
		signingSubnet,
		defaultQuorumPercentage,
		defaultQuorumPercentageBuf,
		d.localPChainHeight,
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
	old, new_ []*Validator,
) ([]ValidatorChange, uint32) {
	var changes []ValidatorChange
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
			changes = append(changes, ValidatorChange{
				UncompressedPublicKeyBytes: old[oi].UncompressedPublicKeyBytes,
				Weight:                     0,
			})
			oi++
		case cmp > 0:
			// Addition: exists in new but not in old
			changes = append(changes, ValidatorChange{
				UncompressedPublicKeyBytes: new_[ni].UncompressedPublicKeyBytes,
				Weight:                     new_[ni].Weight,
			})
			numAdded++
			ni++
		default:
			// Same key: check if weight changed
			if old[oi].Weight != new_[ni].Weight {
				changes = append(changes, ValidatorChange{
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
	currentSet []*Validator,
	changes []ValidatorChange,
) []*Validator {
	var result []*Validator
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
			result = append(result, &Validator{
				UncompressedPublicKeyBytes: changes[ci].UncompressedPublicKeyBytes,
				Weight:                     changes[ci].Weight,
			})
			ci++
		default:
			if changes[ci].Weight != 0 {
				result = append(result, &Validator{
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
	oldValidators []*Validator,
	changes []ValidatorChange,
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

		diff, err := NewValidatorSetDiff(
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
	if receipt.Status == types.ReceiptStatusFailed {
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
		if receipt.Status == types.ReceiptStatusFailed {
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
) ([]*Validator, error) {
	allValidatorSets, err := d.pChainClient.GetAllValidatorSets(ctx, pChainHeight)
	if err != nil {
		return nil, fmt.Errorf("failed to get validator sets: %w", err)
	}

	vdrSet, ok := allValidatorSets[d.subnetID]
	if !ok {
		return nil, fmt.Errorf("subnet %s not found in validator sets at height %d", d.subnetID, pChainHeight)
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

func buildDiffICMMessage(signedMsg *avalancheWarp.Message) (diffupdater.ICMMessage, error) {
	addressedCall, err := warppayload.ParseAddressedCall(signedMsg.UnsignedMessage.Payload)
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

// ShardValidatorsAsDiff creates ValidatorSetDiff (type ID 5) shards suitable
// for bootstrapping a DiffUpdater contract from an empty validator set.
// All validators are treated as additions; each shard carries a subset of changes.
func ShardValidatorsAsDiff(
	validators []*Validator,
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

		changes := make([]ValidatorChange, len(shardValidators))
		for j, v := range shardValidators {
			changes[j] = ValidatorChange{
				UncompressedPublicKeyBytes: v.UncompressedPublicKeyBytes,
				Weight:                     v.Weight,
			}
		}
		numAdded := uint32(len(changes))

		diff, err := NewValidatorSetDiff(
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
