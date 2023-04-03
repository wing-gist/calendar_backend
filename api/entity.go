package api

import (
	"time"
)

type User struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

type Todo struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type UserDatabase struct {
	_id      string
	Nickname string
	Email    string
}

type TodoDatabase struct {
	_id         string
	Title       string
	Description string
	Date        time.Time
	author_id   string
}
