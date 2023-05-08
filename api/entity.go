package api

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Nickname string             `json:"nickname" bson:"nickname"`
	Email    string             `json:"email" bson:"email"`
	Password []byte             `json:"password" bson:"password"`
}

type Todo struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	DueDate     time.Time          `json:"due_date" bson:"due_date"`
	AuthorID    string             `json:"author_id" bson:"author_id"`
}
