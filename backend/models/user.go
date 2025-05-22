package models

type User struct {
	Email    string `json:"email"`
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

type IsAdmin struct {
	UserId string `json:"user-id"`
}

type UserInfo struct {
	UserId string `json:"user-id"`
}

type UserInfoResponse struct {
	Email    string `json:"email"`
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}

type IsReallyAdmin struct {
	IsAdmin bool `json:"is-admin"`
}
