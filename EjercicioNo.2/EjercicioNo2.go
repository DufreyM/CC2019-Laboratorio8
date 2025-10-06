package main

import (
	"fmt"
	"time"
)

func function(n int) {
	if n <= 1 {
		return
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			// Simula el printf
			_ = "Sequence"
			break
		}
	}
}

func main() {
	inputs := []int{1, 10, 100, 1000, 10000, 100000, 1000000}
	fmt.Printf("%-10s %-15s\n", "n", "Tiempo (s)")

	for _, n := range inputs {
		start := time.Now()
		function(n)
		elapsed := time.Since(start).Seconds()
		fmt.Printf("%-10d %-15f\n", n, elapsed)
	}
}
