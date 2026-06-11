# Graceful Shutdown

Stop a Go service cleanly: refuse new work on SIGTERM, drain in-flight requests within a deadline, then exit zero.

## The problem

When Kubernetes (or any orchestrator) sends SIGTERM, the process has a finite grace period to finish. A naive service that just calls `os.Exit` drops in-flight HTTP connections, leaves database writes half-applied, and turns "rolling deploys" into "occasional 500s." You need a coordinated stop: signal in, no-new-work, drain, exit.

## The shape

```go
ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
defer stop()

go runWorkers(ctx)            // workers loop on ctx.Done()
go srv.Serve(ln)              // http.Server in its own goroutine

<-ctx.Done()                  // wait for signal

shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
srv.Shutdown(shutdownCtx)     // stops accepting, drains handlers
wg.Wait()                     // wait for workers
```

Three contexts: the signal-derived one cancels workers, the shutdown-deadline one bounds the drain, and the original parent is the root of both.

## When to use

- Long-running HTTP / gRPC servers in production.
- Background workers (queue consumers, schedulers) that hold partial state.
- Any service that runs under systemd, Kubernetes, or a process supervisor.

## When NOT to use

- Short-lived CLIs — normal `return` from `main` is enough.
- Pure compute jobs with no external state — process exit is fine.

## Common pitfall

Calling `srv.Shutdown(ctx)` with the **same** signal-cancelled context. That context is already cancelled (that's why you're shutting down), so Shutdown returns immediately and aborts the drain. Always pass a **fresh** `context.WithTimeout(context.Background(), ...)` to Shutdown.

Second pitfall: forgetting that `srv.Serve` returns `http.ErrServerClosed` on a clean shutdown — treat it as success, not an error.

## Real-world

- `http.Server.Shutdown` (stdlib) — the canonical drain primitive.
- `signal.NotifyContext` (Go 1.16+) — replaces the older `signal.Notify(ch, ...)` boilerplate.
- Kubernetes `terminationGracePeriodSeconds` defines the upper bound of your shutdown ctx's timeout.

## Interview frame

"How do you implement graceful shutdown in a Go HTTP service?" — this is the canonical Go server interview question. The expected answer hits four beats: `signal.NotifyContext`, workers respect `ctx.Done()`, `http.Server.Shutdown` with a fresh timeout context, and `wg.Wait()` (or errgroup) to join everything before returning.
