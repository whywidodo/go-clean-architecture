package config

import (
	"context"
	"fmt"
	"go-clean-architecture/constants"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(ctx context.Context, DBCollection ...string) *mongo.Database {
	connection := fmt.Sprintf("mongodb://%s:%s", MONGOHost, MONGOPort)
	fmt.Println("Connection Mongo: ", connection)

	clientOptions := options.Client().ApplyURI(connection)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return nil
	}

	col := MONGODb
	if len(DBCollection) > constants.EMPTY_VALUE_INT {
		col = DBCollection[constants.EMPTY_VALUE_INT]
	}

	return client.Database(col)
}
