// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/units"
	avalanchevalidatorsetregistry "github.com/ava-labs/icm-services/abi-bindings/go/AvalancheValidatorSetRegistry"
	poamanager "github.com/ava-labs/icm-services/abi-bindings/go/validator-manager/PoAManager"
	"github.com/ava-labs/icm-services/config"
	"github.com/ava-labs/icm-services/icm-contracts/tests/interfaces"
	"github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	relayercfg "github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

// ValidatorSetUpdater tests the relayer's validator set updater functionality.
// This test:
// 1. Deploys an AvalancheValidatorSetRegistry contract on the Ethereum network
// 2. Configures the relayer with external EVM destinations (which automatically enables the validator set updater)
// 3. Waits for the relayer to automatically detect and post validator set updates
// 4. Verifies the registry contract was updated by the relayer
//
// Note: This test requires a local Ethereum network to be running.
func ValidatorSetUpdater(
	ctx context.Context,
	log logging.Logger,
	avalancheNetwork *network.LocalNetwork,
	ethereumNetwork *network.LocalEthereumNetwork,
	teleporter utils.TeleporterTestInfo,
) {
	log.Info("Starting ValidatorSetUpdater test")

	_, ethereumFundedKey := ethereumNetwork.GetFundedAccountInfo()
	fundedAddress, fundedKey := avalancheNetwork.GetFundedAccountInfo()

	// Get L1 info from Avalanche network
	l1Infos := avalancheNetwork.GetL1Infos()
	Expect(len(l1Infos)).Should(BeNumerically(">", 0), "No L1s found in Avalanche network")
	l1Info := l1Infos[0]

	// Create transaction options for Ethereum
	opts, err := bind.NewKeyedTransactorWithChainID(ethereumFundedKey, ethereumNetwork.ChainID)
	Expect(err).Should(BeNil())

	// Deploy AvalancheValidatorSetRegistry contract on the Ethereum network
	log.Info("Deploying AvalancheValidatorSetRegistry to Ethereum network",
		zap.Uint32("avalancheNetworkID", avalancheNetwork.GetNetworkID()),
		zap.Stringer("l1BlockchainID", l1Info.BlockchainID),
	)

	registryAddress, tx, registry, err := avalanchevalidatorsetregistry.DeployAvalancheValidatorSetRegistry(
		opts,
		ethereumNetwork.Client,
		avalancheNetwork.GetNetworkID(),
		l1Info.BlockchainID,
	)
	Expect(err).Should(BeNil())

	// Wait for deployment transaction
	receipt := utils.WaitForTransactionSuccessWithClient(ctx, ethereumNetwork.Client, tx.Hash())
	log.Info("Deployed AvalancheValidatorSetRegistry contract",
		zap.Stringer("address", registryAddress),
		zap.Stringer("txHash", receipt.TxHash),
	)

	// Verify initial contract state
	callOpts := &bind.CallOpts{Context: ctx}
	nextValidatorSetID, err := registry.NextValidatorSetID(callOpts)
	Expect(err).Should(BeNil())
	Expect(nextValidatorSetID).Should(Equal(uint32(0)), "Registry should start with no validator sets")

	log.Info("Registry deployed and verified",
		zap.Stringer("address", registryAddress),
		zap.Uint32("nextValidatorSetID", nextValidatorSetID),
	)

	// =========================================================================
	// Step 2: Configure and start the relayer with external EVM destinations
	// =========================================================================
	log.Info("Configuring relayer with external EVM destinations")

	// Clear any existing relayer storage
	err = utils.ClearRelayerStorage()
	Expect(err).Should(BeNil())

	// Generate a relayer key for signing transactions
	relayerKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())

	// Fund the relayer on Avalanche
	utils.FundRelayers(ctx, []interfaces.L1TestInfo{l1Info}, fundedKey, relayerKey)

	// Create the relayer config with external EVM destinations
	relayerConfig := createRelayerConfigWithValidatorSetUpdater(
		log,
		teleporter,
		l1Info,
		fundedAddress,
		relayerKey,
		ethereumNetwork,
		registryAddress,
	)

	relayerConfigPath := utils.WriteRelayerConfig(log, relayerConfig, utils.DefaultRelayerCfgFname)

	// Start the relayer
	log.Info("Starting relayer")
	relayerCleanup, readyChan := utils.RunRelayerExecutable(
		ctx,
		log,
		relayerConfigPath,
		relayerConfig,
	)
	defer relayerCleanup()

	// Wait for relayer to start up
	startupCtx, startupCancel := context.WithTimeout(ctx, 60*time.Second)
	defer startupCancel()
	utils.WaitForChannelClose(startupCtx, readyChan)

	log.Info("Relayer started, waiting for validator set update...")

	// =========================================================================
	// Step 3: Wait for the relayer to update the registry
	// =========================================================================
	// The ValidatorSetUpdater should automatically detect validators and post an update
	// The poll interval is 60 seconds, and the first attempt may fail while P2P connections
	// are being established, so we need to wait long enough for at least 2 poll cycles.
	// Wait up to 150 seconds for the update to be posted

	var validatorSetRegistered bool
	pollCtx, pollCancel := context.WithTimeout(ctx, 60*time.Second)
	defer pollCancel()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-pollCtx.Done():
			Expect(validatorSetRegistered).Should(BeTrue(), "Validator set should have been registered by relayer")
		case <-ticker.C:
			nextID, err := registry.NextValidatorSetID(callOpts)
			if err != nil {
				log.Warn("Failed to query registry", zap.Error(err))
				continue
			}

			if nextID > 0 {
				validatorSetRegistered = true
				log.Info("Validator set registered by relayer!",
					zap.Uint32("nextValidatorSetID", nextID),
				)

				// Verify the registered validator set
				validatorSet, err := registry.GetValidatorSet(callOpts, big.NewInt(0))
				Expect(err).Should(BeNil())

				log.Info("Verified registered validator set",
					zap.Int64("validatorSetID", 0),
					zap.Uint64("pChainHeight", validatorSet.PChainHeight),
					zap.Uint64("pChainTimestamp", validatorSet.PChainTimestamp),
					zap.Int("validatorCount", len(validatorSet.Validators)),
					zap.Uint64("totalWeight", validatorSet.TotalWeight),
				)

				// Assert on the validator set contents
				Expect(validatorSet.PChainHeight).Should(BeNumerically(">", 0), "P-Chain height should be positive")
				Expect(validatorSet.PChainTimestamp).Should(BeNumerically(">", 0), "P-Chain timestamp should be positive")
				Expect(len(validatorSet.Validators)).Should(BeNumerically(">", 0), "Should have at least one validator")
				Expect(validatorSet.TotalWeight).Should(BeNumerically(">", 0), "Total weight should be positive")

				// Verify each validator has valid data
				var calculatedTotalWeight uint64
				for i, validator := range validatorSet.Validators {
					// BLS public keys are stored as uncompressed G1 points (128 bytes)
					Expect(len(validator.BlsPublicKey)).Should(Equal(128), "BLS public key should be 128 bytes (uncompressed G1)")
					Expect(validator.Weight).Should(BeNumerically(">", 0), "Validator weight should be positive")
					calculatedTotalWeight += validator.Weight
					log.Debug("Validator details",
						zap.Int("index", i),
						zap.Uint64("weight", validator.Weight),
						zap.Int("blsKeyLen", len(validator.BlsPublicKey)),
					)
				}

				// Verify total weight matches sum of individual weights
				Expect(validatorSet.TotalWeight).Should(Equal(calculatedTotalWeight),
					"Total weight should match sum of validator weights")

				log.Info("Step 1 PASSED: registerValidatorSet verified successfully!",
					zap.Stringer("registryAddress", registryAddress),
					zap.Int("validatorCount", len(validatorSet.Validators)),
					zap.Uint64("totalWeight", validatorSet.TotalWeight),
					zap.Uint64("pChainHeight", validatorSet.PChainHeight),
				)

				// Store the first validator set info for comparison
				firstPChainHeight := validatorSet.PChainHeight
				firstValidatorCount := len(validatorSet.Validators)

				// =========================================================================
				// Step 4: Add a new validator to trigger updateValidatorSet
				// =========================================================================
				log.Info("Adding a new validator to trigger updateValidatorSet...")

				// Get validator manager and PoA manager
				validatorManagerProxy, poaManagerProxy := avalancheNetwork.GetValidatorManager(l1Info.SubnetID)
				poaManager, err := poamanager.NewPoAManager(poaManagerProxy.Address, l1Info.RPCClient)
				Expect(err).Should(BeNil())

				// Get P-chain info and signature aggregator
				// Use port 8082 to avoid conflict with icm-relayer (8080) and its metrics (9090)
				pChainInfo := utils.GetPChainInfo(avalancheNetwork.GetPrimaryNetworkInfo())
				aggregator := avalancheNetwork.GetSignatureAggregatorWithPort(8082)
				defer aggregator.Shutdown()

				// Get an extra node to add as validator
				newNodes := avalancheNetwork.GetExtraNodes(1)
				Expect(len(newNodes)).Should(Equal(1), "Should have at least 1 extra node available")

				// Add the node as a subnet validator
				log.Info("Adding extra node as subnet validator",
					zap.Stringer("nodeID", newNodes[0].NodeID),
				)
				l1Info = avalancheNetwork.AddSubnetValidators(newNodes, l1Info, true)

				// Register the validator with PoA manager
				addValidatorCtx, addValidatorCancel := context.WithTimeout(ctx, 120*time.Second)
				defer addValidatorCancel()

				expiry := uint64(time.Now().Add(24 * time.Hour).Unix())
				pop, err := newNodes[0].GetProofOfPossession()
				Expect(err).Should(BeNil())

				// Use a smaller weight to stay within the 20% max churn rate
				// Total weight is ~197852, so max churn is ~39570 (20%)
				// Using 1/5 of Schmeckle to stay safely under the limit
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

				log.Info("New validator added successfully, waiting for relayer to detect and call updateValidatorSet...")

				// Advance the chain to ensure the validator set change is visible
				err = utils.IssueTxsToAdvanceChain(ctx, l1Info.EVMChainID, fundedKey, l1Info.WSClient, 5)
				Expect(err).Should(BeNil())

				// =========================================================================
				// Step 5: Wait for updateValidatorSet to be called
				// Note: updateValidatorSet updates the existing validator set in place
				// (doesn't increment nextValidatorSetID), so we poll the validator set
				// at ID 0 and check if it has been updated with more validators.
				// =========================================================================
				updateCtx, updateCancel := context.WithTimeout(ctx, 90*time.Second)
				defer updateCancel()

				updateTicker := time.NewTicker(2 * time.Second)
				defer updateTicker.Stop()

				for {
					select {
					case <-updateCtx.Done():
						Expect(false).Should(BeTrue(), "Timeout waiting for updateValidatorSet - validator set change should have been detected")
					case <-updateTicker.C:
						// Check if the validator set at ID 0 has been updated
						// updateValidatorSet updates in place, so we look for:
						// - More validators than before
						// - Higher P-chain height
						updatedValidatorSet, err := registry.GetValidatorSet(callOpts, big.NewInt(0))
						if err != nil {
							log.Warn("Failed to query registry for update check", zap.Error(err))
							continue
						}

						// Check if the validator set has been updated with more validators
						if len(updatedValidatorSet.Validators) > firstValidatorCount &&
							updatedValidatorSet.PChainHeight > firstPChainHeight {
							log.Info("updateValidatorSet was called!",
								zap.Int("validatorCount", len(updatedValidatorSet.Validators)),
								zap.Uint64("pChainHeight", updatedValidatorSet.PChainHeight),
							)

							log.Info("Verified updated validator set",
								zap.Int64("validatorSetID", 0),
								zap.Uint64("pChainHeight", updatedValidatorSet.PChainHeight),
								zap.Uint64("pChainTimestamp", updatedValidatorSet.PChainTimestamp),
								zap.Int("validatorCount", len(updatedValidatorSet.Validators)),
								zap.Uint64("totalWeight", updatedValidatorSet.TotalWeight),
							)

							// Verify the updated validator set has more validators
							Expect(len(updatedValidatorSet.Validators)).Should(BeNumerically(">", firstValidatorCount),
								"Updated validator set should have more validators after adding one")
							Expect(updatedValidatorSet.PChainHeight).Should(BeNumerically(">", firstPChainHeight),
								"Updated validator set should have higher P-chain height")
							Expect(updatedValidatorSet.PChainTimestamp).Should(BeNumerically(">", 0),
								"Updated validator set should have positive timestamp")

							log.Info("Step 2 PASSED: updateValidatorSet verified successfully!",
								zap.Int("firstValidatorCount", firstValidatorCount),
								zap.Int("updatedValidatorCount", len(updatedValidatorSet.Validators)),
								zap.Uint64("firstPChainHeight", firstPChainHeight),
								zap.Uint64("updatedPChainHeight", updatedValidatorSet.PChainHeight),
							)

							log.Info("ValidatorSetUpdater test completed successfully - both registerValidatorSet and updateValidatorSet verified!",
								zap.Stringer("registryAddress", registryAddress),
							)
							return
						}

						log.Debug("Waiting for updateValidatorSet...",
							zap.Int("currentValidatorCount", len(updatedValidatorSet.Validators)),
							zap.Int("expectedValidatorCount", firstValidatorCount+1),
							zap.Uint64("currentPChainHeight", updatedValidatorSet.PChainHeight),
						)
					}
				}
			}

			log.Debug("Waiting for validator set registration...",
				zap.Uint32("nextValidatorSetID", nextID),
			)
		}
	}
}

