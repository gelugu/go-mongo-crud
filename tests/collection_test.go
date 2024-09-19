package gomongocrud_test

import (
	"errors"
	"testing"

	"github.com/gelugu/go-mongo-crud"
	"go.mongodb.org/mongo-driver/bson"
)

type TestModel struct {
	ID    string `bson:"_id,omitempty"`
	Name  string `bson:"name"`
	Value int    `bson:"value"`
}

func TestCollectionCRUD(t *testing.T) {
	uri := "mongodb://localhost:27017"
	err := gomongocrud.InitDatabase(uri)
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer gomongocrud.CloseDatabase()

	db, err := gomongocrud.GetDatabase("testdb")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	collection := gomongocrud.NewCollection[TestModel](db, "testcollection")

	// TODO: Clear collection before running tests

	// **Create**
	testItem := TestModel{Name: "TestName", Value: 42}
	insertedID, err := collection.Create(testItem)
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	// **Read**
	filter := bson.M{"_id": insertedID}
	readItem, err := collection.Read(filter)
	if err != nil {
		t.Fatalf("Read failed: %v", err)
	}
	if readItem.Name != testItem.Name || readItem.Value != testItem.Value {
		t.Error("Read item does not match created item")
	}

	// **Update**
	update := bson.M{"value": 100}
	err = collection.Update(filter, update)
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}

	// Проверяем обновление
	updatedItem, err := collection.Read(filter)
	if err != nil {
		t.Fatalf("Read after update failed: %v", err)
	}
	if updatedItem.Value != 100 {
		t.Error("Item value was not updated")
	}

	// **Delete**
	err = collection.Delete(filter)
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}

	// Проверяем удаление
	_, err = collection.Read(filter)
	if !errors.Is(err, gomongocrud.ErrNotFound) {
		t.Error("Expected ErrNotFound after delete")
	}
}
