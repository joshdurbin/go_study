//go:build ignore

package main

import (
	"fmt"
	"time"
)

// Unbuffered:  send blocks until a receive is ready (rendezvous / handshake).
// Buffered(n): send blocks only when n items are unread.

func main() {
	// ─── Case 1: unbuffered — sender blocks until receiver reads ─────
	un := make(chan int)
	start := time.Now()
	go func() {
		time.Sleep(30 * time.Millisecond) // delay the receive
		<-un
	}()
	un <- 1 // blocks ~30ms waiting for the goroutine to receive
	fmt.Printf("unbuffered send took %v (waited for receiver)\n", time.Since(start).Round(time.Millisecond))

	// ─── Case 2: buffered — sender returns immediately while buffer has room ─────
	buf := make(chan int, 3)
	start = time.Now()
	buf <- 1
	buf <- 2
	buf <- 3
	fmt.Printf("buffered (cap=3) 3 sends took %v (no receiver yet)\n", time.Since(start).Round(time.Microsecond))

	// ─── Case 3: buffer full → sender blocks ─────
	start = time.Now()
	go func() {
		time.Sleep(20 * time.Millisecond)
		<-buf // make room
	}()
	buf <- 4 // blocks until receiver drains one slot
	fmt.Printf("buffered send when full took %v\n", time.Since(start).Round(time.Millisecond))

	// drain
	for len(buf) > 0 {
		<-buf
	}

	// ─── Case 4: synchronization handshake (unbuffered) ─────
	done := make(chan struct{})
	go func() {
		fmt.Println("worker: doing thing")
		done <- struct{}{} // blocks until main receives — guaranteed handoff
	}()
	<-done
	fmt.Println("main: worker finished (synchronized via unbuffered chan)")
}
