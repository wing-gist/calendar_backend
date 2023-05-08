package api

import "go.mongodb.org/mongo-driver/bson/primitive"

type SignIn struct {
	Nickname string `json:"nickname" bson:"nickname"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type Login struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type ReturnUser struct {
	Nickname string `json:"nickname" bson:"nickname"`
	Email    string `json:"email" bson:"email"`
}

type DeleteTodo struct {
	ID primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
}
