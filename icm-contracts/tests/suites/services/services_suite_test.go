// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package services_test

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/tests/fixture/e2e"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/units"
	servicesFlows "github.com/ava-labs/icm-services/icm-contracts/tests/flows/services"
	"github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

const (
	warpGenesisTemplateFile   = "./tests/utils/warp-genesis-template.json"
	servicesLabel             = "ICMServices"
	minimumL1ValidatorBalance = 2048 * units.NanoAvax
	defaultBalance            = 100 * units.Avax
)

var (
	log logging.Logger

	localAvalancheNetworkInstance *network.LocalAvalancheNetwork
	localEthereumNetworkInstance  *network.LocalEthereumNetwork
	teleporterInfo                utils.TeleporterTestInfo

	decider *exec.Cmd

	e2eFlags *e2e.FlagVars
)

func TestMain(m *testing.M) {
	e2eFlags = e2e.RegisterFlags()
	flag.Parse()
	os.Exit(m.Run())
}

func TestServices(t *testing.T) {
	// Handle SIGINT and SIGTERM signals.
	signalChan := make(chan os.Signal, 2)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-signalChan
		fmt.Printf("Caught signal %s: Shutting down...\n", sig)
		cleanup()
		os.Exit(1)
	}()

	if os.Getenv("RUN_E2E") == "" {
		t.Skip("Environment variable RUN_E2E not set; skipping E2E tests")
	}

	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Relayer e2e test")
}

// Define the Relayer before and after suite functions.
var _ = ginkgo.BeforeSuite(func(ctx context.Context) {
	log = logging.NewLogger(
		"signature-aggregator",
		logging.NewWrappedCore(
			logging.Info,
			os.Stdout,
			logging.JSON.ConsoleEncoder(),
		),
	)

	log.Info("Building all ICM service executables")
	utils.BuildAllExecutables(ctx, log)

	teleporterContractAddress,
		teleporterDeployerAddress,
		teleporterDeployedByteCode := utils.TeleporterDeploymentValues()

	teleporterDeployerTransaction := utils.TeleporterDeployerTransaction()

	networkStartCtx, networkStartCancel := context.WithTimeout(ctx, 240*2*time.Second)
	defer networkStartCancel()
	localAvalancheNetworkInstance = network.NewLocalAvalancheNetwork(
		networkStartCtx,
		"icm-off-chain-services-e2e-test",
		warpGenesisTemplateFile,
		[]network.L1Spec{
			{
				Name:                         "A",
				EVMChainID:                   12345,
				TeleporterContractAddress:    teleporterContractAddress,
				TeleporterDeployedBytecode:   teleporterDeployedByteCode,
				TeleporterDeployerAddress:    teleporterDeployerAddress,
				NodeCount:                    2,
				RequirePrimaryNetworkSigners: true,
			},
			{
				Name:                         "B",
				EVMChainID:                   54321,
				TeleporterContractAddress:    teleporterContractAddress,
				TeleporterDeployedBytecode:   teleporterDeployedByteCode,
				TeleporterDeployerAddress:    teleporterDeployerAddress,
				NodeCount:                    2,
				RequirePrimaryNetworkSigners: true,
			},
		},
		4,
		6, // Extra nodes: 4 for L1 validators + 2 for validator set updater test
		e2eFlags,
	)

	// Only need to deploy Teleporter on the C-Chain since it is included in the genesis of the L1 chains.
	_, fundedKey := localAvalancheNetworkInstance.GetFundedAccountInfo()
	utils.DeployWithNicksMethod(
		networkStartCtx,
		localAvalancheNetworkInstance.GetPrimaryNetworkInfo(),
		teleporterDeployerTransaction,
		teleporterDeployerAddress,
		teleporterContractAddress,
		fundedKey,
	)

	teleporterInfo = utils.NewTeleporterTestInfo(localAvalancheNetworkInstance.GetAllL1Infos())
	// Deploy the Teleporter registry contracts to all subnets and the C-Chain.
	for _, subnet := range localAvalancheNetworkInstance.GetAllL1Infos() {
		teleporterInfo.SetTeleporter(teleporterContractAddress, subnet)
		teleporterInfo.DeployTeleporterRegistry(subnet, fundedKey)
	}

	// Convert the subnets to sovereign L1s
	for _, subnet := range localAvalancheNetworkInstance.GetL1Infos() {
		localAvalancheNetworkInstance.ConvertSubnet(
			networkStartCtx,
			subnet,
			utils.PoAValidatorManager,
			[]uint64{units.Schmeckle, units.Schmeckle, units.Schmeckle, units.Schmeckle},
			[]uint64{defaultBalance, defaultBalance, defaultBalance, minimumL1ValidatorBalance - 1},
			fundedKey,
			false,
		)
	}

	// Restart the network to attempt to refresh TLS connections
	networkRestartCtx, cancel := context.WithTimeout(ctx, time.Duration(60*len(localAvalancheNetworkInstance.Nodes))*time.Second)
	defer cancel()

	err := localAvalancheNetworkInstance.Restart(networkRestartCtx)
	Expect(err).Should(BeNil())

	decider = exec.CommandContext(ctx, "./tests/cmd/decider/decider")
	decider.Start()
	go func() {
		err := decider.Wait()
		// Context cancellation is the only expected way for the process to exit
		// otherwise log an error but don't panic to allow for easier cleanup
		if !errors.Is(ctx.Err(), context.Canceled) {
			log.Error("Decider exited abnormally: ", zap.Error(err))
		}
	}()
	log.Info("Started decider service")

	localEthereumNetworkInstance = network.NewLocalEthereumNetwork(ctx)
	log.Info("Connected to local Ethereum network")

	log.Info("Set up ginkgo before suite")

	ginkgo.AddReportEntry(
		"network directory with node logs & configs; useful in the case of failures",
		localAvalancheNetworkInstance.Dir(),
		ginkgo.ReportEntryVisibilityFailureOrVerbose,
	)
})

