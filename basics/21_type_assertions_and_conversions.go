//go:build ignore

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Explicit type conversions (not implicit — Go requires explicitness)
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Println(i, f, u)

	// String <-> numeric conversions (use strconv, not fmt)
	s := strconv.Itoa(42)          // int to string
	fmt.Println(s, fmt.Sprintf("%T", s)) // "42" string

	n, err := strconv.Atoi("123")
	fmt.Println(n, err)  // 123 <nil>

	_, err = strconv.Atoi("abc")
	fmt.Println(err) // strconv.Atoi: parsing "abc": invalid syntax

	f2, _ := strconv.ParseFloat("3.14", 64)
	fmt.Println(f2) // 3.14

	b, _ := strconv.ParseBool("true")
	fmt.Println(b) // true

	// byte <-> string
	bytes := []byte("hello")
	bytes[0] = 'H'
	fmt.Println(string(bytes)) // Hello

	// rune <-> string
	r := 'G'
	fmt.Println(string(r)) // G

	// Type assertion on interfaces
	var iface interface{} = "gopher"

	// Panics if wrong type:
	// s2 := iface.(int)

	// Safe form:
	s2, ok := iface.(string)
	fmt.Println(s2, ok) // gopher true

	i2, ok := iface.(int)
	fmt.Println(i2, ok) // 0 false

	// Type switch for multiple possible types
	values := []interface{}{42, "hello", 3.14, true, []int{1, 2}}
	for _, v := range values {
		switch x := v.(type) {
		case int:
			fmt.Printf("int: %d\n", x)
		case string:
			fmt.Printf("string: %q\n", x)
		case float64:
			fmt.Printf("float64: %g\n", x)
		case bool:
			fmt.Printf("bool: %v\n", x)
		default:
			fmt.Printf("other: %T\n", x)
		}
	}
}
