// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"context"
	"encoding/hex"
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
	"github.com/ava-labs/icm-services/icm-contracts/tests/network"
	teleporterTestUtils "github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	testUtils "github.com/ava-labs/icm-services/tests/utils"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/libevm/common"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

const (
	warpGenesisTemplateFile   = "./tests/utils/warp-genesis-template.json"
	minimumL1ValidatorBalance = 2048 * units.NanoAvax
	defaultBalance            = 100 * units.Avax
)

var (
	log logging.Logger

	localNetworkInstance *network.LocalNetwork
	teleporterInfo       teleporterTestUtils.TeleporterTestInfo

	decider  *exec.Cmd
	cancelFn context.CancelFunc

	e2eFlags *e2e.FlagVars
)

func TestMain(m *testing.M) {
	e2eFlags = e2e.RegisterFlags()
	flag.Parse()
	os.Exit(m.Run())
}

func TestE2E(t *testing.T) {
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
	ctx, cancelFn = context.WithCancel(ctx)

	log = logging.NewLogger(
		"signature-aggregator",
		logging.NewWrappedCore(
			logging.Info,
			os.Stdout,
			logging.JSON.ConsoleEncoder(),
		),
	)

	log.Info("Building all ICM service executables")
	testUtils.BuildAllExecutables(ctx, log)

	// Generate the Teleporter deployment values
	teleporterContractAddress := common.HexToAddress(
		testUtils.ReadHexTextFile("./tests/utils/UniversalTeleporterMessengerContractAddress.txt"),
	)
	teleporterDeployerAddress := common.HexToAddress(
		testUtils.ReadHexTextFile("./tests/utils/UniversalTeleporterDeployerAddress.txt"),
	)
	teleporterDeployedByteCode := testUtils.ReadHexTextFile(
		"./tests/utils/UniversalTeleporterDeployedBytecode.txt",
	)
	teleporterDeployerTransactionStr := testUtils.ReadHexTextFile(
		"./tests/utils/UniversalTeleporterDeployerTransaction.txt",
	)
	teleporterDeployerTransaction, err := hex.DecodeString(
		utils.SanitizeHexString(teleporterDeployerTransactionStr),
	)
	Expect(err).Should(BeNil())
	networkStartCtx, networkStartCancel := context.WithTimeout(ctx, 240*2*time.Second)
	defer networkStartCancel()
	localNetworkInstance = network.NewLocalNetwork(
		networkStartCtx,
		log,
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
		4,
		e2eFlags,
	)
	teleporterInfo = teleporterTestUtils.NewTeleporterTestInfo(localNetworkInstance.GetAllL1Infos())

	// Only need to deploy Teleporter on the C-Chain since it is included in the genesis of the subnet chains.
	_, fundedKey := localNetworkInstance.GetFundedAccountInfo()
	teleporterInfo.DeployTeleporterMessenger(
		networkStartCtx,
		localNetworkInstance.GetPrimaryNetworkInfo(),
		teleporterDeployerTransaction,
		teleporterDeployerAddress,
		teleporterContractAddress,
		fundedKey,
	)

	// Deploy the Teleporter registry contracts to all subnets and the C-Chain.
	for _, subnet := range localNetworkInstance.GetAllL1Infos() {
		teleporterInfo.SetTeleporter(teleporterContractAddress, subnet)
		teleporterInfo.InitializeBlockchainID(subnet, fundedKey)
		teleporterInfo.DeployTeleporterRegistry(subnet, fundedKey)
	}

	// Convert the subnets to sovereign L1s
	for _, subnet := range localNetworkInstance.GetL1Infos() {
		localNetworkInstance.ConvertSubnet(
			networkStartCtx,
			subnet,
			teleporterTestUtils.PoAValidatorManager,
			[]uint64{units.Schmeckle, units.Schmeckle, units.Schmeckle, units.Schmeckle},
			[]uint64{defaultBalance, defaultBalance, defaultBalance, minimumL1ValidatorBalance - 1},
			fundedKey,
			false,
		)
	}

	// Restart the network to attempt to refresh TLS connections
	networkRestartCtx, cancel := context.WithTimeout(ctx, time.Duration(60*len(localNetworkInstance.Nodes))*time.Second)
	defer cancel()

	err = localNetworkInstance.Restart(networkRestartCtx)
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

	log.Info("Set up ginkgo before suite")

	ginkgo.AddReportEntry(
		"network directory with node logs & configs; useful in the case of failures",
		localNetworkInstance.Dir(),
		ginkgo.ReportEntryVisibilityFailureOrVerbose,
	)
})

func cleanup() {
	cancelFn()
	if decider != nil {
		decider = nil
	}
	if localNetworkInstance != nil {
		localNetworkInstance.TearDownNetwork()
		localNetworkInstance = nil
	}
}

var _ = ginkgo.AfterSuite(cleanup)

var _ = ginkgo.Describe("[ICM Relayer Integration Tests", func() {
	ginkgo.It("Basic Relay", func(ctx context.Context) {
		BasicRelay(ctx, log, localNetworkInstance, teleporterInfo)
	})
	ginkgo.It("Manually Provided Message", func(ctx context.Context) {
		ManualMessage(ctx, log, localNetworkInstance, teleporterInfo)
	})
	ginkgo.It("Shared Database", func(ctx context.Context) {
		SharedDatabaseAccess(ctx, log, localNetworkInstance, teleporterInfo)
	})
	ginkgo.It("Allowed Addresses", func(ctx context.Context) {
		AllowedAddresses(ctx, log, localNetworkInstance, teleporterInfo)
	})
	ginkgo.It("Batch Message", func(ctx context.Context) {
		BatchRelay(ctx, log, localNetworkInstance, teleporterInfo)
	})
	ginkgo.It("Relay Message API", func(ctx context.Context) {
		RelayMessageAPI(ctx, log, localNetworkInstance, teleporterInfo)
	})
	ginkgo.It("Warp API", func(ctx context.Context) {
		WarpAPIRelay(ctx, log, localNetworkInstance, teleporterInfo)
	})
	ginkgo.It("Signature Aggregator", func(ctx context.Context) {
		SignatureAggregatorAPI(ctx, log, localNetworkInstance, teleporterInfo)
	})
	ginkgo.It("Signature Aggregator Epoch Validators", func(ctx context.Context) {
		SignatureAggregatorEpochAPI(ctx, log, localNetworkInstance, teleporterInfo)
	})
	ginkgo.It("Validators Only Network", func(ctx context.Context) {
		ValidatorsOnlyNetwork(ctx, log, localNetworkInstance, teleporterInfo)
	})
})
