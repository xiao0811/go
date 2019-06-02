package main

import (
	"fmt"
	"time"
)

const (
	max     = 50000000
	block   = 500
	budSize = 100
)

func test() {
	done := make(chan struct{})
	c := make(chan int, budSize)

	go func() {
		//count := 0
		for value := range c {
			//count += value
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

func testBlock() {
	done := make(chan struct{})
	c := make(chan [block]int, budSize)

	go func() {
		//count := 0
		for value := range c {
			for _, val := range value {
				//count += val
				fmt.Println(val)
			}
		}

		close(done)
	}()

	for i := 0; i < max; i += block {
		var b [block]int
		for n := 0; n < block; n++ {
			b[n] = i + n
			if b[n] == max-1 {
				break
			}
		}

		c <- b
	}

	close(c)
	<-done
}

func main() {
	start := time.Now()
	testBlock()
	fmt.Println(time.Now().Sub(start))
}
