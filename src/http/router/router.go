package router

import (
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"http/controllers"
	"http/handlers"
	"net/http"
)

func InitRoute() *iris.Application {
	app := iris.New()

	app.Get("/", controllers.Index)
	app.Post("/login", controllers.Login)
	app.Post("/createUser", controllers.CreateUser)

	// 刷新 token
	app.Get("/refreshToken", controllers.RefreshToken)


	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (i interface{}, e error) {
			return []byte(handlers.ValidationKeyGetter), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		ErrorHandler: func(context iris.Context, s string) {
			context.StatusCode(http.StatusUnauthorized)
			context.JSON(iris.Map{
				"status" : 401,
				"message": "token失效或未登录",
			})
		},
	})

	app.Use(jwtHandler.Serve)


	app.Get("/user", controllers.GetUser)
	return app
}
