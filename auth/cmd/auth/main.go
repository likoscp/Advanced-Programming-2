package main

import (
	"log"
	"log/slog"
	"net"

	authv1 "github.com/barcek2281/finalProto/gen/go/auth"
	"github.com/likoscp/Advanced-Programming-2/auth/internal/configs"
	grpcserver "github.com/likoscp/Advanced-Programming-2/auth/internal/gRPCserver"
	"google.golang.org/grpc"
)

func main() {
	config := configs.NewConfig()

	lis, err := net.Listen("tcp", config.ConfigServer.Addr)
	if err != nil {
		log.Fatalf("cannot listen port, err: %v", err)
	}

	gRPC := grpcserver.NEWgrpcserver(config)

	s := grpc.NewServer()

	authv1.RegisterAuthServer(s, gRPC)

	slog.Info("server is starting", "port", config.ConfigServer.Addr)

	if err := s.Serve(lis); err != nil {
		slog.Error("error with server", "err", err)
	}

}
