// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/message"
	subsetupdater "github.com/ava-labs/icm-services/abi-bindings/go/SubsetUpdater"
	poamanager "github.com/ava-labs/icm-services/abi-bindings/go/validator-manager/PoAManager"
	"github.com/ava-labs/icm-services/config"
	"github.com/ava-labs/icm-services/icm-contracts/tests/network"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/icm-services/peers/clients"
	"github.com/ava-labs/icm-services/relayer"
	relayercfg "github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/libevm/ethclient"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

const (
	testShardSize           uint32 = 2
	testPollIntervalSeconds uint64 = 5
)

// SubsetUpdater tests the relayer's SubsetSetUpdater end-to-end:
//  1. Deploys a SubsetUpdater contract on the external Ethereum network.
//  2. Configures the relayer with an ExternalEVMDestination pointing at the contract.
//  3. Starts the relayer and waits for it to automatically detect validators,
//     build a SubsetUpdate warp message, aggregate signatures, and submit the
//     registerValidatorSet / updateValidatorSet transactions.
//  4. Verifies the on-chain validator set matches expectations.
//  5. Adds a new validator to the L1 to trigger a re-registration (non-first-time).
//  6. Waits for the relayer to detect the validator set change and re-register.
//  7. Verifies the updated on-chain validator set includes the new validator.
func SubsetUpdater(
	ctx context.Context,
	log logging.Logger,
	avalancheNetwork *network.LocalAvalancheNetwork,
	ethereumNetwork *network.LocalEthereumNetwork,
	teleporter utils.TeleporterTestInfo,
) {
	log.Info("Starting SubsetUpdater e2e test")

	l1Info := avalancheNetwork.GetL1Infos()[0]
	blockchainID := l1Info.BlockchainID
	networkID := avalancheNetwork.GetNetworkID()

	log.Info("Test configuration",
		zap.Stringer("blockchainID", blockchainID),
		zap.Stringer("subnetID", l1Info.SubnetID),
		zap.Uint32("networkID", networkID),
	)

	ethClient := ethereumNetwork.EthClient()
	_, ethFundedKey := ethereumNetwork.GetFundedAccountInfo()
	chainID := ethereumNetwork.ChainID()
	fundedAddress, fundedKey := avalancheNetwork.GetFundedAccountInfo()

	// =========================================================================
	// Step 1: Fetch primary network validators for P-chain bootstrap
	// =========================================================================
	primaryNetworkInfo := avalancheNetwork.GetPrimaryNetworkInfo()
	pChainClient := clients.NewCanonicalValidatorClient(&config.APIConfig{
		BaseURL: primaryNetworkInfo.NodeURIs[0],
	})
	pChainHeight, err := pChainClient.GetLatestHeight(ctx)
	Expect(err).Should(BeNil())

	pChainWarpSet, err := pChainClient.GetProposedValidators(ctx, ids.Empty)
	Expect(err).Should(BeNil())

	pChainValidators := make([]*message.Validator, len(pChainWarpSet.Validators))
	for i, vdr := range pChainWarpSet.Validators {
		pChainValidators[i] = &message.Validator{
			UncompressedPublicKeyBytes: [96]byte(vdr.PublicKey.Serialize()),
			Weight:                     vdr.Weight,
		}
	}
	sort.Slice(pChainValidators, func(i, j int) bool {
		return string(pChainValidators[i].UncompressedPublicKeyBytes[:]) <
			string(pChainValidators[j].UncompressedPublicKeyBytes[:])
	})

	pChainShardBytesList, pChainShardHashes, err := relayer.ShardValidators(pChainValidators, testShardSize)
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
	// Step 2: Deploy ValidatorSets library, then SubsetUpdater contract
	// =========================================================================
	txOpts, err := bind.NewKeyedTransactorWithChainID(ethFundedKey, chainID)
	Expect(err).Should(BeNil())

	libAddr := deployValidatorSetsLibrary(ctx, log, txOpts, ethClient)

	const libPlaceholder = "__$aaf4ae346b84a712cc43f25bb66199d6fb$__"
	libAddrHex := strings.ToLower(libAddr.Hex()[2:])
	origBin := subsetupdater.SubsetUpdaterBin
	subsetupdater.SubsetUpdaterBin = strings.ReplaceAll(origBin, libPlaceholder, libAddrHex)
	defer func() { subsetupdater.SubsetUpdaterBin = origBin }()

	var pChainID [32]byte // all zeros = PlatformChainID
	initialMetadata := subsetupdater.ValidatorSetMetadata{
		AvalancheBlockchainID: pChainID,
		PChainHeight:          pChainHeight,
		PChainTimestamp:       uint64(time.Now().Unix()),
		ShardHashes:           pChainShardHashesBytes,
	}
	contractAddr, deployTx, contract, err := subsetupdater.DeploySubsetUpdater(
		txOpts, ethClient, networkID, initialMetadata,
	)
	Expect(err).Should(BeNil())

	deployReceipt, err := bind.WaitMined(ctx, ethClient, deployTx)
	Expect(err).Should(BeNil())
	Expect(deployReceipt.Status).Should(Equal(uint64(1)))

	log.Info("Deployed SubsetUpdater contract",
		zap.String("address", contractAddr.Hex()),
		zap.String("txHash", deployTx.Hash().Hex()),
	)

	// =========================================================================
	// Step 3: Bootstrap P-chain validators via updateValidatorSet
	// =========================================================================
	for i, shardBytes := range pChainShardBytesList {
		shard := subsetupdater.ValidatorSetShard{
			ShardNumber:           uint64(i + 1),
			AvalancheBlockchainID: pChainID,
		}
		tx, err := contract.UpdateValidatorSet(txOpts, shard, shardBytes)
		Expect(err).Should(BeNil())
		receipt, err := bind.WaitMined(ctx, ethClient, tx)
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(uint64(1)),
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

	// Verify L1 validator set not yet registered
	onChainVS, err := contract.GetValidatorSet(callOpts, blockchainID)
	Expect(err).Should(BeNil())
	Expect(onChainVS.TotalWeight).Should(Equal(uint64(0)),
		"L1 validator set should start empty")

	// =========================================================================
	// Step 4: Configure and start the relayer
	// =========================================================================
	log.Info("Configuring relayer with ExternalEVMDestination")

	err = utils.ClearRelayerStorage()
	Expect(err).Should(BeNil())

	relayerKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	utils.FundRelayers(ctx, []testinfo.L1TestInfo{l1Info}, fundedKey, relayerKey)

	relayerConfig := createSubsetUpdaterRelayerConfig(
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
	// Step 5: Wait for the relayer to register the validator set
	// =========================================================================
	pollCtx, pollCancel := context.WithTimeout(ctx, 120*time.Second)
	defer pollCancel()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	var firstPChainHeight uint64
	var firstValidatorCount int

registrationLoop:
	for {
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

			// =========================================================================
			// Step 6: Verify the registered validator set
			// =========================================================================
			log.Info("Validator set registered by relayer!",
				zap.Int("validatorCount", len(vs.Validators)),
				zap.Uint64("totalWeight", vs.TotalWeight),
				zap.Uint64("pChainHeight", vs.PChainHeight),
				zap.Uint64("pChainTimestamp", vs.PChainTimestamp),
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

			firstPChainHeight = vs.PChainHeight
			firstValidatorCount = len(vs.Validators)

			log.Info("First registration verified",
				zap.String("contractAddress", contractAddr.Hex()),
				zap.Int("validatorCount", firstValidatorCount),
				zap.Uint64("totalWeight", vs.TotalWeight),
				zap.Uint64("pChainHeight", firstPChainHeight),
			)
			break registrationLoop
		}
	}

	// =========================================================================
	// Step 7: Add a new validator to trigger re-registration
	// =========================================================================
	log.Info("Adding a new validator to trigger re-registration...")

	validatorManagerProxy, poaManagerProxy := avalancheNetwork.GetValidatorManager(l1Info.SubnetID)
	poaManager, err := poamanager.NewPoAManager(poaManagerProxy.Address, l1Info.EthClient)
	Expect(err).Should(BeNil())

	pChainInfo := utils.GetPChainInfo(avalancheNetwork.GetPrimaryNetworkInfo())
	aggregator := avalancheNetwork.GetSignatureAggregator()
	defer aggregator.Shutdown()

	newNodes := avalancheNetwork.GetExtraNodes(1)
	Expect(len(newNodes)).Should(Equal(1), "Should have at least 1 extra node available")

	log.Info("Adding extra node as subnet validator",
		zap.Stringer("nodeID", newNodes[0].NodeID),
	)
	l1Info = avalancheNetwork.AddSubnetValidators(newNodes, l1Info, true)

	addValidatorCtx, addValidatorCancel := context.WithTimeout(ctx, 120*time.Second)
	defer addValidatorCancel()

	expiry := uint64(time.Now().Add(24 * time.Hour).Unix())
	pop, err := newNodes[0].GetProofOfPossession()
	Expect(err).Should(BeNil())

	node := utils.Node{
		NodeID:  newNodes[0].NodeID,
		NodePoP: pop,
		Weight:  units.Schmeckle / 5,
	}

	log.Info("Initiating PoA validator registration",
		zap.Stringer("nodeID", node.NodeID),
		zap.Uint64("weight", node.Weight),
	)

	utils.InitiateAndCompletePoAValidatorRegistration(
		addValidatorCtx,
		aggregator,
		fundedKey,
		l1Info,
		pChainInfo,
		poaManager,
		poaManagerProxy.Address,
		validatorManagerProxy.Address,
		expiry,
		node,
		avalancheNetwork.GetPChainWallet(),
		avalancheNetwork.GetNetworkID(),
	)

	log.Info("New validator added, waiting for relayer to detect and re-register...")

	err = utils.IssueTxsToAdvanceChain(ctx, l1Info.EVMChainID, fundedKey, l1Info.EthClient, 5)
	Expect(err).Should(BeNil())

	// =========================================================================
	// Step 8: Wait for re-registration with updated validator set
	// =========================================================================
	updateCtx, updateCancel := context.WithTimeout(ctx, 120*time.Second)
	defer updateCancel()

	updateTicker := time.NewTicker(2 * time.Second)
	defer updateTicker.Stop()

	for {
		select {
		case <-updateCtx.Done():
			Expect(false).Should(BeTrue(),
				"Timed out waiting for relayer to re-register validator set after validator change")
		case <-updateTicker.C:
			vs, err := contract.GetValidatorSet(callOpts, blockchainID)
			if err != nil {
				log.Warn("Failed to query on-chain validator set", zap.Error(err))
				continue
			}

			if len(vs.Validators) > firstValidatorCount && vs.PChainHeight > firstPChainHeight {
				log.Info("Validator set re-registered by relayer!",
					zap.Int("validatorCount", len(vs.Validators)),
					zap.Uint64("totalWeight", vs.TotalWeight),
					zap.Uint64("pChainHeight", vs.PChainHeight),
					zap.Uint64("pChainTimestamp", vs.PChainTimestamp),
				)

				Expect(len(vs.Validators)).Should(BeNumerically(">", firstValidatorCount),
					"Updated validator set should have more validators after adding one")
				Expect(vs.PChainHeight).Should(BeNumerically(">", firstPChainHeight),
					"Updated validator set should have higher P-chain height")
				Expect(vs.PChainTimestamp).Should(BeNumerically(">", 0),
					"Updated validator set should have positive timestamp")

				log.Info("SubsetUpdater e2e test PASSED",
					zap.String("contractAddress", contractAddr.Hex()),
					zap.Int("firstValidatorCount", firstValidatorCount),
					zap.Int("updatedValidatorCount", len(vs.Validators)),
					zap.Uint64("firstPChainHeight", firstPChainHeight),
					zap.Uint64("updatedPChainHeight", vs.PChainHeight),
				)
				return
			}

			log.Debug("Waiting for validator set re-registration...",
				zap.Int("currentValidatorCount", len(vs.Validators)),
				zap.Int("expectedMinValidatorCount", firstValidatorCount+1),
				zap.Uint64("currentPChainHeight", vs.PChainHeight),
			)
		}
	}
}

func createSubsetUpdaterRelayerConfig(
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

	baseConfig.APIPort = 8082

	baseConfig.ExternalEVMDestinations = []*relayercfg.ExternalEVMDestination{
		{
			RPCEndpoint:         ethereumNetwork.RPCURL(),
			PrivateKey:          hex.EncodeToString(crypto.FromECDSA(ethFundedKey)),
			ContractAddress:     contractAddress,
			BlockchainID:        blockchainID,
			SubnetID:            subnetID,
			ShardSize:           testShardSize,
			PollIntervalSeconds: testPollIntervalSeconds,
		},
	}

	return baseConfig
}

// deployValidatorSetsLibrary deploys the ValidatorSets Solidity library
// from the forge build artifact and returns its on-chain address.
func deployValidatorSetsLibrary(
	ctx context.Context,
	log logging.Logger,
	txOpts *bind.TransactOpts,
	client *ethclient.Client,
) common.Address {
	artifactBytes, err := os.ReadFile("out/ValidatorSets.sol/ValidatorSets.json")
	Expect(err).Should(BeNil(), "forge artifact not found; run `FOUNDRY_PROFILE=ethereum forge build`")

	var artifact struct {
		Bytecode struct {
			Object string `json:"object"`
		} `json:"bytecode"`
	}
	Expect(json.Unmarshal(artifactBytes, &artifact)).Should(BeNil())

	bytecodeHex := strings.TrimPrefix(artifact.Bytecode.Object, "0x")
	libAddr, libTx, _, err := bind.DeployContract(
		txOpts, abi.ABI{}, common.FromHex(bytecodeHex), client,
	)
	Expect(err).Should(BeNil())

	receipt, err := bind.WaitMined(ctx, client, libTx)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(uint64(1)))

	log.Info("Deployed ValidatorSets library",
		zap.String("address", libAddr.Hex()),
	)
	return libAddr
}
