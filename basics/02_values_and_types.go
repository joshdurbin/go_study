//go:build ignore

package main

import "fmt"

func main() {
	// Basic types
	var b bool = true
	var i int = 42
	var f float64 = 3.14
	var s string = "gopher"

	fmt.Println(b, i, f, s)

	// Type inference with :=
	x := 100
	y := "inferred"
	fmt.Println(x, y)

	// Zero values (what you get without initialization)
	var zeroInt int
	var zeroBool bool
	var zeroStr string
	fmt.Printf("zeros: %d %v %q\n", zeroInt, zeroBool, zeroStr)

	// Constants
	const Pi = 3.14159
	const Greeting = "hello"
	fmt.Println(Pi, Greeting)

	// iota for enumerated constants
	const (
		Sunday = iota // 0
		Monday        // 1
		Tuesday       // 2
	)
	fmt.Println(Sunday, Monday, Tuesday)
}
