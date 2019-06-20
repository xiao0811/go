package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	app := gin.Default()

	app.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"name": "xiaosha",
			"age":  20,
		})
	})
	app.POST("/abc", func(context *gin.Context) {
		name, _ := context.GetPostForm("name")
		context.JSON(http.StatusOK, gin.H{
			"abc": context.Query("abc"),
			"bbb": name,
		})
	})
	app.Run(":9090")
}
