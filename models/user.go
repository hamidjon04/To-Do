package models

import "time"

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string // hashed password
	CreatedAt time.Time
}

type CreateUserReq struct{
	Name string
	Email string
	Password string
}

type ErrorResponse struct {
	Error string `json:"error"` // Error message
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
