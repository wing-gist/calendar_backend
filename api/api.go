package api

import (
	"encoding/json"
	"net/http"

	"calendar/api/database"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

func MainPageHander(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func UserHander(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		client, ctx, cancel := database.Connect()
		defer cancel()

		Cursor, err := database.Find(client, "users", bson.D{{}}, ctx)
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
	case http.MethodPost:
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		client, ctx, cancel := database.Connect()
		defer cancel()

		InsertOneResult, err := database.InsertOne(client, "users", bson.D{{"Nickname", user.Nickname}, {"Email", user.Email}}, ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(InsertOneResult)
	}
}
