//go:build ignore

package main

import "fmt"

// countComponents counts the number of connected components in an undirected
// graph with n nodes labeled 0..n-1 and the given edges.
func countComponents(n int, edges [][]int) int {
	// TODO: implement — build adjacency list, DFS/BFS from each unvisited node.
	return 0
}

func main() {
	fmt.Println(countComponents(5, [][]int{{0, 1}, {1, 2}, {3, 4}})) // expect 2
	fmt.Println(countComponents(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}})) // expect 1
	fmt.Println(countComponents(4, nil)) // expect 4
}
