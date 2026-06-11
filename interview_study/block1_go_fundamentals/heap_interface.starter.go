//go:build ignore

package main

import (
	"container/heap"
	"fmt"
)

// IntHeap is a min-heap of ints implementing heap.Interface.
// Implement the five required methods: Len, Less, Swap, Push, Pop.
type IntHeap []int

// TODO: implement Len, Less, Swap (from sort.Interface)
// TODO: implement Push(x any) and Pop() any
//        — they operate on the BACKING SLICE (append/truncate at the end).

func main() {
	h := &IntHeap{5, 3, 8, 1, 4}
	heap.Init(h)
	heap.Push(h, 2)
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	fmt.Println()
	// expect: 1 2 3 4 5 8
}
