//go:build ignore

package main

import "fmt"

// leastInterval returns the minimum number of CPU intervals required to run
// every task in tasks, where the same task cannot run twice within n intervals.
// tasks contains only uppercase letters A-Z.
func leastInterval(tasks []byte, n int) int {
	// TODO: implement
	return 0
}

func main() {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))                          // expect 8
	fmt.Println(leastInterval([]byte{'A', 'C', 'A', 'B', 'D', 'B'}, 1))                          // expect 6
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2)) // expect 16
}
