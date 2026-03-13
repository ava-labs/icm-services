package network

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
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
	BaseURL    string
	EthClient  *ethclient.Client
	fundedKey  *ecdsa.PrivateKey
	fundedAddr common.Address
	ChainID    *big.Int
}

const (
	localEthereumNetworkBaseURL   = "http://127.0.0.1:5050"
	localEthereumNetworkFundedKey = "764A4A322753120B4667A20B6309CB5EC754A22BDBCBD62398BE8F803B255337"
)

func NewLocalEthereumNetwork(ctx context.Context) *LocalEthereumNetwork {
	return NewLocalEthereumNetworkFromURL(ctx, localEthereumNetworkBaseURL)
}

func NewLocalEthereumNetworkFromURL(ctx context.Context, rpcURL string) *LocalEthereumNetwork {
	ethClient, err := ethclient.DialContext(ctx, rpcURL)
	Expect(err).Should(BeNil())

	fundedKey, err := crypto.HexToECDSA(localEthereumNetworkFundedKey)
	Expect(err).Should(BeNil())
	fundedAddr := crypto.PubkeyToAddress(fundedKey.PublicKey)

	chainID, err := ethClient.ChainID(ctx)
	Expect(err).Should(BeNil())

	return &LocalEthereumNetwork{
		BaseURL:    rpcURL,
		EthClient:  ethClient,
		fundedKey:  fundedKey,
		fundedAddr: fundedAddr,
		ChainID:    chainID,
	}
}

func (n *LocalEthereumNetwork) GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	return n.fundedAddr, n.fundedKey
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
	if n.EthClient != nil {
		n.EthClient.Close()
	}
}
