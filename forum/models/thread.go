package models

import "time"

type Reply struct {
	ID        string    `bson:"id" json:"id"`
	Author    string    `bson:"author" json:"author"`
	Content   string    `bson:"content" json:"content"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

type Thread struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	Author    string    `bson:"author" json:"author"`
	Content   string    `bson:"content" json:"content"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	Replies   []Reply   `bson:"replies" json:"replies"`
}
