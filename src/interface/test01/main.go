package main

import "fmt"

type tester interface {
	test()
	string() string
}

type data struct {}

func (*data) test() {

}

func (data) string() string {
	return ""
}

func main() {
	var d data
	//var t tester = d

	t := &d
	t.test()
	fmt.Println(t.string())
}
