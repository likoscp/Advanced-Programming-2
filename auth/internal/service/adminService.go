package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/likoscp/Advanced-Programming-2/auth/internal/config"
	"github.com/likoscp/Advanced-Programming-2/auth/internal/store"
	"github.com/likoscp/Advanced-Programming-2/auth/models"
	log "github.com/sirupsen/logrus"
)

type AdminService struct {
	db     *store.MongoDB
	config *config.Config
}

func NewAdminrService(db *store.MongoDB, config *config.Config) *AdminService {
	return &AdminService{
		db:     db,
		config: config,
	}
}

func (us *AdminService) Register(admin *models.Admin) (string, *http.Cookie, error) {
	// validation part
	if !admin.IsValid() {
		log.Warn("incorrect user property")
		return "", nil, ErrIncorrectData
	}
	if err := admin.CryptPassword(); err != nil {
		log.Warn("cannot encrypt user: ", err)
		return "", nil, ErrEncrypt
	}

	admin.RegisterAt = time.Now()

	// create part
	if err := us.db.AdminRepo().Register(admin); err != nil {
		log.Error("cannot save user into db: ", err)
		return "", nil, err
	}

	// sending email part
	type Req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	req := Req{
		Username: admin.Username,
		Email:    admin.Email,
	}
	resB2, _ := json.Marshal(&req)
	resp, err := http.Post(us.config.MailURI+"/mail/register", "application/json", bytes.NewBuffer(resB2))

	if err != nil || resp.StatusCode != http.StatusOK {
		log.Warn(err)
		return "", nil, ErrRequestEmail
	}

	// jwt token part
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"_id":  admin.ID,
			"role": admin.Role(),
			"exp":  time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(us.config.SECRET))
	if err != nil {
		log.Error("cannot save cookie: ", err)
		return "", nil, ErrJWT
	}

	cookie := http.Cookie{Name: "token", Value: tokenString}

	return "admin register successfully", &cookie, nil
}
