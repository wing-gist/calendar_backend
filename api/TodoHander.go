package api

import (
	"encoding/json"
	"net/http"

	"calendar/api/database"

	"go.mongodb.org/mongo-driver/bson"
)

func TodoGetHandler(w http.ResponseWriter, r *http.Request) {
	client, ctx, cancel := database.Connect()
	coll := database.GetCollection(client, "todos")
	defer cancel()

	Cursor, err := coll.Find(ctx, bson.D{{}})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var todos = []*Todo{}
	for Cursor.Next(ctx) {
		var todo Todo
		Cursor.Decode(&todo)
		todos = append(todos, &todo)
	}
	json.NewEncoder(w).Encode(todos)
}

func TodoPostHander(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)

	client, ctx, cancel := database.Connect()
	coll := database.GetCollection(client, "todos")
	defer cancel()

	InsertOneResult, err := coll.InsertOne(ctx, bson.D{{"Title", todo.Title}, {"Description", todo.Description}, {"Date", todo.Date}})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(InsertOneResult)
}
