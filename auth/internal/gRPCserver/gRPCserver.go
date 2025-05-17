package grpcserver

import (
	"context"
	"errors"
	"log/slog"
	"time"

	authv1 "github.com/barcek2281/finalProto/gen/go/auth"
	"github.com/likoscp/Advanced-Programming-2/auth/internal/configs"
	"github.com/likoscp/Advanced-Programming-2/auth/internal/lib/jwt"
	"github.com/likoscp/Advanced-Programming-2/auth/internal/repository"
	"github.com/likoscp/Advanced-Programming-2/auth/internal/store/postgresql"
	"github.com/likoscp/Advanced-Programming-2/auth/models"
)

var (
	ErrNotValidEmail    = errors.New("invalid email")
	ErrNotValidPassword = errors.New("invalid password")
)

type GRPCserver struct {
	authv1.UnimplementedAuthServer
	authRepository *repository.AuthRepository
}

func NEWgrpcserver(config *configs.Config) (*GRPCserver, error) {
	store, err := postgresql.NewStore(*config.ConfigDB)
	if err != nil {
		return nil, err
	}
	authRepo := repository.NewAuthRepository(store)
	return &GRPCserver{
		authRepository: authRepo,
	}, nil
}

func (g *GRPCserver) Register(ctx context.Context, in *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	user := &models.User{
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	}

	if err := user.IsValid(); err != nil {
		slog.Warn("invalid email")
		return nil, err
	}
	if err := user.HashPassword(); err != nil {
		slog.Warn("invalid password")
		return nil, err
	}

	id, err := g.authRepository.Create(user)
	if err != nil {
		slog.Error("error to create user", "err", err)
		return nil, err
	}
	user.ID = id

	token, err := jwt.NewToken("SEX", *user, time.Hour*24)

	return &authv1.RegisterResponse{Token: token}, nil
}
