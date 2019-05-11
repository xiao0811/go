package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	c := make(chan int)

	go func() {
		defer wg.Add(-1)

		for {
			var v int
			var ok bool

			select {
			case v, ok = <-c:
				println("a1", v)
			case v, ok = <-c:
				println("a2", v)
			}

			if !ok {
				return
			}
		}
	}()

	go func() {
		defer wg.Add(-1)
		defer close(c)

		for i := 0; i < 10; i++ {
			select {
			case c <- i:
			case c <- i * 10:
			}
		}
	}()

	wg.Wait()
}
