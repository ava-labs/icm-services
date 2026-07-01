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
	// oracleSidecarPort is the port the mock gRPC sidecar listens on.
	oracleSidecarPort = 9900
	// realSidecarPort is the port the real solanarpc gRPC sidecar listens on when
	// SOLANA_RPC_URL is set. Validators for the real-Solana L1 point here.
	realSidecarPort = 9901
)

var (
	log                  logging.Logger
	localNetworkInstance *network.LocalAvalancheNetwork
	oracleSidecar        *exec.Cmd // mock gRPC sidecar, always running
	realSidecar          *exec.Cmd // real solanarpc sidecar, non-nil only when SOLANA_RPC_URL set
	e2eFlags             *e2e.FlagVars
	solanaRPCURL         string // non-empty when SOLANA_RPC_URL is set
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

	// Always start the mock gRPC sidecar on port 9900 (unconditional accept).
	mockEndpoint := fmt.Sprintf("127.0.0.1:%d", oracleSidecarPort)
	log.Info("Starting oracle mock sidecar", zap.String("endpoint", mockEndpoint))
	oracleSidecar = exec.Command(filepath.Join(repoRoot, "build/oracle-sidecar"),
		"--port", fmt.Sprintf("%d", oracleSidecarPort),
	)
	oracleSidecar.Stdout = os.Stdout
	oracleSidecar.Stderr = os.Stderr
	Expect(oracleSidecar.Start()).Should(BeNil())
	go func() {
		if waitErr := oracleSidecar.Wait(); waitErr != nil {
			log.Error("oracle-sidecar exited abnormally", zap.Error(waitErr))
		}
	}()
	waitForTCP(mockEndpoint, 10*time.Second)
	log.Info("Mock oracle sidecar is ready", zap.String("addr", mockEndpoint))

	// When SOLANA_RPC_URL is set, also build and start the real solanarpc sidecar on port 9901.
	if solanaRPCURL != "" {
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

		realEndpoint := fmt.Sprintf("127.0.0.1:%d", realSidecarPort)
		log.Info("Starting real solanarpc sidecar",
			zap.String("endpoint", realEndpoint),
			zap.String("solanaRPC", solanaRPCURL),
		)
		realSidecar = exec.Command(solanarpcBin,
			"--addr", fmt.Sprintf(":%d", realSidecarPort),
			"--verifier-type", "solanarpc",
			"--config-path", configPath,
		)
		realSidecar.Stdout = os.Stdout
		realSidecar.Stderr = os.Stderr
		Expect(realSidecar.Start()).Should(BeNil())
		go func() {
			if waitErr := realSidecar.Wait(); waitErr != nil {
				log.Error("solanarpc-sidecar exited abnormally", zap.Error(waitErr))
			}
		}()
		waitForTCP(realEndpoint, 10*time.Second)
		log.Info("Real solanarpc sidecar is ready", zap.String("addr", realEndpoint))
	}

	// Build chain configs pointing each L1 at its sidecar.
	mockChainConfig := utils.DefaultChainConfig()
	mockChainConfig["oracle"] = map[string]any{
		"endpoint": mockEndpoint,
	}

	l1Specs := []network.L1Spec{
		{
			Name:        "mock",
			EVMChainID:  12345,
			NodeCount:   2,
			ChainConfig: mockChainConfig,
		},
	}

	if solanaRPCURL != "" {
		realChainConfig := utils.DefaultChainConfig()
		realChainConfig["oracle"] = map[string]any{
			"endpoint": fmt.Sprintf("127.0.0.1:%d", realSidecarPort),
		}
		l1Specs = append(l1Specs, network.L1Spec{
			Name:        "real",
			EVMChainID:  12346,
			NodeCount:   2,
			ChainConfig: realChainConfig,
		})
	}

	networkStartCtx, networkStartCancel := context.WithTimeout(ctx, 240*time.Second)
	defer networkStartCancel()

	localNetworkInstance = network.NewLocalAvalancheNetwork(
		networkStartCtx,
		"oracle-attestation-e2e",
		filepath.Join(repoRoot, "tests/utils/warp-genesis-template.json"),
		l1Specs,
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
	for _, cmd := range []*exec.Cmd{oracleSidecar, realSidecar} {
		if cmd != nil && cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
	}
	oracleSidecar = nil
	realSidecar = nil
	if localNetworkInstance != nil {
		localNetworkInstance.TearDownNetwork()
		localNetworkInstance = nil
	}
}

// waitForTCP polls addr until a TCP connection succeeds or deadline is exceeded.
func waitForTCP(addr string, timeout time.Duration) {
	deadline := time.Now().Add(timeout)
	for {
		conn, err := net.DialTimeout("tcp", addr, 200*time.Millisecond)
		if err == nil {
			conn.Close()
			return
		}
		if time.Now().After(deadline) {
			Expect(fmt.Errorf("process did not bind %s within %s: %w", addr, timeout, err)).Should(BeNil())
		}
		time.Sleep(50 * time.Millisecond)
	}
}

var _ = ginkgo.AfterSuite(cleanup)

var _ = ginkgo.Describe("[Oracle Attestation E2E Tests]", func() {
	ginkgo.It("Oracle Attestation (mock sidecar — unconditional accept)",
		ginkgo.Label(oracleLabel),
		func(ctx context.Context) {
			l1Infos := localNetworkInstance.GetL1Infos()
			oracleFlows.OracleAttestation(ctx, log, localNetworkInstance, l1Infos[0], "")
		})

	ginkgo.It("Oracle Attestation (real solanarpc sidecar)",
		ginkgo.Label(oracleLabel),
		func(ctx context.Context) {
			if solanaRPCURL == "" {
				ginkgo.Skip("SOLANA_RPC_URL not set; skipping real Solana sidecar test")
			}
			l1Infos := localNetworkInstance.GetL1Infos()
			oracleFlows.OracleAttestation(ctx, log, localNetworkInstance, l1Infos[1], solanaRPCURL)
		})
})
