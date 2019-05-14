package main

import "github.com/kataras/iris"

func main() {
	//app := gin.New()
	//app.Use(gin.Logger())
	//app.GET("/xiaosha", func(context *gin.Context) {
	//	context.JSON(200, gin.H{
	//		"name": "xiaosha",
	//		"age": 20,
	//	})
	//})
	//app.Run()

	app := iris.New()
	app.Logger()

	app.Get("/xiaosha", func(context iris.Context) {
		_, _ = context.JSON(iris.Map{
			"name": "xiaosha",
			"age": 20,
		})
	})

	app.Run(iris.Addr(":8080"))
}
