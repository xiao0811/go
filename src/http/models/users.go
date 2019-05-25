package models

import (
	"http/database"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int8   `json:"role"`
}

func (user *User) GetUser() {
	database.Eloquent.Where("username = ?", user.Username).First(user)
}

func (user *User) CreateUser() {
	database.Eloquent.Create(user)
}
