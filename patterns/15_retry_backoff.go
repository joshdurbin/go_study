//go:build ignore

package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// RETRY WITH EXPONENTIAL BACKOFF + JITTER
// =======================================
// Problem: a transient failure (network blip, throttled API) shouldn't kill the
// caller. Retry — but blindly retrying immediately stampedes the dependency.
//
// Recipe:
//   1. Exponential backoff: wait 2^n × base after the n-th failure.
//   2. Jitter: randomize within the window to prevent thundering-herd (everyone
//      retrying at the same instant after a shared outage).
//   3. Cap the backoff so it doesn't grow unbounded.
//   4. Respect context cancellation — never sleep past it.
//   5. Distinguish retryable from terminal errors — don't retry 4xx, do retry 5xx.

// Retryable wraps an error to signal "retry me." Sentinel + errors.Is would work
// equally well; the wrapper is more flexible when the inner error already has its
// own type identity.
type Retryable struct{ Err error }

func (r *Retryable) Error() string { return r.Err.Error() }
func (r *Retryable) Unwrap() error { return r.Err }

// Do retries fn until it succeeds, hits maxAttempts, or ctx is cancelled.
// Only errors marked Retryable trigger a retry.
func Do(ctx context.Context, maxAttempts int, base, cap time.Duration, fn func() error) error {
	var last error
	for attempt := 0; attempt < maxAttempts; attempt++ {
		if err := fn(); err == nil {
			return nil
		} else {
			last = err
			var r *Retryable
			if !errors.As(err, &r) {
				return err // terminal — don't retry
			}
		}
		// Exponential window: base * 2^attempt, capped at `cap`.
		window := base << attempt
		if window > cap {
			window = cap
		}
		// Full jitter: pick uniformly in [0, window).
		delay := time.Duration(rand.Int63n(int64(window)))
		select {
		case <-time.After(delay):
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return fmt.Errorf("after %d attempts: %w", maxAttempts, last)
}

func main() {
	// Simulate a flaky call that succeeds on the 3rd attempt.
	calls := 0
	flaky := func() error {
		calls++
		if calls < 3 {
			return &Retryable{Err: fmt.Errorf("transient error #%d", calls)}
		}
		return nil
	}

	err := Do(context.Background(), 5, 10*time.Millisecond, 1*time.Second, flaky)
	fmt.Printf("flaky: err=%v, calls=%d\n", err, calls)

	// Terminal error: do NOT retry.
	calls = 0
	terminal := func() error {
		calls++
		return fmt.Errorf("auth failed (404)") // not wrapped in Retryable
	}
	err = Do(context.Background(), 5, 10*time.Millisecond, 1*time.Second, terminal)
	fmt.Printf("terminal: err=%v, calls=%d\n", err, calls)
}
