//go:build ignore

package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// A goroutine blocked on a channel that nobody will send to (or close) leaks.
// It holds memory and any captured resources forever. Detect with runtime.NumGoroutine
// or pprof/goroutine; prevent with select + ctx.Done() or guaranteed close.

// LEAKING: consumer waits on ch, but producer never sends and never closes.
func leakyConsumer() {
	ch := make(chan int)
	go func() {
		<-ch // blocks forever; goroutine leaks
		fmt.Println("never prints")
	}()
	// the goroutine survives this function returning
}

// FIXED: cancel via context. The select unblocks when ctx is done.
func safeConsumer(ctx context.Context) {
	ch := make(chan int)
	go func() {
		select {
		case v := <-ch:
			fmt.Println("got:", v)
		case <-ctx.Done():
			fmt.Println("ctx canceled, goroutine exits cleanly")
		}
	}()
}

func main() {
	before := runtime.NumGoroutine()

	for i := 0; i < 5; i++ {
		leakyConsumer()
	}
	time.Sleep(10 * time.Millisecond) // let them start
	leaked := runtime.NumGoroutine()
	fmt.Printf("goroutines after 5 leaks: %d (was %d)\n", leaked, before)

	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		safeConsumer(ctx)
	}
	time.Sleep(10 * time.Millisecond)
	fmt.Printf("goroutines after 5 safe starts: %d\n", runtime.NumGoroutine())

	cancel() // tell the safe consumers to exit
	time.Sleep(10 * time.Millisecond)
	fmt.Printf("goroutines after cancel:        %d\n", runtime.NumGoroutine())
	// The leaked ones never exit. Only the safe ones do.
}
