// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// oracle-sidecar is a minimal gRPC sidecar for E2E testing. It implements the
// oracle.OracleSidecar service and accepts every Verify request unconditionally,
// standing in for a real Solana RPC sidecar.
// Usage: oracle-sidecar --port 9900
package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// acceptAll is a dummy interface so grpc.RegisterService type-checks pass.
type acceptAll interface {
	accept()
}

type mockSidecar struct{}

func (mockSidecar) accept() {}

// verifyHandler accepts every oracle message unconditionally.
// It decodes the request into emptypb.Empty (proto3 discards unknown fields)
// and returns an empty response matching the wire format of oracle.VerifyResponse.
func verifyHandler(
	_ interface{},
	ctx context.Context,
	dec func(interface{}) error,
	interceptor grpc.UnaryServerInterceptor,
) (interface{}, error) {
	var req emptypb.Empty
	if err := dec(&req); err != nil {
		return nil, err
	}
	h := func(_ context.Context, _ interface{}) (interface{}, error) {
		return &emptypb.Empty{}, nil
	}
	if interceptor == nil {
		return h(ctx, &req)
	}
	return interceptor(ctx, &req, &grpc.UnaryServerInfo{Server: nil, FullMethod: "/oracle.OracleSidecar/Verify"}, h)
}

var oracleServiceDesc = grpc.ServiceDesc{
	ServiceName: "oracle.OracleSidecar",
	HandlerType: (*acceptAll)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Verify",
			Handler:    verifyHandler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oracle/oracle.proto",
}

func main() {
	port := flag.Uint("port", 9900, "port to listen on")
	flag.Parse()

	addr := fmt.Sprintf(":%d", *port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "oracle-sidecar: listen: %v\n", err)
		os.Exit(1)
	}

	s := grpc.NewServer()
	s.RegisterService(&oracleServiceDesc, &mockSidecar{})

	fmt.Fprintf(os.Stdout, "oracle-sidecar listening on %s\n", addr)
	if err := s.Serve(lis); err != nil {
		fmt.Fprintf(os.Stderr, "oracle-sidecar: serve: %v\n", err)
		os.Exit(1)
	}
}
