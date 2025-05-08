package models

import "time"

type CreateToDoReq struct {
	UserId      int    `json:"user_id"`
	Title       string `json:"titile"`
	Description string `json:"description"`
}

type Todo struct {
	ID          int
	UserID      int
	Title       string
	Description string
	IsCompleted bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UpdateToDoReq struct {
	ID          int
	UserId      int
	Title       string
	Description string
	IsCompleted bool
}

type DeleteToDoReq struct {
	ID     int
	UserID int 
}
