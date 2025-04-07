package db

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient wraps *mongo.Client for type safety
type MongoClient struct {
    Client *mongo.Client // Capitalize 'client' to 'Client' to make it exported
}

func ConnectMongo(uri, dbName string) (*MongoClient, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        return nil, err
    }

    log.Println("Connected to MongoDB successfully")
    return &MongoClient{Client: client}, nil
}

func (mc *MongoClient) GetUserByEmail(dbName, collection, email string) (map[string]interface{}, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var user map[string]interface{}
    err := mc.Client.Database(dbName).Collection(collection).FindOne(ctx, map[string]string{"email": email}).Decode(&user)
    if err != nil {
        return nil, err
    }

    return user, nil
}