//go:build ignore

package main

import "fmt"

// fib: iterative O(n) time, O(1) space. Roll two variables forward.
func fib(n int) int {
	if n < 2 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("fib(%d)=%d\n", i, fib(i))
	}
}
