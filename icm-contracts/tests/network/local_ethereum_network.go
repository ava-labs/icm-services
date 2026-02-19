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
	"time"

	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/icm-services/icm-contracts/tests/testinfo"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/libevm/ethclient"
	"github.com/ava-labs/libevm/log"
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

// StartLocalEthereumNetwork starts a local Ethereum network using the docker script
// and waits for it to be ready.
func StartLocalEthereumNetwork(ctx context.Context) *LocalEthereumNetwork {
	log.Info("Starting local Ethereum network")

	// Get the repository root
	repoRoot, err := getRepoRoot()
	Expect(err).Should(BeNil())

	scriptPath := filepath.Join(repoRoot, "ethereum", "docker", "run_ethereum_network.sh")

	// Start the Ethereum network script in the background
	cmd := exec.CommandContext(ctx, scriptPath)
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

	return &LocalEthereumNetwork{
		BaseURL:         localEthereumNetworkBaseURL,
		EthClient:       client,
		ChainID:         chainID,
		globalFundedKey: globalFundedKey,
		cmd:             cmd,
	}
}

// NewLocalEthereumNetwork connects to an already-running local Ethereum network.
// Use StartLocalEthereumNetwork to start a new network instance.
func NewLocalEthereumNetwork(ctx context.Context) *LocalEthereumNetwork {
	// The local Ethereum network must already be running and accessible at a known endpoint.
	// We test that it is accessible here.
	client, err := ethclient.Dial(localEthereumNetworkBaseURL)
	Expect(err).Should(BeNil())

	// Get the chain ID of the local Ethereum network
	chainID, err := client.ChainID(ctx)
	Expect(err).Should(BeNil())

	fundedKeyBytes, err := hex.DecodeString(localEthereumNetworkFundedKey)
	Expect(err).Should(BeNil())
	globalFundedKey, err := secp256k1.ToPrivateKey(fundedKeyBytes)
	Expect(err).Should(BeNil())

	return &LocalEthereumNetwork{
		BaseURL:         localEthereumNetworkBaseURL,
		EthClient:       client,
		ChainID:         chainID,
		globalFundedKey: globalFundedKey,
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

func (n *LocalEthereumNetwork) TearDownNetwork() {
	log.Info("Tearing down local Ethereum network")
	Expect(n).ShouldNot(BeNil())
	Expect(n.EthClient).ShouldNot(BeNil())

	// Stop the Ethereum network process if it was started by us
	if n.cmd != nil && n.cmd.Process != nil {
		log.Info(fmt.Sprintf("Stopping Ethereum network process (PID %d)", n.cmd.Process.Pid))

		// Send interrupt signal to allow graceful shutdown
		err := n.cmd.Process.Signal(os.Interrupt)
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
				log.Info(fmt.Sprintf("Ethereum network process exited with error: %v", err))
			} else {
				log.Info("Ethereum network process stopped successfully")
			}
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
				}
			}
			// Continue polling
		}
	}
}

// getRepoRoot finds the repository root by looking for the go.mod file
func getRepoRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("could not find repository root (no go.mod found)")
		}
		dir = parent
	}
}
