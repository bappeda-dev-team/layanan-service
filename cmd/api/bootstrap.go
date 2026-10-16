package main

import (
	"database/sql"
	"log/slog"

	"github.com/bappeda-dev-team/layanan-service/internal/api"
)

// buildApplication adalah fungsi untuk injeksi repo dan service
// tambahkan service dan repository baru kesini
func buildApplication(
	cfg api.Config,
	logger *slog.Logger,
	db *sql.DB,
) *api.Application {

	// repository

	// service

	// application
	app := &api.Application{
		Config: cfg,
		Logger: logger,

	}

	return app
}
