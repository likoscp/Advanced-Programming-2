package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (a *AuthHandler) Configure(mux *http.ServeMux) {
	mux.HandleFunc("POST /auth/login", a.Login())
	mux.HandleFunc("POST /auth/register", a.Register())
	// mux.HandleFunc("GET /auth/is-admin", nil)
	// mux.HandleFunc("POST /auth/login-admin", nil)
	// mux.HandleFunc("POST /auth/register-admin", nil)
}

func (a *AuthHandler) Login() http.HandlerFunc {
	type User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		user := User{}

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			slog.Warn("error to json", "error", err)
			return
		}

	}
}

func (a *AuthHandler) Register() http.HandlerFunc {
	type User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		user := User{}

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			slog.Warn("error to json", "error", err)
			return
		}

	}
}
