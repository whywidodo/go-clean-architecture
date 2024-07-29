package repositories

import (
	"context"
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	DB      *sql.DB
	Context context.Context
	MongoDB *mongo.Database
}

func NewRepository(conn *sql.DB, ctx context.Context, MongoDB *mongo.Database) Repository {
	return Repository{
		DB:      conn,
		Context: ctx,
		MongoDB: MongoDB,
	}
}
