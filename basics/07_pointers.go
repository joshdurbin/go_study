//go:build ignore

package main

import "fmt"

func increment(n *int) {
	*n++
}

func newInt(v int) *int {
	// Returning a pointer to a local variable is safe in Go —
	// the compiler moves it to the heap automatically (escape analysis).
	return &v
}

func main() {
	x := 10
	p := &x          // p holds the address of x
	fmt.Println(p)   // address, e.g. 0xc000014090
	fmt.Println(*p)  // dereference: 10

	*p = 99
	fmt.Println(x)   // 99 — x was changed through the pointer

	increment(&x)
	fmt.Println(x)   // 100

	// new() allocates zeroed memory and returns a pointer
	q := new(int)
	fmt.Println(*q)  // 0
	*q = 42
	fmt.Println(*q)  // 42

	ptr := newInt(7)
	fmt.Println(*ptr) // 7

	// Go has no pointer arithmetic — intentional safety choice
}
