package main

import (
	"github.com/kataras/iris"
)

type Student struct {
	Name string `json:"name"`
	Age  int `json:"age"`
}

//var students map[string]*Student

func main() {
	app := iris.New()
	students := make(map[string]*Student)
	students["no1"] = &Student{Name:"xiaosha"}
	students["no2"] = &Student{Name:"xiaozang"}
	students["no3"] = &Student{Name:"xiaoxiong"}



	app.Get("/", func(context iris.Context) {
		res := iris.Map{}
		for key, value := range students {
			res[key] = *value
		}
		context.JSON(res)
	})
	app.Run(iris.Addr(":9099"))
	//fmt.Println(students)

}
