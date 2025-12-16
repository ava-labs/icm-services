// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporter_test

import (
	"context"
	"flag"
	"os"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/tests/fixture/e2e"
	"github.com/ava-labs/avalanchego/utils/units"
	teleporterFlows "github.com/ava-labs/icm-services/icm-contracts/tests/flows/teleporter"
	registryFlows "github.com/ava-labs/icm-services/icm-contracts/tests/flows/teleporter/registry"
	"github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/icm-services/log"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	warpGenesisTemplateFile = "./tests/utils/warp-genesis-template.json"

	teleporterMessengerLabel = "TeleporterMessenger"
	upgradabilityLabel       = "upgradability"
	utilsLabel               = "utils"

	teleporterRegistryAddressFile = "TeleporterRegistryAddress.json"
	validatorAddressesFile        = "ValidatorAddresses.json"
)

var (
	LocalNetworkInstance *network.LocalNetwork
	TeleporterInfo       utils.TeleporterTestInfo
	e2eFlags             *e2e.FlagVars
)

func TestMain(m *testing.M) {
	e2eFlags = e2e.RegisterFlags()
	flag.Parse()
	os.Exit(m.Run())
}

func TestTeleporter(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("Environment variable RUN_E2E not set; skipping E2E tests")
	}

	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Teleporter e2e test")
}

// Define the Teleporter before and after suite functions.
var _ = ginkgo.BeforeSuite(func(ctx context.Context) {
	teleporterContractAddress,
		teleporterDeployerAddress,
		teleporterDeployedByteCode := utils.TeleporterDeploymentValues()

	teleporterDeployerTransaction := utils.TeleporterDeployerTransaction()

	// Create the local network instance
	ctx, cancel := context.WithTimeout(ctx, 240*2*time.Second)
	defer cancel()

	LocalNetworkInstance = network.NewLocalNetwork(
		ctx,
		"teleporter-test-local-network",
		warpGenesisTemplateFile,
		[]network.L1Spec{
			{
				Name:                         "A",
				EVMChainID:                   12345,
				TeleporterContractAddress:    teleporterContractAddress,
				TeleporterDeployedBytecode:   teleporterDeployedByteCode,
				TeleporterDeployerAddress:    teleporterDeployerAddress,
				NodeCount:                    5,
				RequirePrimaryNetworkSigners: true,
			},
			{
				Name:                         "B",
				EVMChainID:                   54321,
				TeleporterContractAddress:    teleporterContractAddress,
				TeleporterDeployedBytecode:   teleporterDeployedByteCode,
				TeleporterDeployerAddress:    teleporterDeployerAddress,
				NodeCount:                    5,
				RequirePrimaryNetworkSigners: true,
			},
		},
		2,
		2,
		e2eFlags,
	)
	TeleporterInfo = utils.NewTeleporterTestInfo(LocalNetworkInstance.GetAllL1Infos())
	log.Info("Started local network")

	// Only need to deploy Teleporter on the C-Chain since it is included in the genesis of the l1 chains.
	_, fundedKey := LocalNetworkInstance.GetFundedAccountInfo()
	if e2eFlags.NetworkDir() == "" {
		utils.DeployTeleporterMessenger(
			ctx,
			LocalNetworkInstance.GetPrimaryNetworkInfo(),
			teleporterDeployerTransaction,
			teleporterDeployerAddress,
			teleporterContractAddress,
			fundedKey,
		)
		balance := 100 * units.Avax
		for _, subnet := range LocalNetworkInstance.GetL1Infos() {
			// Choose weights such that we can test validator churn
			LocalNetworkInstance.ConvertSubnet(
				ctx,
				subnet,
				utils.PoAValidatorManager,
				[]uint64{units.Schmeckle, units.Schmeckle, units.Schmeckle, units.Schmeckle, units.Schmeckle},
				[]uint64{balance, balance, balance, balance, balance},
				fundedKey,
				false,
			)
		}

		for _, l1 := range LocalNetworkInstance.GetAllL1Infos() {
			TeleporterInfo.SetTeleporter(teleporterContractAddress, l1)
			TeleporterInfo.InitializeBlockchainID(l1, fundedKey)
			TeleporterInfo.DeployTeleporterRegistry(l1, fundedKey)
		}

		// Save the Teleporter registry address and validator addresses to files
		utils.SaveRegistyAddress(TeleporterInfo, teleporterRegistryAddressFile)

		LocalNetworkInstance.SaveValidatorAddress(validatorAddressesFile)
	} else {
		// Read the Teleporter registry address from the file
		utils.SetTeleporterInfoFromFile(
			teleporterRegistryAddressFile,
			teleporterContractAddress,
			TeleporterInfo,
			LocalNetworkInstance.GetAllL1Infos(),
		)

		// Read the validator addresses from the file
		LocalNetworkInstance.SetValidatorAddressFromFile(validatorAddressesFile)
	}

	log.Info("Set up ginkgo before suite")
})

