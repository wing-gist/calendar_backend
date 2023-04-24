package main

import (
	"calendar/api"
	"fmt"
	"net/http"
)

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mainpage := http.HandlerFunc(api.MainPageHander)
	userHander := api.ValidateJWTGaurd(http.HandlerFunc(api.UserHander))
	authHander := http.HandlerFunc(api.AuthHander)
	todoHander := api.ValidateJWTGaurd(http.HandlerFunc(api.TodoHander))

	mux.Handle("/", mainpage)
	mux.Handle("/users", jsonContentTypeMiddleware(userHander))
	mux.Handle("/auth", jsonContentTypeMiddleware(authHander))
	mux.Handle("/todos", jsonContentTypeMiddleware(todoHander))

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", mux)
}
