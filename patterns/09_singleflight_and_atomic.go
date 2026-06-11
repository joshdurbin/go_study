//go:build ignore

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// SINGLEFLIGHT + ATOMIC PATTERNS
// ================================
// Two distinct but related patterns for handling concurrent access efficiently.
//
// SINGLEFLIGHT: deduplicate concurrent requests for the same work.
// Problem: 100 goroutines all call fetchUser(42) simultaneously during a cache miss.
// Without singleflight: 100 identical DB queries fire.
// With singleflight: only 1 query fires; all 100 goroutines get the same result.
//
// Use it when: cache stampede, thundering herd, expensive idempotent operations.
// Real-world: your ProxySQL config fetches, Redis topology reads, GCP API calls.
//
// ATOMIC: lock-free operations on single numeric values.
// Faster than mutex for simple counters/flags because it uses CPU instructions
// (CAS — compare-and-swap) rather than OS-level locking.
// Use it for: request counters, feature flags, state machine state words.

// ── Singleflight (manual implementation — production use golang.org/x/sync/singleflight) ──

// Group ensures only one execution per key runs at a time.
// Concurrent callers with the same key wait for the in-flight call and share its result.
type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

// Do executes fn if no call for key is in flight; otherwise waits for and returns
// the result of the in-flight call. The second return value is true if the result
// was shared (not the one that actually ran fn).
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, bool, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		// In-flight call exists — wait for it
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, true, c.err // shared=true
	}
	c := &call{}
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	// This goroutine is the one that actually does the work
	c.val, c.err = fn()
	c.wg.Done()

	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, false, c.err // shared=false — this was the runner
}

// Simulate an expensive DB fetch
var dbFetchCount int64 // atomic counter

func fetchFromDB(id string) (string, error) {
	atomic.AddInt64(&dbFetchCount, 1)
	time.Sleep(20 * time.Millisecond) // simulate latency
	return fmt.Sprintf("data-for-%s", id), nil
}

// ── Atomic patterns ──────────────────────────────────────────────────────────

// AtomicCounter wraps sync/atomic for a clean API
type AtomicCounter struct {
	val int64
}

func (c *AtomicCounter) Inc()         { atomic.AddInt64(&c.val, 1) }
func (c *AtomicCounter) Dec()         { atomic.AddInt64(&c.val, -1) }
func (c *AtomicCounter) Load() int64  { return atomic.LoadInt64(&c.val) }
func (c *AtomicCounter) Store(v int64){ atomic.StoreInt64(&c.val, v) }

// CompareAndSwap: set value to new only if current matches expected.
// Core primitive for lock-free state machines.
func (c *AtomicCounter) CAS(expected, new int64) bool {
	return atomic.CompareAndSwapInt64(&c.val, expected, new)
}

// AtomicFlag: a boolean flag safe for concurrent read/write without mutex
type AtomicFlag struct{ val int32 }

func (f *AtomicFlag) Set()      { atomic.StoreInt32(&f.val, 1) }
func (f *AtomicFlag) Clear()    { atomic.StoreInt32(&f.val, 0) }
func (f *AtomicFlag) IsSet() bool { return atomic.LoadInt32(&f.val) == 1 }

func main() {
	fmt.Println("=== Singleflight ===")
	var g Group
	var wg sync.WaitGroup
	sharedCount := int64(0)

	// 10 concurrent goroutines all request the same key simultaneously
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			val, shared, err := g.Do("user-42", func() (interface{}, error) {
				return fetchFromDB("user-42")
			})
			if shared {
				atomic.AddInt64(&sharedCount, 1)
			}
			_ = val
			_ = err
		}(i)
	}
	wg.Wait()

	fmt.Printf("10 concurrent requests → %d actual DB calls, %d shared results\n",
		atomic.LoadInt64(&dbFetchCount), sharedCount)
	// Expected: 1 DB call, 9 shared results

	fmt.Println("\n=== Atomic Counter ===")
	var counter AtomicCounter
	var wg2 sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			counter.Inc()
		}()
	}
	wg2.Wait()
	fmt.Printf("counter after 1000 concurrent increments: %d\n", counter.Load())

	// CAS for lock-free state transition
	// Only one goroutine can win the transition from 0 → 1
	wins := int64(0)
	var wg3 sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg3.Add(1)
		go func() {
			defer wg3.Done()
			if counter.CAS(1000, 9999) {
				atomic.AddInt64(&wins, 1) // exactly one goroutine wins
			}
		}()
	}
	wg3.Wait()
	fmt.Printf("CAS winners (expected 1): %d, counter=%d\n", wins, counter.Load())

	fmt.Println("\n=== Atomic Flag ===")
	var initialized AtomicFlag
	fmt.Println("before init:", initialized.IsSet())
	initialized.Set()
	fmt.Println("after Set:", initialized.IsSet())
	initialized.Clear()
	fmt.Println("after Clear:", initialized.IsSet())
}
