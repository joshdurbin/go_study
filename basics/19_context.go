//go:build ignore

package main

import (
	"context"
	"fmt"
	"time"
)

// Context carries deadlines, cancellation signals, and request-scoped values.
// Always the first parameter by convention.

func fetchData(ctx context.Context, id int) (string, error) {
	// Simulate work that respects cancellation
	select {
	case <-time.After(50 * time.Millisecond): // work takes 50ms
		return fmt.Sprintf("data-%d", id), nil
	case <-ctx.Done():
		return "", ctx.Err() // context.DeadlineExceeded or context.Canceled
	}
}

func processWithValue(ctx context.Context) {
	// Extract a value stored in context (use sparingly — prefer explicit params)
	if reqID, ok := ctx.Value("requestID").(string); ok {
		fmt.Println("request ID:", reqID)
	}
}

func main() {
	// WithTimeout: automatically cancels after duration
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel() // always call cancel to release resources

	result, err := fetchData(ctx, 1)
	fmt.Println(result, err) // data-1 <nil>

	// WithTimeout that expires
	ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel2()

	result, err = fetchData(ctx2, 2)
	fmt.Println(result, err) // "" context deadline exceeded

	// WithCancel: manual cancellation
	ctx3, cancel3 := context.WithCancel(context.Background())
	go func() {
		time.Sleep(5 * time.Millisecond)
		cancel3() // signal cancellation
	}()
	result, err = fetchData(ctx3, 3)
	fmt.Println(result, err) // "" context canceled

	// WithValue: attach request-scoped data
	ctx4 := context.WithValue(context.Background(), "requestID", "abc-123")
	processWithValue(ctx4)
}
