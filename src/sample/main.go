package main

import (
	"crypto/md5"
	"fmt"
	"strconv"

	_ "sample/matchers"
)

// init is called prior to main.
// init 在 main之前调用
//func init() {
//	// Change the device for logging to stdout.
//	// 将日志输出到标准输出
//	//log.SetOutput(os.Stdout)
//}

// main is the entry point for the program.
// main 整个程序的入口
func main() {
	// Perform the search for the specified term.
	// 使用特定的项做搜索
	aims := "1104959d53dc3b60f2d40cd4a47d79e7"
	for i := 100000; i < 1000000; i++ {
		mima := fmt.Sprintf("%x", md5.Sum([]byte(strconv.Itoa(i))))
		if mima == aims {
			fmt.Println(i)
		}
	}
}
