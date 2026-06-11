//go:build ignore

package main

import "fmt"

// This file covers generic SYNTAX: type parameters, constraints, and generic
// data structures. For Map / Filter / Reduce / GroupBy / FlatMap as a reusable
// toolkit (and when NOT to reach for them), see patterns/19_generic_helpers.go.

// Type constraint built from an interface listing concrete underlying types.
// ~int means "int or any type whose underlying type is int" — this is what
// lets you write generics over user-defined types like `type UserID int`.
type Number interface {
	~int | ~int64 | ~float64
}

// Sum: a generic function constrained to numeric types so `+=` is valid.
func Sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

// Generic data structure
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T)       { s.items = append(s.items, v) }
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}
func (s *Stack[T]) Len() int { return len(s.items) }

func main() {
	// Same Sum works for any underlying numeric type — including custom ones.
	fmt.Println(Sum([]int{1, 2, 3, 4, 5}))     // 15
	fmt.Println(Sum([]float64{1.1, 2.2, 3.3})) // 6.6

	// Generic data structure: zero-value-friendly, type-safe at every call site.
	var s Stack[string]
	s.Push("a")
	s.Push("b")
	s.Push("c")
	fmt.Println(s.Len()) // 3
	v, ok := s.Pop()
	fmt.Println(v, ok) // c true
}
