// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"sort"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/units"
	diffupdater "github.com/ava-labs/icm-services/abi-bindings/go/DiffUpdater"
	poamanager "github.com/ava-labs/icm-services/abi-bindings/go/validator-manager/PoAManager"
	"github.com/ava-labs/icm-services/config"
	"github.com/ava-labs/icm-services/icm-contracts/tests/network"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/icm-services/peers/clients"
	relayercfg "github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/icm-services/relayer/validatorupdater"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/crypto"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

// Primary-network bootstrap shards for DiffUpdater can exceed the default
// transactor gas cap (~10M); observed estimate ~11.5M on local Ethereum.
const diffUpdaterBootstrapGasLimit uint64 = 16_000_000

const (
	// Must be below the ValidatorManager's MaximumChurnPercentage (20%) so
	// that a weight large enough to trigger the threshold doesn't violate
	// the churn limit.
	thresholdWeightChangeThresholdPct float64 = 0.05
	// Must be larger than Phase 1 wall-clock time (registration ~90s on CI +
	// 20s observation). 120s gives headroom so staleness never fires during
	// the below-threshold observation window.
	thresholdMaxUpdateIntervalSeconds uint64 = 120
)

// DiffUpdater tests the relayer's DiffSetUpdater end-to-end with threshold-based
// gas optimization enabled:
//
//  1. Deploys a DiffUpdater contract on the external Ethereum network.
//  2. Starts the relayer with threshold config, waits for initial validator set
//     registration (totalWeight=0 always triggers an immediate update).
//  3. Adds a small-weight validator below the threshold — verifies NO update.
//  4. Waits for the staleness interval to force the update.
//  5. Adds a large-weight validator above the threshold — verifies fast update.
func DiffUpdater(
	ctx context.Context,
	log logging.Logger,
	avalancheNetwork *network.LocalAvalancheNetwork,
	ethereumNetwork *network.LocalEthereumNetwork,
	teleporter utils.TeleporterTestInfo,
) {
	log.Info("Starting DiffUpdater e2e test")

	l1Info := avalancheNetwork.GetL1Infos()[0]
	blockchainID := l1Info.BlockchainID
	networkID := avalancheNetwork.GetNetworkID()

	log.Info("Test configuration",
		zap.Stringer("blockchainID", blockchainID),
		zap.Stringer("subnetID", l1Info.SubnetID),
		zap.Uint32("networkID", networkID),
		zap.Float64("weightChangeThresholdPct", thresholdWeightChangeThresholdPct),
		zap.Uint64("maxUpdateIntervalSeconds", thresholdMaxUpdateIntervalSeconds),
	)

	ethClient := ethereumNetwork.EthClient
	_, ethFundedKey := ethereumNetwork.GetFundedAccountInfo()
	chainID := ethereumNetwork.ChainID
	fundedAddress, fundedKey := avalancheNetwork.GetFundedAccountInfo()

	// =========================================================================
	// Setup: Fetch primary network validators for P-chain bootstrap
	// =========================================================================
	primaryNetworkInfo := avalancheNetwork.GetPrimaryNetworkInfo()
	pChainClient := clients.NewCanonicalValidatorClient(&config.APIConfig{
		BaseURL: primaryNetworkInfo.NodeURIs[0],
	})
	pChainHeight, err := pChainClient.GetLatestHeight(ctx)
	Expect(err).Should(BeNil())

	pChainWarpSet, err := pChainClient.GetProposedValidators(ctx, ids.Empty)
	Expect(err).Should(BeNil())

	pChainValidators := make([]*validatorupdater.Validator, len(pChainWarpSet.Validators))
	for i, vdr := range pChainWarpSet.Validators {
		pChainValidators[i] = &validatorupdater.Validator{
			UncompressedPublicKeyBytes: [96]byte(vdr.PublicKey.Serialize()),
			Weight:                     vdr.Weight,
		}
	}
	sort.Slice(pChainValidators, func(i, j int) bool {
		return string(pChainValidators[i].UncompressedPublicKeyBytes[:]) <
			string(pChainValidators[j].UncompressedPublicKeyBytes[:])
	})

	var pChainID [32]byte // all zeros = PlatformChainID
	pChainTimestamp, err := pChainClient.GetBlockTimestampAtHeight(ctx, pChainHeight)
	Expect(err).Should(BeNil())

	bootstrapHeight := pChainHeight + 1
	bootstrapTimestamp := pChainTimestamp + 1

	pChainShardBytesList, pChainShardHashes, err := validatorupdater.ShardValidatorsAsDiff(
		pChainValidators,
		testShardSize,
		ids.ID(pChainID),
		pChainHeight,
		pChainTimestamp,
		bootstrapHeight,
		bootstrapTimestamp,
	)
	Expect(err).Should(BeNil())

	pChainShardHashesBytes := make([][32]byte, len(pChainShardHashes))
	for i, h := range pChainShardHashes {
		pChainShardHashesBytes[i] = h
	}

	log.Info("Fetched primary network validators",
		zap.Int("numValidators", len(pChainValidators)),
		zap.Int("numShards", len(pChainShardBytesList)),
		zap.Uint64("pChainHeight", pChainHeight),
	)

	// =========================================================================
	// Setup: Deploy DiffUpdater contract
	// =========================================================================
	txOpts, err := bind.NewKeyedTransactorWithChainID(ethFundedKey, chainID)
	Expect(err).Should(BeNil())
	txOpts.GasLimit = diffUpdaterBootstrapGasLimit

	initialMetadata := diffupdater.ValidatorSetMetadata{
		AvalancheBlockchainID: pChainID,
		PChainHeight:          pChainHeight,
		PChainTimestamp:       pChainTimestamp,
		ShardHashes:           pChainShardHashesBytes,
	}
	contractAddr := utils.DeployDiffUpdaterWithMetadata(
		ctx,
		ethereumNetwork.EthereumTestInfo(),
		ethFundedKey,
		networkID,
		initialMetadata,
	)
	contract, err := diffupdater.NewDiffUpdater(contractAddr, ethClient)
	Expect(err).Should(BeNil())

	log.Info("Deployed DiffUpdater contract",
		zap.String("address", contractAddr.Hex()),
	)

	// =========================================================================
	// Setup: Bootstrap P-chain validators via updateValidatorSet
	// =========================================================================
	for i, shardBytes := range pChainShardBytesList {
		shard := diffupdater.ValidatorSetShard{
			ShardNumber:           uint64(i + 1),
			AvalancheBlockchainID: pChainID,
		}
		tx, err := contract.UpdateValidatorSet(txOpts, shard, shardBytes)
		Expect(err).Should(BeNil())
		receipt, err := bind.WaitMined(ctx, ethClient, tx)
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful),
			"updateValidatorSet shard %d failed", i+1)
		log.Info("Bootstrapped P-chain shard",
			zap.Int("shardNumber", i+1),
			zap.Int("totalShards", len(pChainShardBytesList)),
		)
	}

	callOpts := &bind.CallOpts{Context: ctx}
	pChainInitialized, err := contract.PChainInitialized(callOpts)
	Expect(err).Should(BeNil())
	Expect(pChainInitialized).Should(BeTrue(), "P-chain validator set should be initialized")

	log.Info("P-chain validators bootstrapped successfully")

	onChainVS, err := contract.GetValidatorSet(callOpts, blockchainID)
	Expect(err).Should(BeNil())
	Expect(onChainVS.TotalWeight).Should(Equal(uint64(0)),
		"L1 validator set should start empty")

	// =========================================================================
	// Setup: Configure and start the relayer with threshold config
	// =========================================================================
	log.Info("Configuring relayer with ExternalEVMDestination (diff + threshold)")

	err = utils.ClearRelayerStorage()
	Expect(err).Should(BeNil())

	relayerKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	utils.FundRelayers(ctx, []testinfo.L1TestInfo{l1Info}, fundedKey, relayerKey)

	relayerConfig := createDiffUpdaterRelayerConfig(
		log,
		teleporter,
		l1Info,
		fundedAddress,
		relayerKey,
		ethereumNetwork,
		contractAddr.Hex(),
		blockchainID.String(),
		l1Info.SubnetID.String(),
	)
	relayerConfigPath := utils.WriteRelayerConfig(log, relayerConfig, utils.DefaultRelayerCfgFname)

	log.Info("Starting relayer")
	relayerCleanup, readyChan := utils.RunRelayerExecutable(
		ctx,
		log,
		relayerConfigPath,
		relayerConfig,
	)
	defer relayerCleanup()

	startupCtx, startupCancel := context.WithTimeout(ctx, 60*time.Second)
	defer startupCancel()
	utils.WaitForChannelClose(startupCtx, readyChan)

	log.Info("Relayer started, waiting for validator set registration...")

	// =========================================================================
	// Wait for the relayer to register the initial validator set.
	// Even with threshold config, totalWeight=0 triggers an immediate update.
	// =========================================================================
	pollCtx, pollCancel := context.WithTimeout(ctx, 120*time.Second)
	defer pollCancel()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	var firstTotalWeight uint64

	for done := false; !done; {
		select {
		case <-pollCtx.Done():
			Expect(pollCtx.Err()).ShouldNot(HaveOccurred(),
				"Timed out waiting for relayer to register validator set")
		case <-ticker.C:
			vs, err := contract.GetValidatorSet(callOpts, blockchainID)
			if err != nil {
				log.Warn("Failed to query on-chain validator set", zap.Error(err))
				continue
			}
			if vs.TotalWeight == 0 {
				log.Debug("Validator set not yet registered, waiting...")
				continue
			}

			log.Info("Initial validator set registered",
				zap.Int("validatorCount", len(vs.Validators)),
				zap.Uint64("totalWeight", vs.TotalWeight),
				zap.Uint64("pChainHeight", vs.PChainHeight),
			)

			Expect(vs.PChainHeight).Should(BeNumerically(">", 0),
				"P-Chain height should be positive")
			Expect(vs.PChainTimestamp).Should(BeNumerically(">", 0),
				"P-Chain timestamp should be positive")
			Expect(len(vs.Validators)).Should(BeNumerically(">", 0),
				"Should have at least one validator")
			Expect(vs.TotalWeight).Should(BeNumerically(">", 0),
				"Total weight should be positive")

			var calculatedWeight uint64
			for i, v := range vs.Validators {
				Expect(len(v.BlsPublicKey)).Should(Equal(128),
					"BLS public key should be 128 bytes (uncompressed G1)")
				Expect(v.Weight).Should(BeNumerically(">", 0),
					"Validator weight should be positive")
				calculatedWeight += v.Weight
				log.Debug("Validator details",
					zap.Int("index", i),
					zap.Uint64("weight", v.Weight),
				)
			}
			Expect(vs.TotalWeight).Should(Equal(calculatedWeight),
				"Total weight should match sum of individual validator weights")

			firstTotalWeight = vs.TotalWeight
			done = true
		}
	}

	firstRegistrationTime := time.Now()
	log.Info("Initial registration complete",
		zap.Uint64("firstTotalWeight", firstTotalWeight),
	)

	// =========================================================================
	// Phase 1: Add a small validator (below threshold) — verify NO update
	// =========================================================================
	log.Info("Phase 1: Adding small validator (below threshold)...")

	validatorManagerProxy, poaManagerProxy := avalancheNetwork.GetValidatorManager(l1Info.SubnetID)
	poaManager, err := poamanager.NewPoAManager(poaManagerProxy.Address, l1Info.EthClient)
	Expect(err).Should(BeNil())

	pChainInfo := utils.GetPChainInfo(avalancheNetwork.GetPrimaryNetworkInfo())
	aggregator := avalancheNetwork.GetSignatureAggregator()
	defer aggregator.Shutdown()

	newNodes := avalancheNetwork.GetExtraNodes(2)
	Expect(len(newNodes)).Should(Equal(2))

	smallWeight := units.Schmeckle / 10
	expectedDeltaPct := float64(smallWeight) / float64(firstTotalWeight)
	log.Info("Phase 1: Small validator weight details",
		zap.Uint64("smallWeight", smallWeight),
		zap.Uint64("firstTotalWeight", firstTotalWeight),
		zap.Float64("expectedDeltaPct", expectedDeltaPct),
		zap.Float64("threshold", thresholdWeightChangeThresholdPct),
	)
	Expect(expectedDeltaPct).Should(BeNumerically("<", thresholdWeightChangeThresholdPct),
		"Small validator weight should be below the threshold")

	l1Info = avalancheNetwork.AddSubnetValidators(newNodes[:1], l1Info, true)

	addSmallCtx, addSmallCancel := context.WithTimeout(ctx, 150*time.Second)
	defer addSmallCancel()

	expiry := uint64(time.Now().Add(24 * time.Hour).Unix())
	pop, err := newNodes[0].GetProofOfPossession()
	Expect(err).Should(BeNil())

	smallNode := utils.Node{
		NodeID:  newNodes[0].NodeID,
		NodePoP: pop,
		Weight:  smallWeight,
	}

	utils.InitiateAndCompletePoAValidatorRegistration(
		addSmallCtx,
		aggregator,
		fundedKey,
		l1Info,
		pChainInfo,
		poaManager,
		poaManagerProxy.Address,
		validatorManagerProxy.Address,
		expiry,
		smallNode,
		avalancheNetwork.GetPChainWallet(),
		avalancheNetwork.GetNetworkID(),
	)

	err = utils.IssueTxsToAdvanceChain(ctx, l1Info.EVMChainID, fundedKey, l1Info.EthClient, 5)
	Expect(err).Should(BeNil())

	// Snapshot on-chain state after the P-chain registration completes.
	baselineVS, err := contract.GetValidatorSet(callOpts, blockchainID)
	Expect(err).Should(BeNil())
	baselineValidatorCount := len(baselineVS.Validators)
	baselinePChainHeight := baselineVS.PChainHeight

	log.Info("Phase 1: Baseline snapshot after registration",
		zap.Int("baselineValidatorCount", baselineValidatorCount),
		zap.Uint64("baselinePChainHeight", baselinePChainHeight),
		zap.Uint64("baselineTotalWeight", baselineVS.TotalWeight),
	)

	log.Info("Phase 1: Verifying NO on-chain update for 20s...")

	noUpdateCtx, noUpdateCancel := context.WithTimeout(ctx, 20*time.Second)
	defer noUpdateCancel()

	noUpdateTicker := time.NewTicker(2 * time.Second)
	defer noUpdateTicker.Stop()

	for {
		select {
		case <-noUpdateCtx.Done():
			log.Info("Phase 1 PASSED: No update within 20s (below threshold)")
		case <-noUpdateTicker.C:
			vs, err := contract.GetValidatorSet(callOpts, blockchainID)
			if err != nil {
				log.Warn("Failed to query on-chain validator set", zap.Error(err))
				continue
			}
			Expect(len(vs.Validators)).Should(Equal(baselineValidatorCount),
				"Phase 1: Validator count should NOT change while below threshold")
			Expect(vs.PChainHeight).Should(Equal(baselinePChainHeight),
				"Phase 1: P-chain height should NOT change while below threshold")
			continue
		}
		break
	}

	// =========================================================================
	// Phase 2: Wait for staleness to force the update
	// =========================================================================
	log.Info("Phase 2: Waiting for staleness-forced update...")

	elapsed := time.Since(firstRegistrationTime)
	stalenessTimeout := time.Duration(thresholdMaxUpdateIntervalSeconds)*time.Second + 90*time.Second
	remainingWait := stalenessTimeout - elapsed
	if remainingWait < 30*time.Second {
		remainingWait = 30 * time.Second
	}

	log.Info("Phase 2: Timing",
		zap.Duration("elapsed", elapsed),
		zap.Duration("stalenessTimeout", stalenessTimeout),
		zap.Duration("remainingWait", remainingWait),
	)

	stalenessCtx, stalenessCancel := context.WithTimeout(ctx, remainingWait)
	defer stalenessCancel()

	stalenessTicker := time.NewTicker(2 * time.Second)
	defer stalenessTicker.Stop()

	var secondPChainHeight uint64
	var secondValidatorCount int
	var secondTotalWeight uint64

	for {
		select {
		case <-stalenessCtx.Done():
			Expect(false).Should(BeTrue(),
				"Phase 2: Timed out waiting for staleness-forced update")
		case <-stalenessTicker.C:
			vs, err := contract.GetValidatorSet(callOpts, blockchainID)
			if err != nil {
				log.Warn("Failed to query on-chain validator set", zap.Error(err))
				continue
			}

			if vs.PChainHeight <= baselinePChainHeight {
				log.Debug("Phase 2: Still waiting for staleness update...",
					zap.Duration("elapsed", time.Since(firstRegistrationTime)),
				)
				continue
			}

			log.Info("Phase 2: Staleness-forced update detected!",
				zap.Int("validatorCount", len(vs.Validators)),
				zap.Uint64("totalWeight", vs.TotalWeight),
				zap.Uint64("pChainHeight", vs.PChainHeight),
			)

			Expect(len(vs.Validators)).Should(BeNumerically(">", baselineValidatorCount),
				"Phase 2: Staleness update should include the small validator")
			Expect(vs.PChainHeight).Should(BeNumerically(">", baselinePChainHeight))

			secondPChainHeight = vs.PChainHeight
			secondValidatorCount = len(vs.Validators)
			secondTotalWeight = vs.TotalWeight
		}
		break
	}

	stalenessUpdateTime := time.Now()
	log.Info("Phase 2 PASSED: Staleness-forced update confirmed",
		zap.Int("secondValidatorCount", secondValidatorCount),
		zap.Uint64("secondTotalWeight", secondTotalWeight),
		zap.Uint64("secondPChainHeight", secondPChainHeight),
	)

	// =========================================================================
	// Phase 3: Add a large validator (above threshold) — verify fast update
	// =========================================================================
	log.Info("Phase 3: Adding large validator (above threshold)...")

	largeWeight := units.Schmeckle / 3
	expectedLargeDeltaPct := float64(largeWeight) / float64(secondTotalWeight)
	log.Info("Phase 3: Large validator weight details",
		zap.Uint64("largeWeight", largeWeight),
		zap.Uint64("secondTotalWeight", secondTotalWeight),
		zap.Float64("expectedDeltaPct", expectedLargeDeltaPct),
		zap.Float64("threshold", thresholdWeightChangeThresholdPct),
	)
	Expect(expectedLargeDeltaPct).Should(BeNumerically(">=", thresholdWeightChangeThresholdPct),
		"Large validator weight should exceed the threshold")

	l1Info = avalancheNetwork.AddSubnetValidators(newNodes[1:2], l1Info, true)

	addLargeCtx, addLargeCancel := context.WithTimeout(ctx, 150*time.Second)
	defer addLargeCancel()

	expiry = uint64(time.Now().Add(24 * time.Hour).Unix())
	pop, err = newNodes[1].GetProofOfPossession()
	Expect(err).Should(BeNil())

	largeNode := utils.Node{
		NodeID:  newNodes[1].NodeID,
		NodePoP: pop,
		Weight:  largeWeight,
	}

	utils.InitiateAndCompletePoAValidatorRegistration(
		addLargeCtx,
		aggregator,
		fundedKey,
		l1Info,
		pChainInfo,
		poaManager,
		poaManagerProxy.Address,
		validatorManagerProxy.Address,
		expiry,
		largeNode,
		avalancheNetwork.GetPChainWallet(),
		avalancheNetwork.GetNetworkID(),
	)

	err = utils.IssueTxsToAdvanceChain(ctx, l1Info.EVMChainID, fundedKey, l1Info.EthClient, 5)
	Expect(err).Should(BeNil())

	log.Info("Phase 3: Large validator added, waiting for threshold-triggered update...")

	// 90s timeout — under the 120s staleness cap. If the update arrives
	// in this window it was threshold-triggered, not staleness-triggered.
	thresholdCtx, thresholdCancel := context.WithTimeout(ctx, 90*time.Second)
	defer thresholdCancel()

	thresholdTicker := time.NewTicker(2 * time.Second)
	defer thresholdTicker.Stop()

	for {
		select {
		case <-thresholdCtx.Done():
			Expect(false).Should(BeTrue(),
				"Phase 3: Timed out waiting for threshold-triggered update")
		case <-thresholdTicker.C:
			vs, err := contract.GetValidatorSet(callOpts, blockchainID)
			if err != nil {
				log.Warn("Failed to query on-chain validator set", zap.Error(err))
				continue
			}

			if vs.PChainHeight <= secondPChainHeight {
				log.Debug("Phase 3: Waiting for threshold-triggered update...",
					zap.Int("currentCount", len(vs.Validators)),
					zap.Uint64("currentPChainHeight", vs.PChainHeight),
				)
				continue
			}

			updateLatency := time.Since(stalenessUpdateTime)
			log.Info("Phase 3: Threshold-triggered update detected!",
				zap.Int("validatorCount", len(vs.Validators)),
				zap.Uint64("totalWeight", vs.TotalWeight),
				zap.Uint64("pChainHeight", vs.PChainHeight),
				zap.Duration("updateLatency", updateLatency),
			)

			Expect(len(vs.Validators)).Should(BeNumerically(">", secondValidatorCount))
			Expect(vs.PChainHeight).Should(BeNumerically(">", secondPChainHeight))
			Expect(vs.TotalWeight).Should(BeNumerically(">", secondTotalWeight))

			Expect(updateLatency).Should(BeNumerically("<",
				time.Duration(thresholdMaxUpdateIntervalSeconds)*time.Second),
				"Phase 3: Update should arrive before staleness interval")

			log.Info("Phase 3 PASSED: Threshold-triggered update confirmed",
				zap.Duration("latency", updateLatency),
			)
		}
		break
	}

	log.Info("DiffUpdater e2e test PASSED")
}

