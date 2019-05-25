package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"http/handlers"
	"http/models"
	"net/http"
)

func Index(ctx iris.Context) {
	ctx.HTML("<p>xiaosha</p>")
}

func Login(ctx iris.Context) {
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")

	if username == "" || password == "" {
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.JSON(iris.Map{
			"status":  http.StatusUnauthorized,
			"message": "用户名/密码错误",
		})
		return
	}

	user := models.User{Username: username}
	user.GetUser()

	if user.ID == 0 || handlers.MYSHA256(password) != user.Password {
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.JSON(iris.Map{
			"status":  http.StatusUnauthorized,
			"message": "用户名/密码错误",
		})
		return
	}

	token := handlers.NewToken(user.Username, user.ID)

	ctx.JSON(iris.Map{
		"status": http.StatusOK,
		"token":  token,
	})
}

func GetUser(ctx iris.Context) {
	userInfo := ctx.Values().Get("userInfo")
	ctx.JSON(iris.Map{
		"status":  http.StatusOK,
		"message": "成功",
		"data":    userInfo,
	})
}

func CreateUser(ctx iris.Context) {
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")

	if username == "" || password == "" {
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.JSON(iris.Map{
			"status":  http.StatusUnauthorized,
			"message": "用户名/密码错误",
		})
		return
	}

	user := models.User{Username: username}
	user.GetUser()

	if user.ID != 0 {
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.JSON(iris.Map{
			"status":  http.StatusUnauthorized,
			"message": "用户名已存在",
		})
		return
	}

	user.Password = handlers.MYSHA256(password)
	user.CreateUser()

	ctx.JSON(iris.Map{
		"status":  http.StatusOK,
		"message": "成功",
		"data":    user,
	})
}

func RefreshToken(ctx iris.Context) {
	userInfo := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	username := userInfo["username"].(string)
	userId := int(userInfo["id"].(float64))

	token := handlers.NewToken(username, userId)
	ctx.JSON(iris.Map{
		"status":  http.StatusOK,
		"message": "成功",
		"token":   token,
	})
}
