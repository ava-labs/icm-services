package testinfo

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	subnetevmRPC "github.com/ava-labs/avalanchego/graft/subnet-evm/rpc"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/icm-services/icm-contracts/tests/rpc_client"
	"github.com/ava-labs/libevm/ethclient"
	libevmRPC "github.com/ava-labs/libevm/rpc"
	. "github.com/onsi/gomega"
)

// Abstraction over the info for Avalanche L1s and Ethereum networks.
type NetworkTestInfo interface {
	GetEVMTestInfo() *EVMTestInfo
	ChainID() string
	RPCClient(ctx context.Context) rpc_client.RpcClient
}

type EVMTestInfo struct {
	EVMChainID *big.Int
	EthClient  *ethclient.Client
}

// Tracks information about a test L1 used for executing tests against.
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

func (l1 *L1TestInfo) ChainID() string {
	return l1.BlockchainID.String()
}

func (l1 *L1TestInfo) RPCClient(ctx context.Context) rpc_client.RpcClient {
	rpcClient, err := subnetevmRPC.DialContext(
		ctx,
		fmt.Sprintf(
			"http://%s/ext/bc/%s/rpc",
			strings.TrimPrefix(l1.NodeURIs[0], "http://"),
			l1.BlockchainID.String(),
		),
	)
	Expect(err).Should(BeNil())
	return &rpc_client.SubnetEvmRpcClient{Client: rpcClient}
}

// Tracks information about a test Ethereum network used for executing tests against.
type EthereumTestInfo struct {
	EVMTestInfo
	BaseURL string
}

func (e *EthereumTestInfo) GetEVMTestInfo() *EVMTestInfo {
	return &e.EVMTestInfo
}

func (e *EthereumTestInfo) ChainID() string {
	return e.EVMChainID.String()
}

func (e *EthereumTestInfo) RPCClient(ctx context.Context) rpc_client.RpcClient {
	rpcClient, err := libevmRPC.DialContext(ctx, e.BaseURL)
	Expect(err).Should(BeNil())
	return &rpc_client.LibevmRPC{Client: rpcClient}
}
