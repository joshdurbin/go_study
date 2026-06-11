//go:build ignore

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // decrement counter when this goroutine returns
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(10 * time.Millisecond)
	fmt.Printf("worker %d done\n", id)
}

func main() {
	// go keyword launches a goroutine — lightweight thread managed by the Go runtime
	go func() {
		fmt.Println("goroutine running")
	}()

	// WaitGroup: wait for a collection of goroutines to finish
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait() // blocks until all goroutines call Done()
	fmt.Println("all workers done")

	// Mutex: protect shared state from concurrent access
	var mu sync.Mutex
	counter := 0
	var wg2 sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg2.Wait()
	fmt.Println("counter:", counter) // always 100

	// Once: run something exactly once across goroutines
	var once sync.Once
	for i := 0; i < 5; i++ {
		once.Do(func() {
			fmt.Println("initialized once")
		})
	}
}
