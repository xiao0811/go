package main

import "fmt"

func main() {
	map1 := make(map[string]struct{
		Name string
		Age int
	})

	map1["aa"] = struct {
		Name string
		Age  int
	}{Name: "xiaosha", Age: 20}

	fmt.Println(map1)
}



































