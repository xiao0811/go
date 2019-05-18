package main

import "fmt"

func main() {
	done := make(chan struct{})
	c := make(chan string)

	go func() {
		s := <-c
		fmt.Println(s)
		close(done)
	}()

	c <- "nibaba"
	<-done
}
