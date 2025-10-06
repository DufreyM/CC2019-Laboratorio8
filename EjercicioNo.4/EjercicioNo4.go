package main

import (
	"fmt"
	"time"
)

func linearSearch(arr []int, x int) int {
	for i, v := range arr {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	sizes := []int{10, 100, 1000, 10000, 100000, 1000000}

	fmt.Printf("%-10s %-15s %-15s %-15s\n", "n", "Mejor (s)", "Promedio (s)", "Peor (s)")

	for _, n := range sizes {
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = i
		}

		// --- Mejor caso
		start := time.Now()
		linearSearch(arr, arr[0])
		best := time.Since(start).Seconds()

		// --- Caso promedio
		start = time.Now()
		linearSearch(arr, arr[n/2])
		average := time.Since(start).Seconds()

		// --- Peor caso
		start = time.Now()
		linearSearch(arr, arr[n-1])
		worst := time.Since(start).Seconds()

		fmt.Printf("%-10d %-15f %-15f %-15f\n", n, best, average, worst)
	}
}
