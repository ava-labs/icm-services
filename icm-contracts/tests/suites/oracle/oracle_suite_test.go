// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package oracle_test is the E2E test suite for oracle attestation.
//
// Run via the e2e test script, which handles building executables and sets the
// working directory to the repo root (required by the shared test utilities):
//
//	AVALANCHEGO_PATH=/abs/path/to/avalanchego ./scripts/e2e_test.sh --components oracle
//
// Requirements to run:
//   - AVALANCHEGO_PATH pointing to a binary built with oracle sidecar support
//     (merged upstream via the oracle handler wiring); e2e_test.sh skips the
//     release avalanchego install when it is set
//   - RUN_E2E=true environment variable (set automatically by e2e_test.sh)
//
// Optional — enables Solana verification via the solanarpc sidecar:
//   - SOLANA_RPC_URL set to a Solana JSON-RPC endpoint (e.g. https://api.devnet.solana.com)
//     When set, the suite builds and runs the solanarpc sidecar (from the avalanchego
//     source tree) in addition to the mock, and the test flow fetches a live Memo Program
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
	// solanarpcSidecarPort is the port the solanarpc gRPC sidecar listens on when
	// SOLANA_RPC_URL is set.
	solanarpcSidecarPort = 9901
)

var (
	log                  logging.Logger
	localNetworkInstance *network.LocalAvalancheNetwork
	oracleSidecar        *exec.Cmd // mock gRPC sidecar, always running
	solanarpcSidecar     *exec.Cmd // solanarpc sidecar, non-nil only when SOLANA_RPC_URL set
	e2eFlags             *e2e.FlagVars
	solanaRPCURL         string // non-empty when SOLANA_RPC_URL is set
	solanarpcConfigPath  string // path to the solanarpc sidecar's config file; also read by the validator
	solanaKeypairPath    string // non-empty when SOLANA_KEYPAIR is set (funded devnet keypair)
	runObserverDemo      bool   // true when RUN_OBSERVER_DEMO is set
	observerBinaryPath   string // path to the built solana-observer binary
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
	solanaKeypairPath = os.Getenv("SOLANA_KEYPAIR")
	runObserverDemo = os.Getenv("RUN_OBSERVER_DEMO") != ""

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
	waitForTCP(ctx, mockEndpoint, 10*time.Second)
	log.Info("Mock oracle sidecar is ready", zap.String("addr", mockEndpoint))

	// When SOLANA_RPC_URL is set, also build and start the solanarpc sidecar on port 9901.
	if solanaRPCURL != "" {
		avalancheGoPath := os.Getenv("AVALANCHEGO_PATH")
		Expect(filepath.IsAbs(avalancheGoPath)).Should(
			BeTrue(),
			"AVALANCHEGO_PATH must be an absolute path when SOLANA_RPC_URL is set (got %q); "+
				"go test sets the working directory to the package directory, making relative paths resolve incorrectly",
			avalancheGoPath,
		)
		avalancheGoRoot := filepath.Dir(filepath.Dir(avalancheGoPath))
		solanarpcBin := filepath.Join(repoRoot, "build/solanarpc-sidecar")
		log.Info("Building solanarpc sidecar", zap.String("avalancheGoRoot", avalancheGoRoot))
		buildCmd := exec.Command("go", "build", "-o", solanarpcBin, "./sidecar/")
		buildCmd.Dir = avalancheGoRoot
		buildOut, buildErr := buildCmd.CombinedOutput()
		log.Info(string(buildOut))
		Expect(buildErr).Should(BeNil())

		solanarpcConfigPath = filepath.Join(repoRoot, "build/solanarpc-config.json")
		// Include allowed_programs listing the Memo program so the observer,
		// which reads this same file for its subscription filter, has
		// something to watch. Harmless when the observer isn't running.
		configJSON := fmt.Sprintf(
			`{"verifiers": {"solana": {"rpc_url": %q, "allowed_programs": [%q]}}}`,
			solanaRPCURL,
			"MemoSq4gqABAXKb96qnH8TysNcWxMyWCqXgDLGmfcHr",
		)
		Expect(os.WriteFile(solanarpcConfigPath, []byte(configJSON), 0o600)).Should(BeNil())

		solanarpcEndpoint := fmt.Sprintf("127.0.0.1:%d", solanarpcSidecarPort)
		log.Info("Starting solanarpc sidecar",
			zap.String("endpoint", solanarpcEndpoint),
			zap.String("solanaRPC", solanaRPCURL),
		)
		solanarpcSidecar = exec.Command(solanarpcBin,
			"--addr", fmt.Sprintf(":%d", solanarpcSidecarPort),
			"--config-path", solanarpcConfigPath,
		)
		solanarpcSidecar.Stdout = os.Stdout
		solanarpcSidecar.Stderr = os.Stderr
		Expect(solanarpcSidecar.Start()).Should(BeNil())
		go func() {
			if waitErr := solanarpcSidecar.Wait(); waitErr != nil {
				log.Error("solanarpc-sidecar exited abnormally", zap.Error(waitErr))
			}
		}()
		waitForTCP(ctx, solanarpcEndpoint, 10*time.Second)
		log.Info("solanarpc sidecar is ready", zap.String("addr", solanarpcEndpoint))
	}

	// Build the solana-observer binary when the auto-relay demo is requested.
	// The spec that consumes it is skipped unless RUN_OBSERVER_DEMO,
	// SOLANA_RPC_URL, and SOLANA_KEYPAIR are all set.
	if runObserverDemo {
		observerBinaryPath = filepath.Join(repoRoot, "build/solana-observer")
		log.Info("Building solana-observer", zap.String("path", observerBinaryPath))
		buildCmd := exec.Command("go", "build", "-o", observerBinaryPath, "./solana-observer/")
		buildCmd.Dir = repoRoot
		buildOut, buildErr := buildCmd.CombinedOutput()
		log.Info(string(buildOut))
		Expect(buildErr).Should(BeNil())
	}

	// Build chain configs pointing each L1 at its sidecar. As of the upstream
	// oracle handler wiring, validators declare their allowed source types
	// directly via oracle.allowed-sources instead of reading the sidecar's
	// config file.
	mockChainConfig := utils.DefaultChainConfig()
	mockChainConfig["oracle"] = map[string]any{
		"endpoint":        mockEndpoint,
		"allowed-sources": []string{"solana"},
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
		solanarpcChainConfig := utils.DefaultChainConfig()
		solanarpcChainConfig["oracle"] = map[string]any{
			"endpoint":        fmt.Sprintf("127.0.0.1:%d", solanarpcSidecarPort),
			"allowed-sources": []string{"solana"},
		}
		l1Specs = append(l1Specs, network.L1Spec{
			Name:        "solanarpc",
			EVMChainID:  12346,
			NodeCount:   2,
			ChainConfig: solanarpcChainConfig,
		})
	}

	networkStartCtx, networkStartCancel := context.WithTimeout(ctx, 240*time.Second)
	defer networkStartCancel()

	localNetworkInstance = network.NewLocalAvalancheNetwork(
		networkStartCtx,
		"oracle-attestation-e2e",
		filepath.Join(repoRoot, "tests/utils/warp-genesis-template.json"),
		l1Specs,
		len(l1Specs)+1, // numPrimaryNetworkValidators; +1 keeps one spare beyond the L1 count
		0,              // extraNodeCount
		e2eFlags,
	)

	ginkgo.AddReportEntry(
		"network directory with node logs & configs",
		localNetworkInstance.Dir(),
		ginkgo.ReportEntryVisibilityFailureOrVerbose,
	)
})

