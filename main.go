package main

import (
	"fmt"
)

func getNonFibonacci(limit int) []int {
	fib := make(map[int]bool)
	nonFib := []int{}

	a, b := 0, 1
	for b <= limit {
		fib[b] = true
		a, b = b, a+b

		for num := a + 1; num < b; num++ {
			nonFib = append(nonFib, num)
		}
	}

	return nonFib
}

func main() {
	var limit int

	fmt.Print("Enter the limit: ")
	fmt.Scan(&limit)

	result := getNonFibonacci(limit)

	fmt.Printf("Numbers not in the Fibonacci series up to %d: %v\n", limit, result)
}
