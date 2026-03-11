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

var _ LocalNetwork = (*LocalEthereumNetwork)(nil)

type LocalEthereumNetwork struct {
	rpcURL     string
	ethClient  *ethclient.Client
	fundedKey  *ecdsa.PrivateKey
	fundedAddr common.Address
	chainID    *big.Int
	cmd        *exec.Cmd
}

const (
	defaultGethURL                = "http://127.0.0.1:5050"
	gethFundedKeyHex              = "764a4a322753120b4667a20b6309cb5ec754a22bdbcbd62398be8f803b255337"
	ethereumNetworkStartupTimeout = 60 * time.Second
)

func NewLocalEthereumNetwork(ctx context.Context) *LocalEthereumNetwork {
	return NewLocalEthereumNetworkFromURL(ctx, defaultGethURL)
}

func NewLocalEthereumNetworkFromURL(ctx context.Context, rpcURL string) *LocalEthereumNetwork {
	ethClient, err := ethclient.DialContext(ctx, rpcURL)
	Expect(err).Should(BeNil())

	fundedKey, err := crypto.HexToECDSA(gethFundedKeyHex)
	Expect(err).Should(BeNil())
	fundedAddr := crypto.PubkeyToAddress(fundedKey.PublicKey)

	chainID, err := ethClient.ChainID(ctx)
	Expect(err).Should(BeNil())

	return &LocalEthereumNetwork{
		rpcURL:     rpcURL,
		ethClient:  ethClient,
		fundedKey:  fundedKey,
		fundedAddr: fundedAddr,
		chainID:    chainID,
	}
}

// StartLocalEthereumNetwork starts a local Ethereum network
// and waits for it to be ready.
func StartLocalEthereumNetwork(ctx context.Context) *LocalEthereumNetwork {
	log.Info("Starting local Ethereum network")

	repoRoot, err := utils.GetRepoRoot()
	Expect(err).Should(BeNil())

	scriptPath := filepath.Join(repoRoot, "ethereum", "run_ethereum_network.sh")

	// Use exec.Command instead of CommandContext to prevent context cancellation
	// from sending SIGKILL before we can gracefully shutdown with SIGTERM
	cmd := exec.Command(scriptPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	Expect(err).Should(BeNil(), "Failed to start Ethereum network script")

	log.Info(fmt.Sprintf("Started Ethereum network process with PID %d", cmd.Process.Pid))

	startTime := time.Now()
	client, err := waitForEthereumNetwork(ctx, ethereumNetworkStartupTimeout)
	Expect(err).Should(BeNil(), "Ethereum network failed to become ready")

	log.Info(fmt.Sprintf("Ethereum network is ready after %s", time.Since(startTime)))

	chainID, err := client.ChainID(ctx)
	Expect(err).Should(BeNil())

	fundedKeyBytes, err := hex.DecodeString(gethFundedKeyHex)
	Expect(err).Should(BeNil())
	secp256k1Key, err := secp256k1.ToPrivateKey(fundedKeyBytes)
	Expect(err).Should(BeNil())
	fundedKey := secp256k1Key.ToECDSA()
	fundedAddr := crypto.PubkeyToAddress(fundedKey.PublicKey)

	return &LocalEthereumNetwork{
		rpcURL:     defaultGethURL,
		ethClient:  client,
		fundedKey:  fundedKey,
		fundedAddr: fundedAddr,
		chainID:    chainID,
		cmd:        cmd,
	}
}

func (n *LocalEthereumNetwork) GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	return n.fundedAddr, n.fundedKey
}

func (n *LocalEthereumNetwork) EthClient() *ethclient.Client {
	return n.ethClient
}

func (n *LocalEthereumNetwork) ChainID() *big.Int {
	return n.chainID
}

func (n *LocalEthereumNetwork) RPCURL() string {
	return n.rpcURL
}

func (n *LocalEthereumNetwork) EthereumTestInfo() *testinfo.EthereumTestInfo {
	return &testinfo.EthereumTestInfo{
		EVMTestInfo: testinfo.EVMTestInfo{
			EthClient:  n.ethClient,
			EVMChainID: n.chainID,
		},
		BaseURL: n.rpcURL,
	}
}

func (n *LocalEthereumNetwork) TearDownNetwork() {
	if n.cmd == nil || n.cmd.Process == nil {
		if n.ethClient != nil {
			n.ethClient.Close()
		}
		return
	}

	log.Info("Tearing down local Ethereum network")
	log.Info(fmt.Sprintf("Stopping Ethereum network process (PID %d)", n.cmd.Process.Pid))

	err := n.cmd.Process.Signal(syscall.SIGTERM)
	if err != nil {
		log.Warn(fmt.Sprintf("Failed to send interrupt signal to Ethereum network: %v", err))
	}

	done := make(chan error, 1)
	go func() {
		done <- n.cmd.Wait()
	}()

	select {
	case <-time.After(10 * time.Second):
		log.Warn("Ethereum network process did not exit gracefully, killing it")
		_ = n.cmd.Process.Kill()
		<-done
	case err := <-done:
		if err != nil {
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
			resp, err := http.Post(
				defaultGethURL,
				"application/json",
				nil,
			)
			if err == nil {
				resp.Body.Close()
				client, err := ethclient.Dial(defaultGethURL)
				if err == nil {
					return client, nil
				} else {
					log.Warn("Failed to create Ethereum client, retrying in 1 second")
				}
			} else {
				log.Debug("Failed to connect to Ethereum RPC endpoint, retrying in 1 second")
			}
		}
	}
}
