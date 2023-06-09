package api

import (
	"calendar/database"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TodoGetHandler(w http.ResponseWriter, r *http.Request) {
	claim := r.Context().Value("user").(*AuthtokenClaims)
	if claim == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	filter := bson.D{{"author_id", claim.UserID}}
	Cursor, err := database.Find("todos", filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var todos []*Todo
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
	if todo.Title == "" || todo.Description == "" || time.Time.IsZero(todo.DueDate) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var claim = r.Context().Value("user").(*AuthtokenClaims)
	if claim == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	todo.AuthorID = claim.UserID

	InsertOneResult, err := database.InsertOne("todos", todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(InsertOneResult)
}

func TodoDeleteHander(w http.ResponseWriter, r *http.Request) {
	var deleteTodo DeleteTodo
	json.NewDecoder(r.Body).Decode(&deleteTodo)
	if deleteTodo.ID == primitive.NilObjectID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var claim = r.Context().Value("user").(*AuthtokenClaims)
	if claim == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	filter := bson.D{{"author_id", claim.UserID}, {"_id", deleteTodo.ID}}
	DeleteResult, err := database.DeleteOne("todos", filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if DeleteResult.DeletedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(DeleteResult)
}
