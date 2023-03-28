package api

import (
	"net/http"
	"encoding/json"
)

type User struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

var users = map[string]*User{}

func MainPageHander (w http.ResponseWriter, r * http.Request) {
	w.Write([]byte("Hello World!"))
}

func UserHander (w http.ResponseWriter, r * http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(users)
	case http.MethodPost:
		var user User
		json.NewDecoder(r.Body).Decode(&user)

		users[user.Email] = &user

		json.NewEncoder(w).Encode(user)
	}
}