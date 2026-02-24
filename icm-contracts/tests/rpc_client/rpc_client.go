package rpc_client

import (
	"context"

	subnetevmRPC "github.com/ava-labs/avalanchego/graft/subnet-evm/rpc"
	libevmRPC "github.com/ava-labs/libevm/rpc"
)

var _ RpcClient = (*SubnetEvmRpcClient)(nil)
var _ RpcClient = (*LibevmRPC)(nil)

// Common abstraction over the RPC interfaces from libevm and subnet-evm
type RpcClient interface {
	CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error
	Close()
}

type SubnetEvmRpcClient struct {
	*subnetevmRPC.Client
}

type LibevmRPC struct {
	*libevmRPC.Client
}
