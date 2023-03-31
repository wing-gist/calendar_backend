package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Client, context.Context, context.CancelFunc) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_DB_CONNECTION")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client, ctx, cancel
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database(os.Getenv("DATABASE_NAME")).Collection(collectionName)
}

func InsertOne(client *mongo.Client, collectionName string, document interface{}, ctx context.Context) (*mongo.InsertOneResult, error) {
	collection := GetCollection(client, collectionName)

	return collection.InsertOne(ctx, document)
}

func FindOne(client *mongo.Client, collectionName string, filter interface{}, ctx context.Context) *mongo.SingleResult {
	collection := GetCollection(client, collectionName)

	return collection.FindOne(ctx, filter)
}

func Find(client *mongo.Client, collectionName string, filter interface{}, ctx context.Context) (*mongo.Cursor, error) {
	collection := GetCollection(client, collectionName)

	return collection.Find(ctx, filter)
}

func UpdateOne(client *mongo.Client, collectionName string, filter interface{}, update interface{}, ctx context.Context) (*mongo.UpdateResult, error) {
	collection := GetCollection(client, collectionName)

	return collection.UpdateOne(ctx, filter, update)
}

func DeleteOne(client *mongo.Client, collectionName string, filter interface{}, ctx context.Context) (*mongo.DeleteResult, error) {
	collection := GetCollection(client, collectionName)

	return collection.DeleteOne(ctx, filter)
}
