//go:build ignore

package main

import "fmt"

// reverse converts to []rune so multi-byte runes stay intact, then swaps in place.
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	fmt.Println(reverse("hello"))  // "olleh"
	fmt.Println(reverse("héllo"))  // "olléh" — multi-byte rune preserved
	fmt.Println(reverse("go 🚀"))  // "🚀 og"
}
