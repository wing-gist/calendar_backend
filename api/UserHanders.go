package api

import (
	"context"
	"encoding/json"
	"net/http"

	"calendar/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func UserGetHandler(w http.ResponseWriter, r *http.Request) {
	filter := bson.D{}
	opts := options.Find().SetProjection(bson.D{{"password", 0}, {"_id", 0}})

	Cursor, err := database.Find("users", filter, opts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var users = []*ReturnUser{}
	for Cursor.Next(context.Background()) {
		var user ReturnUser
		Cursor.Decode(&user)
		users = append(users, &user)
	}
	json.NewEncoder(w).Encode(users)
}

func UserPostHandler(w http.ResponseWriter, r *http.Request) {
	var signIn SignIn
	json.NewDecoder(r.Body).Decode(&signIn)
	if signIn.Email == "" || signIn.Nickname == "" || signIn.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := User{
		Nickname: signIn.Nickname,
		Email:    signIn.Email,
	}
	user.Password, _ = bcrypt.GenerateFromPassword([]byte(signIn.Password), bcrypt.DefaultCost)

	InsertOneResult, err := database.InsertOne("users", user)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	json.NewEncoder(w).Encode(InsertOneResult)
}

func UserDeleteHander(w http.ResponseWriter, r *http.Request) {
	var login Login
	json.NewDecoder(r.Body).Decode(&login)

	filter := bson.D{{"email", login.Email}}

	SingleResult := database.FindOne("users", filter)

	var UserFromDB User
	SingleResult.Decode(&UserFromDB)

	err := bcrypt.CompareHashAndPassword(UserFromDB.Password, []byte(login.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	TodoFilter := bson.D{{"author_id", UserFromDB.ID}}
	database.DeleteMany("todos", TodoFilter)

	DeleteResult, err := database.DeleteOne("users", filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(DeleteResult)
}
