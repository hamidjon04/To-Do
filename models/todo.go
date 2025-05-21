package models

import "time"

type CreateToDoReq struct {
	UserId      int    `json:"user_id"`
	Title       string `json:"titile"`
	Description string `json:"description"`
}

type Todo struct {
	ID          int       `json:"id"`
	UserId      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateToDoReq struct {
	ID          int    `josn:"id"`
	UserId      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}

type DeleteToDoReq struct {
	ID     int `josn:"id"`
	UserId int `json:"user_id"`
}
