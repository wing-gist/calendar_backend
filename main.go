package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

var users = map[string]*User{}

func main() {
	mux := http.NewServeMux()

	mainpage := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	userHander := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			json.NewEncoder(w).Encode(users)
		case http.MethodPost:
			var user User
			json.NewDecoder(r.Body).Decode(&user)

			users[user.Email] = &user

			json.NewEncoder(w).Encode(user)
		}
	})

	mux.Handle("/", mainpage);
	mux.Handle("/users", jsonContentTypeMiddleware(userHander));
	http.ListenAndServe(":8080", mux)
}
