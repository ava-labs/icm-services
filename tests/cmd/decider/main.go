// a CLI command to serve as a gRPC provider of icm-relayer/proto/decider

package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/ava-labs/avalanchego/utils/logging"
	pb "github.com/ava-labs/icm-services/proto/pb/decider"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 50051, "The server port")

type deciderServer struct {
	pb.UnimplementedDeciderServiceServer
}

func (s *deciderServer) ShouldSendMessage(
	ctx context.Context,
	msg *pb.ShouldSendMessageRequest,
) (*pb.ShouldSendMessageResponse, error) {
	return &pb.ShouldSendMessageResponse{
		ShouldSendMessage: true,
	}, nil
}

func main() {
	flag.Parse()

	server := grpc.NewServer()
	pb.RegisterDeciderServiceServer(server, &deciderServer{})

	log := logging.NewLogger(
		"signature-aggregator",
		logging.NewWrappedCore(
			logging.Info,
			os.Stdout,
			logging.JSON.ConsoleEncoder(),
		),
	)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal("decider failed to listen", zap.Error(err))
		os.Exit(1)
	}

	log.Info("decider listening at", zap.Stringer("address", listener.Addr()))

	err = server.Serve(listener)
	if err != nil {
		log.Fatal("decider failed to serve", zap.Error(err))
		os.Exit(1)
	}
}
