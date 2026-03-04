// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"os"
	"strings"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	subsetupdater "github.com/ava-labs/icm-services/abi-bindings/go/SubsetUpdater"
	"github.com/ava-labs/icm-services/icm-contracts/tests/network"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
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
	// Step 1: Deploy ValidatorSets library, then SubsetUpdater contract
	// =========================================================================
	txOpts, err := bind.NewKeyedTransactorWithChainID(ethFundedKey, chainID)
	Expect(err).Should(BeNil())

	libAddr := deployValidatorSetsLibrary(ctx, log, txOpts, ethClient)

	// Link the library address into the SubsetUpdater bytecode
	const libPlaceholder = "__$aaf4ae346b84a712cc43f25bb66199d6fb$__"
	libAddrHex := strings.ToLower(libAddr.Hex()[2:])
	origBin := subsetupdater.SubsetUpdaterBin
	subsetupdater.SubsetUpdaterBin = strings.ReplaceAll(origBin, libPlaceholder, libAddrHex)
	defer func() { subsetupdater.SubsetUpdaterBin = origBin }()

	initialMetadata := subsetupdater.ValidatorSetMetadata{
		AvalancheBlockchainID: blockchainID,
		PChainHeight:          0,
		PChainTimestamp:       0,
		ShardHashes:           [][32]byte{},
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

	// Verify initial state: no validator set registered yet
	callOpts := &bind.CallOpts{Context: ctx}
	onChainVS, err := contract.GetValidatorSet(callOpts, blockchainID)
	Expect(err).Should(BeNil())
	Expect(onChainVS.TotalWeight).Should(Equal(uint64(0)),
		"Contract should start with empty validator set")

	// =========================================================================
	// Step 2: Configure and start the relayer
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
	// Step 3: Wait for the relayer to register the validator set
	// =========================================================================
	pollCtx, pollCancel := context.WithTimeout(ctx, 120*time.Second)
	defer pollCancel()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

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
			// Step 4: Verify the registered validator set
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

			log.Info("SubsetUpdater e2e test PASSED",
				zap.String("contractAddress", contractAddr.Hex()),
				zap.Int("validatorCount", len(vs.Validators)),
				zap.Uint64("totalWeight", vs.TotalWeight),
				zap.Uint64("pChainHeight", vs.PChainHeight),
			)
			return
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
