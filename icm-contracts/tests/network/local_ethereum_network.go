package network

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"

	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/icm-services/log"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/libevm/ethclient"
	. "github.com/onsi/gomega"
)

// LocalEthereumNetwork is a separate network type for external Ethereum chains.
// It does not implement the same interface as LocalNetwork (Avalanche) since
// the functionality is different.

var _ LocalNetwork = (*LocalEthereumNetwork)(nil)

type LocalEthereumNetwork struct {
	BaseURL         string
	EthClient       *ethclient.Client
	ChainID         *big.Int
	globalFundedKey *secp256k1.PrivateKey
	cmd             *exec.Cmd
}

const (
	localEthereumNetworkBaseURL   = "http://127.0.0.1:5050"
	localEthereumNetworkFundedKey = "764A4A322753120B4667A20B6309CB5EC754A22BDBCBD62398BE8F803B255337"
	ethereumNetworkStartupTimeout = 60 * time.Second
)

// StartLocalEthereumNetwork starts a local Ethereum network
// and waits for it to be ready.
func StartLocalEthereumNetwork(ctx context.Context) *LocalEthereumNetwork {
	log.Info("Starting local Ethereum network")

	// Get the repository root
	repoRoot, err := utils.GetRepoRoot()
	Expect(err).Should(BeNil())

	scriptPath := filepath.Join(repoRoot, "ethereum", "run_ethereum_network.sh")

	// Start the Ethereum network script in the background
	// Use exec.Command instead of CommandContext to prevent context cancellation
	// from sending SIGKILL before we can gracefully shutdown with SIGTERM
	cmd := exec.Command(scriptPath)
	// Inherit stdout/stderr for debugging
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	Expect(err).Should(BeNil(), "Failed to start Ethereum network script")

	log.Info(fmt.Sprintf("Started Ethereum network process with PID %d", cmd.Process.Pid))

	// Wait for the network to be ready
	startTime := time.Now()
	client, err := waitForEthereumNetwork(ctx, ethereumNetworkStartupTimeout)
	Expect(err).Should(BeNil(), "Ethereum network failed to become ready")

	log.Info(fmt.Sprintf("Ethereum network is ready after %s", time.Since(startTime)))

	// Get the chain ID of the local Ethereum network
	chainID, err := client.ChainID(ctx)
	Expect(err).Should(BeNil())

	fundedKeyBytes, err := hex.DecodeString(localEthereumNetworkFundedKey)
	Expect(err).Should(BeNil())
	globalFundedKey, err := secp256k1.ToPrivateKey(fundedKeyBytes)
	Expect(err).Should(BeNil())

	// Wait for the funded account to have a non-zero balance.
	// The startup script transfers funds asynchronously after geth becomes
	// reachable, so we need to poll until the transfer is mined.
	fundedAddr := crypto.PubkeyToAddress(globalFundedKey.ToECDSA().PublicKey)
	waitForBalance(ctx, client, fundedAddr, ethereumNetworkStartupTimeout)

	return &LocalEthereumNetwork{
		BaseURL:         localEthereumNetworkBaseURL,
		EthClient:       client,
		ChainID:         chainID,
		globalFundedKey: globalFundedKey,
		cmd:             cmd,
	}
}

func (n *LocalEthereumNetwork) GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	ecdsaKey := n.globalFundedKey.ToECDSA()
	fundedAddress := crypto.PubkeyToAddress(ecdsaKey.PublicKey)
	return fundedAddress, ecdsaKey
}

func (n *LocalEthereumNetwork) EthereumTestInfo() *testinfo.EthereumTestInfo {
	return &testinfo.EthereumTestInfo{
		EVMTestInfo: testinfo.EVMTestInfo{
			EthClient:  n.EthClient,
			EVMChainID: n.ChainID,
		},
		BaseURL: n.BaseURL,
	}
}

func (n *LocalEthereumNetwork) GetNetworkInfo() []testinfo.NetworkTestInfo {
	networks := make([]testinfo.NetworkTestInfo, 1)
	networks[0] = n.EthereumTestInfo()
	return networks
}

func (n *LocalEthereumNetwork) TearDownNetwork() {
	log.Info("Tearing down local Ethereum network")

	// Stop the Ethereum network process if it was started by us
	if n.cmd == nil || n.cmd.Process == nil {
		return
	}
	log.Info(fmt.Sprintf("Stopping Ethereum network process (PID %d)", n.cmd.Process.Pid))

	// Send interrupt signal to allow graceful shutdown
	err := n.cmd.Process.Signal(syscall.SIGTERM)
	if err != nil {
		log.Warn(fmt.Sprintf("Failed to send interrupt signal to Ethereum network: %v", err))
	}

	// Wait for the process to exit (with timeout)
	done := make(chan error, 1)
	go func() {
		done <- n.cmd.Wait()
	}()

	select {
	case <-time.After(10 * time.Second):
		log.Warn("Ethereum network process did not exit gracefully, killing it")
		_ = n.cmd.Process.Kill()
		<-done // Wait for Wait() to complete
	case err := <-done:
		if err != nil {
			// Exit codes 130 (SIGINT) and 143 (SIGTERM) are expected when we terminate the process
			if exitErr, ok := err.(*exec.ExitError); ok {
				exitCode := exitErr.ExitCode()
				if exitCode == 130 || exitCode == 143 {
					log.Info("Ethereum network process stopped successfully")
				} else {
					log.Error(fmt.Sprintf("Ethereum network process exited with error: %v", err))
				}
			} else {
				log.Error(fmt.Sprintf("Ethereum network process exited with error: %v", err))
			}
		} else {
			log.Info("Ethereum network process stopped successfully")
		}
	}
}

// waitForEthereumNetwork polls the Ethereum RPC endpoint until it responds successfully
// or the timeout is reached.
func waitForEthereumNetwork(ctx context.Context, timeout time.Duration) (*ethclient.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("timeout waiting for Ethereum network to be ready")
		case <-ticker.C:
			// Try to connect to the RPC endpoint
			resp, err := http.Post(
				localEthereumNetworkBaseURL,
				"application/json",
				nil,
			)
			if err == nil {
				resp.Body.Close()
				// Successfully connected, now try to create a proper client
				client, err := ethclient.Dial(localEthereumNetworkBaseURL)
				if err == nil {
					return client, nil
				} else {
					log.Warn("Failed to create Ethereum client, retrying in 1 second")
				}
			} else {
				log.Debug("Failed to connect to Ethereum RPC endpoint, retrying in 1 second")
			}
			// Continue polling
		}
	}
}

// waitForBalance polls until the given address has a non-zero balance or the
// timeout is reached. This is needed because the startup script funds the test
// address asynchronously after the RPC becomes reachable.
func waitForBalance(ctx context.Context, client *ethclient.Client, addr common.Address, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			Expect(ctx.Err()).ShouldNot(HaveOccurred(),
				"Timed out waiting for funded account to have non-zero balance")
		case <-ticker.C:
			bal, err := client.BalanceAt(ctx, addr, nil)
			if err == nil && bal.Sign() > 0 {
				log.Info(fmt.Sprintf("Funded account %s has balance %s wei", addr.Hex(), bal.String()))
				return
			}
		}
	}
}
