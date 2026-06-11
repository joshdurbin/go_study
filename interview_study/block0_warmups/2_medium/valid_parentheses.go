//go:build ignore

package main

import "fmt"

// isValid: push openers; on a closer, pop and verify the match.
// Empty stack at end means everything paired up.
func isValid(s string) bool {
	pair := map[rune]rune{')': '(', ']': '[', '}': '{'}
	stack := make([]rune, 0, len(s))
	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pair[c] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	for _, s := range []string{"()", "()[]{}", "(]", "([)]", "{[]}"} {
		fmt.Printf("%-8s → %v\n", s, isValid(s))
	}
}
