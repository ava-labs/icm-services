// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"context"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/icm-services/database"
	"github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/testinfo"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/crypto"
	. "github.com/onsi/gomega"
)

// This tests the basic functionality of the relayer, including:
// - Relaying from Subnet A to Subnet B
// - Relaying from Subnet B to Subnet A
// - Relaying an already delivered message
// - Setting ProcessHistoricalBlocksFromHeight in config
func BasicRelay(
	ctx context.Context,
	log logging.Logger,
	network *network.LocalAvalancheNetwork,
	teleporter utils.TeleporterTestInfo,
) {
	l1AInfo := network.GetPrimaryNetworkInfo()
	l1BInfo, _ := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()
	err := utils.ClearRelayerStorage()
	Expect(err).Should(BeNil())

	//
	// Fund the relayer address on all subnets
	//

	log.Info("Funding relayer address on all subnets")
	relayerKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	utils.FundRelayers(ctx, []testinfo.L1TestInfo{l1AInfo, l1BInfo}, fundedKey, relayerKey)

	//
	// Set up relayer config
	//
	relayerConfig := utils.CreateDefaultRelayerConfig(
		log,
		teleporter,
		[]testinfo.L1TestInfo{l1AInfo, l1BInfo},
		[]testinfo.L1TestInfo{l1AInfo, l1BInfo},
		fundedAddress,
		relayerKey,
	)

	// The config needs to be validated in order to be passed to database.GetConfigRelayerIDs
	relayerConfig.Validate()

	relayerConfigPath := utils.WriteRelayerConfig(log, relayerConfig, utils.DefaultRelayerCfgFname)

	//
	// Test Relaying from Subnet A to Subnet B
	//
	log.Info("Test Relaying from Subnet A to Subnet B")

	log.Info("Starting the relayer")
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

	log.Info("Sending transaction from Subnet A to Subnet B")
	utils.RelayBasicMessage(
		ctx,
		log,
		teleporter,
		l1AInfo,
		l1BInfo,
		fundedKey,
		fundedAddress,
	)

	//
	// Test Relaying from Subnet B to Subnet A
	//
	log.Info("Test Relaying from Subnet B to Subnet A")
	utils.RelayBasicMessage(
		ctx,
		log,
		teleporter,
		l1BInfo,
		l1AInfo,
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
	relayerConfig.ProcessMissedBlocks = true
	relayerConfigPath = utils.WriteRelayerConfig(log, relayerConfig, utils.DefaultRelayerCfgFname)

	jsonDB, err := database.NewJSONFileStorage(
		log,
		relayerConfig.StorageLocation,
		database.GetConfigRelayerIDs(&relayerConfig),
	)
	Expect(err).Should(BeNil())

	// Create relayer keys that allow all source and destination addresses
	relayerIDA := database.CalculateRelayerID(
		l1AInfo.BlockchainID,
		l1BInfo.BlockchainID,
		database.AllAllowedAddress,
		database.AllAllowedAddress,
	)
	relayerIDB := database.CalculateRelayerID(
		l1BInfo.BlockchainID,
		l1AInfo.BlockchainID,
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
	sub, err := l1BInfo.WSClient.SubscribeNewHead(ctx, newHeadsB)
	Expect(err).Should(BeNil())
	defer sub.Unsubscribe()

	// Run the relayer
	log.Info("Creating new relayer instance to test already delivered message")
	relayerCleanup, readyChan = utils.RunRelayerExecutable(
		ctx,
		log,
		relayerConfigPath,
		relayerConfig,
	)
	defer relayerCleanup()

	// Wait for relayer to start up
	log.Info("Waiting for the relayer to start up")
	startupCtx, startupCancel = context.WithTimeout(ctx, 60*time.Second)
	defer startupCancel()
	utils.WaitForChannelClose(startupCtx, readyChan)

	// We should not receive a new block on subnet B, since the relayer should have
	// seen the Teleporter message was already delivered.
	log.Info("Waiting for 10s to ensure no new block confirmations on destination chain")
	Consistently(newHeadsB, 10*time.Second, 500*time.Millisecond).ShouldNot(Receive())

	//
	// Set ProcessHistoricalBlocksFromHeight in config
	//
	log.Info("Test Setting ProcessHistoricalBlocksFromHeight in config")
	utils.TriggerProcessMissedBlocks(
		ctx,
		log,
		teleporter,
		l1AInfo,
		l1BInfo,
		relayerCleanup,
		relayerConfig,
		fundedAddress,
		fundedKey,
	)
}
