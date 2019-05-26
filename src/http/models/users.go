package models

import (
	"http/database"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int8   `json:"role"`
	TokenUUID string `json:"token_uuid"`
}

func (user *User) GetUser() {
	database.Eloquent.Where("username = ?", user.Username).First(user)
}

func (user *User) CreateUser() {
	database.Eloquent.Create(user)
}

func (user *User) UpdateUser(key, value string) {
	database.Eloquent.Model(user).Update(key, value)
}
