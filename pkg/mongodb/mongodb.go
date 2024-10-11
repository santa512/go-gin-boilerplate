package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Client

func InitMongoDB(uri string) {
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping the MongoDB primary to ensure connection is alive
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	MongoClient = client
	log.Println("Connected to MongoDB successfully")
}

func GetCollection(database, collection string) *mongo.Collection {
	return MongoClient.Database(database).Collection(collection)
}
