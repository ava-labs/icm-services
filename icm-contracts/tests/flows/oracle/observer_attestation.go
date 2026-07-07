// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package oracle

import (
	"bufio"
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	mockoraclereceiver "github.com/ava-labs/icm-services/abi-bindings/go/mocks/MockOracleReceiver"
	oracleadapter "github.com/ava-labs/icm-services/abi-bindings/go/teleporterV2/OracleAdapter"
	"github.com/ava-labs/icm-services/icm-contracts/tests/network"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

// ObserverAttestation exercises the auto-relay pipeline end-to-end. It deploys
// the standard oracle contract set on the L1, starts the solana-observer as a
// subprocess, submits a fresh spl-memo transaction from the configured Solana
// keypair, and polls MockOracleReceiver until its last payload equals the memo.
//
// The observer is torn down via a Ginkgo DeferCleanup registered inside the
// function, so the spec can be re-run back-to-back without leaking processes.
//
// Assumes the running solanarpc sidecar was launched with a config whose
// verifiers.solana.allowed_programs includes the Memo program ID — the suite
// writes this file.
func ObserverAttestation(
	ctx context.Context,
	log logging.Logger,
	avalancheNetwork *network.LocalAvalancheNetwork,
	l1Info testinfo.L1TestInfo,
	observerBinaryPath string,
	sidecarConfigPath string,
	solanaKeypairPath string,
) {
	ginkgo.By("Step 1: Deploy OracleAdapter")
	_, fundedKey := avalancheNetwork.GetFundedAccountInfo()
	deployOpts, err := bind.NewKeyedTransactorWithChainID(fundedKey, l1Info.EVMChainID)
	Expect(err).Should(BeNil())

	adapterAddress, adapterDeployTx, adapterContract, err := oracleadapter.DeployOracleAdapter(
		deployOpts, l1Info.EthClient, deployOpts.From,
	)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, l1Info.EthClient, adapterDeployTx.Hash())
	log.Info("Deployed OracleAdapter", zap.Stringer("address", adapterAddress))

	ginkgo.By("Step 1b: Deploy TeleporterMessengerV2 with OracleAdapter as verifier")
	teleporterAddress := utils.DeployTeleporterV2(ctx, &l1Info, adapterAddress, fundedKey)
	log.Info("Deployed TeleporterMessengerV2", zap.Stringer("address", teleporterAddress))

	ginkgo.By("Step 1c: Deploy MockOracleReceiver pointing to TeleporterMessengerV2")
	mockAddress, mockDeployTx, mockContract, err := mockoraclereceiver.DeployMockOracleReceiver(
		deployOpts, l1Info.EthClient, teleporterAddress,
	)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, l1Info.EthClient, mockDeployTx.Hash())
	log.Info("Deployed MockOracleReceiver", zap.Stringer("address", mockAddress))

	ginkgo.By("Step 2: Start signature aggregator")
	sigAggConfig := utils.CreateDefaultSignatureAggregatorConfig(log, []testinfo.L1TestInfo{l1Info})
	sigAggConfigPath := utils.WriteSignatureAggregatorConfig(log, sigAggConfig, "sig-agg-observer-config.json")
	sigAggCancel, readyChan := utils.RunSignatureAggregatorExecutable(ctx, log, sigAggConfigPath, sigAggConfig)
	ginkgo.DeferCleanup(sigAggCancel)
	startupCtx, startupCancel := context.WithTimeout(ctx, 20*time.Second)
	defer startupCancel()
	utils.WaitForChannelClose(startupCtx, readyChan)

	ginkgo.By("Step 3: Allowlist the Memo program on OracleAdapter")
	allowTx, err := adapterContract.SetAllowedSource(deployOpts, "solana", memoProgram, true)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, l1Info.EthClient, allowTx.Hash())

	ginkgo.By("Step 4: Resolve Solana keypair address (for memo self-transfer)")
	solanaAddress := runCmd(ctx, "solana", "address", "--keypair", solanaKeypairPath)
	solanaAddress = strings.TrimSpace(solanaAddress)
	log.Info("Solana keypair address", zap.String("addr", solanaAddress))

	ginkgo.By("Step 5: Write observer config with deployed addresses")
	repoRoot, err := utils.GetRepoRoot()
	Expect(err).Should(BeNil())
	nonceFile := filepath.Join(repoRoot, "build/observer-nonces.json")
	// Delete any stale nonce file from a previous run so the delivered
	// message's nonce starts at 1, matching the operator's expectation for a
	// fresh demo.
	_ = os.Remove(nonceFile)

	observerConfigPath := filepath.Join(repoRoot, "build/observer-config.json")
	observerCfgBytes := buildObserverConfig(
		l1Info,
		avalancheNetwork.GetNetworkID(),
		teleporterAddress,
		mockAddress,
		fundedKey,
		sigAggConfig.APIPort,
		sidecarConfigPath,
		nonceFile,
	)
	Expect(os.WriteFile(observerConfigPath, observerCfgBytes, 0o600)).Should(BeNil())

	ginkgo.By("Step 6: Start solana-observer subprocess")
	observerCmd := exec.CommandContext(ctx, observerBinaryPath, "--config-path", observerConfigPath)
	observerReady := make(chan struct{})
	observerLogs := &teeReader{prefix: "[observer] "}
	observerCmd.Stdout = observerLogs
	observerCmd.Stderr = observerLogs
	observerLogs.notifyOnLine("subscribed to Solana logs", observerReady)
	Expect(observerCmd.Start()).Should(BeNil())

	ginkgo.DeferCleanup(func() {
		if observerCmd.Process == nil {
			return
		}
		log.Info("Sending SIGTERM to observer")
		_ = observerCmd.Process.Signal(os.Interrupt)
		done := make(chan struct{})
		go func() { _ = observerCmd.Wait(); close(done) }()
		select {
		case <-done:
		case <-time.After(5 * time.Second):
			log.Warn("Observer did not exit within 5s of SIGTERM; killing")
			_ = observerCmd.Process.Kill()
			<-done
		}
	})

	select {
	case <-observerReady:
		log.Info("Observer reported subscribed")
	case <-time.After(15 * time.Second):
		ginkgo.Fail("solana-observer did not print `subscribed to Solana logs` within 15s")
	}

	ginkgo.By("Step 7: Submit an spl-memo transaction with a fresh payload")
	memoText := fmt.Sprintf("Hello Avalanche %d", time.Now().UnixNano())
	log.Info("Submitting memo tx", zap.String("memo", memoText))
	runCmd(ctx,
		"solana", "transfer",
		"--url", "devnet",
		"--keypair", solanaKeypairPath,
		"--allow-unfunded-recipient",
		"--with-memo", memoText,
		"--fee-payer", solanaKeypairPath,
		"--commitment", "finalized",
		solanaAddress, "0.00001",
	)

	ginkgo.By("Step 8: Poll MockOracleReceiver.LastPayload until memo delivered (up to 3m)")
	pollDeadline := time.Now().Add(3 * time.Minute)
	want := []byte(memoText)
	for time.Now().Before(pollDeadline) {
		got, callErr := mockContract.LastPayload(&bind.CallOpts{Context: ctx})
		if callErr == nil && bytes.Equal(got, want) {
			log.Info("Auto-relay delivered", zap.String("memo", memoText))
			return
		}
		select {
		case <-ctx.Done():
			ginkgo.Fail("context cancelled while polling for delivery")
			return
		case <-time.After(3 * time.Second):
		}
	}
	ginkgo.Fail(fmt.Sprintf("timed out waiting for auto-relay delivery of memo %q", memoText))
}

