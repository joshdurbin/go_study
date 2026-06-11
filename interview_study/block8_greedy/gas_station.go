//go:build ignore

package main

import "fmt"

// canCompleteCircuit: if total gas < total cost, impossible. Otherwise the unique
// start is the station immediately after the lowest running tank.
// O(n) time, O(1) space.
func canCompleteCircuit(gas, cost []int) int {
	total, tank, start := 0, 0, 0
	for i := range gas {
		diff := gas[i] - cost[i]
		total += diff
		tank += diff
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}
	if total < 0 {
		return -1
	}
	return start
}

func main() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})) // 3
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))             // -1
	fmt.Println(canCompleteCircuit([]int{5}, []int{4}))                         // 0
}
