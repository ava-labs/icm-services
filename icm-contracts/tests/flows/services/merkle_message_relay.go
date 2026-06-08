// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/logging"
	merkleregistry "github.com/ava-labs/icm-services/abi-bindings/go/MerkleValidatorSetRegistry"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	"github.com/ava-labs/icm-services/config"
	"github.com/ava-labs/icm-services/icm-contracts/tests/network"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	teleporterUtils "github.com/ava-labs/icm-services/icm-contracts/utils/teleporter-utils"
	"github.com/ava-labs/icm-services/peers/clients"
	relayercfg "github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/icm-services/relayer/validatorupdater"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

// MerkleMessageRelay verifies that the relayer process itself (not the test) delivers a
// TeleporterV2 application message from an Avalanche L1 to an external EVM chain (the local
// "Ethereum" network) using the Merkle verification path.
//
// Flow:
//  1. Deploy a MerkleValidatorSetRegistry (the TeleporterMessengerV2 adapter) with the P-chain
//     genesis commitment on both the L1 and Ethereum at the same address (universal deployer).
//  2. Deploy TeleporterMessengerV2 bound to that registry on both chains (same address).
//  3. Start the relayer configured with: a source L1 "teleporter-v2" message-contracts entry and an
//     external-EVM delivery destination (contract-type "merkle"). The same external destination
//     also runs the Merkle validator-set updater that registers/refreshes the L1 set on Ethereum.
//  4. Wait for the relayer to register the L1 validator set on Ethereum.
//  5. Send a TeleporterV2 message from the L1 to Ethereum and assert the relayer delivers it
//     (MessageReceived becomes true on the Ethereum TeleporterMessengerV2).
func MerkleMessageRelay(
	ctx context.Context,
	log logging.Logger,
	avalancheNetwork *network.LocalAvalancheNetwork,
	ethereumNetwork *network.LocalEthereumNetwork,
	teleporter utils.TeleporterTestInfo,
) {
	log.Info("Starting MerkleMessageRelay e2e test")

	l1Info := avalancheNetwork.GetL1Infos()[2]
	l1BlockchainID := l1Info.BlockchainID
	networkID := avalancheNetwork.GetNetworkID()
	ethInfo := ethereumNetwork.EthereumTestInfo()
	ethBlockchainID := ethInfo.ChainID()

	fundedAddress, fundedKey := avalancheNetwork.GetFundedAccountInfo()
	ethFundedAddress, ethFundedKey := ethereumNetwork.GetFundedAccountInfo()

	// =========================================================================
	// Compute the P-chain genesis commitment for the registry constructor
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
	utils.SortValidators(pChainValidators)
	pChainGenesisRoot := validatorupdater.BuildMerkleRoot(pChainValidators)
	var pChainTotalWeight uint64
	for _, v := range pChainValidators {
		pChainTotalWeight += v.Weight
	}

	// =========================================================================
	// Deploy the MerkleValidatorSetRegistry (adapter) on both chains, then the
	// TeleporterMessengerV2 bound to it. The universal deployer yields identical
	// addresses across chains for both contracts.
	// =========================================================================
	registryAddr := utils.DeployMerkleValidatorSetRegistry(
		ctx, ethInfo, ethFundedKey, networkID, constants.PlatformChainID,
		pChainGenesisRoot, pChainTotalWeight, pChainHeight, pChainTimestamp, true,
	)
	registryAddrL1 := utils.DeployMerkleValidatorSetRegistry(
		ctx, &l1Info, fundedKey, networkID, constants.PlatformChainID,
		pChainGenesisRoot, pChainTotalWeight, pChainHeight, pChainTimestamp, true,
	)
	Expect(registryAddrL1).Should(Equal(registryAddr))

	ethTeleporterAddr := utils.DeployTeleporterV2(ctx, ethInfo, registryAddr, ethFundedKey)
	l1TeleporterAddr := utils.DeployTeleporterV2(ctx, &l1Info, registryAddr, fundedKey)
	Expect(l1TeleporterAddr).Should(Equal(ethTeleporterAddr))

	registry, err := merkleregistry.NewMerkleValidatorSetRegistry(registryAddr, ethereumNetwork.EthClient)
	Expect(err).Should(BeNil())

	log.Info("Deployed merkle registry and TeleporterMessengerV2",
		zap.String("registryAddress", registryAddr.Hex()),
		zap.String("teleporterAddress", l1TeleporterAddr.Hex()),
	)

	// =========================================================================
	// Configure and start the relayer
	// =========================================================================
	Expect(utils.ClearRelayerStorage()).Should(BeNil())

	relayerKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	utils.FundRelayers(ctx, []testinfo.L1TestInfo{l1Info}, fundedKey, relayerKey)

	relayerConfig := createMerkleMessageRelayConfig(
		log, teleporter, l1Info, fundedAddress, relayerKey, ethereumNetwork,
		registryAddr, l1TeleporterAddr, ethFundedKey, ethBlockchainID,
	)
	Expect(relayerConfig.Validate()).Should(BeNil())
	relayerConfigPath := utils.WriteRelayerConfig(log, relayerConfig, utils.DefaultRelayerCfgFname)

	log.Info("Starting relayer")
	relayerCleanup, readyChan := utils.RunRelayerExecutable(ctx, log, relayerConfigPath, relayerConfig)
	defer relayerCleanup()

	startupCtx, startupCancel := context.WithTimeout(ctx, 60*time.Second)
	defer startupCancel()
	utils.WaitForChannelClose(startupCtx, readyChan)

	// =========================================================================
	// Wait for the relayer's Merkle updater to register the L1 validator set on Ethereum
	// =========================================================================
	callOpts := &bind.CallOpts{Context: ctx}
	log.Info("Waiting for relayer to register L1 validator set on Ethereum")
	pollForCommitmentUpdate(
		ctx, log, registry, callOpts, l1BlockchainID, 120*time.Second,
		"Timed out waiting for relayer to register L1 validator set",
		func(cmt merkleregistry.ValidatorSetMerkleCommitment) bool {
			return cmt.TotalWeight != 0
		},
		func(cmt merkleregistry.ValidatorSetMerkleCommitment) {
			log.Info("L1 validator set registered on Ethereum",
				zap.Uint64("totalWeight", cmt.TotalWeight),
				zap.Uint64("pChainHeight", cmt.PChainHeight),
				zap.String("root", hex.EncodeToString(cmt.Root[:])),
			)
		},
	)

	// =========================================================================
	// Send a TeleporterV2 message from the L1 to Ethereum
	// =========================================================================
	l1Teleporter, err := teleportermessengerv2.NewTeleporterMessengerV2(l1TeleporterAddr, l1Info.EthClient)
	Expect(err).Should(BeNil())
	l1Opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, l1Info.EVMChainID)
	Expect(err).Should(BeNil())

	input := teleportermessengerv2.TeleporterMessageInput{
		DestinationBlockchainID: ethBlockchainID,
		DestinationAddress:      ethFundedAddress,
		FeeInfo: teleportermessengerv2.TeleporterFeeInfo{
			FeeTokenAddress: common.Address{},
			Amount:          big.NewInt(0),
		},
		RequiredGasLimit:        big.NewInt(100_000),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte("Hello Ethereum from the relayer!"),
	}
	tx, err := l1Teleporter.SendCrossChainMessage(l1Opts, input)
	Expect(err).Should(BeNil())
	receipt := utils.WaitForTransactionSuccess(ctx, l1Info.EthClient, tx.Hash())

	sendEvent, err := utils.GetEventFromLogs(receipt.Logs, l1Teleporter.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(sendEvent.DestinationBlockchainID[:]).Should(Equal(ethBlockchainID[:]))

	messageID, err := teleporterUtils.CalculateMessageID(
		l1TeleporterAddr, l1BlockchainID, ethBlockchainID, sendEvent.Message.MessageNonce,
	)
	Expect(err).Should(BeNil())
	log.Info("Sent TeleporterV2 message from L1 to Ethereum", zap.Stringer("messageID", messageID))

	// Advance the source chain a few blocks so the relayer's WS subscription reliably surfaces the
	// block containing the message. subnet-evm only produces blocks when there are transactions, so
	// without follow-on activity the message block's new-head notification can be delayed.
	Expect(utils.IssueTxsToAdvanceChain(ctx, l1Info.EVMChainID, fundedKey, l1Info.EthClient, 5)).Should(BeNil())

	// =========================================================================
	// Assert the relayer delivers the message to the Ethereum TeleporterMessengerV2
	// =========================================================================
	ethTeleporter, err := teleportermessengerv2.NewTeleporterMessengerV2(ethTeleporterAddr, ethereumNetwork.EthClient)
	Expect(err).Should(BeNil())

	deliveryCtx, deliveryCancel := context.WithTimeout(ctx, 120*time.Second)
	defer deliveryCancel()
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	delivered := false
	for !delivered {
		select {
		case <-deliveryCtx.Done():
			Expect(false).Should(BeTrue(), "Timed out waiting for relayer to deliver message to Ethereum")
		case <-ticker.C:
			received, recvErr := ethTeleporter.MessageReceived(&bind.CallOpts{Context: ctx}, messageID)
			if recvErr != nil {
				log.Warn("Failed to query MessageReceived", zap.Error(recvErr))
				continue
			}
			if received {
				delivered = true
			}
		}
	}

	log.Info("MerkleMessageRelay e2e test PASSED: relayer delivered L1 -> Ethereum message")
}

