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
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/units"
	merkleregistry "github.com/ava-labs/icm-services/abi-bindings/go/MerkleValidatorSetRegistry"
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
	"github.com/ava-labs/libevm/crypto"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

const merkleMaxUpdateIntervalSeconds uint64 = 60

// MerkleUpdater tests the relayer's MerkleSetUpdater end-to-end:
//
//  1. Deploys a MerkleValidatorSetRegistry with the current P-chain genesis Merkle root.
//  2. Starts the relayer and waits for the initial L1 validator set registration.
//  3. Verifies the on-chain commitment matches the expected Merkle root and total weight.
//  4. Adds a validator to the L1 — verifies the relayer immediately updates the commitment
//     (the Merkle updater fires on any root change, no threshold delay).
//  5. Waits for a staleness-forced update after MaxUpdateIntervalSeconds with no validator
//     changes, verifying the commitment is refreshed to a newer P-chain height.
func MerkleUpdater(
	ctx context.Context,
	log logging.Logger,
	avalancheNetwork *network.LocalAvalancheNetwork,
	ethereumNetwork *network.LocalEthereumNetwork,
	teleporter utils.TeleporterTestInfo,
) {
	log.Info("Starting MerkleUpdater e2e test")

	l1Info := avalancheNetwork.GetL1Infos()[0]
	l1BlockchainID := l1Info.BlockchainID
	networkID := avalancheNetwork.GetNetworkID()

	log.Info("Test configuration",
		zap.Stringer("l1BlockchainID", l1BlockchainID),
		zap.Stringer("subnetID", l1Info.SubnetID),
		zap.Uint32("networkID", networkID),
		zap.Uint64("maxUpdateIntervalSeconds", merkleMaxUpdateIntervalSeconds),
	)

	ethClient := ethereumNetwork.EthClient
	_, ethFundedKey := ethereumNetwork.GetFundedAccountInfo()
	chainID := ethereumNetwork.ChainID
	fundedAddress, fundedKey := avalancheNetwork.GetFundedAccountInfo()
	_ = chainID

	// =========================================================================
	// Setup: Fetch P-chain validators and compute the genesis Merkle root
	// =========================================================================
	primaryNetworkInfo := avalancheNetwork.GetPrimaryNetworkInfo()
	pChainClient := clients.NewCanonicalValidatorClient(&config.APIConfig{
		BaseURL: primaryNetworkInfo.NodeURIs[0],
	})
	pChainHeight, err := pChainClient.GetLatestHeight(ctx)
	Expect(err).Should(BeNil())
	pChainTimestamp, err := pChainClient.GetBlockTimestampAtHeight(ctx, pChainHeight)
	Expect(err).Should(BeNil())

	pChainWarpSet, err := pChainClient.GetProposedValidators(ctx, ids.Empty)
	Expect(err).Should(BeNil())

	pChainValidators := make([]*validatorupdater.Validator, len(pChainWarpSet.Validators))
	for i, val := range pChainWarpSet.Validators {
		pChainValidators[i] = &validatorupdater.Validator{
			UncompressedPublicKeyBytes: [96]byte(val.PublicKey.Serialize()),
			Weight:                     val.Weight,
		}
	}
	sort.Slice(pChainValidators, func(i, j int) bool {
		return string(pChainValidators[i].UncompressedPublicKeyBytes[:]) <
			string(pChainValidators[j].UncompressedPublicKeyBytes[:])
	})

	pChainGenesisRoot := validatorupdater.BuildMerkleRoot(pChainValidators)
	var pChainTotalWeight uint64
	for _, v := range pChainValidators {
		pChainTotalWeight += v.Weight
	}

	log.Info("Fetched primary network validators",
		zap.Int("numValidators", len(pChainValidators)),
		zap.Uint64("pChainHeight", pChainHeight),
		zap.Uint64("pChainTimestamp", pChainTimestamp),
		zap.String("pChainGenesisRoot", hex.EncodeToString(pChainGenesisRoot[:])),
		zap.Uint64("pChainTotalWeight", pChainTotalWeight),
	)

	// =========================================================================
	// Setup: Deploy MerkleValidatorSetRegistry with the P-chain genesis root
	// =========================================================================
	contractAddr := utils.DeployMerkleValidatorSetRegistry(
		ctx,
		ethereumNetwork.EthereumTestInfo(),
		ethFundedKey,
		networkID,
		constants.PlatformChainID,
		pChainGenesisRoot,
		pChainTotalWeight,
		pChainHeight,
		pChainTimestamp,
	)
	contract, err := merkleregistry.NewMerkleValidatorSetRegistry(contractAddr, ethClient)
	Expect(err).Should(BeNil())

	log.Info("Deployed MerkleValidatorSetRegistry contract",
		zap.String("address", contractAddr.Hex()),
	)

	// =========================================================================
	// Verify P-chain is initialized (done in constructor) and L1 is not yet registered
	// =========================================================================
	callOpts := &bind.CallOpts{Context: ctx}

	pChainInitialized, err := contract.PChainInitialized(callOpts)
	Expect(err).Should(BeNil())
	Expect(pChainInitialized).Should(BeTrue(), "P-chain should be initialized by the constructor")

	isRegistered, err := contract.IsRegistered(callOpts, l1BlockchainID)
	Expect(err).Should(BeNil())
	Expect(isRegistered).Should(BeFalse(), "L1 validator set should not be registered yet")

	initialCommitment, err := contract.GetValidatorSetCommitment(callOpts, l1BlockchainID)
	Expect(err).Should(BeNil())
	Expect(initialCommitment.TotalWeight).Should(Equal(uint64(0)),
		"L1 commitment should start empty")

	log.Info("Contract state verified: P-chain initialized, L1 not yet registered")

	// =========================================================================
	// Setup: Configure and start the relayer
	// =========================================================================
	log.Info("Configuring relayer with ExternalEVMDestination (merkle)")

	err = utils.ClearRelayerStorage()
	Expect(err).Should(BeNil())

	relayerKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	utils.FundRelayers(ctx, []testinfo.L1TestInfo{l1Info}, fundedKey, relayerKey)

	relayerConfig := createMerkleUpdaterRelayerConfig(
		log,
		teleporter,
		l1Info,
		fundedAddress,
		relayerKey,
		ethereumNetwork,
		contractAddr.Hex(),
		l1BlockchainID.String(),
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

	log.Info("Relayer started, waiting for initial L1 validator set registration...")

	// =========================================================================
	// Wait for the relayer to register the initial L1 validator set.
	// totalWeight == 0 always triggers an immediate registration.
	// =========================================================================
	pollCtx, pollCancel := context.WithTimeout(ctx, 120*time.Second)
	defer pollCancel()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	var firstTotalWeight uint64
	var firstPChainHeight uint64

	for done := false; !done; {
		select {
		case <-pollCtx.Done():
			Expect(pollCtx.Err()).ShouldNot(HaveOccurred(),
				"Timed out waiting for relayer to register L1 validator set")
		case <-ticker.C:
			cmt, err := contract.GetValidatorSetCommitment(callOpts, l1BlockchainID)
			if err != nil {
				log.Warn("Failed to query on-chain commitment", zap.Error(err))
				continue
			}
			if cmt.TotalWeight == 0 {
				log.Debug("Validator set not yet registered, waiting...")
				continue
			}

			log.Info("Initial validator set registered",
				zap.Uint64("totalWeight", cmt.TotalWeight),
				zap.Uint64("pChainHeight", cmt.PChainHeight),
				zap.String("root", hex.EncodeToString(cmt.Root[:])),
			)

			Expect(cmt.PChainHeight).Should(BeNumerically(">", 0),
				"P-chain height should be positive")
			Expect(cmt.PChainTimestamp).Should(BeNumerically(">", 0),
				"P-chain timestamp should be positive")
			Expect(cmt.TotalWeight).Should(BeNumerically(">", 0),
				"Total weight should be positive")
			Expect(cmt.Root).ShouldNot(Equal([32]byte{}),
				"Merkle root should not be empty")

			// Verify the on-chain root matches the expected L1 validator set at the recorded height.
			expectedValidators := fetchSortedL1ValidatorsAtHeight(
				ctx, pChainClient, l1Info.SubnetID, cmt.PChainHeight,
			)
			expectedRoot := validatorupdater.BuildMerkleRoot(expectedValidators)
			var expectedWeight uint64
			for _, v := range expectedValidators {
				expectedWeight += v.Weight
			}
			Expect(cmt.Root).Should(Equal(expectedRoot),
				"Merkle root should match the P-chain validator set at the recorded height")
			Expect(cmt.TotalWeight).Should(Equal(expectedWeight),
				"Total weight should match the sum of validator weights")

			firstTotalWeight = cmt.TotalWeight
			firstPChainHeight = cmt.PChainHeight
			done = true
		}
	}

	firstRegistrationTime := time.Now()
	log.Info("Initial registration complete",
		zap.Uint64("firstTotalWeight", firstTotalWeight),
		zap.Uint64("firstPChainHeight", firstPChainHeight),
	)

	// =========================================================================
	// Phase 1: Add a validator — verify the relayer immediately updates the root.
	// The Merkle updater fires on any root change with no threshold delay.
	// =========================================================================
	log.Info("Phase 1: Adding validator, expecting immediate root update...")

	validatorManagerProxy, poaManagerProxy := avalancheNetwork.GetValidatorManager(l1Info.SubnetID)
	poaManager, err := poamanager.NewPoAManager(poaManagerProxy.Address, l1Info.EthClient)
	Expect(err).Should(BeNil())

	pChainInfo := utils.GetPChainInfo(avalancheNetwork.GetPrimaryNetworkInfo())
	aggregator := avalancheNetwork.GetSignatureAggregator()
	defer aggregator.Shutdown()

	newNodes := avalancheNetwork.GetExtraNodes(1)
	Expect(len(newNodes)).Should(Equal(1))

	newValidatorWeight := units.Schmeckle / 3
	l1Info = avalancheNetwork.AddSubnetValidators(newNodes[:1], l1Info, true)

	addCtx, addCancel := context.WithTimeout(ctx, 150*time.Second)
	defer addCancel()

	expiry := uint64(time.Now().Add(24 * time.Hour).Unix())
	pop, err := newNodes[0].GetProofOfPossession()
	Expect(err).Should(BeNil())

	newNode := utils.Node{
		NodeID:  newNodes[0].NodeID,
		NodePoP: pop,
		Weight:  newValidatorWeight,
	}

	utils.InitiateAndCompletePoAValidatorRegistration(
		addCtx,
		aggregator,
		fundedKey,
		l1Info,
		pChainInfo,
		poaManager,
		poaManagerProxy.Address,
		validatorManagerProxy.Address,
		expiry,
		newNode,
		avalancheNetwork.GetPChainWallet(),
		avalancheNetwork.GetNetworkID(),
	)

	err = utils.IssueTxsToAdvanceChain(ctx, l1Info.EVMChainID, fundedKey, l1Info.EthClient, 5)
	Expect(err).Should(BeNil())

	log.Info("Phase 1: Validator added, waiting for root update (no threshold delay)...")

	// 90s timeout — well under the staleness cap. If the update arrives in this window
	// it was triggered by the root change, not by staleness.
	rootChangeCtx, rootChangeCancel := context.WithTimeout(ctx, 90*time.Second)
	defer rootChangeCancel()

	rootChangeTicker := time.NewTicker(2 * time.Second)
	defer rootChangeTicker.Stop()

	var secondPChainHeight uint64
	var secondTotalWeight uint64
	var secondRoot [32]byte

	for {
		select {
		case <-rootChangeCtx.Done():
			Expect(false).Should(BeTrue(),
				"Phase 1: Timed out waiting for root-change-triggered update")
		case <-rootChangeTicker.C:
			cmt, err := contract.GetValidatorSetCommitment(callOpts, l1BlockchainID)
			if err != nil {
				log.Warn("Failed to query on-chain commitment", zap.Error(err))
				continue
			}

			if cmt.PChainHeight <= firstPChainHeight {
				log.Debug("Phase 1: Waiting for root-change update...",
					zap.Uint64("currentPChainHeight", cmt.PChainHeight),
				)
				continue
			}

			log.Info("Phase 1: Root-change update detected!",
				zap.Uint64("totalWeight", cmt.TotalWeight),
				zap.Uint64("pChainHeight", cmt.PChainHeight),
				zap.String("root", hex.EncodeToString(cmt.Root[:])),
			)

			Expect(cmt.PChainHeight).Should(BeNumerically(">", firstPChainHeight))
			Expect(cmt.TotalWeight).Should(BeNumerically(">", firstTotalWeight),
				"Phase 1: New commitment should include the added validator's weight")
			Expect(cmt.Root).ShouldNot(Equal([32]byte{}))

			// Verify the new root matches the updated validator set at the recorded height.
			expectedValidators := fetchSortedL1ValidatorsAtHeight(
				ctx, pChainClient, l1Info.SubnetID, cmt.PChainHeight,
			)
			expectedRoot := validatorupdater.BuildMerkleRoot(expectedValidators)
			Expect(cmt.Root).Should(Equal(expectedRoot),
				"Phase 1: Merkle root should match the updated P-chain validator set")

			secondPChainHeight = cmt.PChainHeight
			secondTotalWeight = cmt.TotalWeight
			secondRoot = cmt.Root
		}
		break
	}

	secondUpdateTime := time.Now()
	log.Info("Phase 1 PASSED: Root-change-triggered update confirmed",
		zap.Uint64("secondTotalWeight", secondTotalWeight),
		zap.Uint64("secondPChainHeight", secondPChainHeight),
	)

	// =========================================================================
	// Phase 2: Wait for staleness-forced update.
	// No new validators are added; after MaxUpdateIntervalSeconds the relayer
	// refreshes the commitment to a newer P-chain height.
	// =========================================================================
	log.Info("Phase 2: Waiting for staleness-forced update (no validator changes)...")

	// Issue a transaction toAdvance the P-chain by one block so the relayer has
	// a new height to commit to when its staleness timer fires.
	_, err = avalancheNetwork.GetPChainWallet().IssueBaseTx(nil)
	Expect(err).Should(BeNil())

	elapsed := time.Since(firstRegistrationTime)
	stalenessTimeout := time.Duration(merkleMaxUpdateIntervalSeconds)*time.Second + 90*time.Second
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

	for {
		select {
		case <-stalenessCtx.Done():
			Expect(false).Should(BeTrue(),
				"Phase 2: Timed out waiting for staleness-forced update")
		case <-stalenessTicker.C:
			cmt, err := contract.GetValidatorSetCommitment(callOpts, l1BlockchainID)
			if err != nil {
				log.Warn("Failed to query on-chain commitment", zap.Error(err))
				continue
			}

			if cmt.PChainHeight <= secondPChainHeight {
				log.Debug("Phase 2: Still waiting for staleness update...",
					zap.Duration("elapsed", time.Since(secondUpdateTime)),
				)
				continue
			}

			log.Info("Phase 2: Staleness-forced update detected!",
				zap.Uint64("totalWeight", cmt.TotalWeight),
				zap.Uint64("pChainHeight", cmt.PChainHeight),
				zap.String("root", hex.EncodeToString(cmt.Root[:])),
			)

			Expect(cmt.PChainHeight).Should(BeNumerically(">", secondPChainHeight))
			Expect(cmt.TotalWeight).Should(Equal(secondTotalWeight),
				"Phase 2: Total weight should be unchanged (no new validators)")
			Expect(cmt.Root).Should(Equal(secondRoot),
				"Phase 2: Merkle root should be unchanged (no new validators)")

			log.Info("Phase 2 PASSED: Staleness-forced update confirmed")
		}
		break
	}

	log.Info("MerkleUpdater e2e test PASSED")
}

func createMerkleUpdaterRelayerConfig(
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

	baseConfig.APIPort = 8084

	baseConfig.ExternalEVMDestinations = []*relayercfg.ExternalEVMDestination{
		{
			RPCEndpoint:              ethereumNetwork.BaseURL,
			PrivateKey:               hex.EncodeToString(crypto.FromECDSA(ethFundedKey)),
			ContractAddress:          contractAddress,
			BlockchainID:             blockchainID,
			SubnetID:                 subnetID,
			PollIntervalSeconds:      testPollIntervalSeconds,
			ContractType:             "merkle",
			MaxUpdateIntervalSeconds: merkleMaxUpdateIntervalSeconds,
		},
	}

	return baseConfig
}
