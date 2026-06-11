//go:build ignore

package main

import (
	"fmt"
	"sync"
	"time"
)

// MEMOIZATION (Generic, Thread-Safe)
// ==================================
// Problem: an expensive deterministic function gets called repeatedly with the
// same arguments. Cache by argument.
//
// Bridges naturally to the DP problems in the interview practice section —
// top-down DP with memoization is just this pattern applied to a recursive call.
//
// Key design choices:
//   - Generic in K (must be comparable) and V (any).
//   - Thread-safe: any number of concurrent callers.
//   - Double-checked: cheap RLock on the read path, full Lock only on miss.
//
// Caveat: this is a permanent cache. For LRU or TTL semantics see the LRU
// problem in interview practice + the rate limiter pattern.

type Memo[K comparable, V any] struct {
	mu    sync.RWMutex
	cache map[K]V
	fn    func(K) V
}

func NewMemo[K comparable, V any](fn func(K) V) *Memo[K, V] {
	return &Memo[K, V]{cache: make(map[K]V), fn: fn}
}

func (m *Memo[K, V]) Get(k K) V {
	m.mu.RLock()
	if v, ok := m.cache[k]; ok {
		m.mu.RUnlock()
		return v
	}
	m.mu.RUnlock()

	m.mu.Lock()
	defer m.mu.Unlock()
	// Re-check after upgrading: another goroutine may have filled it.
	if v, ok := m.cache[k]; ok {
		return v
	}
	v := m.fn(k)
	m.cache[k] = v
	return v
}

// Note: while one goroutine computes fn(k), other goroutines requesting the
// same key BLOCK on the write lock — they don't run fn(k) again. For full
// "compute once, share result" with concurrent waiters, see singleflight.

func main() {
	calls := 0
	slow := func(x int) int {
		calls++
		time.Sleep(20 * time.Millisecond)
		return x * x
	}

	m := NewMemo(slow)

	start := time.Now()
	fmt.Println(m.Get(5)) // 25 — computed
	fmt.Println(m.Get(5)) // 25 — cached
	fmt.Println(m.Get(6)) // 36 — computed
	fmt.Println(m.Get(5)) // 25 — cached
	fmt.Printf("calls=%d in %v\n", calls, time.Since(start).Round(time.Millisecond))

	// Concurrent access: 100 goroutines hitting 5 distinct keys.
	calls = 0
	m2 := NewMemo(slow)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_ = m2.Get(i % 5)
		}(i)
	}
	wg.Wait()
	fmt.Printf("concurrent: calls=%d (between 5 and ~5+races)\n", calls)
}