// createMerkleMessageRelayConfig builds a relayer config whose source L1 routes its TeleporterV2
// messages (emitted via the Merkle registry adapter) to an external EVM delivery destination, and
// which also runs the Merkle validator-set updater against the same registry.
func createMerkleMessageRelayConfig(
	log logging.Logger,
	teleporter utils.TeleporterTestInfo,
	l1Info testinfo.L1TestInfo,
	fundedAddress common.Address,
	relayerKey *ecdsa.PrivateKey,
	ethereumNetwork *network.LocalEthereumNetwork,
	registryAddr common.Address,
	teleporterAddr common.Address,
	ethFundedKey *ecdsa.PrivateKey,
	ethBlockchainID ids.ID,
) relayercfg.Config {
	baseConfig := utils.CreateDefaultRelayerConfig(
		log,
		teleporter,
		[]testinfo.L1TestInfo{l1Info},
		[]testinfo.L1TestInfo{l1Info},
		fundedAddress,
		relayerKey,
	)
	baseConfig.APIPort = 8086

	// Register the TeleporterV2 (Merkle) message protocol for the source L1, keyed by the registry
	// adapter address (the warp message origin sender for these messages).
	source := baseConfig.SourceBlockchains[0]
	source.MessageContracts[registryAddr.Hex()] = relayercfg.MessageProtocolConfig{
		MessageFormat: relayercfg.TELEPORTER_V2.String(),
		Settings: map[string]interface{}{
			"reward-address":     fundedAddress.Hex(),
			"registry-address":   registryAddr.Hex(),
			"teleporter-address": teleporterAddr.Hex(),
		},
	}
	// Only deliver to the external Ethereum destination in this test.
	source.SupportedDestinations = []*relayercfg.SupportedDestination{
		{BlockchainID: ethBlockchainID.String()},
	}

	baseConfig.ExternalEVMDestinations = []*relayercfg.ExternalEVMDestination{
		{
			RPCEndpoint:              ethereumNetwork.BaseURL,
			PrivateKey:               hex.EncodeToString(crypto.FromECDSA(ethFundedKey)),
			ContractAddress:          registryAddr.Hex(),
			BlockchainID:             l1Info.BlockchainID.String(),
			SubnetID:                 l1Info.SubnetID.String(),
			ContractType:             "merkle",
			PollIntervalSeconds:      testPollIntervalSeconds,
			MaxUpdateIntervalSeconds: merkleMaxUpdateIntervalSeconds,

			// Message delivery
			Deliver:                 true,
			DestinationBlockchainID: ethBlockchainID.String(),
			TeleporterAddress:       teleporterAddr.Hex(),
		},
	}

	return baseConfig
}
