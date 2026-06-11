//go:build ignore

package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// ERRGROUP PATTERN
// ================
// Problem: launch N goroutines, wait for all to finish, propagate the FIRST error
// and cancel any still in-flight when it happens.
//
// sync.WaitGroup alone doesn't carry errors or cancel siblings.
// golang.org/x/sync/errgroup is the real production answer, but the pattern is
// only ~30 lines — here we implement it inline so you can recognize and recreate
// it on a whiteboard.
//
// Use it for: parallel fetches, fan-out aggregation, anything where "all must
// succeed or we abort early."

type group struct {
	wg      sync.WaitGroup
	once    sync.Once
	err     error
	cancel  context.CancelFunc
}

// withContext returns a group bound to a derived context. Calling cancel either
// happens manually OR fires automatically on the first goroutine error.
func withContext(parent context.Context) (*group, context.Context) {
	ctx, cancel := context.WithCancel(parent)
	return &group{cancel: cancel}, ctx
}

func (g *group) Go(fn func() error) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		if err := fn(); err != nil {
			g.once.Do(func() { // record only the first error
				g.err = err
				if g.cancel != nil {
					g.cancel() // signal siblings to bail
				}
			})
		}
	}()
}

func (g *group) Wait() error {
	g.wg.Wait()
	if g.cancel != nil {
		g.cancel()
	}
	return g.err
}

// fetch simulates an I/O call that respects context cancellation.
func fetch(ctx context.Context, id int, fail bool) error {
	select {
	case <-time.After(50 * time.Millisecond):
		if fail {
			return fmt.Errorf("worker %d failed", id)
		}
		fmt.Printf("worker %d ok\n", id)
		return nil
	case <-ctx.Done():
		fmt.Printf("worker %d cancelled: %v\n", id, ctx.Err())
		return ctx.Err()
	}
}

func main() {
	// Happy path: all succeed.
	g, ctx := withContext(context.Background())
	for i := 1; i <= 3; i++ {
		i := i
		g.Go(func() error { return fetch(ctx, i, false) })
	}
	fmt.Println("happy err:", g.Wait())

	// Failure: one worker fails; the rest see ctx.Done and exit early.
	g2, ctx2 := withContext(context.Background())
	for i := 1; i <= 4; i++ {
		i := i
		g2.Go(func() error { return fetch(ctx2, i, i == 2) })
	}
	err := g2.Wait()
	fmt.Println("first error captured:", err)
	fmt.Println("is context-cancelled?:", errors.Is(err, context.Canceled))
}
