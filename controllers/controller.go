package controllers


import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mongo_api/model"
)

const connectionString string = "mongodb+srv://mern_websites:lP3gKsjzyDYiFqkV@cluster0.oqerjlv.mongodb.net/"
const dbName string = "netflix"
const collectionName string = "watchlists"

var collection *mongo.Collection

func init() {
	// client setup for the connection string
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to the mongodb with options (connection string)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	// When you're confident that no special context features
	// (like cancellation or timeouts) are needed or you haven't yet
	// extended the surrounding function to accept a context parameter.
	// then use context.Background()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection instance/ reference created for export")

}


func createOneMovie (movie model.Netflix){
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted one movie", inserted)
}

