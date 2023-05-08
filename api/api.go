package api

import (
	"encoding/json"
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
	case http.MethodDelete:
		UserDeleteHander(w, r)
	}
}

func AuthHander(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		Claims := ValidateJWT(w, r)
		if Claims == nil {
			return
		}
		json.NewEncoder(w).Encode(Claims)
	case http.MethodPost:
		ValidateUser(w, r)
	case http.MethodDelete:
		Logout(w, r)
	}
}

func TodoHander(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		TodoGetHandler(w, r)
	case http.MethodPost:
		TodoPostHander(w, r)
	case http.MethodDelete:
		TodoDeleteHander(w, r)
	}
}
