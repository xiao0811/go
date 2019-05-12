package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.New()

	app.GET("/xiaosha", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"name": "xiaosha",
			"age": 20,
		})
	})
	app.Run()
}
