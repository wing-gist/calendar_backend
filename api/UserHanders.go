package api

import (
	"encoding/json"
	"net/http"

	"calendar/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UserGetHandler(w http.ResponseWriter, r *http.Request) {
	client, ctx, cancel := database.Connect()
	coll := database.GetCollection(client, "users")
	defer cancel()

	filter := bson.D{}
	opts := options.Find().SetProjection(bson.D{{Key: "Email", Value: 0}})
	Cursor, err := coll.Find(ctx, filter, opts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var users = []*User{}
	for Cursor.Next(ctx) {
		var user User
		Cursor.Decode(&user)
		users = append(users, &user)
	}
	json.NewEncoder(w).Encode(users)
}

func UserPostHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	client, ctx, cancel := database.Connect()
	coll := database.GetCollection(client, "users")
	defer cancel()

	InsertOneResult, err := coll.InsertOne(ctx, bson.D{{Key: "Nickname", Value: user.Nickname}, {Key: "Email", Value: user.Email}})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(InsertOneResult)
}
