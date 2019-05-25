package router

import (
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"http/controllers"
	"http/handlers"
)

func InitRoute() *iris.Application {
	app := iris.New()

	app.Get("/", controllers.Index)
	app.Post("/login", controllers.Login)
	app.Post("/createUser", controllers.CreateUser)
	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (i interface{}, e error) {
			return []byte(handlers.ValidationKeyGetter), nil
		},

		SigningMethod: jwt.SigningMethodHS256,
	})

	app.Use(jwtHandler.Serve)

	// 刷新 token
	app.Get("/refreshToken", controllers.RefreshToken)

	app.Get("/user", handlers.Checkjwt, controllers.GetUser)
	return app
}
