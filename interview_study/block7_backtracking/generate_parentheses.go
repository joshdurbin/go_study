//go:build ignore

package main

import "fmt"

// generateParenthesis: track open and close counts. We can add '(' while
// open < n; we can add ')' while close < open (otherwise we'd unbalance).
// O(C_n) results where C_n is the n-th Catalan number; ~4^n / (n*sqrt(n)).
func generateParenthesis(n int) []string {
	res := []string{}
	cur := make([]byte, 0, 2*n)
	var backtrack func(open, closed int)
	backtrack = func(open, closed int) {
		if len(cur) == 2*n {
			res = append(res, string(cur))
			return
		}
		if open < n {
			cur = append(cur, '(')
			backtrack(open+1, closed)
			cur = cur[:len(cur)-1]
		}
		if closed < open {
			cur = append(cur, ')')
			backtrack(open, closed+1)
			cur = cur[:len(cur)-1]
		}
	}
	backtrack(0, 0)
	return res
}

func main() {
	fmt.Println(generateParenthesis(3)) // [((())) (()()) (())() ()(()) ()()()]
	fmt.Println(generateParenthesis(1)) // [()]
}
