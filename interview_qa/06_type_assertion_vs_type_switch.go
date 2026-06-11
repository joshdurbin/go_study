//go:build ignore

package main

import "fmt"

// Type assertion:  v, ok := i.(T)         — single target, comma-ok form is safe.
// Type switch:     switch v := i.(type) — multiple targets in one statement.

type Cat struct{ Name string }
type Dog struct{ Name string }

func describe(i interface{}) string {
	switch v := i.(type) {
	case nil:
		return "nil"
	case string:
		return "string: " + v
	case int:
		return fmt.Sprintf("int: %d", v)
	case *Cat:
		return "cat: " + v.Name
	case *Dog:
		return "dog: " + v.Name
	default:
		return fmt.Sprintf("unknown: %T", v)
	}
}

func main() {
	var i interface{} = "hello"

	// ─── Case 1: assertion with comma-ok (safe) ─────
	s, ok := i.(string)
	fmt.Println("string?", ok, s) // true hello

	n, ok := i.(int)
	fmt.Println("int?   ", ok, n) // false 0 (zero value of int)

	// ─── Case 2: assertion without comma-ok (panics on mismatch) ─────
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
		}
	}()

	// ─── Case 3: type switch ─────
	for _, x := range []interface{}{42, "go", &Cat{"Mittens"}, &Dog{"Rex"}, 3.14, nil} {
		fmt.Println(describe(x))
	}

	// Now the panic — kept last so the deferred recover catches it
	_ = i.(int) // panics: interface conversion: interface {} is string, not int
}
