package main

import (
	"fmt"
	"os"
)

type A interface {
}

type Cat struct {
	name string
	age  uint8
}

type Person struct {
	name string
	sex  string
}

func test1(a A) {

}

func test2(a interface{}) {

}

func test3(slice2 []interface{}) {
	for i := 0; i < (len(slice2)); i++ {
		fmt.Printf("第%d个数据\n", i+1)
		switch ins := slice2[i].(type) {
		case Cat:
			fmt.Println("\tCat对象:", ins.name, ins.age)
		case Person:
			fmt.Println("\tPerson对象:", ins.name, ins.sex)
		case int:
			fmt.Println("\tInt类型:", ins)
		case string:
			fmt.Println("\tString类型", ins)
		}
	}
}

func main() {
	fmt.Println(os.Args[0])
	a1 := Cat{"花猫", 1}
	a2 := Person{"王二狗", "men"}
	a3 := "xiaosha"
	a4 := 100

	test1(3.14)
	test1(a1)

	map1 := make(map[string]interface{})
	map1["name"] = "王二狗"
	map1["age"] = 10
	fmt.Println(map1)

	fmt.Println(100, "xiaosha", "hello", Cat{"小飞", 20})

	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1, a1, a2, a3, a4)

	fmt.Println(slice1)
	test3(slice1)
}
