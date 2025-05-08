package handler

import (
	postgres "todo/storage"
)

type Handler struct {
	Storage *postgres.DB
}

func NewHandler(storage *postgres.DB) *Handler {
	return &Handler{
		Storage: storage,
	}
}
