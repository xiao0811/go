package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	jwtmiddlerware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"time"
)

func myHandler(ctx iris.Context) {
	token := ctx.Values().Get("jwt").(*jwt.Token)

	ctx.Writef("This is an authenticated request\n")
	ctx.Writef("Claim content:\n")

	ctx.Writef("%s", token.Signature)

	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)

	fmt.Println(userMsg["id"])
	fmt.Println(userMsg["nick_name"])
	fmt.Println(userMsg["exp"])
}

func main() {
	app := iris.New()

	jwtHandler := jwtmiddlerware.New(jwtmiddlerware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (i interface{}, e error) {
			return []byte("My Secret"), nil
		},

		SigningMethod: jwt.SigningMethodHS256,
	})

	app.Get("login", LoginHandler)

	app.Use(jwtHandler.Serve)


	app.Get("ping", myHandler)

	app.Run(iris.Addr(":3001"))
}

func LoginHandler(ctx iris.Context)  {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nick_name": "iris",
		"email":     "go-iris@qq.com",
		"id":        1,
		"iss":       "Iris",
		"iat":       time.Now().Unix(),
		"jti":       "9527",
		"exp":       time.Now().Add(10 * time.Hour * time.Duration(1)).Unix(),
	})

	tokenString, _ := token.SignedString([]byte("My Secret"))

	ctx.JSON(iris.Map{
		"status": 200,
		"token": tokenString,
	})
}
