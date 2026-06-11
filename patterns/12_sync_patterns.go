//go:build ignore

package main

import (
	"fmt"
	"sync"
	"time"
)

// SYNC PATTERNS: sync.Map, sync.Pool, sync.Cond
// ==============================================
// These are the less-obvious parts of the sync package that come up
// in interviews and real infra code. Knowing when to use each vs
// a plain mutex demonstrates real concurrency experience.

// ── sync.Map: concurrent map with no external locking ───────────────────────
// Use when: high read contention, keys are written once and read many times.
// Don't use when: you need iteration guarantees or atomic multi-key operations.
// Internally: uses two maps — a read-only "dirty" map for fast reads and a
// mutex-protected map for writes. Reads that hit the read map need no lock.

func syncMapDemo() {
	fmt.Println("=== sync.Map ===")
	var m sync.Map

	// Store
	m.Store("host", "db1.internal")
	m.Store("port", 3306)
	m.Store("replica_count", 3)

	// Load
	if v, ok := m.Load("host"); ok {
		fmt.Println("host:", v)
	}

	// LoadOrStore: atomic get-or-set
	actual, loaded := m.LoadOrStore("host", "db2.internal")
	fmt.Printf("LoadOrStore: value=%v, already_existed=%v\n", actual, loaded)

	// LoadAndDelete: atomic get-then-delete
	v, ok := m.LoadAndDelete("replica_count")
	fmt.Printf("LoadAndDelete: value=%v, existed=%v\n", v, ok)

	// Range: iterate all key-value pairs (order not guaranteed)
	// Return false from the callback to stop iteration.
	m.Range(func(k, v interface{}) bool {
		fmt.Printf("  %v = %v\n", k, v)
		return true // continue
	})

	// Concurrent writes — safe without any external locking
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			m.Store(fmt.Sprintf("worker-%d", n), n*n)
		}(i)
	}
	wg.Wait()
	count := 0
	m.Range(func(k, v interface{}) bool { count++; return true })
	fmt.Println("total keys:", count)
}

// ── sync.Pool: reuse temporary objects to reduce GC pressure ────────────────
// Use when: you allocate many short-lived objects of the same type.
// Common uses: byte buffers, encoders, database connection parameters.
// The pool may evict objects between GC cycles — never store state across calls.
// Internally: per-P (processor) local pools to avoid contention.

func syncPoolDemo() {
	fmt.Println("\n=== sync.Pool ===")
	allocations := 0

	bufPool := sync.Pool{
		New: func() interface{} {
			allocations++
			return make([]byte, 0, 1024) // pre-allocate 1KB
		},
	}

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			buf := bufPool.Get().([]byte)    // Get from pool (or allocate via New)
			buf = append(buf[:0], fmt.Sprintf("worker-%d", n)...)
			_ = buf
			bufPool.Put(buf[:0])             // Return to pool — reset slice length
		}(i)
	}
	wg.Wait()
	fmt.Printf("5 goroutines, %d actual allocations (pool reused the rest)\n", allocations)
}

// ── sync.Cond: condition variable for signaling between goroutines ───────────
// Use when: goroutines need to wait for a condition to become true.
// More expressive than a channel when: multiple goroutines wait, and you need
// both Signal (wake one) and Broadcast (wake all).
// Classic use: producer-consumer with a bounded buffer.

type BoundedQueue struct {
	mu    sync.Mutex
	cond  *sync.Cond
	items []int
	cap   int
}

func NewBoundedQueue(cap int) *BoundedQueue {
	q := &BoundedQueue{cap: cap}
	q.cond = sync.NewCond(&q.mu)
	return q
}

func (q *BoundedQueue) Push(item int) {
	q.mu.Lock()
	defer q.mu.Unlock()
	for len(q.items) >= q.cap {
		q.cond.Wait() // releases lock, sleeps, reacquires lock on wake
	}
	q.items = append(q.items, item)
	q.cond.Broadcast() // wake all waiters (consumers waiting for items)
}

func (q *BoundedQueue) Pop() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	for len(q.items) == 0 {
		q.cond.Wait() // wait until Push adds something
	}
	item := q.items[0]
	q.items = q.items[1:]
	q.cond.Broadcast() // wake producers waiting for space
	return item
}

func syncCondDemo() {
	fmt.Println("\n=== sync.Cond (bounded queue) ===")
	q := NewBoundedQueue(3)
	var wg sync.WaitGroup

	// Producer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 6; i++ {
			q.Push(i)
			fmt.Printf("pushed %d\n", i)
			time.Sleep(5 * time.Millisecond)
		}
	}()

	// Consumer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 6; i++ {
			v := q.Pop()
			fmt.Printf("popped %d\n", v)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	wg.Wait()
}

func main() {
	syncMapDemo()
	syncPoolDemo()
	syncCondDemo()
}
