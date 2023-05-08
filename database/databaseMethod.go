package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
	client, _, _ = Connect()
}

func GetCollection(collectionName string) *mongo.Collection {
	return client.Database("calendar").Collection(collectionName)
}

func Find(collectionName string, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	coll := GetCollection(collectionName)
	return coll.Find(context.TODO(), filter, opts...)
}

func FindOne(collectionName string, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	coll := GetCollection(collectionName)
	return coll.FindOne(context.TODO(), filter, opts...)
}

func InsertOne(collectionName string, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	coll := GetCollection(collectionName)
	return coll.InsertOne(context.TODO(), document, opts...)
}

func DeleteOne(collectionName string, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	coll := GetCollection(collectionName)
	return coll.DeleteOne(context.Background(), filter, opts...)
}

func DeleteMany(collectionName string, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	coll := GetCollection(collectionName)
	return coll.DeleteMany(context.Background(), filter, opts...)
}
