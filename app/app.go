package app

import (
	"database/sql"
	"go-clean-arhitecture/repositories"
	"go-clean-arhitecture/services"
)

func SetupApp(DB *sql.DB, repo repositories.Repository) services.UsecaseService {
	usecaseSvc := services.NewUsecaseService(DB)

	return usecaseSvc
}
