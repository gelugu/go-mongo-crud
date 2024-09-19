# go-mongo-crud

`go-mongo-crud` is a simple Go library for performing CRUD (Create, Read, Update, Delete) operations with MongoDB using generics. It provides a convenient and type-safe way to interact with MongoDB collections in Go applications.

## Features

- Simplified CRUD operations using generics
- Easy initialization and connection management
- Supports custom data models
- Lightweight and easy to integrate

## Installation

To install the library, run:

```bash
Copy code
go get github.com/gelugu/go-mongo-crud
```

## Usage

### Initialization

First, initialize the MongoDB client and connect to your database:

```go
import (
    "github.com/gelugu/go-mongo-crud"
)

func main() {
    gomongocrud.InitDatabase("mongodb://localhost:27017")
    defer gomongocrud.CloseDatabase()

    db, err := gomongocrud.GetDatabase("your_database")
    if err != nil {
        // Handle error
    }
    // ...
}
```

### Defining Your Model
Create your data model as a Go struct. For example:

```go
type User struct {
    ID    primitive.ObjectID `bson:"_id,omitempty"`
    Name  string             `bson:"name"`
    Email string             `bson:"email"`
}
```

### Creating a Collection Instance

Create a collection instance for your model:

```go
collection := gomongocrud.NewCollection[User](db, "users")
```

### CRUD Operations

#### Create

Insert a new document into the collection:

```go
user := User{Name: "Alice", Email: "alice@example.com"}
insertedID, err := collection.Create(user)
if err != nil {
// Handle error
}
```

#### Read

Retrieve a single document:

```go
filter := bson.M{"email": "alice@example.com"}
foundUser, err := collection.Read(filter)
if err != nil {
    if errors.Is(err, gomongocrud.ErrNotFound) {
        // Document not found
    } else {
        // Handle other errors
    }
}
```

Retrieve multiple documents:

```go
filter := bson.M{"name": "Alice"}
users, err := collection.ReadAll(filter)
if err != nil {
    // Handle error
}
```

#### Update

Update an existing document:

```go
filter := bson.M{"email": "alice@example.com"}
update := bson.M{"email": "alice.new@example.com"}
err = collection.Update(filter, update)
if err != nil {
    // Handle error
}
```

#### Delete

Delete a document:

```go
filter := bson.M{"email": "alice.new@example.com"}
err = collection.Delete(filter)
if err != nil {
    // Handle error
}
```

## Testing
To run the tests execute:

```bash
make tests
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.
