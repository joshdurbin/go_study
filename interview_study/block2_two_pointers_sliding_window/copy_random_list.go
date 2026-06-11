//go:build ignore

package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// copyRandomList deep-copies a list whose nodes have Next AND Random pointers.
// Interleave-clone-then-split: weave A→A'→B→B'→C→C' so each clone sits right
// after its original; this gives O(1) lookup from original to clone when wiring
// Random pointers. Then unzip the two lists. O(n) time, O(1) extra space.
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	// 1. interleave clones: A → A' → B → B' → ...
	for cur := head; cur != nil; cur = cur.Next.Next {
		clone := &Node{Val: cur.Val, Next: cur.Next}
		cur.Next = clone
	}
	// 2. wire Random pointers on the clones
	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
	}
	// 3. unzip the two interleaved lists
	cloneHead := head.Next
	for cur := head; cur != nil; cur = cur.Next {
		clone := cur.Next
		cur.Next = clone.Next
		if clone.Next != nil {
			clone.Next = clone.Next.Next
		}
	}
	return cloneHead
}

// buildWithRandom: vals[i] and randomIdx[i] = index that node i's Random points to
// (-1 for nil).
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
	// 7 → 13 → 11 → 10 → 1   random: nil, 0, 4, 2, 0
	h := buildWithRandom([]int{7, 13, 11, 10, 1}, []int{-1, 0, 4, 2, 0})
	dump(h)                  // [7, r=-1] [13, r=0] [11, r=4] [10, r=2] [1, r=0]
	dump(copyRandomList(h))  // same shape on the deep copy
	dump(copyRandomList(nil)) // (empty)
}
