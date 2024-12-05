package app

import (
	"database/sql"
	"go-clean-architecture/repositories"
	"go-clean-architecture/services"
)

func SetupApp(DB *sql.DB, repo repositories.Repository) services.UsecaseService {
	usecaseSvc := services.NewUsecaseService(DB)

	return usecaseSvc
}
