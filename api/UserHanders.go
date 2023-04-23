package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"calendar/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UserGetHandler(w http.ResponseWriter, r *http.Request) {
	coll := database.GetCollection("users")
	start := time.Now()

	filter := bson.D{}
	opts := options.Find().SetProjection(bson.D{{Key: "Email", Value: 0}})
	Cursor, err := coll.Find(context.Background(), filter, opts)
	fmt.Println("Time to get data from database: ", time.Since(start))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var users = []*User{}
	for Cursor.Next(context.Background()) {
		var user User
		Cursor.Decode(&user)
		users = append(users, &user)
	}
	json.NewEncoder(w).Encode(users)
}

func UserPostHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	coll := database.GetCollection("users")

	InsertOneResult, err := coll.InsertOne(context.Background(), bson.D{{Key: "Nickname", Value: user.Nickname}, {Key: "Email", Value: user.Email}})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(InsertOneResult)
}
