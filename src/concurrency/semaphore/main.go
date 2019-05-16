package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup

	sem := make(chan struct{}, 2)

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Add(-1)

			sem <- struct{}{}
			defer func() { <-sem }()

			time.Sleep(time.Second * 2)
			fmt.Println(id, time.Now())
		}(i)
	}

	wg.Wait()
}