//go:build ignore

package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// ERROR WRAPPING + SENTINELS + TYPED ERRORS
// =========================================
// Three idioms working together:
//
//   1. Sentinel errors (`var ErrFoo = errors.New(...)`) — equality check via
//      errors.Is. Use when callers need a stable identity to switch on.
//   2. Typed errors (`type ValidationError struct{ Field string }`) — when
//      callers need DATA from the error, not just identity. Match via errors.As.
//   3. Wrapping with `%w` — preserves the chain so errors.Is / errors.As can
//      still find a sentinel or type buried under contextual messages.
//
// Anti-patterns to avoid:
//   - Using `%s` or `fmt.Errorf` without `%w` — the chain breaks.
//   - String-matching error.Error() — fragile and breaks on rewording.
//   - Returning generic `errors.New("not found")` when callers need to
//     differentiate from other not-found cases.

// 1. Sentinel
var ErrNotFound = errors.New("not found")

// 2. Typed error carrying data
type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation: %s: %s", e.Field, e.Msg)
}

// Layer 1: low-level lookup.
func dbLookup(id string) error {
	if id == "" {
		return &ValidationError{Field: "id", Msg: "required"}
	}
	if id == "missing" {
		return ErrNotFound
	}
	return nil
}

// Layer 2: business logic — adds context via %w.
func getUser(id string) error {
	if err := dbLookup(id); err != nil {
		return fmt.Errorf("getUser %q: %w", id, err)
	}
	return nil
}

// Layer 3: HTTP handler-style caller — pattern-matches on the chain.
func handle(id string) {
	err := getUser(id)
	if err == nil {
		fmt.Printf("[%s] ok\n", id)
		return
	}

	// errors.Is: walks the chain comparing to a sentinel.
	if errors.Is(err, ErrNotFound) {
		fmt.Printf("[%s] 404: %v\n", id, err)
		return
	}

	// errors.As: walks the chain looking for a type, binding it for data access.
	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Printf("[%s] 400 (field=%s): %v\n", id, ve.Field, err)
		return
	}

	fmt.Printf("[%s] 500: %v\n", id, err)
}

func main() {
	handle("alice")   // ok
	handle("missing") // 404 sentinel through 1 wrap
	handle("")        // 400 typed error through 1 wrap

	// Bonus: stdlib errors mix in cleanly. io.EOF is a sentinel.
	_, err := os.Open("/no/such/file")
	fmt.Println("os.Open ErrNotExist?:", errors.Is(err, os.ErrNotExist))

	wrapped := fmt.Errorf("reading config: %w", io.EOF)
	fmt.Println("wrapped is EOF?:", errors.Is(wrapped, io.EOF))
}
