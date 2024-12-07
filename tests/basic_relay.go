// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"context"
	"os"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/icm-services/database"
	testUtils "github.com/ava-labs/icm-services/tests/utils"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

// This tests the basic functionality of the relayer, including:
// - Relaying from Subnet A to Subnet B
// - Relaying from Subnet B to Subnet A
// - Relaying an already delivered message
// - Setting ProcessHistoricalBlocksFromHeight in config
func BasicRelay(network *network.LocalNetwork, teleporter utils.TeleporterTestInfo) {
	subnetAInfo := network.GetPrimaryNetworkInfo()
	subnetBInfo, _ := network.GetTwoSubnets()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()
	err := testUtils.ClearRelayerStorage()
	Expect(err).Should(BeNil())

	//
	// Fund the relayer address on all subnets
	//
	ctx := context.Background()

	log.Info("Funding relayer address on all subnets")
	relayerKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	testUtils.FundRelayers(ctx, []interfaces.SubnetTestInfo{subnetAInfo, subnetBInfo}, fundedKey, relayerKey)

	//
	// Set up relayer config
	//
	relayerConfig := testUtils.CreateDefaultRelayerConfig(
		teleporter,
		[]interfaces.SubnetTestInfo{subnetAInfo, subnetBInfo},
		[]interfaces.SubnetTestInfo{subnetAInfo, subnetBInfo},
		fundedAddress,
		relayerKey,
	)

	// The config needs to be validated in order to be passed to database.GetConfigRelayerIDs
	relayerConfig.Validate()

	relayerConfigPath := testUtils.WriteRelayerConfig(relayerConfig, testUtils.DefaultRelayerCfgFname)

	//
	// Test Relaying from Subnet A to Subnet B
	//
	log.Info("Test Relaying from Subnet A to Subnet B")

	log.Info("Starting the relayer")
	relayerCleanup, readyChan := testUtils.RunRelayerExecutable(
		ctx,
		relayerConfigPath,
		relayerConfig,
	)
	defer relayerCleanup()

	// Wait for relayer to start up
	startupCtx, startupCancel := context.WithTimeout(ctx, 60*time.Second)
	defer startupCancel()
	testUtils.WaitForChannelClose(startupCtx, readyChan)

	log.Info("Sending transaction from Subnet A to Subnet B")
	testUtils.RelayBasicMessage(
		ctx,
		teleporter,
		subnetAInfo,
		subnetBInfo,
		fundedKey,
		fundedAddress,
	)

	//
	// Test Relaying from Subnet B to Subnet A
	//
	log.Info("Test Relaying from Subnet B to Subnet A")
	testUtils.RelayBasicMessage(
		ctx,
		teleporter,
		subnetBInfo,
		subnetAInfo,
		fundedKey,
		fundedAddress,
	)

	log.Info("Finished sending warp message, closing down output channel")
	// Cancel the command and stop the relayer
	relayerCleanup()

	//
	// Try Relaying Already Delivered Message
	//
	log.Info("Test Relaying Already Delivered Message")
	logger := logging.NewLogger(
		"icm-relayer",
		logging.NewWrappedCore(
			logging.Info,
			os.Stdout,
			logging.JSON.ConsoleEncoder(),
		),
	)
	jsonDB, err := database.NewJSONFileStorage(
		logger,
		relayerConfig.StorageLocation,
		database.GetConfigRelayerIDs(&relayerConfig),
	)
	Expect(err).Should(BeNil())

	// Create relayer keys that allow all source and destination addresses
	relayerIDA := database.CalculateRelayerID(
		subnetAInfo.BlockchainID,
		subnetBInfo.BlockchainID,
		database.AllAllowedAddress,
		database.AllAllowedAddress,
	)
	relayerIDB := database.CalculateRelayerID(
		subnetBInfo.BlockchainID,
		subnetAInfo.BlockchainID,
		database.AllAllowedAddress,
		database.AllAllowedAddress,
	)
	// Modify the JSON database to force the relayer to re-process old blocks
	err = jsonDB.Put(relayerIDA, database.LatestProcessedBlockKey, []byte("0"))
	Expect(err).Should(BeNil())
	err = jsonDB.Put(relayerIDB, database.LatestProcessedBlockKey, []byte("0"))
	Expect(err).Should(BeNil())

	// Subscribe to the destination chain
	newHeadsB := make(chan *types.Header, 10)
	sub, err := subnetBInfo.WSClient.SubscribeNewHead(ctx, newHeadsB)
	Expect(err).Should(BeNil())
	defer sub.Unsubscribe()

	// Run the relayer
	log.Info("Creating new relayer instance to test already delivered message")
	relayerCleanup, readyChan = testUtils.RunRelayerExecutable(
		ctx,
		relayerConfigPath,
		relayerConfig,
	)
	defer relayerCleanup()

	// Wait for relayer to start up
	log.Info("Waiting for the relayer to start up")
	startupCtx, startupCancel = context.WithTimeout(ctx, 60*time.Second)
	defer startupCancel()
	testUtils.WaitForChannelClose(startupCtx, readyChan)

	// We should not receive a new block on subnet B, since the relayer should have
	// seen the Teleporter message was already delivered.
	log.Info("Waiting for 10s to ensure no new block confirmations on destination chain")
	Consistently(newHeadsB, 10*time.Second, 500*time.Millisecond).ShouldNot(Receive())

	//
	// Set ProcessHistoricalBlocksFromHeight in config
	//
	log.Info("Test Setting ProcessHistoricalBlocksFromHeight in config")
	testUtils.TriggerProcessMissedBlocks(
		ctx,
		teleporter,
		subnetAInfo,
		subnetBInfo,
		relayerCleanup,
		relayerConfig,
		fundedAddress,
		fundedKey,
	)
}
