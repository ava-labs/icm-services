package testinfo

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	subnetevmRPC "github.com/ava-labs/avalanchego/graft/subnet-evm/rpc"
	"github.com/ava-labs/avalanchego/ids"
	rpcclient "github.com/ava-labs/icm-services/icm-contracts/tests/rpc-client"
	"github.com/ava-labs/libevm/ethclient"
	. "github.com/onsi/gomega"
)

// NetworkTestInfo Abstraction over the info for Avalanche L1s and Ethereum networks.
type NetworkTestInfo interface {
	GetEVMTestInfo() *EVMTestInfo
	ChainID() string
	RPCClient(ctx context.Context) rpcclient.RpcClient
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

func (l1 *L1TestInfo) ChainID() string {
	return l1.BlockchainID.String()
}

func (l1 *L1TestInfo) RPCClient(ctx context.Context) rpcclient.RpcClient {
	rpcClient, err := subnetevmRPC.DialContext(
		ctx,
		fmt.Sprintf(
			"http://%s/ext/bc/%s/rpc",
			strings.TrimPrefix(l1.NodeURIs[0], "http://"),
			l1.BlockchainID.String(),
		),
	)
	Expect(err).Should(BeNil())
	return &rpcclient.SubnetEvmRpcClient{Client: rpcClient}
}
