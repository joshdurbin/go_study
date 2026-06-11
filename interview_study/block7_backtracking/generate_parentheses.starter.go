//go:build ignore

package main

import "fmt"

// generateParenthesis returns every well-formed string of n pairs of parens.
// Track open and close counts so you only place a character when the resulting
// prefix can still extend to a valid string.
func generateParenthesis(n int) []string {
	// TODO: implement
	return nil
}

func main() {
	fmt.Println(generateParenthesis(3)) // expect [((())) (()()) (())() ()(()) ()()()]
	fmt.Println(generateParenthesis(1)) // expect [()]
}
