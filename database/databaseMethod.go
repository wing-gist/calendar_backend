package database

import "go.mongodb.org/mongo-driver/mongo"

var client *mongo.Client

func init() {
	client, _, _ = Connect()
}

func GetCollection(collectionName string) *mongo.Collection {
	return client.Database("calendar").Collection(collectionName)
}
