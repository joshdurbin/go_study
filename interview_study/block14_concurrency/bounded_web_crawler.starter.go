//go:build ignore

package main

import (
	"fmt"
	"sort"
	"time"
)

// fetch simulates an I/O call. Do not modify.
func fetch(url string) string {
	time.Sleep(50 * time.Millisecond)
	return "body:" + url
}

// crawl fetches every url, with at most maxConcurrent fetches in flight.
// Returns the bodies and the peak observed concurrency (for the demo to assert
// the bound is honored).
func crawl(urls []string, maxConcurrent int) (results []string, peak int32) {
	// TODO: implement using a semaphore channel + sync.WaitGroup
	return nil, 0
}

func main() {
	urls := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	start := time.Now()
	results, peak := crawl(urls, 3)
	elapsed := time.Since(start)

	bounded := peak >= 1 && peak <= 3
	parallel := elapsed < 400*time.Millisecond

	sort.Strings(results)
	fmt.Printf("fetched=%d bounded=%v parallel=%v\n",
		len(results), bounded, parallel)
	if len(results) > 0 {
		fmt.Println("first:", results[0], "last:", results[len(results)-1])
	}
	// expect: fetched=10 bounded=true parallel=true
	// expect: first: body:a last: body:j
}
