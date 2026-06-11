//go:build ignore

package main

import "fmt"

// dailyTemperatures returns, for each day, the number of days until a warmer
// temperature. 0 means no warmer day exists.
// Hint: monotonic decreasing stack of indices, O(n).
func dailyTemperatures(temps []int) []int {
	// TODO: implement
	return nil
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})) // expect [1 1 4 2 1 1 0 0]
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))                 // expect [1 1 1 0]
	fmt.Println(dailyTemperatures([]int{30, 60, 90}))                     // expect [1 1 0]
}
