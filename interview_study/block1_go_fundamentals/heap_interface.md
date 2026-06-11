PROBLEM: Implement heap.Interface — Min-Heap of Ints
=====================================================
Implement a min-heap using Go's container/heap package.
Then use it to extract the K smallest elements from an unsorted slice.

This is the single most commonly fumbled Go interview pattern.
You MUST know how to implement Len, Less, Swap, Push, Pop.

Key insight: heap.Pop returns interface{} — you must type-assert.
Key gotcha:  Push/Pop work on *MinHeap (pointer receiver via heap.Interface).

Example:
  Input:  [5, 3, 8, 1, 9, 2, 7], K=3
  Output: [1, 2, 3]
