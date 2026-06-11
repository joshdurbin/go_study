## Hint 1
Use a buffered channel as a counting semaphore. Capacity = max concurrency. Acquire before launching, release in defer.

```go
sem := make(chan struct{}, maxConcurrent)
for _, u := range urls {
    sem <- struct{}{} // blocks when N are in flight
    go func(u string) { defer func() { <-sem }(); /* fetch */ }(u)
}
```

## Hint 2
You also need a `sync.WaitGroup` so the caller knows when every fetch finished. Acquire the semaphore BEFORE the goroutine starts; that's what bounds the launch rate.

```go
var wg sync.WaitGroup
for _, u := range urls {
    wg.Add(1)
    sem <- struct{}{}
    go func(u string) { defer wg.Done(); defer func() { <-sem }(); /* fetch */ }(u)
}
wg.Wait()
```

## Hint 3
Collect results without locks by writing into a pre-sized slice at each goroutine's own index. Different indices = no race.

```go
out := make([]string, len(urls))
for i, u := range urls {
    i, u := i, u
    // ... inside goroutine:
    out[i] = fetch(u)
}
```
