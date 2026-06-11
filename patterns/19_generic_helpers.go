//go:build ignore

package main

import (
	"fmt"
	"strings"
)

// GENERIC MAP / FILTER / REDUCE
// =============================
// Problem: every codebase ends up writing the same loops over and over.
// Go 1.18+ generics let you express them once, type-safely.
//
// Note: the stdlib `slices` package now provides many of these. Writing them
// here is for understanding — and for cases where you need behavior the stdlib
// doesn't ship (e.g., FlatMap, GroupBy, Reduce with custom accumulator type).
//
// Style note: in real Go code, an explicit for-loop is usually clearer than
// chained Map/Filter/Reduce calls. Use these when they genuinely simplify;
// don't reach for them just because they exist.

func Map[T, R any](s []T, fn func(T) R) []R {
	out := make([]R, len(s))
	for i, v := range s {
		out[i] = fn(v)
	}
	return out
}

func Filter[T any](s []T, pred func(T) bool) []T {
	out := make([]T, 0, len(s))
	for _, v := range s {
		if pred(v) {
			out = append(out, v)
		}
	}
	return out
}

// Reduce is generic in both T (element) and R (accumulator) — the killer
// feature vs. languages where reduce's accumulator must match element type.
func Reduce[T, R any](s []T, init R, fn func(R, T) R) R {
	acc := init
	for _, v := range s {
		acc = fn(acc, v)
	}
	return acc
}

// GroupBy: classic case where a for-loop IS the readable answer, but the
// generic version pays off when you're calling it three times in a file.
func GroupBy[T any, K comparable](s []T, key func(T) K) map[K][]T {
	out := make(map[K][]T)
	for _, v := range s {
		k := key(v)
		out[k] = append(out[k], v)
	}
	return out
}

// FlatMap: useful when each element maps to 0..N outputs.
func FlatMap[T, R any](s []T, fn func(T) []R) []R {
	var out []R
	for _, v := range s {
		out = append(out, fn(v)...)
	}
	return out
}

type User struct {
	Name string
	Age  int
	Team string
}

func main() {
	users := []User{
		{"Ada", 35, "platform"},
		{"Bo", 19, "ml"},
		{"Cy", 41, "platform"},
		{"Di", 28, "ml"},
	}

	names := Map(users, func(u User) string { return u.Name })
	fmt.Println(names)

	adults := Filter(users, func(u User) bool { return u.Age >= 21 })
	fmt.Println("adults:", len(adults))

	totalAge := Reduce(users, 0, func(acc int, u User) int { return acc + u.Age })
	fmt.Println("total age:", totalAge)

	// Note R != T: building a single string from User slice.
	joined := Reduce(users, "", func(acc string, u User) string {
		if acc == "" {
			return u.Name
		}
		return acc + ", " + u.Name
	})
	fmt.Println("joined:", joined)

	byTeam := GroupBy(users, func(u User) string { return u.Team })
	for team, ms := range byTeam {
		fmt.Printf("  %s (%d): %v\n", team, len(ms), Map(ms, func(u User) string { return u.Name }))
	}

	words := []string{"go is", "expressive and", "explicit"}
	tokens := FlatMap(words, func(s string) []string { return strings.Fields(s) })
	fmt.Println("tokens:", tokens)
}
