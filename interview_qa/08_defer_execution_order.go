//go:build ignore

package main

import "fmt"

// Defers run in LIFO order. Arguments are evaluated AT defer time, not at return.
// Defers can modify named return values. Defer in a loop stacks up.

func deferOrder() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	// prints 3, 2, 1
}

func argsEvaluatedAtDeferTime() {
	x := 10
	defer fmt.Println("deferred x:", x) // captures x = 10
	x = 20
	fmt.Println("current x:", x) // 20
	// prints "current x: 20" then "deferred x: 10"
}

func deferredClosureSeesLatest() {
	x := 10
	defer func() { fmt.Println("closure x:", x) }() // captures variable, not value
	x = 20
	// prints "closure x: 20"
}

// Named return value — deferred function can modify it.
func wrap() (result string) {
	defer func() {
		result = "wrapped(" + result + ")"
	}()
	return "inner"
	// 1) result = "inner"
	// 2) deferred runs, sets result = "wrapped(inner)"
	// 3) function returns "wrapped(inner)"
}

func deferInLoopWrong() {
	for i := 0; i < 3; i++ {
		// All three defers stack and only fire when the function returns,
		// not at the end of each iteration. For files in a loop, this leaks
		// file handles until the function exits.
		defer fmt.Println("loop defer", i)
	}
	fmt.Println("loop body done")
	// prints: loop body done, loop defer 2, loop defer 1, loop defer 0
}

func main() {
	fmt.Println("─── LIFO ───")
	deferOrder()

	fmt.Println("─── args captured at defer time ───")
	argsEvaluatedAtDeferTime()

	fmt.Println("─── closure captures variable ───")
	deferredClosureSeesLatest()

	fmt.Println("─── defer mutates named return ───")
	fmt.Println("wrap() =", wrap())

	fmt.Println("─── defer in loop ───")
	deferInLoopWrong()
}
