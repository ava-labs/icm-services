// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// oracle-devnet starts a local Avalanche L1 with the full Solana oracle pipeline
// and keeps it running until interrupted. It is the long-lived counterpart to
// the E2E test suite: same network setup, same contract deployment, but the
// network stays up so you can interact with it from a UI or CLI.
//
// Usage:
//
//	oracle-devnet \
//	  --avalanchego-path /abs/path/to/avalanchego             \
//	  --solana-rpc-url   https://api.devnet.solana.com        \
//	  --solana-keypair   ~/.config/solana/demo.json           \
//	  --oracle-token-sol /path/to/oracle-demo/contracts/OracleToken.sol \
//	  --observer         /path/to/build/solana-observer       \
//
// On startup the script writes oracle-demo/local-config.json with all
// deployed addresses and RPC URL for the UI to consume.
package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"math/big"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	avago_tests "github.com/ava-labs/avalanchego/tests"
	"github.com/ava-labs/avalanchego/tests/fixture/e2e"
	"github.com/ava-labs/avalanchego/utils/logging"
	oracleadapter "github.com/ava-labs/icm-services/abi-bindings/go/teleporterV2/OracleAdapter"
	"github.com/ava-labs/icm-services/icm-contracts/tests/network"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/libevm/ethclient"
	"github.com/onsi/gomega"
	"go.uber.org/zap"
)

const (
	solanarpcSidecarPort = 9901
	// memoProgram is the SPL Memo program. The sidecar verifies Memo instructions
	// (which carry the hex-encoded ABI payload), so the sidecar and OracleAdapter
	// allowlist must include this program.
	memoProgram = "MemoSq4gqABAXKb96qnH8TysNcWxMyWCqXgDLGmfcHr"
	// escrowProgram is the Anchor escrow program on Solana devnet. The observer
	// subscribes to it so only our deposit transactions are relayed, not arbitrary
	// devnet Memo traffic.
	escrowProgram = "EZe1KdrZWWqS4PexH9pQeEn4iTeE3TzHe8fs5aLW2nqb"
	l1EVMChainID  = 12346
)

// die logs at Fatal level and exits — avalanchego's Logger.Fatal does not exit by itself.
func die(log logging.Logger, msg string, fields ...zap.Field) {
	log.Fatal(msg, fields...)
	os.Exit(1)
}

