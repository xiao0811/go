package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
		fmt.Println("xiaosha")

	end := time.Now()
	fmt.Println(end.Sub(start))
}
