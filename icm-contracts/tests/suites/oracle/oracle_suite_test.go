// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package oracle_test is the E2E test suite for oracle attestation.
//
// Requirements to run:
//   - RUN_E2E=true environment variable
//   - AVALANCHEGO_PATH pointing to a binary built from the
//     boraplusplus/sidecar-verifier branch (oracle handler ID 4 support)
//
// Optional — enables real Solana verification:
//   - SOLANA_RPC_URL set to a Solana JSON-RPC endpoint (e.g. https://api.devnet.solana.com)
//     When set, the suite builds and runs the real solanarpc sidecar (from the avalanchego
//     source tree) instead of the mock, and the test flow fetches a live Memo Program
//     transaction to use as the oracle payload.
package oracle_test

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
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
	oracleLabel = "OracleAttestation"
	// oracleSidecarPort is the port the mock sidecar listens on. Validators'
	// oracle.endpoint chain config must reference this port.
	oracleSidecarPort = 9900
)

var (
	log                  logging.Logger
	localNetworkInstance *network.LocalAvalancheNetwork
	oracleSidecar        *exec.Cmd
	e2eFlags             *e2e.FlagVars
	solanaRPCURL         string // non-empty when SOLANA_RPC_URL is set; selects real sidecar mode
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

	repoRoot, err := utils.GetRepoRoot()
	Expect(err).Should(BeNil())

	log.Info("Building all ICM service executables (includes oracle-sidecar)")
	utils.BuildAllExecutables(ctx, log)

	solanaRPCURL = os.Getenv("SOLANA_RPC_URL")
	sidecarEndpoint := fmt.Sprintf("127.0.0.1:%d", oracleSidecarPort)

	if solanaRPCURL != "" {
		// Real mode: build the solanarpc sidecar from the avalanchego source tree
		// (derived from AVALANCHEGO_PATH) and start it pointing at the given RPC endpoint.
		avalancheGoRoot := filepath.Dir(filepath.Dir(os.Getenv("AVALANCHEGO_PATH")))
		solanarpcBin := filepath.Join(repoRoot, "build/solanarpc-sidecar")
		log.Info("Building solanarpc sidecar", zap.String("avalancheGoRoot", avalancheGoRoot))
		buildCmd := exec.Command("go", "build", "-o", solanarpcBin, "./sidecar/")
		buildCmd.Dir = avalancheGoRoot
		buildOut, buildErr := buildCmd.CombinedOutput()
		log.Info(string(buildOut))
		Expect(buildErr).Should(BeNil())

		configPath := filepath.Join(repoRoot, "build/solanarpc-config.json")
		configJSON := fmt.Sprintf(`{"rpc_url": %q}`, solanaRPCURL)
		Expect(os.WriteFile(configPath, []byte(configJSON), 0o600)).Should(BeNil())

		log.Info("Starting real solanarpc sidecar",
			zap.String("endpoint", sidecarEndpoint),
			zap.String("solanaRPC", solanaRPCURL),
		)
		oracleSidecar = exec.Command(solanarpcBin,
			"--addr", fmt.Sprintf(":%d", oracleSidecarPort),
			"--verifier-type", "solanarpc",
			"--config-path", configPath,
		)
	} else {
		// Mock mode: start the unconditional accept sidecar (no Solana RPC needed).
		log.Info("Starting oracle mock sidecar", zap.String("endpoint", sidecarEndpoint))
		oracleSidecar = exec.Command(filepath.Join(repoRoot, "build/oracle-sidecar"),
			"--port", fmt.Sprintf("%d", oracleSidecarPort),
		)
	}
	oracleSidecar.Stdout = os.Stdout
	oracleSidecar.Stderr = os.Stderr
	err = oracleSidecar.Start()
	Expect(err).Should(BeNil())
	go func() {
		if waitErr := oracleSidecar.Wait(); waitErr != nil {
			log.Error("oracle-sidecar exited abnormally", zap.Error(waitErr))
		}
	}()
	// Wait until the sidecar actually accepts TCP connections (not just a fixed sleep).
	sidecarAddr := fmt.Sprintf("127.0.0.1:%d", oracleSidecarPort)
	readyDeadline := time.Now().Add(10 * time.Second)
	for {
		conn, dialErr := net.DialTimeout("tcp", sidecarAddr, 200*time.Millisecond)
		if dialErr == nil {
			conn.Close()
			break
		}
		if time.Now().After(readyDeadline) {
			Expect(fmt.Errorf("oracle-sidecar did not bind %s within 10s: %w", sidecarAddr, dialErr)).Should(BeNil())
		}
		time.Sleep(50 * time.Millisecond)
	}
	log.Info("Oracle sidecar is ready", zap.String("addr", sidecarAddr))

	// Build chain config for the oracle L1: enable warp and point the oracle
	// handler at the sidecar. Allowlist enforcement lives in the sidecar config,
	// not in the node config.
	oracleChainConfig := utils.DefaultChainConfig()
	oracleChainConfig["oracle"] = map[string]any{
		"endpoint": sidecarEndpoint,
	}

	networkStartCtx, networkStartCancel := context.WithTimeout(ctx, 240*time.Second)
	defer networkStartCancel()

	localNetworkInstance = network.NewLocalAvalancheNetwork(
		networkStartCtx,
		"oracle-attestation-e2e",
		filepath.Join(repoRoot, "tests/utils/warp-genesis-template.json"),
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
		if oracleSidecar.Process != nil {
			_ = oracleSidecar.Process.Kill()
		}
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
			oracleFlows.OracleAttestation(ctx, log, localNetworkInstance, solanaRPCURL)
		})
})