func main() {
	// e2e.RegisterFlags registers --avalanchego-path into the global flag set.
	e2eFlags := e2e.RegisterFlags()

	solanaRPCURL := flag.String("solana-rpc-url", "https://api.devnet.solana.com", "Solana JSON-RPC URL")
	solanaKeypair := flag.String("solana-keypair", "", "path to funded Solana devnet keypair (enables observer)")
	oracleTokenSol := flag.String("oracle-token-sol", "", "path to OracleToken.sol (compiled and deployed if set)")
	observerBin := flag.String("observer", "", "path to solana-observer binary (built if empty and --solana-keypair set)")
	// avalanchego-path is already registered by e2e.RegisterFlags; we read it
	// directly from the env for the sidecar build step.
	avalancheGoPathEnv := flag.String(
		"avalanchego-path-env",
		os.Getenv("AVALANCHEGO_PATH"),
		"path to avalanchego binary (defaults to $AVALANCHEGO_PATH)",
	)
	flag.Parse()

	log := logging.NewLogger(
		"oracle-local",
		logging.NewWrappedCore(logging.Info, os.Stdout, logging.JSON.ConsoleEncoder()),
	)

	// Allow Gomega assertions outside of Ginkgo by routing failures to Fatal.
	gomega.RegisterFailHandler(func(message string, _ ...int) {
		log.Error("assertion failed", zap.String("message", message))
		os.Exit(1)
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repoRoot, err := utils.GetRepoRoot()
	if err != nil {
		die(log, "could not find repo root", zap.Error(err))
	}

	// ── build ICM executables ────────────────────────────────────────────────
	log.Info("Building ICM service executables")
	utils.BuildAllExecutables(ctx, log)

	// ── build + start solanarpc sidecar ─────────────────────────────────────
	if *avalancheGoPathEnv == "" {
		die(log, "AVALANCHEGO_PATH env var or --avalanchego-path-env flag is required")
	}
	avalancheGoRoot := filepath.Dir(filepath.Dir(*avalancheGoPathEnv))
	solanarpcBin := filepath.Join(repoRoot, "build/solanarpc-sidecar")

	log.Info("Building solanarpc sidecar", zap.String("root", avalancheGoRoot))
	if out, err := runIn(avalancheGoRoot, "go", "build", "-o", solanarpcBin, "./sidecar/"); err != nil {
		die(log, "solanarpc-sidecar build failed", zap.String("output", out), zap.Error(err))
	}

	solanarpcConfigPath := filepath.Join(repoRoot, "build/solanarpc-config.json")
	sidecarCfgJSON := fmt.Sprintf(
		`{"verifiers":{"solana":{"rpc_url":%q,"allowed_programs":[%q]}}}`,
		*solanaRPCURL, memoProgram,
	)
	if err := os.WriteFile(solanarpcConfigPath, []byte(sidecarCfgJSON), 0o600); err != nil {
		die(log, "write solanarpc config", zap.Error(err))
	}

	solanarpcEndpoint := fmt.Sprintf("127.0.0.1:%d", solanarpcSidecarPort)
	solanarpcSidecar := startProcess(log, "solanarpc-sidecar",
		exec.Command(solanarpcBin,
			"--addr", fmt.Sprintf(":%d", solanarpcSidecarPort),
			"--config-path", solanarpcConfigPath,
		),
	)
	waitForTCP(ctx, log, solanarpcEndpoint, 15*time.Second)
	log.Info("solanarpc sidecar ready", zap.String("addr", solanarpcEndpoint))

	// ── optionally build solana-observer ─────────────────────────────────────
	if *observerBin == "" && *solanaKeypair != "" {
		*observerBin = filepath.Join(repoRoot, "build/solana-observer")
		log.Info("Building solana-observer")
		if out, err := runIn(repoRoot, "go", "build", "-o", *observerBin, "./solana-observer/"); err != nil {
			die(log, "solana-observer build failed", zap.String("output", out), zap.Error(err))
		}
	}

	// ── start Avalanche network ──────────────────────────────────────────────
	chainConfig := utils.DefaultChainConfig()
	chainConfig["oracle"] = map[string]any{
		"endpoint":        solanarpcEndpoint,
		"allowed-sources": []string{"solana"},
	}

	l1Specs := []network.L1Spec{{
		Name:        "solanarpc",
		EVMChainID:  l1EVMChainID,
		NodeCount:   2,
		ChainConfig: chainConfig,
	}}

	simpleTc := avago_tests.NewTestContext(log)
	log.Info("Starting local Avalanche network (takes ~3 min)…")
	networkCtx, networkCancel := context.WithTimeout(ctx, 4*time.Minute)
	defer networkCancel()
	localNet := network.NewLocalAvalancheNetwork(
		networkCtx,
		"oracle-local",
		filepath.Join(repoRoot, "tests/utils/warp-genesis-template.json"),
		l1Specs,
		len(l1Specs)+1,
		0,
		e2eFlags,
		simpleTc,
	)
	log.Info("Network started", zap.String("dir", localNet.Dir()))

	l1Info := localNet.GetL1Infos()[0]
	_, fundedKey := localNet.GetFundedAccountInfo()
	deployOpts, err := bind.NewKeyedTransactorWithChainID(fundedKey, l1Info.EVMChainID)
	if err != nil {
		die(log, "new keyed transactor", zap.Error(err))
	}

	// ── deploy OracleAdapter ─────────────────────────────────────────────────
	log.Info("Deploying OracleAdapter")
	adapterAddr, adapterTx, adapterContract, err := oracleadapter.DeployOracleAdapter(
		deployOpts, l1Info.EthClient, deployOpts.From,
	)
	if err != nil {
		die(log, "deploy OracleAdapter", zap.Error(err))
	}
	utils.WaitForTransactionSuccess(ctx, l1Info.EthClient, adapterTx.Hash())
	log.Info("OracleAdapter deployed", zap.Stringer("address", adapterAddr))

	// ── deploy TeleporterMessengerV2 ─────────────────────────────────────────
	log.Info("Deploying TeleporterMessengerV2")
	teleporterAddr := utils.DeployTeleporterV2(ctx, &l1Info, adapterAddr, fundedKey)
	log.Info("TeleporterMessengerV2 deployed", zap.Stringer("address", teleporterAddr))

	// ── allowlist Memo program on OracleAdapter ──────────────────────────────
	// The sidecar verifies Memo instructions, so OracleAdapter must allow
	// messages sourced from the Memo program.
	allowTx, err := adapterContract.SetAllowedSource(deployOpts, "solana", memoProgram, true)
	if err != nil {
		die(log, "SetAllowedSource", zap.Error(err))
	}
	utils.WaitForTransactionSuccess(ctx, l1Info.EthClient, allowTx.Hash())
	log.Info("Memo program allowlisted on OracleAdapter")

	// ── compile + deploy OracleToken (optional) ──────────────────────────────
	var oracleTokenAddr common.Address
	if *oracleTokenSol != "" {
		log.Info("Compiling OracleToken", zap.String("path", *oracleTokenSol))
		bytecode, err := compileSol(*oracleTokenSol, "OracleToken")
		if err != nil {
			die(log, "OracleToken compile", zap.Error(err))
		}
		log.Info("Deploying OracleToken")
		oracleTokenAddr, err = deployWithAddressArg(
			ctx, l1Info.EthClient, fundedKey, l1Info.EVMChainID, bytecode, teleporterAddr,
		)
		if err != nil {
			die(log, "OracleToken deploy", zap.Error(err))
		}
		log.Info("OracleToken deployed", zap.Stringer("address", oracleTokenAddr))
	}

	// ── start signature aggregator ───────────────────────────────────────────
	log.Info("Starting signature aggregator")
	sigAggCfg := utils.CreateDefaultSignatureAggregatorConfig(log, []testinfo.L1TestInfo{l1Info})
	sigAggCfgPath := utils.WriteSignatureAggregatorConfig(log, sigAggCfg, "sig-agg-local-config.json")
	sigAggCancel, sigAggReady := utils.RunSignatureAggregatorExecutable(ctx, log, sigAggCfgPath, sigAggCfg)
	readyCtx, readyCancel := context.WithTimeout(ctx, 30*time.Second)
	utils.WaitForChannelClose(readyCtx, sigAggReady)
	readyCancel()
	log.Info("Signature aggregator ready")

	// ── start solana-observer (optional) ─────────────────────────────────────
	var observerProc *exec.Cmd
	if *observerBin != "" && *solanaKeypair != "" && oracleTokenAddr != (common.Address{}) {
		nonceFile := filepath.Join(repoRoot, "build/observer-local-nonces.json")
		_ = os.Remove(nonceFile)

		observerCfg := buildObserverConfig(
			l1Info, localNet.GetNetworkID(),
			teleporterAddr, oracleTokenAddr,
			fundedKey, sigAggCfg.APIPort,
			solanarpcConfigPath, nonceFile,
		)
		observerCfgPath := filepath.Join(repoRoot, "build/observer-local-config.json")
		if err := os.WriteFile(observerCfgPath, observerCfg, 0o600); err != nil {
			die(log, "write observer config", zap.Error(err))
		}

		observerProc = exec.CommandContext(ctx, *observerBin, "--config-path", observerCfgPath)
		observerReady := make(chan struct{})
		tee := &logTee{prefix: "[observer] "}
		tee.notifyOn("subscribed to Solana logs", observerReady)
		observerProc.Stdout = tee
		observerProc.Stderr = tee
		if err := observerProc.Start(); err != nil {
			die(log, "start observer", zap.Error(err))
		}
		select {
		case <-observerReady:
			log.Info("Observer subscribed to Solana logs")
		case <-time.After(20 * time.Second):
			log.Warn("Observer did not print ready message within 20s — continuing anyway")
		}
	} else if *solanaKeypair != "" && oracleTokenAddr == (common.Address{}) {
		log.Warn("Observer not started: pass --oracle-token-sol to deploy OracleToken first")
	}

	// ── write local-config.json ─────────────────────────────────────────────
	rpc := buildRPC(l1Info)
	devnetCfg := map[string]any{
		"avalanche_rpc":          rpc,
		"chain_id":               l1EVMChainID,
		"teleporter_address":     teleporterAddr.Hex(),
		"oracle_adapter_address": adapterAddr.Hex(),
	}
	if oracleTokenAddr != (common.Address{}) {
		devnetCfg["oracle_token_address"] = oracleTokenAddr.Hex()
	}
	devnetCfgBytes, _ := json.MarshalIndent(devnetCfg, "", "  ")
	outPath := filepath.Join(repoRoot, "../oracle-demo/local-config.json")
	if err := os.WriteFile(outPath, devnetCfgBytes, 0o600); err != nil {
		log.Warn("could not write local-config.json", zap.Error(err), zap.String("path", outPath))
	} else {
		log.Info("Wrote local-config.json", zap.String("path", outPath))
	}

	// ── print summary ─────────────────────────────────────────────────────────
	fmt.Println("\n╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║                   oracle-local ready                     ║")
	fmt.Println("╠══════════════════════════════════════════════════════════╣")
	fmt.Printf("║  L1 RPC:           %s\n", rpc)
	fmt.Printf("║  Chain ID:         %d\n", l1EVMChainID)
	fmt.Printf("║  OracleAdapter:    %s\n", adapterAddr.Hex())
	fmt.Printf("║  TeleporterV2:     %s\n", teleporterAddr.Hex())
	if oracleTokenAddr != (common.Address{}) {
		fmt.Printf("║  OracleToken:      %s\n", oracleTokenAddr.Hex())
	} else {
		fmt.Println("║  OracleToken:      (not deployed — pass --oracle-token-sol)")
	}
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Print("\nPress Ctrl-C to shut down.\n\n")

	// ── block until signal ────────────────────────────────────────────────────
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	fmt.Println("\nShutting down…")
	sigAggCancel()
	if observerProc != nil && observerProc.Process != nil {
		_ = observerProc.Process.Signal(os.Interrupt)
		done := make(chan struct{})
		go func() { _ = observerProc.Wait(); close(done) }()
		select {
		case <-done:
		case <-time.After(5 * time.Second):
			_ = observerProc.Process.Kill()
		}
	}
	if solanarpcSidecar != nil && solanarpcSidecar.Process != nil {
		_ = solanarpcSidecar.Process.Kill()
	}
	localNet.TearDownNetwork()
	log.Info("Done")
}

// ── helpers ───────────────────────────────────────────────────────────────────

func buildRPC(l1 testinfo.L1TestInfo) string {
	host, port, err := utils.GetURIHostAndPort(l1.NodeURIs[0])
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("http://%s:%d/ext/bc/%s/rpc", host, port, l1.BlockchainID.String())
}

// compileSol shells out to solc and returns raw deploy bytecode for contractName.
func compileSol(solPath, contractName string) ([]byte, error) {
	out, err := exec.Command("solc", "--optimize", "--combined-json", "bin", solPath).Output()
	if err != nil {
		return nil, fmt.Errorf("solc: %w", err)
	}
	var result struct {
		Contracts map[string]struct {
			Bin string `json:"bin"`
		} `json:"contracts"`
	}
	if err := json.Unmarshal(out, &result); err != nil {
		return nil, fmt.Errorf("parse solc output: %w", err)
	}
	for key, c := range result.Contracts {
		if strings.HasSuffix(key, ":"+contractName) {
			return hex.DecodeString(c.Bin)
		}
	}
	return nil, fmt.Errorf("%s not found in solc output", contractName)
}

// deployWithAddressArg deploys bytecode with a single address constructor arg.
// ABI encoding for (address): 12 zero bytes + 20 address bytes = 32 bytes.
func deployWithAddressArg(
	ctx context.Context,
	client *ethclient.Client,
	key *ecdsa.PrivateKey,
	chainID *big.Int,
	bytecode []byte,
	arg common.Address,
) (common.Address, error) {
	encoded := make([]byte, 32)
	copy(encoded[12:], arg.Bytes())
	data := append(bytecode, encoded...)

	fromAddr := crypto.PubkeyToAddress(key.PublicKey)
	nonce, err := client.PendingNonceAt(ctx, fromAddr)
	if err != nil {
		return common.Address{}, err
	}
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return common.Address{}, err
	}

	tx := types.NewContractCreation(nonce, big.NewInt(0), 3_000_000, gasPrice, data)
	signed, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), key)
	if err != nil {
		return common.Address{}, err
	}
	if err := client.SendTransaction(ctx, signed); err != nil {
		return common.Address{}, err
	}
	receipt := utils.WaitForTransactionSuccess(ctx, client, signed.Hash())
	return receipt.ContractAddress, nil
}

