package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/likoscp/Advanced-Programming-2/auth/internal/store"
	"github.com/likoscp/Advanced-Programming-2/auth/models"
	"github.com/likoscp/Advanced-Programming-2/auth/utils"
	log "github.com/sirupsen/logrus"
)

type UserHandler struct {
	db *store.MongoDB
}

func NewUserHandler(db *store.MongoDB) *UserHandler {
	return &UserHandler{
		db: db,
	}
}

func (u *UserHandler) RegisterUser() http.HandlerFunc {
	type Request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type Response struct {
		Msg string `json:"msg"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := Request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			log.Warn("register, cannot read from json: ", err)
			return
		}
		user := models.User{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
		}

		if !user.IsValid() {
			log.Warn("incorrect user property")
			utils.Error(w, r, http.StatusBadRequest, errors.New("incorrect data"))
			return
		}
		if err := user.CryptPassword(); err != nil {
			utils.Error(w, r, http.StatusInternalServerError, err)
			log.Warn("cannot encrypt user: ", err)
			return
		}

		user.RegisterAt = time.Now()

		if err := u.db.UserRepo().Register(&user); err != nil {
			utils.Error(w, r, http.StatusInternalServerError, err)
			log.Error("cannot save user into db: ", err)
			return
		}

		res := Response{
			Msg: "user register succesfully",
		}
		utils.Response(w, r, http.StatusCreated, res)
		log.Info("handle users/register")
	}
}

func (u *UserHandler) Login() http.HandlerFunc {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type Res struct {
		Msg string `json:"msg"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := Request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			log.Warn("register, cannot read from json: ", err)
			return
		}

		u, err := u.db.UserRepo().Login(req.Email)
		if err != nil {
			utils.Response(w, r, http.StatusBadRequest, Res{Msg: "incorrect email or password"})
			log.Warn("login, no email:", err)
			return
		}

		if !u.IsCorrectPassword(req.Password) {
			utils.Response(w, r, http.StatusBadRequest, Res{Msg: "incorrect email or password"})
			log.Warn("login, incorrect password")
			return
		}

		utils.Response(w, r, http.StatusAccepted, Res{Msg: "login successfully"})
		log.Info("handle /users/login")
	}
}
