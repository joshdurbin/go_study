//go:build ignore

package main

import (
	"fmt"
	"strings"
)

// fizzbuzz returns the FizzBuzz sequence for 1..n as a slice of strings.
// Using string concatenation makes the rule extensible (e.g., add 7→"Bazz").
func fizzbuzz(n int) []string {
	out := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		var b strings.Builder
		if i%3 == 0 {
			b.WriteString("Fizz")
		}
		if i%5 == 0 {
			b.WriteString("Buzz")
		}
		if b.Len() == 0 {
			fmt.Fprintf(&b, "%d", i)
		}
		out = append(out, b.String())
	}
	return out
}

func main() {
	fmt.Println(fizzbuzz(15))
}
