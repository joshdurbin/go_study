//go:build ignore

package main

import "fmt"

// topologicalSort returns one valid topological order of a DAG with n nodes
// (labeled 0..n-1) and the given edges (u → v). Returns nil if a cycle is present.
// Use Kahn's algorithm (in-degree based BFS).
func topologicalSort(n int, edges [][]int) []int {
	// TODO: implement
	return nil
}

func main() {
	fmt.Println(topologicalSort(4, [][]int{{0, 1}, {1, 2}, {2, 3}})) // expect [0 1 2 3]
	fmt.Println(topologicalSort(3, [][]int{{0, 1}, {1, 2}, {2, 0}})) // expect nil (cycle)
	fmt.Println(topologicalSort(2, nil))                              // expect [0 1] or [1 0]
}