// buildObserverConfig produces the observer-config.json body.
func buildObserverConfig(
	l1 testinfo.L1TestInfo,
	networkID uint32,
	teleporter, dest common.Address,
	fundedKey *ecdsa.PrivateKey,
	sigAggPort uint16,
	sidecarConfigPath, nonceFile string,
) []byte {
	cfg := map[string]any{
		"solana": map[string]any{
			"ws_url":     "wss://api.devnet.solana.com",
			"rpc_url":    "https://api.devnet.solana.com",
			"commitment": "finalized",
		},
		"l1": map[string]any{
			"rpc_url":                  buildRPC(l1),
			"chain_id":                 l1.EVMChainID.Uint64(),
			"blockchain_id":            l1.BlockchainID.String(),
			"network_id":               networkID,
			"teleporter_address":       teleporter.Hex(),
			"dest_contract":            dest.Hex(),
			"subnet_id":                l1.SubnetID.String(),
			"delivery_private_key_hex": hex.EncodeToString(crypto.FromECDSA(fundedKey)),
		},
		"aggregator_url":      fmt.Sprintf("http://127.0.0.1:%d", sigAggPort),
		"sidecar_config_path": sidecarConfigPath,
		"nonce_file":          nonceFile,
		// Subscribe to the escrow program so only our deposit transactions
		// are relayed. The sidecar still verifies the inner Memo instruction.
		"subscription_programs": []string{escrowProgram},
	}
	b, _ := json.MarshalIndent(cfg, "", "  ")
	return b
}

