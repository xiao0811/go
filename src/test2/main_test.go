package main

import (
	"fmt"
	"testing"
)

func TestPrint1to20(t *testing.T) {
	res := Print1to20()
	fmt.Println("hey")

	if res != 210 {
		t.Error("Wrong result of Print1to20")
	}
}
func TestPrint(t *testing.T) {
	t.Run("a1", func(t *testing.T) {
		fmt.Println("a1")
	})
	t.Run("a2", func(t *testing.T) {
		fmt.Println("a2")
	})
	t.Run("a3", func(t *testing.T) {
		fmt.Println("a3")
	})
}

func BenchmarkPrint1to20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Print1to20()
	}
}
