package api

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Nickname  string             `json:"nickname" bson:"nickname"`
	Email     string             `json:"email" bson:"email"`
	Passsword []byte             `json:"password" bson:"password"`
}

type Todo struct {
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Date        time.Time `json:"date" bson:"date"`
	Author_id   string    `json:"author_id" bson:"author_id"`
}
