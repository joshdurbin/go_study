//go:build ignore

package main

import (
	"fmt"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

// crawl: fetch every URL with at most maxConcurrent in flight.
// Semaphore channel of capacity maxConcurrent enforces the limit:
// each goroutine acquires (sem <- struct{}{}) and releases on defer.
// A live-counter (atomic) records the peak observed concurrency so the demo
// can assert the bound is actually honored.
func crawl(urls []string, maxConcurrent int) (results []string, peak int32) {
	sem := make(chan struct{}, maxConcurrent)
	var wg sync.WaitGroup
	var live, hi int32

	out := make([]string, len(urls)) // index-keyed to avoid result-order races

	for i, u := range urls {
		wg.Add(1)
		sem <- struct{}{} // blocks once maxConcurrent are in flight
		go func(i int, u string) {
			defer wg.Done()
			defer func() { <-sem }()

			n := atomic.AddInt32(&live, 1)
			for { // record peak
				h := atomic.LoadInt32(&hi)
				if n <= h || atomic.CompareAndSwapInt32(&hi, h, n) {
					break
				}
			}

			out[i] = fetch(u)

			atomic.AddInt32(&live, -1)
		}(i, u)
	}

	wg.Wait()
	return out, atomic.LoadInt32(&hi)
}

// fetch simulates an I/O call.
func fetch(url string) string {
	time.Sleep(50 * time.Millisecond)
	return "body:" + url
}

func main() {
	urls := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	start := time.Now()
	results, peak := crawl(urls, 3)
	elapsed := time.Since(start)

	// 10 urls / 3 concurrent / 50ms each ≈ 4 batches → ~200ms.
	// bounded: peak goroutine count never exceeded the limit.
	// parallel: total time well under the serial 500ms.
	bounded := peak >= 1 && peak <= 3
	parallel := elapsed < 400*time.Millisecond

	sort.Strings(results)
	fmt.Printf("fetched=%d bounded=%v parallel=%v\n",
		len(results), bounded, parallel)
	fmt.Println("first:", results[0], "last:", results[len(results)-1])
	// expected:
	// fetched=10 bounded=true parallel=true
	// first: body:a last: body:j
}
