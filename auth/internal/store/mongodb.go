package store

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Uri    string
	db     *mongo.Database
	clinet *mongo.Client
}

func NewMongoDB(uri, dbName string) (*MongoDB, error) {

	ClientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), ClientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	db := client.Database(dbName)

	return &MongoDB{
		Uri:    uri,
		db:     db,
		clinet: client,
	}, nil
}