var _ = ginkgo.AfterSuite(func() {
	LocalNetworkInstance.TearDownNetwork()
	LocalNetworkInstance = nil
})

var _ = ginkgo.Describe("[Teleporter integration tests]", func() {
	// Teleporter tests
	ginkgo.It("Send a message from L1 A to L1 B, and one from B to A",
		ginkgo.Label(teleporterMessengerLabel),
		func(ctx context.Context) {
			teleporterFlows.BasicSendReceive(ctx, LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Deliver to the wrong chain",
		ginkgo.Label(teleporterMessengerLabel),
		func(ctx context.Context) {
			teleporterFlows.DeliverToWrongChain(ctx, LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Deliver to non-existent contract",
		ginkgo.Label(teleporterMessengerLabel),
		func(ctx context.Context) {
			teleporterFlows.DeliverToNonExistentContract(ctx, LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Retry successful execution",
		ginkgo.Label(teleporterMessengerLabel),
		func(ctx context.Context) {
			teleporterFlows.RetrySuccessfulExecution(ctx, LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Unallowed relayer",
		ginkgo.Label(teleporterMessengerLabel),
		func(ctx context.Context) {
			teleporterFlows.UnallowedRelayer(ctx, LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Relay message twice",
		ginkgo.Label(teleporterMessengerLabel),
		func(ctx context.Context) {
			teleporterFlows.RelayMessageTwice(ctx, LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Add additional fee amount",
		ginkgo.Label(teleporterMessengerLabel),
		func(ctx context.Context) {
			teleporterFlows.AddFeeAmount(ctx, LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Send specific receipts",
		ginkgo.Label(teleporterMessengerLabel),
		func(ctx context.Context) {
			teleporterFlows.SendSpecificReceipts(ctx, LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Insufficient gas",
		ginkgo.Label(teleporterMessengerLabel),
		func(ctx context.Context) {
			teleporterFlows.InsufficientGas(ctx, LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Resubmit altered message",
		ginkgo.Label(teleporterMessengerLabel),
		func(ctx context.Context) {
			teleporterFlows.ResubmitAlteredMessage(ctx, LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Calculate Teleporter message IDs",
		ginkgo.Label(utilsLabel),
		func(ctx context.Context) {
			teleporterFlows.CalculateMessageID(ctx, LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Relayer modifies message",
		ginkgo.Label(teleporterMessengerLabel),
		func(ctx context.Context) {
			teleporterFlows.RelayerModifiesMessage(ctx, LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Validator churn",
		ginkgo.Label(teleporterMessengerLabel),
		func(ctx context.Context) {
			teleporterFlows.ValidatorChurn(ctx, LocalNetworkInstance, TeleporterInfo)
		})

	// Teleporter Registry tests
	ginkgo.It("Teleporter registry",
		ginkgo.Label(upgradabilityLabel),
		func(ctx context.Context) {
			registryFlows.TeleporterRegistry(ctx, LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Check upgrade access",
		ginkgo.Label(upgradabilityLabel),
		func(ctx context.Context) {
			registryFlows.CheckUpgradeAccess(ctx, LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Pause and Unpause Teleporter",
		ginkgo.Label(upgradabilityLabel),
		func(ctx context.Context) {
			registryFlows.PauseTeleporter(ctx, LocalNetworkInstance, TeleporterInfo)
		})
})
