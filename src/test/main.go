package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	hash := sha256.Sum256([]byte("xiaosha"))
	res := fmt.Sprintf("%x", hash)

	fmt.Println(res)
}