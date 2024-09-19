package gomongocrud

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection represents a MongoDB collection for type T.
type Collection[T any] struct {
	collection *mongo.Collection
}

// ErrNotFound indicates that a document was not found.
var ErrNotFound = errors.New("document not found")

// NewCollection creates a new Collection instance.
func NewCollection[T any](db *mongo.Database, collectionName string) *Collection[T] {
	return &Collection[T]{
		collection: db.Collection(collectionName),
	}
}

// Create inserts a new item into the collection.
func (c *Collection[T]) Create(item T) (interface{}, error) {
	result, err := c.collection.InsertOne(context.Background(), item)
	if err != nil {
		return nil, fmt.Errorf("failed to insert item: %v", err)
	}
	return result.InsertedID, nil
}

// ReadAll retrieves all items matching the filter.
func (c *Collection[T]) ReadAll(filter bson.M, opts ...*options.FindOptions) ([]T, error) {
	cursor, err := c.collection.Find(context.Background(), filter, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to find items: %v", err)
	}
	defer cursor.Close(context.Background())

	var items []T
	if err = cursor.All(context.Background(), &items); err != nil {
		return nil, fmt.Errorf("failed to decode items: %v", err)
	}
	return items, nil
}

// Read retrieves a single item matching the filter.
func (c *Collection[T]) Read(filter bson.M) (*T, error) {
	var item T
	err := c.collection.FindOne(context.Background(), filter).Decode(&item)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, fmt.Errorf("failed to decode item: %v", err)
	}
	return &item, nil
}

// Update modifies an existing item matching the filter.
func (c *Collection[T]) Update(filter bson.M, data T) error {
	result, err := c.collection.UpdateOne(context.Background(), filter, bson.M{"$set": data})
	if err != nil {
		return fmt.Errorf("failed to update item: %v", err)
	}
	if result.MatchedCount == 0 {
		return ErrNotFound
	}
	return nil
}

// Delete removes an item matching the filter.
func (c *Collection[T]) Delete(filter bson.M) error {
	result, err := c.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete item: %v", err)
	}
	if result.DeletedCount == 0 {
		return ErrNotFound
	}
	return nil
}
