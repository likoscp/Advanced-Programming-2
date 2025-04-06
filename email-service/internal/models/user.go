package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Username   string             `bson:"username,omitempty" json:"username"`
	Password   string             `bson:"password,omitempty"`
	Email      string             `bson:"email,omitempty" json:"email"`
	RegisterAt time.Time          `bson:"register_at,omitempty" json:"password"`
}
