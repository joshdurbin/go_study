//go:build ignore

package main

import "fmt"

// Package-level variables (no := allowed here)
var globalMsg = "I am global"

func main() {
	// var declaration forms
	var a int = 10
	var b = 20       // type inferred
	c := 30          // short declaration (most common inside functions)

	fmt.Println(a, b, c)

	// Multi-variable short declaration
	x, y := 1, 2
	fmt.Println(x, y)

	// Swap without temp variable
	x, y = y, x
	fmt.Println(x, y) // 2 1

	// Blank identifier discards a value
	val, _ := divide(10, 3)
	fmt.Println(val)

	fmt.Println(globalMsg)
}

func divide(a, b int) (int, int) {
	return a / b, a % b
}
