package api

import (
	"calendar/database"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type AuthtokenClaims struct {
	TokenUUID string `json:"token_uuid"`
	UserID    string `json:"user_id"`
	Nickname  string `json:"nickname"`
	jwt.StandardClaims
}

func ValidateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	if user.Email == "" || user.Passsword == nil {
		w.WriteHeader(http.StatusBadRequest)
	}

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

	cookie := http.Cookie{Name: "calendar_JWT", Value: signedAuthToken, Expires: time.Now().Add(time.Minute * 15), MaxAge: 0, Secure: false, HttpOnly: false}
	http.SetCookie(w, &cookie)
}

func ValidateJWT(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("calendar_JWT")
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
