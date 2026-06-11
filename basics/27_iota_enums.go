//go:build ignore

package main

import "fmt"

// iota is a compile-time counter that resets at each `const` block and
// increments by 1 on each line. It's how Go expresses enums.

// ─── Simple sequential enum ───
type Status int

const (
	StatusPending Status = iota // 0
	StatusActive                // 1
	StatusClosed                // 2
)

// Stringer interface: any type with String() string gets pretty-printed
// by fmt automatically.
func (s Status) String() string {
	switch s {
	case StatusPending:
		return "pending"
	case StatusActive:
		return "active"
	case StatusClosed:
		return "closed"
	default:
		return fmt.Sprintf("Status(%d)", int(s))
	}
}

// ─── Bit-flag enum: 1 << iota gives powers of two for bitwise OR/AND ───
type Perm uint8

const (
	PermRead  Perm = 1 << iota // 1
	PermWrite                  // 2
	PermExec                   // 4
)

func (p Perm) String() string {
	out := ""
	if p&PermRead != 0 {
		out += "r"
	} else {
		out += "-"
	}
	if p&PermWrite != 0 {
		out += "w"
	} else {
		out += "-"
	}
	if p&PermExec != 0 {
		out += "x"
	} else {
		out += "-"
	}
	return out
}

// ─── Skipping values with _ and starting at 1 ───
// Skipping zero is useful when zero means "unset" and you want a real value.
type Priority int

const (
	_                = iota // discard 0
	PriorityLow             // 1
	PriorityMedium          // 2
	PriorityHigh            // 3
)

func main() {
	// fmt sees the Stringer method and prints "active" instead of "1".
	fmt.Println("status:", StatusActive)
	fmt.Println("zero status:", Status(0))

	// Bitwise composition: combine flags with |
	p := PermRead | PermWrite
	fmt.Println("perm:", p) // rw-

	// Test for a flag with &
	if p&PermWrite != 0 {
		fmt.Println("writable")
	}

	// Add Exec
	p |= PermExec
	fmt.Println("perm:", p) // rwx

	// Priorities start at 1 — zero is "unset"
	var pri Priority
	fmt.Println("default priority:", pri) // 0, meaning unset
	fmt.Println("high priority:", PriorityHigh)

	// `go generate` + golang.org/x/tools/cmd/stringer can emit String() for you:
	//   //go:generate stringer -type=Status
	// Running `go generate ./...` produces status_string.go. Mentioned, not run.
}
