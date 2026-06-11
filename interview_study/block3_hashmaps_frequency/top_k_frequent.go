//go:build ignore

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// ── Approach 1: sort (simpler, interview-safe) ───────────────────────────────

func topKFreqSort(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}
	unique := make([]int, 0, len(freq))
	for n := range freq {
		unique = append(unique, n)
	}
	sort.Slice(unique, func(i, j int) bool {
		return freq[unique[i]] > freq[unique[j]]
	})
	return unique[:k]
}

// ── Approach 2: min-heap of size k, O(n log k) ──────────────────────────────

type pair struct{ val, count int }
type minHeapPair []pair

func (h minHeapPair) Len() int           { return len(h) }
func (h minHeapPair) Less(i, j int) bool { return h[i].count < h[j].count }
func (h minHeapPair) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeapPair) Push(x interface{}) {
	*h = append(*h, x.(pair))
}
func (h *minHeapPair) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func topKFreqHeap(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}
	h := &minHeapPair{}
	heap.Init(h)
	for val, cnt := range freq {
		heap.Push(h, pair{val, cnt})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	result := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).(pair).val
	}
	return result
}

func main() {
	fmt.Println(topKFreqSort([]int{1, 1, 1, 2, 2, 3}, 2)) // [1 2]
	fmt.Println(topKFreqHeap([]int{1, 1, 1, 2, 2, 3}, 2)) // [2 1] or [1 2]
}
