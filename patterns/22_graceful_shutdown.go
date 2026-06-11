//go:build ignore

package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// GRACEFUL SHUTDOWN PATTERN
// =========================
// Problem: a Go service receives SIGTERM (e.g., from Kubernetes) and must stop
// accepting new requests, drain in-flight work, and exit cleanly within a
// bounded deadline — not just os.Exit() and lose connections mid-flight.
//
// The shape: signal.NotifyContext derives a ctx that cancels on SIGINT/SIGTERM.
// Workers loop on ctx.Done(). The http.Server is stopped via Shutdown(timeoutCtx)
// so it refuses new connections but lets active handlers finish.
//
// When to use: long-running HTTP services, background workers, anything that
// holds resources or in-flight work at shutdown.
// When NOT to use: short-lived CLIs — process exit on completion is enough.

// runWorker ticks until the shared ctx cancels, then returns.
func runWorker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	tick := time.NewTicker(150 * time.Millisecond)
	defer tick.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			fmt.Printf("worker %d tick\n", id)
		}
	}
}

func main() {
	// signal.NotifyContext: ctx cancels when SIGINT or SIGTERM arrives.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Bind to an ephemeral port so the demo never collides with a real service.
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	srv := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "ok")
		}),
	}

	fmt.Println("starting")

	// Background workers coordinated by a WaitGroup. They watch the same ctx
	// the signal handler will cancel.
	var wg sync.WaitGroup
	wg.Add(2)
	go runWorker(ctx, 1, &wg)
	go runWorker(ctx, 2, &wg)

	// Server runs in its own goroutine; Serve returns ErrServerClosed on Shutdown.
	serveErr := make(chan error, 1)
	go func() {
		if err := srv.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serveErr <- err
			return
		}
		serveErr <- nil
	}()

	// DEMO ONLY: auto-trigger SIGTERM after 500ms so `go run` completes without
	// a human pressing Ctrl-C. In production this line does not exist.
	time.AfterFunc(500*time.Millisecond, func() {
		_ = syscall.Kill(os.Getpid(), syscall.SIGTERM)
	})

	// Block until the signal-derived ctx cancels.
	<-ctx.Done()
	fmt.Println("received signal")

	// Shutdown has its own bounded deadline. If handlers don't drain in time,
	// Shutdown returns context.DeadlineExceeded and the caller decides what to do.
	fmt.Println("shutting down server")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		fmt.Println("shutdown error:", err)
	}
	<-serveErr

	// Workers see ctx.Done() from the signal context and exit on their own.
	wg.Wait()
	fmt.Println("workers stopped")

	fmt.Println("done")
}
