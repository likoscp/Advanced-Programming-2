package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Username   string             `bson:"username,omitempty"`
	Password   string             `bson:"password,omitempty"`
	Email      string             `bson:"email,omitempty"`
	RegisterAt time.Time          `bson:"register_at,omitempty"`
}

func (u *User) IsValid() bool {
	return true
}
