package gomongocrud_test

import (
	"testing"

	"github.com/gelugu/go-mongo-crud"
)

func TestInitAndCloseDatabase(t *testing.T) {
	uri := "mongodb://localhost:27017"

	gomongocrud.InitDatabase(uri)
	defer gomongocrud.CloseDatabase()

	db, err := gomongocrud.GetDatabase("testdb")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if db == nil {
		t.Fatal("Expected database instance, got nil")
	}
}
