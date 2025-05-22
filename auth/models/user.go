package models

import (
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string
	Username string `validate:"required,min=5,max=35"`
	Password string `validate:"required,min=5,max=25"`
	Email    string `validate:"required,email"`
}

func (u *User) IsValid() error {
	validate := validator.New()

	err := validate.Struct(u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) HashPassword() error {

	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) ComparePassword(u2 User) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u2.Password), []byte(u.Password))
    return err == nil
}