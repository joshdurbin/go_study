//go:build ignore

package main

import (
	"fmt"
	"strings"
)

// ITERATOR AND PIPELINE PATTERNS
// ================================
// Go 1.22 introduced range-over-func (iter package) but the pre-1.22 patterns
// are still widely used and worth knowing. Both are shown here.
//
// ITERATOR: abstract sequential access to a collection without exposing internals.
// In Go, three common approaches:
//   1. Callback-based (most common pre-1.22): func(yield func(T) bool)
//   2. Channel-based: caller ranges over channel (simpler, but goroutine overhead)
//   3. Stateful struct: explicit Next()/Value() (most explicit, mirrors sql.Rows)
//
// PIPELINE: transform a sequence through a series of stages.
// Each stage is a function that reads from one channel and writes to another.
// Stages compose naturally; the compiler ensures type safety at each connection.

// ── Approach 1: Callback iterator ───────────────────────────────────────────
// The yield function returns false to signal early termination (like break).
// This is the pattern the standard library's iter package formalizes.

type IntIterator func(yield func(int) bool)

func RangeOf(start, end int) IntIterator {
	return func(yield func(int) bool) {
		for i := start; i < end; i++ {
			if !yield(i) {
				return // early termination
			}
		}
	}
}

func Filter(it IntIterator, pred func(int) bool) IntIterator {
	return func(yield func(int) bool) {
		it(func(v int) bool {
			if pred(v) {
				return yield(v)
			}
			return true // skip, keep going
		})
	}
}

func Map(it IntIterator, fn func(int) int) IntIterator {
	return func(yield func(int) bool) {
		it(func(v int) bool {
			return yield(fn(v))
		})
	}
}

func Take(it IntIterator, n int) IntIterator {
	return func(yield func(int) bool) {
		count := 0
		it(func(v int) bool {
			if count >= n {
				return false // stop
			}
			count++
			return yield(v)
		})
	}
}

func Collect(it IntIterator) []int {
	var result []int
	it(func(v int) bool {
		result = append(result, v)
		return true
	})
	return result
}

// ── Approach 2: Stateful struct iterator (mirrors sql.Rows, bufio.Scanner) ──

type WordIterator struct {
	words   []string
	current int
}

func NewWordIterator(text string) *WordIterator {
	return &WordIterator{words: strings.Fields(text)}
}

func (w *WordIterator) Next() bool {
	w.current++
	return w.current <= len(w.words)
}

func (w *WordIterator) Value() string {
	return w.words[w.current-1]
}

// ── Channel-based pipeline ───────────────────────────────────────────────────

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func filterChan(in <-chan int, pred func(int) bool) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			if pred(v) {
				out <- v
			}
		}
	}()
	return out
}

func mapChan(in <-chan int, fn func(int) int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			out <- fn(v)
		}
	}()
	return out
}

func collect(in <-chan int) []int {
	var result []int
	for v := range in {
		result = append(result, v)
	}
	return result
}

func main() {
	fmt.Println("=== Callback Iterator ===")

	// Build a lazy pipeline: even numbers from 0..19, squared, first 4
	pipeline := Take(
		Map(
			Filter(RangeOf(0, 20), func(n int) bool { return n%2 == 0 }),
			func(n int) int { return n * n },
		),
		4,
	)
	fmt.Println(Collect(pipeline)) // [0 4 16 36]

	// Early termination: stop as soon as we find first > 10
	RangeOf(0, 100)(func(n int) bool {
		if n*n > 10 {
			fmt.Printf("first n where n²>10: %d\n", n) // 4
			return false // break
		}
		return true
	})

	fmt.Println("\n=== Stateful Iterator ===")
	it := NewWordIterator("the quick brown fox jumps over the lazy dog")
	longWords := []string{}
	for it.Next() {
		w := it.Value()
		if len(w) > 4 {
			longWords = append(longWords, w)
		}
	}
	fmt.Println(longWords) // [quick brown jumps]

	fmt.Println("\n=== Channel Pipeline ===")
	// Same logic, channel-based: even numbers 1-20, squared, collect
	nums := generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	evens := filterChan(nums, func(n int) bool { return n%2 == 0 })
	squared := mapChan(evens, func(n int) int { return n * n })
	fmt.Println(collect(squared)) // [4 16 36 64 100]

	// Tradeoff note:
	// Callback iterator: no goroutine overhead, lazy, early-exit is clean.
	// Channel pipeline: concurrent stages (stages run in parallel), natural
	//   fan-out/fan-in, but goroutine + channel overhead per stage.
	// Choose callback for CPU-bound transforms; channels for I/O-bound stages.
}
