//go:build ignore

package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// copyRandomList returns a deep copy of a list whose nodes have Next AND Random
// pointers. Either solution is fine:
//   - map[*Node]*Node, two passes — O(n) time, O(n) space.
//   - Interleave clones into the original, wire Random, then unzip — O(n) / O(1).
func copyRandomList(head *Node) *Node {
	// TODO: implement
	return nil
}

// buildWithRandom: vals[i] and randomIdx[i] = index that node i's Random points
// to (-1 for nil).
func buildWithRandom(vals []int, randomIdx []int) *Node {
	if len(vals) == 0 {
		return nil
	}
	nodes := make([]*Node, len(vals))
	for i, v := range vals {
		nodes[i] = &Node{Val: v}
	}
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	for i, r := range randomIdx {
		if r >= 0 {
			nodes[i].Random = nodes[r]
		}
	}
	return nodes[0]
}

func dump(h *Node) {
	idx := map[*Node]int{nil: -1}
	i := 0
	for c := h; c != nil; c = c.Next {
		idx[c] = i
		i++
	}
	for c := h; c != nil; c = c.Next {
		fmt.Printf("[%d, r=%d] ", c.Val, idx[c.Random])
	}
	fmt.Println()
}

func main() {
	h := buildWithRandom([]int{7, 13, 11, 10, 1}, []int{-1, 0, 4, 2, 0})
	dump(h)                   // expect [7, r=-1] [13, r=0] [11, r=4] [10, r=2] [1, r=0]
	dump(copyRandomList(h))   // expect same shape on the deep copy
	dump(copyRandomList(nil)) // expect (empty)
}
