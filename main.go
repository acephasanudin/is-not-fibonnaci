package main

import (
	"fmt"
)

func isFibonacci(n int, fib map[int]bool) bool {
	_, found := fib[n]
	return found
}

func getNonFibonacci(limit int) []int {
	fib := make(map[int]bool)
	nonFib := []int{}

	a, b := 0, 1
	for b <= limit {
		fib[b] = true
		a, b = b, a+b
	}

	for num := 0; num <= limit; num++ {
		if !isFibonacci(num, fib) {
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
