package main

import (
	"fmt"
	"time"
)

func function(n int) int {
	counter := 0
	for i := n / 2; i <= n; i++ {
		for j := 1; j+n/2 <= n; j++ {
			for k := 1; k < n; k = k * 2 {
				counter++
			}
		}
	}
	return counter
}

func main() {
	inputs := []int{1, 10, 100, 1000, 10000, 100000, 1000000}
	fmt.Printf("%-10s %-15s %-15s\n", "n", "Tiempo (s)", "Iteraciones")

	for _, n := range inputs {
		start := time.Now()
		counter := function(n)
		elapsed := time.Since(start).Seconds()
		fmt.Printf("%-10d %-15f %-15d\n", n, elapsed, counter)
	}
}
