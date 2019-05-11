package main

import "runtime"

func main() {
	exit := make(chan struct{})

	go func() {
		defer close(exit)
		defer println("a")

		func() {
			defer func() {
				println("b", recover() == nil)
			}()

			func() {
				println("c")
				runtime.Goexit()
				println("c done.")
			}()
			println("b done.")
		}()

		println("a done")
	}()

	<-exit
	println("main exit.")
}
