//go:build ignore

package main

import (
	"fmt"
	"sort"
)

// STRATEGY PATTERN
// ================
// Intent: define a family of algorithms, encapsulate each one, make them interchangeable.
// The context uses a strategy without knowing its concrete type.
//
// In Go, strategy is almost always just a function type or a single-method interface.
// You rarely need a full class hierarchy. This is where Go shines vs OOP languages —
// a func([]int) []int IS a strategy. No wrapper class needed.
//
// Classic OOP would create: Sorter interface, BubbleSortStrategy, QuickSortStrategy...
// Go says: just use func types. Much less ceremony.

// --- Approach 1: Function-based strategy (idiomatic for simple cases) ---

type SortStrategy func([]int) []int

func bubbleSort(nums []int) []int {
	result := append([]int{}, nums...) // copy
	n := len(result)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

func stdSort(nums []int) []int {
	result := append([]int{}, nums...)
	sort.Ints(result)
	return result
}

// Sorter holds a strategy — swap it at runtime
type Sorter struct {
	strategy SortStrategy
}

func (s *Sorter) SetStrategy(fn SortStrategy) { s.strategy = fn }
func (s *Sorter) Sort(nums []int) []int       { return s.strategy(nums) }

// --- Approach 2: Interface-based strategy (better when strategy has state) ---

// Compressor is a strategy interface for compression algorithms
type Compressor interface {
	Compress(data string) string
	Name() string
}

type GzipCompressor struct{ Level int }
func (g GzipCompressor) Compress(data string) string {
	return fmt.Sprintf("[gzip-L%d](%s)", g.Level, data) // simulated
}
func (g GzipCompressor) Name() string { return fmt.Sprintf("gzip-level-%d", g.Level) }

type ZstdCompressor struct{}
func (z ZstdCompressor) Compress(data string) string { return fmt.Sprintf("[zstd](%s)", data) }
func (z ZstdCompressor) Name() string                { return "zstd" }

type NoOpCompressor struct{}
func (n NoOpCompressor) Compress(data string) string { return data }
func (n NoOpCompressor) Name() string                { return "none" }

type Pipeline struct {
	compressor Compressor
}

func NewPipeline(c Compressor) *Pipeline { return &Pipeline{compressor: c} }

func (p *Pipeline) Process(data string) string {
	fmt.Printf("compressing with %s\n", p.compressor.Name())
	return p.compressor.Compress(data)
}

// At runtime, select strategy based on config/environment
func chooseCompressor(env string) Compressor {
	switch env {
	case "prod":
		return ZstdCompressor{}
	case "debug":
		return NoOpCompressor{}
	default:
		return GzipCompressor{Level: 6}
	}
}

func main() {
	// Function-based strategy
	nums := []int{5, 2, 8, 1, 9}
	sorter := &Sorter{strategy: bubbleSort}
	fmt.Println(sorter.Sort(nums))

	sorter.SetStrategy(stdSort) // swap at runtime
	fmt.Println(sorter.Sort(nums))

	// Interface-based strategy
	for _, env := range []string{"prod", "debug", "staging"} {
		p := NewPipeline(chooseCompressor(env))
		out := p.Process("my-data-payload")
		fmt.Printf("  env=%s result=%s\n", env, out)
	}

	// Key insight: Pipeline.Process never changes. Only the Compressor swaps.
	// Adding a new algorithm = new struct implementing Compressor. Zero changes to Pipeline.
}
