//go:build ignore

package main

import "fmt"

// Slices share a backing array. Writes through one slice are visible through
// any other slice that overlaps that region. append() may or may not break
// the alias depending on capacity.

func main() {
	// ─── Case 1: slicing shares storage ─────
	a := []int{1, 2, 3, 4, 5}
	b := a[1:3] // b == [2 3], backed by a's array
	b[0] = 99
	fmt.Println("a:", a) // [1 99 3 4 5]  — a was mutated through b
	fmt.Println("b:", b) // [99 3]

	// ─── Case 2: append within capacity → still aliased ─────
	x := make([]int, 3, 10) // len=3, cap=10
	x[0], x[1], x[2] = 1, 2, 3
	y := x[:2]                  // y shares storage with x
	y = append(y, 999)          // fits in cap → writes into x[2]
	fmt.Println("x:", x)        // [1 2 999] — x[2] clobbered!
	fmt.Println("y:", y)        // [1 2 999]

	// ─── Case 3: append exceeds capacity → new backing array ─────
	p := make([]int, 3, 3) // len=3, cap=3 (full)
	p[0], p[1], p[2] = 1, 2, 3
	q := append(p, 4) // grows → new backing array
	q[0] = 999
	fmt.Println("p:", p) // [1 2 3]   — untouched
	fmt.Println("q:", q) // [999 2 3 4]

	// ─── Case 4: hidden retention ─────
	big := make([]byte, 1<<20) // 1 MiB
	small := big[:4]           // small keeps the whole 1 MiB alive
	_ = small
	// fix: copy to a fresh slice
	clone := make([]byte, 4)
	copy(clone, big[:4])
	_ = clone
	fmt.Println("hidden retention: small aliases big; clone does not")
}
