## Hint 1
The limiter is a single `time.Ticker` that drips into one shared `tokens` channel. Every worker reads from the same channel — that's what makes it global.

```go
tokens := make(chan struct{}, ratePerSec)
tick := time.NewTicker(time.Second / time.Duration(ratePerSec))
go func() { for range tick.C { select { case tokens <- struct{}{}: default: } } }()
```

## Hint 2
Standard worker pool reading from a `jobs` channel. The throttle is one line: `<-tokens` before doing the work.

```go
for w := 0; w < workers; w++ {
    go func() {
        defer wg.Done()
        for job := range jobs {
            <-tokens // global throttle
            process(job)
        }
    }()
}
```

## Hint 3
Stop the ticker goroutine cleanly when work is done — otherwise it leaks. A `stop` channel + `select` in the producer is the idiomatic shutdown.

```go
stop := make(chan struct{})
go func() {
    defer tick.Stop()
    for {
        select {
        case <-tick.C: /* emit token */
        case <-stop: return
        }
    }
}()
// after wg.Wait():
close(stop)
```
