package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

func main() {
	n := runtime.GOMAXPROCS(0)
	//test(n)
	test2(n)
}

func test(n int) {
	start := time.Now()
	for i := 0; i < n; i++ {
		count()
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func test2(n int) {
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			count()
			wg.Add(-1)
		}()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func count() {
	x := 0
	for i := 0; i < math.MaxUint32; i++ {
		x++
	}

	fmt.Println(x)
}
