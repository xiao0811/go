package main

import (
	"github.com/kataras/iris"
	"http/router"
)

func main() {
	app := router.InitRoute()

	app.Run(iris.Addr(":9999/**/"))
}
