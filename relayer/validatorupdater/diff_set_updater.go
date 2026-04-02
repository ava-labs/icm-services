// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validatorupdater

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

	// In-memory cache of the last successfully submitted validator state.
	// Populated once on startup from the contract, then kept in sync after
	// each successful update — eliminating per-poll EVM reads.
	cachedValidators         []*Validator
	cachedPChainHeight       uint64
	cachedPChainTimestamp    uint64
	cacheInitialized         bool

	// weightChangeThresholdBps is the minimum total-weight change (in basis
	// points, 10_000 = 100 %) required to trigger a contract update.
	// 0 means any validator-set difference triggers an update.
	weightChangeThresholdBps uint64
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
	weightChangeThresholdBps uint64,
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
		weightChangeThresholdBps: weightChangeThresholdBps,
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

// initializeCache fetches the current on-chain validator state once so that
// all subsequent polls can compare against the in-memory cache rather than
// issuing a per-poll EVM read.
func (d *DiffSetUpdater) initializeCache(ctx context.Context) error {
	onChainVS, err := d.contract.GetValidatorSet(&bind.CallOpts{Context: ctx}, d.blockchainID)
	if err != nil {
		return fmt.Errorf("failed to initialize validator set cache: %w", err)
	}
	d.cacheInitialized = true
	if onChainVS.TotalWeight > 0 {
		d.cachedValidators = onChainValidatorsToMessage(onChainVS.Validators)
		sort.Slice(d.cachedValidators, func(i, j int) bool {
			return bytes.Compare(
				d.cachedValidators[i].UncompressedPublicKeyBytes[:],
				d.cachedValidators[j].UncompressedPublicKeyBytes[:],
			) < 0
		})
		d.cachedPChainHeight = onChainVS.PChainHeight
		d.cachedPChainTimestamp = onChainVS.PChainTimestamp
	}
	d.logger.Info("Validator set cache initialized from contract",
		zap.Uint64("cachedPChainHeight", d.cachedPChainHeight),
		zap.Int("cachedValidatorCount", len(d.cachedValidators)),
	)
	return nil
}

func (d *DiffSetUpdater) checkAndUpdate(ctx context.Context) error {
	pChainHeight, err := d.pChainClient.GetLatestHeight(ctx)
	if err != nil {
		return fmt.Errorf("failed to get P-chain height: %w", err)
	}

	// Populate the in-memory cache from the contract on first call only.
	if !d.cacheInitialized {
		if err := d.initializeCache(ctx); err != nil {
			return err
		}
	}

	if d.cachedPChainHeight >= pChainHeight {
		d.logger.Debug("P-chain height has not advanced, skipping",
			zap.Uint64("pChainHeight", pChainHeight),
			zap.Uint64("cachedHeight", d.cachedPChainHeight),
		)
		return nil
	}

	isFirstRegistration := d.cachedValidators == nil

	newValidators, err := d.fetchSortedValidators(ctx, pChainHeight)
	if err != nil {
		return fmt.Errorf("failed to fetch validators: %w", err)
	}

	if !isFirstRegistration {
		if !weightDiffExceedsThreshold(d.cachedValidators, newValidators, d.weightChangeThresholdBps) {
			d.logger.Info("Validator set change below threshold, skipping update",
				zap.Uint64("cachedHeight", d.cachedPChainHeight),
				zap.Uint64("pChainHeight", pChainHeight),
				zap.Uint64("weightChangeThresholdBps", d.weightChangeThresholdBps),
			)
			return nil
		}
	}

	d.logger.Info("Validator set update needed",
		zap.Uint64("cachedHeight", d.cachedPChainHeight),
		zap.Uint64("pChainHeight", pChainHeight),
		zap.Bool("isFirstRegistration", isFirstRegistration),
		zap.Int("cachedValidatorCount", len(d.cachedValidators)),
	)

	if isFirstRegistration {
		return d.performFullSetUpdate(ctx, pChainHeight, newValidators)
	}
	return d.performDiffUpdate(ctx, pChainHeight, newValidators)
}

// ---------------------------------------------------------------------------
// First registration: treat all validators as additions in a diff
// ---------------------------------------------------------------------------

// performFullSetUpdate treats all newValidators as additions (first registration).
// Uses d.cachedPChainHeight / d.cachedPChainTimestamp as the "previous" heights
// (both zero when the contract has never been written to).
func (d *DiffSetUpdater) performFullSetUpdate(
	ctx context.Context,
	pChainHeight uint64,
	newValidators []*Validator,
) error {
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
		d.cachedPChainHeight,    // prevHeight  (0 for first registration)
		d.cachedPChainTimestamp, // prevTimestamp (0 for first registration)
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

	signedMsg, err := d.signatureAggregator.CreateSignedMessage(
		ctx,
		d.logger,
		unsignedMsg,
		nil,
		signingSubnet,
		defaultQuorumPercentage,
		defaultQuorumPercentageBuf,
		d.cachedPChainHeight,
	)
	if err != nil {
		return fmt.Errorf("failed to sign message: %w", err)
	}

	if err := d.sendDiffUpdate(ctx, signedMsg, shardBytesList); err != nil {
		return err
	}

	// Update cache after successful submission.
	d.cachedValidators = newValidators
	d.cachedPChainHeight = pChainHeight
	d.cachedPChainTimestamp = pChainTimestamp
	return nil
}

