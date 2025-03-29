package controller

import (
	"context"
	"fmt"
	model "github.com/madhavarora03/mongoapi/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "netflix"
const collectionName = "watchlist"

var collection *mongo.Collection

// Connect wih db
func init() {
	// client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connected successfully!")

	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection instance is ready!")
}

// MONGODB helpers

// insert 1 record
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted one movie with id:", inserted.InsertedID)
}
