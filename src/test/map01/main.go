package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40}

	for key, value := range slice {
		fmt.Printf("Value: %d Value-Addr:%x ElemAddr:%x\n",
			value, &value, &slice[key])
	}
}
