package repository

import (
	"context"
	"time"

	"github.com/likoscp/Advanced-Programming-2/auth/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func (u *UserRepository) RegisterUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	_, err := u.Collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
