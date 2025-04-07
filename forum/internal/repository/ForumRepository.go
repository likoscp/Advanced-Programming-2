package repository

import (
	"context"
	"github.com/likoscp/Advanced-Programming-2/forum/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ForumRepository struct {
	collection *mongo.Collection
}

func NewForumRepository(db *mongo.Database) *ForumRepository {
	return &ForumRepository{
		collection: db.Collection("threads"),
	}
}

func (r *ForumRepository) CreateThread(ctx context.Context, thread models.Thread) error {
	thread.CreatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, thread)
	return err
}

func (r *ForumRepository) GetThread(ctx context.Context, id string) (models.Thread, error) {
	var thread models.Thread
	objID, _ := primitive.ObjectIDFromHex(id)
	err := r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&thread)
	return thread, err
}


func (r *ForumRepository) DeleteThread(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func (r *ForumRepository) UpdateThread(ctx context.Context, id string, thread *models.Thread) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"content": thread.Content}}
	_, err = r.collection.UpdateOne(ctx, filter, update)
	return err
}
