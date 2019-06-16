package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	fmt.Println("xiaosha")
	xiaozang()
}

func xiaozang()  {
	fmt.Println("xiaozang")
}
