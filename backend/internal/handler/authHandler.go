package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	authv1 "github.com/barcek2281/finalProto/gen/go/auth"
	"github.com/likoscp/Advanced-Programming-2/backend/internal/config"
	"github.com/likoscp/Advanced-Programming-2/backend/internal/lib/response"
	"github.com/likoscp/Advanced-Programming-2/backend/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthHandler struct {
	AuthClient authv1.AuthClient
}

func NewAuthHandler(config *config.Config) (*AuthHandler, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("auth:%s", config.ConfigServer.AuthAddr), grpc.WithTransportCredentials(insecure.NewCredentials()))

	authClient := authv1.NewAuthClient(conn)
	return &AuthHandler{
		AuthClient: authClient,
	}, err
}

func (a *AuthHandler) Configure(mux *http.ServeMux) {
	mux.HandleFunc("POST /auth/login", a.Login())
	mux.HandleFunc("POST /auth/register", a.Register())
	mux.HandleFunc("GET /auth/is-admin/{id}", a.IsAdmin())
	mux.HandleFunc("POST /auth/login-admin", a.LoginAdmin())
	mux.HandleFunc("POST /auth/register-admin", a.RegisterAdmin())
	mux.HandleFunc("GET /auth/user-info/{id}", a.UserInfo())
}

// Login Login User
//
//	@Summary		Login User
//	@Description	Login User
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.User	true	"awdasd"
//	@Success		200		{object}	models.Token
//	@Router			/auth/login [post]
func (a *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			slog.Warn("error to json", "error", err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
		defer cancel()

		res, err := a.AuthClient.Login(ctx, &authv1.LoginRequest{
			Email:    user.Email,
			Password: user.Password,
		})

		if err != nil {
			slog.Error("error to auth", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.ResponseJSON(w, http.StatusOK, models.Token{Token: res.Token})
	}
}

// Register Register User
//
//	@Summary		Register User
//	@Description	Register User
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.User	true	"getting"
//	@Success		200		{object}	models.Token
//	@Router			/auth/register [post]
func (a *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			slog.Warn("error to json", "error", err)
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
		defer cancel()
		res, err := a.AuthClient.Register(ctx, &authv1.RegisterRequest{
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
		})

		if err != nil {
			slog.Error("error to atuh", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.ResponseJSON(w, http.StatusOK, models.Token{Token: res.Token})

	}
}

// IsAdmin is Admin?
//
//	@Summary		is Admin?
//	@Description	is Admin?
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200		{object}	models.IsReallyAdmin
//	@Router			/auth/is-admin/{id} [get]
func (a *AuthHandler) IsAdmin() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
		defer cancel()

		res, err := a.AuthClient.IsAdmin(ctx, &authv1.IsAdminRequest{
			UserId: id,
		})
		if err != nil {
			slog.Warn("lox", "error", err)
			response.ResponseError(w, http.StatusBadRequest, err)
			return
		}

		response.ResponseJSON(w, http.StatusAccepted, models.IsReallyAdmin{IsAdmin: res.IsAdmin})
	}
}

// LoginAdmin Login Admin
//
//	@Summary		Login Admin
//	@Description	Login Admin
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.User	true	"getting"
//	@Success		200		{object}	models.Token
//	@Router			/auth/login-admin [post]
func (a *AuthHandler) LoginAdmin() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			slog.Warn("error to json", "error", err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
		defer cancel()

		res, err := a.AuthClient.LoginAdmin(ctx, &authv1.LoginRequest{
			Email:    user.Email,
			Password: user.Password,
		})

		if err != nil {
			slog.Error("error to auth", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.ResponseJSON(w, http.StatusOK,models.Token{Token: res.Token})
	}
}

// Register Register Admin
//
//	@Summary		Register Admin
//	@Description	Register Admin
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.User	true	"getting"
//	@Success		200		{object}	models.Token
//	@Router			/auth/register-admin [post]
func (a *AuthHandler) RegisterAdmin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			slog.Warn("error to json", "error", err)
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
		defer cancel()
		res, err := a.AuthClient.RegisterAdmin(ctx, &authv1.RegisterRequest{
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
		})

		if err != nil {
			slog.Error("error to atuh", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.ResponseJSON(w, http.StatusOK, map[string]string{"token": res.Token})

	}
}
// UserInfo user info
//
//	@Summary		user info
//	@Description	user info
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id"
//	@Success		200		{object}	models.UserInfoResponse
//	@Router			/auth/user-info/{id} [get]
func (a *AuthHandler) UserInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
		defer cancel()
		res, err := a.AuthClient.GetInfoUser(ctx, &authv1.UserInfoRequest{UserId: id})
		if err != nil {
			response.ResponseError(w, http.StatusBadRequest, err)
			slog.Warn("error to json", "error", err)
			return
		}
		response.ResponseJSON(w, http.StatusAccepted, res)
	}
}