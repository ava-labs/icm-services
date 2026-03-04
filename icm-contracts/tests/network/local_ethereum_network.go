package network

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/libevm/ethclient"
	. "github.com/onsi/gomega"
)

const (
	defaultGethURL = "http://127.0.0.1:5050"

	// Pre-funded key for the local geth dev test account.
	// The run_ethereum_network.sh script transfers funds from the geth dev
	// account to the address derived from this key (0x6288dAdf62B57dd9A4ddcd02F88A98d0eb6c2598).
	gethFundedKeyHex = "764a4a322753120b4667a20b6309cb5ec754a22bdbcbd62398be8f803b255337"
)

type LocalEthereumNetwork struct {
	rpcURL     string
	ethClient  *ethclient.Client
	fundedKey  *ecdsa.PrivateKey
	fundedAddr common.Address
	chainID    *big.Int
}

var _ LocalNetwork = (*LocalEthereumNetwork)(nil)

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

func (n *LocalEthereumNetwork) GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	return n.fundedAddr, n.fundedKey
}

func (n *LocalEthereumNetwork) TearDownNetwork() {
	if n.ethClient != nil {
		n.ethClient.Close()
	}
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
