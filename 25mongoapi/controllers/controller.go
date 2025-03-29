package controller

import (
	"context"
	"fmt"
	model "github.com/madhavarora03/mongoapi/models"
	"go.mongodb.org/mongo-driver/v2/bson"
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

// update 1 record
func updateOneMovie(movieId string) {
	id, _ := bson.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count:", result.ModifiedCount)
}

// delete 1 record
func deleteOneMovie(movieId string) {
	id, _ := bson.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie got deleted with delete count:", deleteCount)
}

// delete all records from mongodb
func deleteAllMovies() {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of movies deleted:", deleteResult.DeletedCount)
}

// get all movies from db
func getAllMovies() []bson.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []bson.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())
	return movies
}
