package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	max = 500000
	bufSize = 100
)

func testWg() {
	var wg sync.WaitGroup

	for i := 0; i < max; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Add(-1)
		}(i)
	}

	wg.Wait()
}

func test() {
	done := make(chan struct{})
	c := make(chan int, bufSize)

	go func() {
		for value := range c {
			fmt.Println(value)
		}

		close(done)
	}()

	for i := 0; i < max; i++ {
		c <- i
	}

	close(c)
	<-done
}

func testNo() {
	for i := 0; i < max; i++ {
		fmt.Println(i)
	}
}

func main() {
	start := time.Now()

	//test()
	//testWg()
	testNo()

	end := time.Now()
	fmt.Println(end.Sub(start))
}