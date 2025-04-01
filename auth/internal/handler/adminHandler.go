package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/likoscp/Advanced-Programming-2/auth/internal/config"
	"github.com/likoscp/Advanced-Programming-2/auth/internal/store"
	"github.com/likoscp/Advanced-Programming-2/auth/models"
	"github.com/likoscp/Advanced-Programming-2/auth/utils"
	log "github.com/sirupsen/logrus"
)

type AdminHandler struct {
	db     *store.MongoDB
	config *config.Config
}

func NewAdminHandler(db *store.MongoDB, config *config.Config) *AdminHandler {
	return &AdminHandler{
		db:     db,
		config: config,
	}
}

func (u *AdminHandler) Register() http.HandlerFunc {
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
		admin := models.Admin{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
		}

		if !admin.IsValid() {
			log.Warn("incorrect admin property")
			utils.Error(w, r, http.StatusBadRequest, errors.New("incorrect data"))
			return
		}
		if err := admin.CryptPassword(); err != nil {
			utils.Error(w, r, http.StatusInternalServerError, err)
			log.Warn("cannot encrypt user: ", err)
			return
		}

		admin.RegisterAt = time.Now()

		token := jwt.NewWithClaims(jwt.SigningMethodHS256,
			jwt.MapClaims{
				"_id":  admin.ID,
				"role": admin.Role(),
				"exp":  time.Now().Add(time.Hour * 24).Unix(),
			})

		tokenString, err := token.SignedString([]byte(u.config.SECRET))
		if err != nil {
			utils.Error(w, r, http.StatusInternalServerError, err)
			log.Error("cannot save cookie: ", err)
			return
		}

		cookie := http.Cookie{Name: "token", Value: tokenString}

		http.SetCookie(w, &cookie)

		if err := u.db.AdminRepo().Register(&admin); err != nil {
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

func (u *AdminHandler) Login() http.HandlerFunc {
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

		admin, err := u.db.AdminRepo().Login(req.Email)
		if err != nil {
			utils.Response(w, r, http.StatusBadRequest, Res{Msg: "incorrect email or password"})
			log.Warn("login, no email:", err)
			return
		}

		if !admin.IsCorrectPassword(req.Password) {
			utils.Response(w, r, http.StatusBadRequest, Res{Msg: "incorrect email or password"})
			log.Warn("login, incorrect password")
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256,
			jwt.MapClaims{
				"_id":  admin.ID,
				"role": admin.Role(),
				"exp":  time.Now().Add(time.Hour * 24).Unix(),
			})

		tokenString, err := token.SignedString([]byte(u.config.SECRET))
		if err != nil {
			utils.Error(w, r, http.StatusInternalServerError, err)
			log.Error("cannot save cookie: ", err)
			return
		}

		cookie := http.Cookie{Name: "token", Value: tokenString}

		http.SetCookie(w, &cookie)

		utils.Response(w, r, http.StatusAccepted, Res{Msg: "login successfully"})
		log.Info("handle /users/login")
	}
}
