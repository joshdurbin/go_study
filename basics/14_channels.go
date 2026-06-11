//go:build ignore

package main

import (
	"fmt"
	"time"
)

func produce(ch chan<- int, n int) {
	for i := 0; i < n; i++ {
		ch <- i // send
	}
	close(ch) // signal no more values
}

func main() {
	// Unbuffered channel — sender blocks until receiver is ready (synchronization)
	ch := make(chan int)
	go func() { ch <- 42 }()
	fmt.Println(<-ch) // 42

	// Buffered channel — sender blocks only when buffer is full
	buf := make(chan string, 2)
	buf <- "a"
	buf <- "b"
	// buf <- "c" would block here (buffer full)
	fmt.Println(<-buf) // a
	fmt.Println(<-buf) // b

	// Range over channel until closed
	nums := make(chan int, 5)
	go produce(nums, 5)
	for n := range nums {
		fmt.Print(n, " ")
	}
	fmt.Println()

	// Select: like a switch for channel operations
	c1 := make(chan string)
	c2 := make(chan string)
	go func() { time.Sleep(1 * time.Millisecond); c1 <- "one" }()
	go func() { time.Sleep(2 * time.Millisecond); c2 <- "two" }()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received:", msg1)
		case msg2 := <-c2:
			fmt.Println("received:", msg2)
		}
	}

	// Select with default: non-blocking channel operation
	ch2 := make(chan int, 1)
	select {
	case v := <-ch2:
		fmt.Println("got:", v)
	default:
		fmt.Println("no value ready") // prints this
	}

	// Done channel pattern: cancellation signal
	done := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Millisecond)
		close(done) // broadcast to all receivers
	}()
	select {
	case <-done:
		fmt.Println("done signal received")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timeout")
	}
}