func startProcess(log logging.Logger, name string, cmd *exec.Cmd) *exec.Cmd {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		die(log, "start "+name, zap.Error(err))
	}
	go func() {
		if err := cmd.Wait(); err != nil {
			log.Error(name+" exited", zap.Error(err))
		}
	}()
	return cmd
}

func waitForTCP(ctx context.Context, log logging.Logger, addr string, timeout time.Duration) {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		conn, err := net.DialTimeout("tcp", addr, 200*time.Millisecond)
		if err == nil {
			conn.Close()
			return
		}
		select {
		case <-ctx.Done():
			die(log, "context cancelled waiting for "+addr, zap.Error(ctx.Err()))
		case <-time.After(50 * time.Millisecond):
		}
	}
	die(log, "timed out waiting for "+addr)
}

func runIn(dir, name string, args ...string) (string, error) {
	c := exec.Command(name, args...)
	c.Dir = dir
	out, err := c.CombinedOutput()
	return string(out), err
}

// logTee mirrors subprocess output to stdout and fires a channel on a substring match.
type logTee struct {
	prefix string
	match  string
	fired  bool
	notify chan struct{}
	buf    bytes.Buffer
}

func (t *logTee) notifyOn(substring string, ch chan struct{}) {
	t.match = substring
	t.notify = ch
}

func (t *logTee) Write(p []byte) (int, error) {
	fmt.Fprint(os.Stdout, t.prefix, string(p))
	if t.notify == nil || t.fired {
		return len(p), nil
	}
	t.buf.Write(p)
	sc := bufio.NewScanner(&t.buf)
	for sc.Scan() {
		if strings.Contains(sc.Text(), t.match) {
			t.fired = true
			close(t.notify)
			t.buf.Reset()
			return len(p), nil
		}
	}
	return len(p), nil
}
