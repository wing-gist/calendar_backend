package api

import (
	"context"
	"encoding/json"
	"net/http"

	"calendar/database"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func UserGetHandler(w http.ResponseWriter, r *http.Request) {
	filter := bson.D{}

	Cursor, err := database.Find("users", filter)
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

	user.Passsword, _ = bcrypt.GenerateFromPassword(user.Passsword, bcrypt.DefaultCost)

	InsertOneResult, err := database.InsertOne("users", user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(InsertOneResult)
}

func ValidateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	filter := bson.D{{"email", user.Email}}

	SingleResult := database.FindOne("users", filter)

	var UserFromDB User
	SingleResult.Decode(&UserFromDB)

	err := bcrypt.CompareHashAndPassword(UserFromDB.Passsword, user.Passsword)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(UserFromDB)
}
