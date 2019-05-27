package models

type User struct {
	UserName string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type JwtToken struct {
	Token string `json:"token"`
}
