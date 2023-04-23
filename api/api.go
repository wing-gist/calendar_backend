package api

import (
	"net/http"
)

func MainPageHander(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func UserHander(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		UserGetHandler(w, r)
	case http.MethodPost:
		UserPostHandler(w, r)
	}
}

func AuthHander(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		ValidateUser(w, r)
	}
}

func TodoHander(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		TodoGetHandler(w, r)
	case http.MethodPost:
		TodoPostHander(w, r)
	}
}
