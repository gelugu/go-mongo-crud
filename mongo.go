package gomongocrud

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// InitDatabase initializes the MongoDB client with the given URI.
func InitDatabase(uri string) error {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	return nil
}

// GetDatabase returns a MongoDB database instance.
func GetDatabase(databaseName string) (*mongo.Database, error) {
	if client == nil {
		return nil, fmt.Errorf("MongoDB client is not initialized. Call InitDatabase first.")
	}
	return client.Database(databaseName), nil
}

// CloseDatabase closes the MongoDB connection.
func CloseDatabase() {
	if client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := client.Disconnect(ctx); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
	}
}
