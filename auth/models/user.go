package models

type User struct {
	Email    string
	Password string
}

func (u *User) IsValid() bool {
	return true
}

func (u *User) IsValidPassword() bool {
	return true
}