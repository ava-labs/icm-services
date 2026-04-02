package rpcclient

import (
	"context"

	libevmRPC "github.com/ava-labs/libevm/rpc"
)

var _ RpcClient = (*SubnetEvmRpcClient)(nil)
var _ RpcClient = (*LibevmRPC)(nil)

// RpcClient Common abstraction over the RPC interfaces from libevm and subnet-evm
type RpcClient interface {
	CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error
	Close()
}

type SubnetEvmRpcClient struct {
	*libevmRPC.Client
}

type LibevmRPC struct {
	*libevmRPC.Client
}
