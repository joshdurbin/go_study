//go:build ignore

package main

import "fmt"

// reverse returns s with its runes in reverse order. Must work on multi-byte
// UTF-8 input (e.g., "héllo", "go 🚀") without corrupting runes.
func reverse(s string) string {
	// TODO: implement
	return ""
}

func main() {
	fmt.Println(reverse("hello"))  // expect "olleh"
	fmt.Println(reverse("héllo"))  // expect "olléh"
	fmt.Println(reverse("go 🚀"))  // expect "🚀 og"
}
