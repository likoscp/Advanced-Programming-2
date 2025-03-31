package store

import (
	"context"

	"github.com/likoscp/Advanced-Programming-2/auth/internal/config"
	"github.com/likoscp/Advanced-Programming-2/auth/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	uri            string
	config         config.Config
	userRepository *repository.UserRepository
	db             *mongo.Database
	clinet         *mongo.Client
}

func NewMongoDB(config *config.Config) (*MongoDB, error) {
	ClientOptions := options.Client().ApplyURI(config.MongoUri)
	client, err := mongo.Connect(context.TODO(), ClientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	db := client.Database(config.DBname)

	return &MongoDB{
		uri:    config.MongoUri,
		db:     db,
		clinet: client,
	}, nil
}

func (m *MongoDB) UserRepo() *repository.UserRepository {
	if m.userRepository == nil {
		m.userRepository = &repository.UserRepository{
			Collection: m.db.Collection(m.config.CollectionName),
		}
	}
	return m.userRepository
}
