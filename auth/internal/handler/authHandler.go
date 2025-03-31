package handler

import (
	"encoding/json"
	"net/http"

	"github.com/likoscp/Advanced-Programming-2/auth/internal/store"
	"github.com/likoscp/Advanced-Programming-2/auth/models"
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
	return func(w http.ResponseWriter, r *http.Request) {
		req := Request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("register, cannot read from json: ", err)
			return
		}
		user := models.User{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
		}

		if err := u.db.UserRepo().RegisterUser(&user); err != nil {
			log.Error("cannot save user into db: ", err)
			return
		}
		w.Write([]byte("register user"))
	}
}
