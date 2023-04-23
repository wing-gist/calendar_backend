package api

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"calendar/database"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/google/uuid"
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

	preAuthToken := AuthtokenClaims{
		TokenUUID: uuid.NewString(),
		UserID:    UserFromDB.ID.String(),
		Nickname:  UserFromDB.Nickname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Minute * 15)),
		},
	}

	authToken := jwt.NewWithClaims(jwt.SigningMethodHS256, preAuthToken)
	signedAuthToken, err := authToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{Name: "MyJWT", Value: signedAuthToken, Expires: time.Now().Add(time.Minute * 15), MaxAge: 0, Secure: false, HttpOnly: false}
	http.SetCookie(w, &cookie)
	w.Write([]byte(signedAuthToken))
}

func ValidateJWT(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("MyJWT")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &AuthtokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if token.Valid {
		json.NewEncoder(w).Encode(token.Claims)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}
