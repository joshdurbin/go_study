//go:build ignore

package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// %w wraps; errors.Is walks the chain to compare; errors.As walks the chain to extract.

var ErrNotFound = errors.New("not found")

type ValidationError struct {
	Field string
}

func (v *ValidationError) Error() string { return "validation failed on field: " + v.Field }

// Layer 1: lookup; returns a sentinel.
func lookup(id string) error {
	return ErrNotFound
}

// Layer 2: wraps with context.
func loadUser(id string) error {
	if err := lookup(id); err != nil {
		return fmt.Errorf("loadUser(%q): %w", id, err)
	}
	return nil
}

// Layer 3: wraps again.
func handleRequest(id string) error {
	if err := loadUser(id); err != nil {
		return fmt.Errorf("handleRequest: %w", err)
	}
	return nil
}

func main() {
	// ─── Case 1: %w builds a chain ─────
	err := handleRequest("u-42")
	fmt.Println("err:", err)
	// handleRequest: loadUser("u-42"): not found

	// ─── Case 2: errors.Is checks anywhere in the chain ─────
	fmt.Println("Is ErrNotFound? ", errors.Is(err, ErrNotFound)) // true

	// ─── Case 3: errors.As extracts a typed error from the chain ─────
	wrapped := fmt.Errorf("save failed: %w", &ValidationError{Field: "email"})
	var ve *ValidationError
	if errors.As(wrapped, &ve) {
		fmt.Println("validation error on:", ve.Field) // email
	}

	// ─── Case 4: stdlib sentinels work the same way ─────
	_, openErr := os.Open("/nonexistent/path/xyz")
	fmt.Println("is os.ErrNotExist? ", errors.Is(openErr, os.ErrNotExist)) // true
	fmt.Println("is fs.ErrNotExist? ", errors.Is(openErr, fs.ErrNotExist)) // true (same value)

	// ─── Case 5: %v does NOT wrap ─────
	plain := fmt.Errorf("plain: %v", ErrNotFound)
	fmt.Println("Is ErrNotFound via %v? ", errors.Is(plain, ErrNotFound)) // false
}
