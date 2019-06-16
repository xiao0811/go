package main

import "fmt"

func main() {
	testMap := map[string]struct {
		Name string `json:"name"`
	}{
		"xiaosha":  {"xiaosha"},
		"xiaozang": {"xiaozang"},
	}

	name, exists := testMap["xiaozang"]
	fmt.Println(name, exists)
}
