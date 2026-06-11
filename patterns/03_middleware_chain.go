//go:build ignore

package main

import (
	"fmt"
	"log"
	"time"
)

// MIDDLEWARE / HANDLER CHAIN PATTERN
// ===================================
// This is how net/http middleware works under the hood, and the same pattern
// applies to gRPC interceptors, CLI command hooks, and any pipeline where
// you want to wrap behavior without modifying the core handler.
//
// Core idea: a Middleware is a function that takes a Handler and returns a Handler.
// Wrapping happens at setup time; the chain is just nested function calls at runtime.
//
// Why this matters in interviews: it demonstrates you understand:
//   - Functions as first-class values
//   - Closure-based composition
//   - The Open/Closed principle (add behavior without modifying existing code)

// Handler is the core unit of work — analogous to http.Handler
type Handler func(req string) (string, error)

// Middleware wraps a Handler to add behavior before/after
type Middleware func(Handler) Handler

// Chain applies middlewares right-to-left so the first middleware in the list
// is the outermost wrapper (first to run on the way in, last on the way out).
// This matches net/http convention.
func Chain(h Handler, middlewares ...Middleware) Handler {
	// Apply in reverse so middleware[0] is outermost
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

// --- Concrete middleware implementations ---

// Logger logs the request and response
func Logger(next Handler) Handler {
	return func(req string) (string, error) {
		log.Printf("→ request: %q", req)
		resp, err := next(req)
		if err != nil {
			log.Printf("✗ error: %v", err)
		} else {
			log.Printf("← response: %q", resp)
		}
		return resp, err
	}
}

// Timer measures execution time
func Timer(next Handler) Handler {
	return func(req string) (string, error) {
		start := time.Now()
		resp, err := next(req)
		log.Printf("⏱ took %v", time.Since(start))
		return resp, err
	}
}

// Retry retries on error up to n times
func Retry(n int) Middleware {
	// Middleware factory — parameterized middleware is the common pattern
	return func(next Handler) Handler {
		return func(req string) (string, error) {
			var err error
			for i := 0; i < n; i++ {
				var resp string
				resp, err = next(req)
				if err == nil {
					return resp, nil
				}
				log.Printf("retry %d/%d after error: %v", i+1, n, err)
			}
			return "", fmt.Errorf("failed after %d retries: %w", n, err)
		}
	}
}

// RateLimit is a simple token bucket (illustrative — not thread-safe here)
func RateLimit(maxPerSec int) Middleware {
	ticker := time.NewTicker(time.Second / time.Duration(maxPerSec))
	return func(next Handler) Handler {
		return func(req string) (string, error) {
			<-ticker.C // block until a token is available
			return next(req)
		}
	}
}

// --- Core handler (no knowledge of middleware) ---

var attempts = 0

func coreHandler(req string) (string, error) {
	attempts++
	if attempts < 3 {
		return "", fmt.Errorf("service temporarily unavailable")
	}
	return fmt.Sprintf("hello, %s!", req), nil
}

func main() {
	// Build the chain: Logger → Timer → Retry(3) → coreHandler
	// Logger sees the request first and response last.
	handler := Chain(
		coreHandler,
		Logger,
		Timer,
		Retry(3),
	)

	resp, err := handler("world")
	fmt.Println("result:", resp, err)

	// Key insight: coreHandler has no idea it's being logged, timed, or retried.
	// You can add/remove middleware without touching coreHandler.
	// This is the Open/Closed principle in practice.
}
