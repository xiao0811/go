package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	max     = 50000000
	block   = 500
	bufSize = 100
)
var wg sync.WaitGroup

func test() {
	defer wg.Add(-1)
	start := time.Now()
	done := make(chan struct{})
	c := make(chan int, bufSize)

	go func() {
		count := 0
		for x := range c {
			count += x
		}

		close(done)
	}()

	for i := 0; i < max; i++ {
		c <- i
	}

	close(c)
	<-done

	end := time.Now()
	fmt.Println("test:", end.Sub(start))
}

func testBlock() {
	defer wg.Add(-1)
	start := time.Now()
	done := make(chan struct{})
	c := make(chan [block]int, bufSize)

	go func() {
		count := 0
		for a := range c {
			for _, x := range a {
				count += x
			}
		}

		close(done)
	}()

	for i := 0; i < max; i += block {
		var b [block]int
		for n := 0; n < block; n++ {
			b[n] = i + n
			if i+n == max-1 {
				break
			}
		}
		c <- b
	}
	close(c)
	<-done

	end := time.Now()
	fmt.Println("testBlock:", end.Sub(start))
}

func main() {
	wg.Add(2)
	go test()
	go testBlock()
	wg.Wait()
}
