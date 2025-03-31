package models

import (
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var ()

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Username   string             `bson:"username,omitempty"`
	Password   string             `bson:"password,omitempty"`
	Email      string             `bson:"email,omitempty"`
	RegisterAt time.Time          `bson:"register_at,omitempty"`
}

func (u *User) IsValid() bool {
	email := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !email.MatchString(u.Email) {
		return false
	}
	if len(u.Password) < 3 {
		return false
	}
	username := regexp.MustCompile(`^[a-zA-Z0-9_]{3,16}$`)
	if !username.MatchString(u.Username) {
		return false
	}
	return true
}

func (u *User) CryptPassword() error {
	newPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}
	u.Password = string(newPassword)
	return nil
}

func (u *User) IsCorrectPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}
