package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/likoscp/Advanced-Programming-2/auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func (u *UserRepository) Register(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := u.Collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(user)

	if err == nil {
		return fmt.Errorf("user with email exist already")
	}

	_, err = u.Collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) Login(user *models.User) error {
	
	return nil
}
