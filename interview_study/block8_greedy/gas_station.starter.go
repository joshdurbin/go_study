//go:build ignore

package main

import "fmt"

// canCompleteCircuit returns the starting station index for a full clockwise
// loop, or -1 if no such station exists. gas[i] is fuel available at station i;
// cost[i] is fuel needed to travel from i to i+1 (wrapping).
func canCompleteCircuit(gas, cost []int) int {
	// TODO: implement
	return -1
}

func main() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})) // expect 3
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))             // expect -1
	fmt.Println(canCompleteCircuit([]int{5}, []int{4}))                         // expect 0
}