// ---------------------------------------------------------------------------
// Diff update: compute changes between on-chain and P-chain validator sets
// ---------------------------------------------------------------------------

// performDiffUpdate computes and submits the diff between d.cachedValidators
// and newValidators. Uses d.cachedPChainHeight / d.cachedPChainTimestamp as the
// "previous" heights for the diff message.
func (d *DiffSetUpdater) performDiffUpdate(
	ctx context.Context,
	pChainHeight uint64,
	newValidators []*Validator,
) error {
	changes, _ := computeValidatorDiff(d.cachedValidators, newValidators)

	if len(changes) == 0 {
		d.logger.Info("No validator changes detected, skipping update")
		return nil
	}

	pChainTimestamp, err := d.pChainClient.GetBlockTimestampAtHeight(ctx, pChainHeight)
	if err != nil {
		return fmt.Errorf("failed to get P-chain block timestamp at height %d: %w", pChainHeight, err)
	}

	shardBytesList, shardHashes, err := d.shardDiff(
		d.blockchainID,
		d.cachedPChainHeight,    // prevHeight
		d.cachedPChainTimestamp, // prevTimestamp
		pChainHeight,
		pChainTimestamp,
		d.cachedValidators,
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
		d.cachedPChainHeight,
	)
	if err != nil {
		return fmt.Errorf("failed to sign message: %w", err)
	}

	d.logger.Info("Sending diff update",
		zap.Int("numChanges", len(changes)),
		zap.Int("numShards", len(shardBytesList)),
	)

	if err := d.sendDiffUpdate(ctx, signedMsg, shardBytesList); err != nil {
		return err
	}

	// Update cache after successful submission.
	d.cachedValidators = newValidators
	d.cachedPChainHeight = pChainHeight
	d.cachedPChainTimestamp = pChainTimestamp
	return nil
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

// unPadOnChainBlsPublicKey converts a padded on-chain BLS public key to the 96-byte
// uncompressed form used in warp messages. When padded is 128 bytes, the layout matches
// BLST.unPadUncompressedBlsPublicKey: X occupies padded[16:64], Y padded[80:128] (see
// BLST.padUncompressedBLSPublicKey). Otherwise the first 96 bytes are copied as a
// best-effort fallback for unexpected formats.
func unPadOnChainBlsPublicKey(padded []byte) [96]byte {
	var pk [96]byte
	if len(padded) != 128 {
		copy(pk[:], padded)
		return pk
	}
	copy(pk[0:48], padded[16:64])
	copy(pk[48:96], padded[80:128])
	return pk
}

// onChainValidatorsToMessage converts on-chain Validator structs (with padded
// 128-byte BLS keys) to Validator structs (with uncompressed 96-byte keys).
func onChainValidatorsToMessage(validators []diffupdater.Validator) []*Validator {
	result := make([]*Validator, len(validators))
	for i, v := range validators {
		result[i] = &Validator{
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

// ---------------------------------------------------------------------------
// Threshold helpers
// ---------------------------------------------------------------------------

// weightDiffExceedsThreshold reports whether the validator set should be
// updated given the change from old to new.
//
//   - thresholdBps == 0: any validator-set difference (membership or weight
//     change) triggers an update, equivalent to the pre-threshold behavior.
//   - thresholdBps > 0: only a change in TOTAL weight that exceeds
//     thresholdBps / 10_000 of the previous total triggers an update.
//     Pure validator swaps at identical total weight are suppressed.
func weightDiffExceedsThreshold(old, new []*Validator, thresholdBps uint64) bool {
	if thresholdBps == 0 {
		changes, _ := computeValidatorDiff(old, new)
		return len(changes) > 0
	}

	var oldTotal, newTotal uint64
	for _, v := range old {
		oldTotal += v.Weight
	}
	for _, v := range new {
		newTotal += v.Weight
	}
	if oldTotal == newTotal {
		return false
	}
	if oldTotal == 0 {
		return true
	}
	var diff uint64
	if newTotal > oldTotal {
		diff = newTotal - oldTotal
	} else {
		diff = oldTotal - newTotal
	}
	// diff / oldTotal > thresholdBps / 10_000
	// Rearranged to avoid floating-point: diff * 10_000 > oldTotal * thresholdBps
	return diff*10_000 > oldTotal*thresholdBps
}

// ---------------------------------------------------------------------------
// Test / convenience helpers
// ---------------------------------------------------------------------------

// PerformSingleUpdate is a convenience method for tests: builds, signs, and sends
// a single diff update at the given P-chain height, bypassing the weight threshold.
func (d *DiffSetUpdater) PerformSingleUpdate(ctx context.Context, pChainHeight uint64) error {
	if !d.cacheInitialized {
		if err := d.initializeCache(ctx); err != nil {
			return err
		}
	}

	newValidators, err := d.fetchSortedValidators(ctx, pChainHeight)
	if err != nil {
		return fmt.Errorf("failed to fetch validators: %w", err)
	}

	if d.cachedValidators == nil {
		return d.performFullSetUpdate(ctx, pChainHeight, newValidators)
	}
	return d.performDiffUpdate(ctx, pChainHeight, newValidators)
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