func cleanup() {
	for _, cmd := range []*exec.Cmd{oracleSidecar, solanarpcSidecar} {
		if cmd != nil && cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
	}
	oracleSidecar = nil
	solanarpcSidecar = nil
	if localNetworkInstance != nil {
		localNetworkInstance.TearDownNetwork()
		localNetworkInstance = nil
	}
}

// waitForTCP polls addr until a TCP connection succeeds, the timeout elapses,
// or the parent context is cancelled.
func waitForTCP(ctx context.Context, addr string, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			Expect(ctx.Err()).ShouldNot(HaveOccurred(),
				"process did not bind %s within %s", addr, timeout)
		case <-ticker.C:
			conn, err := net.DialTimeout("tcp", addr, 200*time.Millisecond)
			if err == nil {
				conn.Close()
				return
			}
		}
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

	ginkgo.It("Oracle Attestation (solanarpc sidecar)",
		ginkgo.Label(oracleLabel),
		func(ctx context.Context) {
			if solanaRPCURL == "" {
				ginkgo.Skip("SOLANA_RPC_URL not set; skipping solanarpc sidecar test")
			}
			l1Infos := localNetworkInstance.GetL1Infos()
			oracleFlows.OracleAttestation(ctx, log, localNetworkInstance, l1Infos[1], solanaRPCURL)
		})

	ginkgo.It("Auto-relay demo (solana-observer)",
		ginkgo.Label(oracleLabel),
		func(ctx context.Context) {
			if !runObserverDemo {
				ginkgo.Skip("RUN_OBSERVER_DEMO not set; skipping observer demo")
			}
			if solanaRPCURL == "" || solanaKeypairPath == "" {
				ginkgo.Skip("observer demo requires SOLANA_RPC_URL and SOLANA_KEYPAIR")
			}
			l1Infos := localNetworkInstance.GetL1Infos()
			oracleFlows.ObserverAttestation(
				ctx, log, localNetworkInstance, l1Infos[1],
				observerBinaryPath,
				solanarpcConfigPath,
				solanaKeypairPath,
			)
		})
})
