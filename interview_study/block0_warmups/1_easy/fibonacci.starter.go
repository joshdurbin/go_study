//go:build ignore

package main

import "fmt"

// fib returns the n-th Fibonacci number. F(0)=0, F(1)=1, F(n)=F(n-1)+F(n-2).
// Target O(n) time and O(1) space.
func fib(n int) int {
	// TODO: implement
	return 0
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("fib(%d)=%d\n", i, fib(i))
	}
	// expect: 0 1 1 2 3 5 8 13 21 34
}