// buildObserverConfig produces the observer-config.json body. Fields mirror
// solana-observer/config.go; the sidecar-config-path is the file the sidecar
// binary already reads, so the observer, the sidecar, and the validators all
// agree on which programs are in scope.
func buildObserverConfig(
	l1Info testinfo.L1TestInfo,
	networkID uint32,
	teleporter common.Address,
	dest common.Address,
	fundedKey *ecdsa.PrivateKey,
	sigAggPort uint16,
	sidecarConfigPath string,
	nonceFile string,
) []byte {
	host, port, err := utils.GetURIHostAndPort(l1Info.NodeURIs[0])
	Expect(err).Should(BeNil())
	l1RPC := fmt.Sprintf("http://%s:%d/ext/bc/%s/rpc", host, port, l1Info.BlockchainID.String())

	cfg := map[string]any{
		"solana": map[string]any{
			"ws_url":     "wss://api.devnet.solana.com",
			"rpc_url":    "https://api.devnet.solana.com",
			"commitment": "finalized",
		},
		"l1": map[string]any{
			"rpc_url":                   l1RPC,
			"chain_id":                  l1Info.EVMChainID.Uint64(),
			"blockchain_id":             l1Info.BlockchainID.String(),
			"network_id":                networkID,
			"teleporter_address":        teleporter.Hex(),
			"dest_contract":             dest.Hex(),
			"subnet_id":                 l1Info.SubnetID.String(),
			"delivery_private_key_hex":  hex.EncodeToString(crypto.FromECDSA(fundedKey)),
		},
		"aggregator_url":       fmt.Sprintf("http://127.0.0.1:%d", sigAggPort),
		"sidecar_config_path":  sidecarConfigPath,
		"nonce_file":           nonceFile,
	}
	b, err := json.MarshalIndent(cfg, "", "  ")
	Expect(err).Should(BeNil())
	return b
}

// runCmd runs cmd with args, checks for a zero exit, and returns stdout.
func runCmd(ctx context.Context, name string, args ...string) string {
	c := exec.CommandContext(ctx, name, args...)
	var out, stderr bytes.Buffer
	c.Stdout = &out
	c.Stderr = &stderr
	if err := c.Run(); err != nil {
		ginkgo.Fail(fmt.Sprintf("%s %v failed: %v\nstderr: %s", name, args, err, stderr.String()))
	}
	return out.String()
}

// teeReader implements io.Writer, mirroring the observer subprocess's stdout
// to the test's stdout (with a prefix for readability) and firing a channel
// the first time a target substring is observed.
type teeReader struct {
	prefix string
	scan   struct {
		enabled bool
		match   string
		fired   bool
		notify  chan struct{}
	}
	buf bytes.Buffer
}

func (t *teeReader) notifyOnLine(substring string, ch chan struct{}) {
	t.scan.enabled = true
	t.scan.match = substring
	t.scan.notify = ch
}

func (t *teeReader) Write(p []byte) (int, error) {
	// Mirror to test stdout.
	_, _ = fmt.Fprint(os.Stdout, t.prefix, string(p))
	if !t.scan.enabled || t.scan.fired {
		return len(p), nil
	}
	t.buf.Write(p)
	sc := bufio.NewScanner(&t.buf)
	for sc.Scan() {
		if strings.Contains(sc.Text(), t.scan.match) {
			t.scan.fired = true
			close(t.scan.notify)
			// Drop the rest — we're not consuming further from t.buf.
			t.buf.Reset()
			return len(p), nil
		}
	}
	return len(p), nil
}

var _ io.Writer = (*teeReader)(nil)
