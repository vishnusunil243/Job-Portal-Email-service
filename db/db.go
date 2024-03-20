package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB(connectTo string) (*mongo.Database, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectTo))
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to check if the connection was successful
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	// Return the database object for the specified database name
	return client.Database("notifications"), nil
}
