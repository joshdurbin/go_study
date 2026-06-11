## Hint 1
Track producers with one `WaitGroup`. The channel is shared; only producers may close it. Start with the skeleton.

```go
ch := make(chan int, bufSize)
var producers sync.WaitGroup
for p := 0; p < numProducers; p++ {
    producers.Add(1)
    go func(id int) { defer producers.Done(); /* push items */ }(p)
}
```

## Hint 2
A separate goroutine waits on the producer WaitGroup, then closes `ch`. This is the only safe place to close — closing inside a producer would race the others.

```go
go func() {
    producers.Wait()
    close(ch) // unblocks all consumer range loops
}()
```

## Hint 3
Consumers `range ch` and count atomically. A second WaitGroup tracks them so `main` returns only after everything is drained.

```go
var consumers sync.WaitGroup
for c := 0; c < numConsumers; c++ {
    consumers.Add(1)
    go func() { defer consumers.Done(); for range ch { atomic.AddInt64(&processed, 1) } }()
}
consumers.Wait()
```
