package models


import (
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)


type Admin struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Username   string             `bson:"username,omitempty"`
	Password   string             `bson:"password,omitempty"`
	Email      string             `bson:"email,omitempty"`
	RegisterAt time.Time          `bson:"register_at,omitempty"`
}

func (a *Admin) IsValid() bool {
	email := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !email.MatchString(a.Email) {
		return false
	}
	if len(a.Password) < 3 {
		return false
	}
	username := regexp.MustCompile(`^[a-zA-Z0-9_]{3,16}$`)
	if !username.MatchString(a.Username) {
		return false
	}
	return true
}

func (a *Admin) CryptPassword() error {
	newPassword, err := bcrypt.GenerateFromPassword([]byte(a.Password), 10)
	if err != nil {
		return err
	}
	a.Password = string(newPassword)
	return nil
}

func (a *Admin) IsCorrectPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password)) == nil
}
func (a Admin) Role() string {
	return "admin"
}
