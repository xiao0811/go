package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))

	fmt.Println(c)
}
/**
package io

// Writer 接口封装了基础的写入方法
type Writer interface {
	// Write 从p 向底层数据流写入len(p) 个字节数据
	// 返回实际写入的字节数(0 <= n <= len(p))
	// 如果没有写完, 那么会返回遇到的错误
	// 在Write 返回 n < len(p) 时, err 必须为非nil
	// Write 不允许修改p 的数据, 即使是临时修改
	//
	// 实现时不允许残留 p 的引用
	Write(p []byte) (n int, err error)
}
 */