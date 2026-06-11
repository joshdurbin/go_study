//go:build ignore

package main

import "fmt"

type Edge struct {
	to, w int
}

// dijkstra returns shortest distances from src to every node in a graph with n
// nodes and adjacency list adj. Unreachable nodes get math.MaxInt.
// Use a min-heap keyed by tentative distance. Skip stale heap entries.
func dijkstra(n int, adj [][]Edge, src int) []int {
	// TODO: implement
	return nil
}

func main() {
	adj := [][]Edge{
		0: {{1, 4}, {2, 1}},
		1: {{3, 1}},
		2: {{1, 2}, {3, 5}},
		3: {},
	}
	fmt.Println(dijkstra(4, adj, 0)) // expect [0 3 1 4]
}
