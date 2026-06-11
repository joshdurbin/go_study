//go:build ignore

package main

import (
	"errors"
	"fmt"
)

// This file introduces Go's error model: error-as-interface, sentinel errors,
// custom error types, and basic errors.Is / errors.As usage. For multi-layer
// wrapping with %w, anti-patterns to avoid, and stdlib error interop, see
// patterns/20_error_wrapping.go.

// Sentinel errors — compare with ==  or errors.Is()
var ErrNotFound = errors.New("not found")
var ErrPermission = errors.New("permission denied")

// Custom error type — implement the error interface (Error() string)
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: field=%q message=%q", e.Field, e.Message)
}

func findUser(id int) (string, error) {
	if id <= 0 {
		return "", &ValidationError{Field: "id", Message: "must be positive"}
	}
	if id > 100 {
		return "", fmt.Errorf("findUser: %w", ErrNotFound) // wrapping with %w
	}
	return fmt.Sprintf("user-%d", id), nil
}

func main() {
	// Standard error handling pattern
	user, err := findUser(42)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(user)
	}

	// errors.Is() unwraps and checks sentinel errors
	_, err = findUser(999)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("not found!")
	}

	// errors.As() unwraps into a target type
	_, err = findUser(-1)
	var valErr *ValidationError
	if errors.As(err, &valErr) {
		fmt.Printf("bad field: %s\n", valErr.Field)
	}

	// panic / recover — for truly unexpected situations (not normal control flow)
	safeDiv(10, 2)
	safeDiv(10, 0)
}

func safeDiv(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
		}
	}()
	fmt.Println(a / b)
}