func createDiffUpdaterRelayerConfig(
	log logging.Logger,
	teleporter utils.TeleporterTestInfo,
	l1Info testinfo.L1TestInfo,
	fundedAddress common.Address,
	relayerKey *ecdsa.PrivateKey,
	ethereumNetwork *network.LocalEthereumNetwork,
	contractAddress string,
	blockchainID string,
	subnetID string,
) relayercfg.Config {
	baseConfig := utils.CreateDefaultRelayerConfig(
		log,
		teleporter,
		[]testinfo.L1TestInfo{l1Info},
		[]testinfo.L1TestInfo{l1Info},
		fundedAddress,
		relayerKey,
	)

	_, ethFundedKey := ethereumNetwork.GetFundedAccountInfo()

	baseConfig.APIPort = 8083

	baseConfig.ExternalEVMDestinations = []*relayercfg.ExternalEVMDestination{
		{
			RPCEndpoint:              ethereumNetwork.BaseURL,
			PrivateKey:               hex.EncodeToString(crypto.FromECDSA(ethFundedKey)),
			ContractAddress:          contractAddress,
			BlockchainID:             blockchainID,
			SubnetID:                 subnetID,
			ShardSize:                testShardSize,
			PollIntervalSeconds:      testPollIntervalSeconds,
			ContractType:             "diff",
			WeightChangeThresholdPct: thresholdWeightChangeThresholdPct,
			MaxUpdateIntervalSeconds: thresholdMaxUpdateIntervalSeconds,
		},
	}

	return baseConfig
}
