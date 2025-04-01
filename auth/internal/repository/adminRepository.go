package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/likoscp/Advanced-Programming-2/auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminRepository struct {
	Collection *mongo.Collection
}

func (u *AdminRepository) Register(admin *models.Admin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := u.Collection.FindOne(ctx, bson.M{"email": admin.Email}).Decode(admin)

	if err == nil {
		return fmt.Errorf("admin with email exist already")
	}

	_, err = u.Collection.InsertOne(ctx, admin)
	if err != nil {
		return err
	}

	return nil
}

func (u *AdminRepository) Login(email string) (*models.Admin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var u1 models.Admin
	err := u.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&u1)
	if mongo.ErrNoDocuments == err {
		return nil, fmt.Errorf("no email with this user")
	}

	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}

	return &u1, nil
}
