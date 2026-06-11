//go:build ignore

package main

import (
	"fmt"
	"strings"
)

type Point struct{ X, Y int }

func main() {
	// fmt verbs
	fmt.Printf("%v\n", Point{1, 2})    // {1 2}      — default format
	fmt.Printf("%+v\n", Point{1, 2})   // {X:1 Y:2}  — with field names
	fmt.Printf("%#v\n", Point{1, 2})   // main.Point{X:1, Y:2} — Go syntax
	fmt.Printf("%T\n", Point{1, 2})    // main.Point — type

	fmt.Printf("%d %05d %+d\n", 42, 42, 42)  // 42 00042 +42
	fmt.Printf("%b %o %x %X\n", 255, 255, 255, 255) // 11111111 377 ff FF
	fmt.Printf("%f %.2f %e\n", 3.14159, 3.14159, 3.14159) // 3.141590 3.14 3.141590e+00
	fmt.Printf("%s %q\n", "hello", "hello") // hello "hello"

	// Sprintf returns a formatted string
	msg := fmt.Sprintf("x=%d y=%d", 10, 20)
	fmt.Println(msg)

	// strings package
	s := "  Hello, Gopher!  "
	fmt.Println(strings.TrimSpace(s))          // Hello, Gopher!
	fmt.Println(strings.ToUpper(s))            // upper
	fmt.Println(strings.Contains(s, "Gopher")) // true
	fmt.Println(strings.Replace(s, "o", "0", -1)) // all occurrences
	fmt.Println(strings.Split("a,b,c", ","))   // [a b c]
	fmt.Println(strings.Join([]string{"x", "y", "z"}, "-")) // x-y-z
	fmt.Println(strings.HasPrefix("golang", "go")) // true
	fmt.Println(strings.Count("cheese", "e"))      // 3
	fmt.Println(strings.Index("gopher", "ph"))     // 2

	// strings.Builder: efficient string concatenation (avoid + in loops)
	var b strings.Builder
	for i := 0; i < 5; i++ {
		fmt.Fprintf(&b, "%d", i)
	}
	fmt.Println(b.String()) // 01234
}
