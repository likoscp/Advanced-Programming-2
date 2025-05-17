package grpcserver

import (
	authv1 "github.com/barcek2281/finalProto/gen/go/auth"
	"github.com/likoscp/Advanced-Programming-2/auth/internal/configs"
)

type GRPCserver struct {
	authv1.UnimplementedAuthServer
}

func NEWgrpcserver(config *configs.Config) *GRPCserver {
	return &GRPCserver{}
}

// func (g GRPCserver) Register(ctx context.Context, in *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {

// }