func cleanup() {
	if decider != nil {
		decider = nil
	}
	if localEthereumNetworkInstance != nil {
		localEthereumNetworkInstance.TearDownNetwork()
		localEthereumNetworkInstance = nil
	}
	if localAvalancheNetworkInstance != nil {
		localAvalancheNetworkInstance.TearDownNetwork()
		localAvalancheNetworkInstance = nil
	}
}

var _ = ginkgo.AfterSuite(cleanup)

var _ = ginkgo.Describe("[ICM Relayer & Signature Aggregator Integration Tests", func() {
	ginkgo.It("Basic Relay",
		ginkgo.Label(servicesLabel),
		func(ctx context.Context) {
			servicesFlows.BasicRelay(ctx, log, localNetworkInstance, teleporterInfo)
		})
	ginkgo.It("Manually Provided Message",
		ginkgo.Label(servicesLabel),
		func(ctx context.Context) {
			servicesFlows.ManualMessage(ctx, log, localNetworkInstance, teleporterInfo)
		})
	ginkgo.It("Shared Database",
		ginkgo.Label(servicesLabel),
		func(ctx context.Context) {
			servicesFlows.SharedDatabaseAccess(ctx, log, localNetworkInstance, teleporterInfo)
		})
	ginkgo.It("Allowed Addresses",
		ginkgo.Label(servicesLabel),
		func(ctx context.Context) {
			servicesFlows.AllowedAddresses(ctx, log, localNetworkInstance, teleporterInfo)
		})
	ginkgo.It("Batch Message",
		ginkgo.Label(servicesLabel),
		func(ctx context.Context) {
			servicesFlows.BatchRelay(ctx, log, localNetworkInstance, teleporterInfo)
		})
	ginkgo.It("Relay Message API",
		ginkgo.Label(servicesLabel),
		func(ctx context.Context) {
			servicesFlows.RelayMessageAPI(ctx, log, localNetworkInstance, teleporterInfo)
		})
	ginkgo.It("Warp API",
		ginkgo.Label(servicesLabel),
		func(ctx context.Context) {
			servicesFlows.WarpAPIRelay(ctx, log, localNetworkInstance, teleporterInfo)
		})
	ginkgo.It("Signature Aggregator",
		ginkgo.Label(servicesLabel),
		func(ctx context.Context) {
			servicesFlows.SignatureAggregatorAPI(ctx, log, localNetworkInstance, teleporterInfo)
		})
	ginkgo.It("Signature Aggregator Epoch Validators",
		ginkgo.Label(servicesLabel),
		func(ctx context.Context) {
			servicesFlows.SignatureAggregatorEpochAPI(ctx, log, localNetworkInstance, teleporterInfo)
		})
	ginkgo.It("Validators Only Network",
		ginkgo.Label(servicesLabel),
		func(ctx context.Context) {
			servicesFlows.ValidatorsOnlyNetwork(ctx, log, localNetworkInstance, teleporterInfo)
		})
	ginkgo.It("External EVM Validator Set Updater",
		ginkgo.Label(servicesLabel),
		func(ctx context.Context) {
			servicesFlows.ValidatorSetUpdater(ctx, log, localNetworkInstance, localEthereumNetworkInstance, teleporterInfo)
		})
})
