package grpcserver

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	authv1 "github.com/barcek2281/finalProto/gen/go/auth"
	"github.com/likoscp/Advanced-Programming-2/auth/internal/configs"
	"github.com/likoscp/Advanced-Programming-2/auth/internal/lib/jwt"
	"github.com/likoscp/Advanced-Programming-2/auth/internal/repository"
	"github.com/likoscp/Advanced-Programming-2/auth/internal/store/postgresql"
	"github.com/likoscp/Advanced-Programming-2/auth/models"
)

const (
	adminRole = "admin"
	userRole  = "user"
)

var (
	ErrNotValidEmail    = errors.New("invalid email")
	ErrNotValidPassword = errors.New("invalid password")

	// Prometheus metrics
	userRegistrations = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "auth_user_registrations_total",
		Help: "Total number of user registrations.",
	})
	userLogins = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "auth_user_logins_total",
		Help: "Total number of user logins.",
	})
	adminRegistrations = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "auth_admin_registrations_total",
		Help: "Total number of admin registrations.",
	})
	adminLogins = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "auth_admin_logins_total",
		Help: "Total number of admin logins.",
	})
	adminChecks = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "auth_admin_checks_total",
		Help: "Total number of admin privilege checks.",
	})
)

func init() {
	// Register Prometheus metrics
	prometheus.MustRegister(userRegistrations)
	prometheus.MustRegister(userLogins)
	prometheus.MustRegister(adminRegistrations)
	prometheus.MustRegister(adminLogins)
	prometheus.MustRegister(adminChecks)
}

type GRPCserver struct {
	authv1.UnimplementedAuthServer
	config         *configs.Config
	authRepository *repository.AuthRepository
}

func NEWgrpcserver(config *configs.Config) (*GRPCserver, error) {
	store, err := postgresql.NewStore(config.ConfigDB)
	if err != nil {
		return nil, err
	}
	authRepo := repository.NewAuthRepository(store)

	// Start HTTP server for metrics
	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())
		slog.Info("starting metrics server", "port", "8083")
		if err := http.ListenAndServe(":8083", mux); err != nil {
			slog.Error("failed to start metrics server", "error", err)
		}
	}()

	return &GRPCserver{
		config:         config,
		authRepository: authRepo,
	}, nil
}

func (g *GRPCserver) Register(ctx context.Context, in *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	user := &models.User{
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
		Username: in.GetUsername(),
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

	token, err := jwt.NewToken(g.config.ConfigServer.Secret, userRole, *user, time.Hour*24)
	if err != nil {
		return nil, err
	}

	// Send registration event to email microservice
	emailReq := map[string]string{
		"username": user.Username,
		"email":    user.Email,
	}
	reqBody, _ := json.Marshal(emailReq)
	resp, err := http.Post("http://email-service:8082/mail/register", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		slog.Warn("failed to notify email service", "error", err)
		// Continue even if email fails to avoid blocking registration
	} else {
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			slog.Warn("email service returned non-OK status", "status", resp.StatusCode)
		}
	}

	// Increment metric
	userRegistrations.Inc()

	return &authv1.RegisterResponse{Token: token}, nil
}

func (g *GRPCserver) Login(ctx context.Context, in *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	user := &models.User{
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	}

	user2, err := g.authRepository.Get(user.Email)
	if err != nil {
		slog.Error("error with db", "error", err)
		return nil, err
	}

	if !user.ComparePassword(user2) {
		return nil, fmt.Errorf("password doesn't match")
	}
	token, err := jwt.NewToken(g.config.ConfigServer.Secret, userRole, *user, time.Hour*24)
	if err != nil {
		return nil, err
	}

	// Increment metric
	userLogins.Inc()

	return &authv1.LoginResponse{Token: token}, nil
}

func (g *GRPCserver) IsAdmin(ctx context.Context, in *authv1.IsAdminRequest) (*authv1.IsAdminResponse, error) {
	id := in.GetUserId()

	yes, err := g.authRepository.GetAdminId(id)
	if err != nil {
		return nil, err
	}

	// Increment metric
	adminChecks.Inc()

	return &authv1.IsAdminResponse{
		IsAdmin: yes,
	}, nil
}

func (g *GRPCserver) RegisterAdmin(ctx context.Context, in *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	user := &models.User{
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
		Username: in.GetUsername(),
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

	token, err := jwt.NewToken(g.config.ConfigServer.Secret, adminRole, *user, time.Hour*24)
	if err != nil {
		return nil, err
	}

	// Increment metric
	adminRegistrations.Inc()

	return &authv1.RegisterResponse{Token: token}, nil
}

func (g *GRPCserver) LoginAdmin(ctx context.Context, in *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	user := &models.User{
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	}

	user2, err := g.authRepository.GetAdmin(user.Email)
	if err != nil {
		slog.Error("error with db", "error", err)
		return nil, err
	}

	if !user.ComparePassword(user2) {
		return nil, fmt.Errorf("password doesn't match")
	}
	token, err := jwt.NewToken(g.config.ConfigServer.Secret, adminRole, *user, time.Hour*24)
	if err != nil {
		return nil, err
	}

	// Increment metric
	adminLogins.Inc()

	return &authv1.LoginResponse{Token: token}, nil
}

func (g *GRPCserver) GetInfoUser(ctx context.Context, in *authv1.UserInfoRequest) (*authv1.UserInfoResponse, error) {
	user, err := g.authRepository.GetById(in.UserId)
	if err != nil {
		return nil, err
	}

	return &authv1.UserInfoResponse{
		Email:    user.Email,
		Username: user.Username,
	}, nil
}
