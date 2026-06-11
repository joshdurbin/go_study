//go:build ignore

package main

import "fmt"

type Point struct {
	X, Y int
}

type Rectangle struct {
	TopLeft     Point
	BottomRight Point
}

// Value receiver — works on a copy, cannot mutate
func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

// Pointer receiver — can mutate, also avoids copying large structs
func (p *Point) Scale(factor int) {
	p.X *= factor
	p.Y *= factor
}

func (r Rectangle) Area() int {
	w := r.BottomRight.X - r.TopLeft.X
	h := r.BottomRight.Y - r.TopLeft.Y
	return w * h
}

func main() {
	// Struct literal
	p := Point{X: 3, Y: 4}
	fmt.Println(p.String()) // (3, 4)

	p.Scale(2)
	fmt.Println(p) // {6 8}

	// Anonymous struct (useful for one-off shapes, test data)
	person := struct {
		Name string
		Age  int
	}{Name: "Josh", Age: 40}
	fmt.Println(person)

	// Nested structs
	rect := Rectangle{
		TopLeft:     Point{0, 0},
		BottomRight: Point{10, 5},
	}
	fmt.Println(rect.Area()) // 50

	// Struct embedding (composition over inheritance)
	type Animal struct {
		Name string
	}
	type Dog struct {
		Animal        // embedded — Dog "inherits" Animal's fields and methods
		Breed string
	}
	d := Dog{Animal: Animal{Name: "Rex"}, Breed: "Husky"}
	fmt.Println(d.Name, d.Breed) // Rex Husky (d.Name promoted from Animal)

	// Struct comparison — valid if all fields are comparable
	p1 := Point{1, 2}
	p2 := Point{1, 2}
	fmt.Println(p1 == p2) // true
}
