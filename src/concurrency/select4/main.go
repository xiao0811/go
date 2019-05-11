package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)

		for {
			select {
			case x, ok := <-c:
				if !ok {
					return
				}
				fmt.Println("data:", x)
			// 当所有通道都不可用时
			// select会执行default语句
			default:
				fmt.Println("nil")
			}
			fmt.Println(time.Now())
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second * 5)
	c <- 100
	close(c)

	<-done
}
