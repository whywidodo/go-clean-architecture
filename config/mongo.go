package config

import (
	"context"
	"fmt"
	"go-clean-arhitecture/constants"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(ctx context.Context, DBCollection ...string) *mongo.Database {
	connection := fmt.Sprintf("mongodb://%s:%s", MONGOHost, MONGOPort)
	fmt.Println("Connection Mongo: ", connection)

	clientOptions := options.Client()
	clientOptions.ApplyURI(connection)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil
	}

	col := MONGODb

	if len(DBCollection) > constants.EMPTY_VALUE_INT {
		col = DBCollection[constants.EMPTY_VALUE_INT]
	}

	return client.Database(col)
}
