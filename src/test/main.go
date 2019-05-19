package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	pas := sha256.Sum256([]byte("panxiao"))
	fmt.Printf("%x", pas)
}