// createRelayerConfigWithValidatorSetUpdater creates a relayer config with external EVM destinations
func createRelayerConfigWithValidatorSetUpdater(
	log logging.Logger,
	teleporter utils.TeleporterTestInfo,
	l1Info interfaces.L1TestInfo,
	fundedAddress common.Address,
	relayerKey *ecdsa.PrivateKey,
	ethereumNetwork *network.LocalEthereumNetwork,
	registryAddress common.Address,
) relayercfg.Config {
	// Get the base relayer config
	baseConfig := utils.CreateDefaultRelayerConfig(
		log,
		teleporter,
		[]interfaces.L1TestInfo{l1Info},
		[]interfaces.L1TestInfo{l1Info},
		fundedAddress,
		relayerKey,
	)

	// Get Ethereum funded key for signing transactions on Ethereum
	_, ethereumFundedKey := ethereumNetwork.GetFundedAccountInfo()

	// Add external EVM destinations - this automatically enables the validator set updater
	baseConfig.ExternalEVMDestinations = []*relayercfg.ExternalEVMDestination{
		{
			ChainID: ethereumNetwork.ChainID.String(),
			RPCEndpoint: config.APIConfig{
				BaseURL: ethereumNetwork.BaseURL,
			},
			RegistryAddress:           registryAddress.Hex(),
			AccountPrivateKey:         common.Bytes2Hex(crypto.FromECDSA(ethereumFundedKey)),
			BlockGasLimit:             15_000_000,
			TxInclusionTimeoutSeconds: 60,
		},
	}

	return baseConfig
}
