package main

import (
	"flag"
	"fmt"
)

var (
	n   = flag.Bool("n", false, "omit trailing newline")
	sep = flag.String("s", "xiaosha", "separator")
)

func main() {
	fmt.Println(*sep)
}
