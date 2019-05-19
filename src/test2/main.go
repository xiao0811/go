package main

import (
	"io"
	"net/http"
)

func Print1to20() int {
	res := 0
	for i := 1; i < 21; i++ {
		res += i
	}
	return res
}

func firstPage(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "<p>Hello, xiaosha</p>")
}

func main() {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":8080", nil)
}
