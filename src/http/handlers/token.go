package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"net/http"
	"time"
)

const ValidationKeyGetter = "Xiaosha"

func NewToken(username string, id int) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"id":       id,
		"iss":      "Xiaosha",
		"iat":      time.Now().Unix(),
		"jti":      "9527",
		"exp":      time.Now().Add(10 * time.Hour * time.Duration(1)).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(ValidationKeyGetter))
	return tokenString
}

func Checkjwt(ctx iris.Context) {
	userMsg := ctx.Values().Get("jwt").(*jwt.Token)
	userInfo := userMsg.Claims.(jwt.MapClaims)
	ctx.Values().Set("userInfo", userInfo)
	exp := int64(userInfo["exp"].(float64))

	if exp < time.Now().Unix() {
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.JSON(iris.Map{
			"status":  http.StatusUnauthorized,
			"message": "token 过期",
		})
		return
	}
	ctx.Next()
}
