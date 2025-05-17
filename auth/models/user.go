package models

import (
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string
	Password string `validate:"required,email"`
	Email    string `validate:"required,min=5,max=25"`
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
