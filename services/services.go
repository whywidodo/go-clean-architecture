package services

import "database/sql"

type UsecaseService struct {
	RepoDB *sql.DB
}

func NewUsecaseService(
	repoDB *sql.DB,
) UsecaseService {
	return UsecaseService{
		RepoDB: repoDB,
	}
}
