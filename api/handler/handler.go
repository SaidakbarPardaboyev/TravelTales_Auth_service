package handler

import (
	"database/sql"
	"log/slog"
	"travel/pkg/logger"
	"travel/storage/postgres"
)

type Handler struct {
	Logger   *slog.Logger
	AuthRepo *postgres.AuthRepo
}

func NewHandler(db *sql.DB) *Handler {
	logger := logger.NewLogger()
	auth := postgres.NewAuthRepo(db)
	return &Handler{
		Logger:   logger,
		AuthRepo: auth,
	}
}
