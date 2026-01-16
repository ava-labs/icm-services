// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/icm-services/icm-contracts/tests/interfaces"
	"github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	offchainregistry "github.com/ava-labs/icm-services/messages/off-chain-registry"
	"github.com/ava-labs/icm-services/relayer/api"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	. "github.com/onsi/gomega"
)

// Tests relayer support for off-chain Teleporter Registry updates
// - Configures the relayer to send an off-chain message to the Teleporter Registry
// - Verifies that the Teleporter Registry is updated
func ManualMessage(
	ctx context.Context,
	log logging.Logger,
	network *network.LocalAvalancheNetwork,
	teleporter utils.TeleporterTestInfo,
) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	l1AInfo, l1BInfo := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()
	err := utils.ClearRelayerStorage()
	Expect(err).Should(BeNil())

	//
	// Get the current Teleporter Registry version
	//
	currentVersion, err := teleporter.TeleporterRegistry(cChainInfo).LatestVersion(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	expectedNewVersion := currentVersion.Add(currentVersion, big.NewInt(1))

	//
	// Fund the relayer address on all subnets
	//

	log.Info("Funding relayer address on all subnets")
	relayerKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	utils.FundRelayers(ctx, []interfaces.L1TestInfo{cChainInfo}, fundedKey, relayerKey)

	//
	// Define the off-chain Warp message
	//
	log.Info("Creating off-chain Warp message")
	newProtocolAddress := common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567")
	networkID := network.GetNetworkID()

	//
	// Set up the nodes to accept the off-chain message
	//
	// Create chain config file with off chain message for each chain
	unsignedMessage, warpEnabledChainConfigC := utils.InitOffChainMessageChainConfig(
		networkID,
		cChainInfo,
		teleporter.TeleporterRegistryAddress(cChainInfo),
		newProtocolAddress,
		2,
	)

	_, warpEnabledChainConfigA := utils.InitOffChainMessageChainConfig(
		networkID,
		l1AInfo,
		teleporter.TeleporterRegistryAddress(l1AInfo),
		newProtocolAddress,
		2,
	)

	_, warpEnabledChainConfigB := utils.InitOffChainMessageChainConfig(
		networkID,
		l1BInfo,
		teleporter.TeleporterRegistryAddress(l1BInfo),
		newProtocolAddress,
		2,
	)

	// Create chain config with off chain messages

	chainConfigs := make(utils.ChainConfigMap)
	chainConfigs.Add(cChainInfo, warpEnabledChainConfigC)
	chainConfigs.Add(l1BInfo, warpEnabledChainConfigB)
	chainConfigs.Add(l1AInfo, warpEnabledChainConfigA)

	// Restart nodes with new chain config
	log.Info("Restarting nodes with new chain config")
	network.SetChainConfigs(chainConfigs)

	// Refresh the subnet info to get the new clients
	cChainInfo = network.GetPrimaryNetworkInfo()

	//
	// Set up relayer config
	//
	relayerConfig := utils.CreateDefaultRelayerConfig(
		log,
		teleporter,
		[]interfaces.L1TestInfo{cChainInfo},
		[]interfaces.L1TestInfo{cChainInfo},
		fundedAddress,
		relayerKey,
	)
	relayerConfigPath := utils.WriteRelayerConfig(
		log,
		relayerConfig,
		utils.DefaultRelayerCfgFname,
	)

	log.Info("Starting the relayer")
	relayerCleanup, readyChan := utils.RunRelayerExecutable(
		ctx,
		log,
		relayerConfigPath,
		relayerConfig,
	)
	defer relayerCleanup()

	// Wait for relayer to startup.
	log.Info("Waiting for the relayer to start up")
	startupCtx, startupCancel := context.WithTimeout(ctx, 15*time.Second)
	defer startupCancel()
	utils.WaitForChannelClose(startupCtx, readyChan)

	reqBody := api.ManualWarpMessageRequest{
		UnsignedMessageBytes: unsignedMessage.Bytes(),
		SourceAddress:        offchainregistry.OffChainRegistrySourceAddress.Hex(),
	}

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	requestURL := fmt.Sprintf("http://localhost:%d%s", relayerConfig.APIPort, api.RelayMessageAPIPath)

	// Send request to API
	{
		b, err := json.Marshal(reqBody)
		Expect(err).Should(BeNil())
		bodyReader := bytes.NewReader(b)

		req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
		Expect(err).Should(BeNil())
		req.Header.Set("Content-Type", "application/json")

		res, err := client.Do(req)
		Expect(err).Should(BeNil())
		Expect(res.Status).Should(Equal("200 OK"))

		// Wait for all nodes to see new transaction
		time.Sleep(1 * time.Second)

		newVersion, err := teleporter.TeleporterRegistry(cChainInfo).LatestVersion(&bind.CallOpts{})
		Expect(err).Should(BeNil())
		Expect(newVersion.Uint64()).Should(Equal(expectedNewVersion.Uint64()))
	}
}
