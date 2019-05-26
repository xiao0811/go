package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"http/models"
	"time"
)

const ValidationKeyGetter = "Xiaosha"

func NewToken(username string, id int) string {
	uuid := NewUUID()
	user := models.User{Username:username, ID:id}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"id":       id,
		"iss":      "Xiaosha",
		"iat":      time.Now().Unix(),
		"jti":      uuid,
		"exp":      time.Now().Add(10 * time.Hour * time.Duration(1)).Unix(),
		//"exp": time.Now().Add(time.Second * 20 * time.Duration(1)).Unix(),
	})
	// 如果要限制登录 把UUID放入用户表中  再和表中token_uuid 对比
	user.TokenUUID = uuid
	user.UpdateUser("token_uuid", uuid)
	tokenString, _ := token.SignedString([]byte(ValidationKeyGetter))
	return tokenString
}
