package grpcserver

import (
	"context"
	"errors"
	"fmt"
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
		slog.Warn("invalid email", "error", err)
		return nil, err
	}
	if err := user.HashPassword(); err != nil {
		slog.Warn("invalid password", "error", err)
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
// UNIT TESTS

func (g *GRPCserver) Login(ctx context.Context, in *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	user := &models.User{
		Email: in.GetEmail(),
		Password: in.GetPassword(),
	}

	user2, err := g.authRepository.Get(user.Email)
	if err != nil {
		slog.Error("error with db", "error", err)
		return nil, err
	}

	if !user.ComparePassword(user2) {
		return nil, fmt.Errorf("wadawdiorvoijhaevv[huQRNHNH[ERH[R]MJCEAVRB['0H0UJ9 J [IOJEW[IHESAH01`WLJJ.KZ,JLS;/Zk/; OK]]]]]")
	}
	token, err := jwt.NewToken("SEX", *user, time.Hour*24)

	return &authv1.LoginResponse{Token: token}, nil
}

func (g *GRPCserver) IsAdmin(ctx context.Context, in *authv1.IsAdminRequest) (*authv1.IsAdminResponse, error) {

	id := in.GetUserId()

	yes, err := g.authRepository.GetAdminId(id)
	if err != nil {
		return nil, err
	}
	return &authv1.IsAdminResponse{
		IsAdmin: yes,
	}, nil
}

func (g *GRPCserver) RegisterAdmin(ctx context.Context, in *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	user := &models.User{
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	}

	if err := user.IsValid(); err != nil {
		slog.Warn("invalid email", "error", err)
		return nil, err
	}
	if err := user.HashPassword(); err != nil {
		slog.Warn("invalid password", "error", err)
		return nil, err
	}

	id, err := g.authRepository.CreateAdmin(user)
	if err != nil {
		slog.Error("error to create user", "err", err)
		return nil, err
	}
	user.ID = id

	token, err := jwt.NewToken("SEX", *user, time.Hour*24)

	return &authv1.RegisterResponse{Token: token}, nil
}

func (g *GRPCserver) LoginAdmin(ctx context.Context, in *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	user := &models.User{
		Email: in.GetEmail(),
		Password: in.GetPassword(),
	}

	user2, err := g.authRepository.GetAdmin(user.Email)
	if err != nil {
		slog.Error("error with db", "error", err)
		return nil, err
	}

	if !user.ComparePassword(user2) {
		return nil, fmt.Errorf("wadawdiorvoijhaevv[huQRNHNH[ERH[R]MJCEAVRB['0H0UJ9 J [IOJEW[IHESAH01`WLJJ.KZ,JLS;/Zk/; OK]]]]]")
	}
	token, err := jwt.NewToken("SEX", *user, time.Hour*24)

	return &authv1.LoginResponse{Token: token}, nil
}