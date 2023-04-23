package api

import (
	"context"
	"encoding/json"
	"net/http"

	"calendar/database"

	"go.mongodb.org/mongo-driver/bson"
)

func TodoGetHandler(w http.ResponseWriter, r *http.Request) {
	coll := database.GetCollection("todos")

	Cursor, err := coll.Find(context.Background(), bson.D{{}})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var todos = []*Todo{}
	for Cursor.Next(context.Background()) {
		var todo Todo
		Cursor.Decode(&todo)
		todos = append(todos, &todo)
	}
	json.NewEncoder(w).Encode(todos)
}

func TodoPostHander(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)

	coll := database.GetCollection("todos")

	InsertOneResult, err := coll.InsertOne(context.Background(), bson.D{{Key: "Title", Value: todo.Title}, {Key: "Description", Value: todo.Description}, {Key: "Date", Value: todo.Date}})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(InsertOneResult)
}
