package handler

import (
	"add/config"
	"add/storage"
)

type Handler struct {
	cfg     config.Config
	storage storage.StorageI
}

func NewHandler(cfg config.Config, storage storage.StorageI) *Handler {
	return &Handler{
		cfg:     cfg,
		storage: storage,
	}
}
