package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"http/handlers"
	"http/models"
	"net/http"
	"strings"
)

func Index(ctx iris.Context) {
	ctx.HTML(handlers.NewUUID())
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
	userMsg := ctx.Values().Get("jwt").(*jwt.Token)
	userInfo := userMsg.Claims.(jwt.MapClaims)
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
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.JSON(iris.Map{
			"status":  http.StatusUnauthorized,
			"message": "没有token",
		})
		return
	}

	// TODO: Make this a bit more robust, parsing-wise
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.JSON(iris.Map{
			"status":  http.StatusUnauthorized,
			"message": "Authorization header format must be Bearer {token}",
		})
		return
	}


	parsedToken, _ := jwt.Parse(authHeaderParts[1], func(token *jwt.Token) (i interface{}, e error) {
		return []byte(handlers.ValidationKeyGetter), nil
	})

	data := parsedToken.Claims.(jwt.MapClaims)
	if data["username"] == nil {
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.JSON(iris.Map{
			"status":  http.StatusUnauthorized,
			"message": "Toke is error",
		})
		return
	}

	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       data["id"],
		"username": data["username"],
		"iss":      data["iss"],
		"iat":      data["iat"],
		"jti":      data["jti"],
		"exp":      data["exp"],
	})
	tokenString, _ := _token.SignedString([]byte(handlers.ValidationKeyGetter))
	if "Bearer "+tokenString != authHeader {
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.JSON(iris.Map{
			"status":  http.StatusUnauthorized,
			"message": "Toke is error",
		})
		return
	}
	token := handlers.NewToken(data["username"].(string), int(data["id"].(float64)))
	ctx.JSON(iris.Map{
		"status": http.StatusOK,
		"token":  token,
	})
}
