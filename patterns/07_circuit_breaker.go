//go:build ignore

package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// CIRCUIT BREAKER PATTERN
// =======================
// Intent: prevent cascading failures by stopping calls to a failing dependency.
// Named after electrical circuit breakers that trip to protect downstream components.
//
// Three states:
//   Closed   → normal operation; failures are counted
//   Open     → dependency is considered down; calls fail fast (no network hit)
//   HalfOpen → after a cooldown, allow one probe call; success resets, failure reopens
//
// Real-world: your ProxySQL health checks, Redis Sentinel failover detection,
// and any service that calls an external API should consider this pattern.
// It's a key resilience primitive in distributed systems.

type State int

const (
	StateClosed   State = iota
	StateOpen
	StateHalfOpen
)

func (s State) String() string {
	return [...]string{"CLOSED", "OPEN", "HALF_OPEN"}[s]
}

var ErrCircuitOpen = errors.New("circuit breaker is open")

type CircuitBreaker struct {
	mu           sync.Mutex
	maxFailures  int
	resetTimeout time.Duration
	state        State
	failures     int
	lastFailure  time.Time
}

func NewCircuitBreaker(maxFailures int, resetTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		maxFailures:  maxFailures,
		resetTimeout: resetTimeout,
		state:        StateClosed,
	}
}

func (cb *CircuitBreaker) Call(fn func() error) error {
	cb.mu.Lock()
	switch cb.state {
	case StateOpen:
		if time.Since(cb.lastFailure) > cb.resetTimeout {
			cb.state = StateHalfOpen
			fmt.Println("circuit: OPEN → HALF_OPEN (probing)")
		} else {
			cb.mu.Unlock()
			return ErrCircuitOpen
		}
	}
	cb.mu.Unlock()

	err := fn() // call outside lock — don't hold mutex during I/O

	cb.mu.Lock()
	defer cb.mu.Unlock()
	if err != nil {
		cb.failures++
		cb.lastFailure = time.Now()
		if cb.failures >= cb.maxFailures {
			if cb.state != StateOpen {
				fmt.Printf("circuit: %s → OPEN (failures=%d)\n", cb.state, cb.failures)
			}
			cb.state = StateOpen
		}
		return err
	}
	if cb.state == StateHalfOpen {
		fmt.Println("circuit: HALF_OPEN → CLOSED (probe succeeded)")
	}
	cb.state = StateClosed
	cb.failures = 0
	return nil
}

func (cb *CircuitBreaker) State() State {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return cb.state
}

func main() {
	cb := NewCircuitBreaker(3, 50*time.Millisecond)

	// Simulated downstream: fails 3 times, then recovers.
	// We track actual downstream calls separately from cb.Call() calls.
	downstreamCalls := 0
	downstream := func() error {
		downstreamCalls++
		if downstreamCalls <= 3 {
			return fmt.Errorf("service unavailable (downstream call %d)", downstreamCalls)
		}
		return nil // recovered
	}

	// Calls 1–3: fail, tripping the breaker on call 3
	for i := 1; i <= 3; i++ {
		err := cb.Call(downstream)
		fmt.Printf("call %d: err=%v  state=%s\n", i, err, cb.State())
	}

	// Calls 4–5: circuit open, fails fast — downstream NOT called
	for i := 4; i <= 5; i++ {
		err := cb.Call(downstream)
		fmt.Printf("call %d: err=%v  state=%s\n", i, err, cb.State())
	}
	fmt.Printf("(downstream was called %d times — fast-fail protected it)\n\n", downstreamCalls)

	// Wait for reset timeout, then probe — downstream has recovered by now
	time.Sleep(60 * time.Millisecond)

	err := cb.Call(downstream)
	fmt.Printf("probe: err=%v  state=%s\n", err, cb.State())

	// Normal operation resumes
	err = cb.Call(downstream)
	fmt.Printf("normal: err=%v  state=%s\n", err, cb.State())
}
