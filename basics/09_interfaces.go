//go:build ignore

package main

import (
	"fmt"
	"math"
)

// Interface: a set of method signatures. Satisfied implicitly.
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

type Rect struct {
	W, H float64
}

func (r Rect) Area() float64      { return r.W * r.H }
func (r Rect) Perimeter() float64 { return 2 * (r.W + r.H) }

func printShape(s Shape) {
	fmt.Printf("area=%.2f perim=%.2f\n", s.Area(), s.Perimeter())
}

// Stringer interface from fmt package — implement String() to control printing
type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%.1f°C", float64(c))
}

// Empty interface accepts any value (use sparingly; prefer generics or concrete types)
func describe(i interface{}) {
	fmt.Printf("value=%v type=%T\n", i, i)
}

// Type assertion extracts the concrete type from an interface
func typeSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("int: %d\n", v)
	case string:
		fmt.Printf("string: %q\n", v)
	case Shape:
		fmt.Printf("shape area: %.2f\n", v.Area())
	default:
		fmt.Printf("unknown: %T\n", v)
	}
}

func main() {
	shapes := []Shape{Circle{5}, Rect{3, 4}}
	for _, s := range shapes {
		printShape(s)
	}

	temp := Celsius(36.6)
	fmt.Println(temp) // 36.6°C

	describe(42)
	describe("hello")

	typeSwitch(10)
	typeSwitch("world")
	typeSwitch(Circle{2})

	// Type assertion with comma-ok (safe, no panic)
	var i interface{} = "hello"
	s, ok := i.(string)
	fmt.Println(s, ok) // hello true
	n, ok := i.(int)
	fmt.Println(n, ok) // 0 false
}
