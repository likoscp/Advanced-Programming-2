package models

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

type IsAdmin struct {
	UserId string `json:"user-id"`
}

type IsReallyAdmin struct {
	IsAdmin bool `json:"is-admin"`
}