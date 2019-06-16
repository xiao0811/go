package main

import (
	"log"
	"os"

	_ "sample/matchers"
	"sample/search"
)

// init is called prior to main.
// init 在 main之前调用
func init() {
	// Change the device for logging to stdout.
	// 将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
// main 整个程序的入口
func main() {
	// Perform the search for the specified term.
	// 使用特定的项做搜索
	search.Run("president")
}
