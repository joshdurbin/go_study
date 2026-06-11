//go:build ignore

package main

import "fmt"

// An interface is a (type, value) pair. It only equals nil when BOTH halves are nil.
// A non-nil interface wrapping a nil pointer is NOT nil.

type Error struct {
	Msg string
}

func (e *Error) Error() string { return e.Msg }

// Returns a *Error declared nil — but the interface return value carries
// the type *Error, so the returned `error` interface is NOT nil.
func doWork() error {
	var e *Error = nil
	return e
}

func main() {
	// ─── Case 1: plain nil pointer ─────
	var p *Error = nil
	fmt.Println("p == nil:", p == nil) // true

	// ─── Case 2: nil pointer assigned to interface ─────
	var i interface{} = p
	fmt.Println("i == nil:", i == nil) // false — interface holds (type=*Error, value=nil)
	fmt.Printf("i type=%T value=%v\n", i, i)

	// ─── Case 3: the classic error footgun ─────
	err := doWork()
	if err != nil {
		fmt.Println("ERROR returned (but the underlying pointer is nil!):", err)
	} else {
		fmt.Println("no error")
	}

	// ─── Case 4: the fix — return untyped nil ─────
	err2 := func() error { return nil }()
	fmt.Println("err2 == nil:", err2 == nil) // true
}
