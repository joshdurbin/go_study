//go:build ignore

package main

import "fmt"

// isValid reports whether the bracket sequence s is balanced and correctly nested.
// Brackets: ()[]{}.
func isValid(s string) bool {
	// TODO: implement
	return false
}

func main() {
	for _, s := range []string{"()", "()[]{}", "(]", "([)]", "{[]}", ""} {
		fmt.Printf("%-8s → %v\n", s, isValid(s))
	}
	// expect: () → true, ()[]{} → true, (] → false, ([)] → false, {[]} → true, "" → true
}
