package testinfo

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/libevm/ethclient"
	"github.com/ava-labs/libevm/rpc"
	. "github.com/onsi/gomega"
)

// NetworkTestInfo Abstraction over the info for Avalanche L1s and Ethereum networks.
type NetworkTestInfo interface {
	GetEVMTestInfo() *EVMTestInfo
	ChainID() ids.ID
	RPCClient(ctx context.Context) *rpc.Client
}

type EVMTestInfo struct {
	EVMChainID *big.Int
	EthClient  *ethclient.Client
}

// L1TestInfo Tracks information about a test Avalanche L1 used for executing tests against.
type L1TestInfo struct {
	EVMTestInfo
	SubnetID                     ids.ID
	BlockchainID                 ids.ID
	NodeURIs                     []string
	WSClient                     *ethclient.Client
	RequirePrimaryNetworkSigners bool
}

func (l1 *L1TestInfo) GetEVMTestInfo() *EVMTestInfo {
	return &l1.EVMTestInfo
}

func (l1 *L1TestInfo) ChainID() ids.ID {
	return l1.BlockchainID
}

func (l1 *L1TestInfo) RPCClient(ctx context.Context) *rpc.Client {
	rpcClient, err := rpc.DialContext(
		ctx,
		fmt.Sprintf(
			"http://%s/ext/bc/%s/rpc",
			strings.TrimPrefix(l1.NodeURIs[0], "http://"),
			l1.BlockchainID.String(),
		),
	)
	Expect(err).Should(BeNil())
	return rpcClient
}

// EthereumTestInfo Tracks information about a test Ethereum network used for executing tests against.
type EthereumTestInfo struct {
	EVMTestInfo
	BaseURL string
}

func (e *EthereumTestInfo) GetEVMTestInfo() *EVMTestInfo {
	return &e.EVMTestInfo
}

func (e *EthereumTestInfo) ChainID() ids.ID {
	blockchainID, err := ids.ToID(e.EVMChainID.FillBytes(make([]byte, 32)))
	Expect(err).Should(BeNil())
	return blockchainID
}

func (e *EthereumTestInfo) RPCClient(ctx context.Context) *rpc.Client {
	rpcClient, err := rpc.DialContext(ctx, e.BaseURL)
	Expect(err).Should(BeNil())
	return rpcClient
}
