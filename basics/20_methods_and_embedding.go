//go:build ignore

package main

import (
	"fmt"
	"math"
)

// Value vs pointer receivers — the most common source of confusion in Go.
//
// Use pointer receiver when:
//   - The method mutates the receiver
//   - The struct is large (avoid copying)
//   - Consistency: if any method uses pointer receiver, use it for all
//
// Use value receiver when:
//   - The method is read-only and the struct is small
//   - The type is a basic type (int, string) alias

type Vector struct {
	X, Y float64
}

func (v Vector) Length() float64 {          // value receiver — read-only
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector) Scale(factor float64) {    // pointer receiver — mutates
	v.X *= factor
	v.Y *= factor
}

func (v Vector) Add(other Vector) Vector {  // returns a new value — functional style
	return Vector{v.X + other.X, v.Y + other.Y}
}

// Embedding composes behavior without inheritance
type Logger struct {
	Prefix string
}

func (l Logger) Log(msg string) {
	fmt.Printf("[%s] %s\n", l.Prefix, msg)
}

type Service struct {
	Logger        // embedded — Service "inherits" Log method
	Name   string
}

// Service can override the embedded method
func (s Service) Log(msg string) {
	s.Logger.Log(fmt.Sprintf("(%s) %s", s.Name, msg)) // call embedded
}

// Interface satisfaction via embedding
type Reader interface{ Read() string }
type Writer interface{ Write(string) }
type ReadWriter interface {
	Reader
	Writer
}

type Buffer struct {
	data string
}
func (b *Buffer) Read() string      { return b.data }
func (b *Buffer) Write(s string)    { b.data += s }

func useRW(rw ReadWriter) {
	rw.Write("hello")
	fmt.Println(rw.Read())
}

func main() {
	v := Vector{3, 4}
	fmt.Println(v.Length()) // 5

	v.Scale(2)
	fmt.Println(v) // {6 8}

	sum := v.Add(Vector{1, 1})
	fmt.Println(sum) // {7 9}

	svc := Service{Logger: Logger{Prefix: "INFO"}, Name: "auth"}
	svc.Log("user logged in")         // [INFO] (auth) user logged in
	svc.Logger.Log("direct log")      // [INFO] direct log

	buf := &Buffer{}
	useRW(buf) // hello
}
