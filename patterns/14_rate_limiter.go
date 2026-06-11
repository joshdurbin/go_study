//go:build ignore

package main

import (
	"fmt"
	"sync"
	"time"
)

// TOKEN BUCKET RATE LIMITER
// =========================
// Problem: enforce "at most R requests per second, with burst capacity B."
//
// Token bucket: a bucket holds up to B tokens. Tokens refill at rate R/sec.
// Each request consumes 1 token. If empty, the request is denied (or waits).
//
// Why this design: leaky-bucket smooths to a strict rate; token-bucket allows
// short bursts up to capacity then enforces the long-term rate. Most APIs
// expose token-bucket semantics.
//
// Implementation trick: don't run a refill goroutine. Compute elapsed time
// on each Allow() call and add the proportional tokens — accurate and lock-free
// of background work.

type tokenBucket struct {
	mu         sync.Mutex
	tokens     float64   // current tokens (fractional allowed for sub-second precision)
	capacity   float64   // max tokens (burst size)
	refillRate float64   // tokens per second
	lastRefill time.Time
}

func newBucket(rps, burst float64) *tokenBucket {
	return &tokenBucket{
		tokens:     burst,
		capacity:   burst,
		refillRate: rps,
		lastRefill: time.Now(),
	}
}

// Allow returns true and consumes a token if one is available; otherwise false.
// O(1), no background goroutines.
func (b *tokenBucket) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(b.lastRefill).Seconds()
	b.tokens += elapsed * b.refillRate
	if b.tokens > b.capacity {
		b.tokens = b.capacity
	}
	b.lastRefill = now

	if b.tokens >= 1 {
		b.tokens--
		return true
	}
	return false
}

// Wait blocks until a token is available. Useful for client-side throttling.
func (b *tokenBucket) Wait() {
	for !b.Allow() {
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	// 5 requests/second, burst of 3.
	b := newBucket(5, 3)

	// First 3 succeed instantly (burst).
	for i := 0; i < 5; i++ {
		fmt.Printf("req %d allowed=%v\n", i, b.Allow())
	}

	// After 400ms we'll have ~2 fresh tokens.
	time.Sleep(400 * time.Millisecond)
	fmt.Println("after 400ms:")
	for i := 0; i < 3; i++ {
		fmt.Printf("req %d allowed=%v\n", i, b.Allow())
	}

	// Wait for a token (blocks ~200ms at 5 rps).
	start := time.Now()
	b.Wait()
	fmt.Printf("blocked until token in %v\n", time.Since(start).Round(10*time.Millisecond))
}
