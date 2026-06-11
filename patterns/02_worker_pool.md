# Worker Pool

Fixed concurrency for unbounded work. N goroutines all read from one jobs channel.

## The problem

- Spawning one goroutine per task can blow up at scale (thousands → millions).
- Doing tasks serially is too slow.
- You want **bounded** concurrency: at most N in flight.

## The shape

```go
jobs := make(chan Job)
results := make(chan Result)
var wg sync.WaitGroup
for i := 0; i < N; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        for j := range jobs { results <- process(j) }
    }()
}
// producer sends to jobs, then close(jobs)
go func() { wg.Wait(); close(results) }()
```

The closer goroutine is essential — without it, consumers block forever on the never-closed `results` channel.

## When to use

- Network calls with rate limits.
- CPU-bound work parallelized across cores.
- Anywhere "for each X, do Y" runs in production at scale.

## When NOT to use

- Single-shot parallelism with a known small N — just use `errgroup` (`patterns/13`).
- IO-bound work where you'd rather use a semaphore (channel-as-token-bucket) than a dedicated pool.

## Real-world

Crawlers, batch processors, anything with a queue of similar tasks.
