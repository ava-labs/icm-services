// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package oracle_test is the E2E test suite for oracle attestation.
//
// Requirements to run:
//   - RUN_E2E=true environment variable
//   - AVALANCHEGO_PATH pointing to a binary built from the
//     boraplusplus/sidecar-verifier branch (oracle handler ID 4 support)
//   - build/oracle-sidecar binary (built by scripts/build_oracle_sidecar.sh)
package oracle_test

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
	oracleFlows "github.com/ava-labs/icm-services/icm-contracts/tests/flows/oracle"
	"github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

const (
	warpGenesisTemplateFile = "./tests/utils/warp-genesis-template.json"
	oracleLabel             = "OracleAttestation"
	// oracleSidecarPort is the port the mock sidecar listens on. Validators'
	// oracle.endpoint chain config must reference this port.
	oracleSidecarPort = 9900
)

var (
	log                  logging.Logger
	localNetworkInstance *network.LocalAvalancheNetwork
	oracleSidecar        *exec.Cmd
	e2eFlags             *e2e.FlagVars
)

func TestMain(m *testing.M) {
	e2eFlags = e2e.RegisterFlags()
	flag.Parse()
	os.Exit(m.Run())
}

func TestOracle(t *testing.T) {
	signalChan := make(chan os.Signal, 2)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-signalChan
		fmt.Printf("Caught signal %s: shutting down...\n", sig)
		cleanup()
		os.Exit(1)
	}()

	if os.Getenv("RUN_E2E") == "" {
		t.Skip("Environment variable RUN_E2E not set; skipping oracle E2E tests")
	}

	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Oracle attestation e2e tests")
}

var _ = ginkgo.BeforeSuite(func(ctx context.Context) {
	log = logging.NewLogger(
		"oracle-e2e",
		logging.NewWrappedCore(
			logging.Info,
			os.Stdout,
			logging.JSON.ConsoleEncoder(),
		),
	)

	log.Info("Building all ICM service executables (includes oracle-sidecar)")
	utils.BuildAllExecutables(ctx, log)

	// Start the mock oracle sidecar before the network so validators can reach
	// it when they boot with oracle.endpoint in their chain config.
	sidecarEndpoint := fmt.Sprintf("http://127.0.0.1:%d", oracleSidecarPort)
	log.Info("Starting oracle mock sidecar", zap.String("endpoint", sidecarEndpoint))
	oracleSidecar = exec.CommandContext(ctx, "./build/oracle-sidecar",
		"--port", fmt.Sprintf("%d", oracleSidecarPort))
	oracleSidecar.Stdout = os.Stdout
	oracleSidecar.Stderr = os.Stderr
	err := oracleSidecar.Start()
	Expect(err).Should(BeNil())
	go func() {
		waitErr := oracleSidecar.Wait()
		if !errors.Is(ctx.Err(), context.Canceled) {
			log.Error("oracle-sidecar exited abnormally", zap.Error(waitErr))
		}
	}()
	// Give the sidecar a moment to bind its port before the network starts.
	time.Sleep(500 * time.Millisecond)

	// Build chain config for the oracle L1: enable warp and point oracle
	// handler at the mock sidecar. The empty allowed-sources slice for "solana"
	// means all source addresses on that source type are permitted.
	oracleChainConfig := utils.DefaultChainConfig()
	oracleChainConfig["oracle"] = map[string]any{
		"endpoint": sidecarEndpoint,
		"allowed-sources": map[string]any{
			"solana": []any{},
		},
	}

	networkStartCtx, networkStartCancel := context.WithTimeout(ctx, 240*time.Second)
	defer networkStartCancel()

	localNetworkInstance = network.NewLocalAvalancheNetwork(
		networkStartCtx,
		"oracle-attestation-e2e",
		warpGenesisTemplateFile,
		[]network.L1Spec{
			{
				Name:        "A",
				EVMChainID:  12345,
				NodeCount:   2,
				ChainConfig: oracleChainConfig,
			},
		},
		3, // numPrimaryNetworkValidators
		0, // extraNodeCount
		e2eFlags,
	)

	ginkgo.AddReportEntry(
		"network directory with node logs & configs",
		localNetworkInstance.Dir(),
		ginkgo.ReportEntryVisibilityFailureOrVerbose,
	)
})

func cleanup() {
	if oracleSidecar != nil {
		oracleSidecar = nil
	}
	if localNetworkInstance != nil {
		localNetworkInstance.TearDownNetwork()
		localNetworkInstance = nil
	}
}

var _ = ginkgo.AfterSuite(cleanup)

var _ = ginkgo.Describe("[Oracle Attestation E2E Tests]", func() {
	ginkgo.It("Oracle Attestation",
		ginkgo.Label(oracleLabel),
		func(ctx context.Context) {
			oracleFlows.OracleAttestation(ctx, log, localNetworkInstance)
		})
})
