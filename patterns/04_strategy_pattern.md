# Strategy Pattern

Swap algorithms at runtime by treating "the algorithm" as an interface with one method.

## The shape

```go
type Compressor interface {
    Compress([]byte) []byte
}

type Gzip struct{}
func (Gzip) Compress(b []byte) []byte { /* ... */ }

type Snappy struct{}
func (Snappy) Compress(b []byte) []byte { /* ... */ }

func send(data []byte, c Compressor) {
    payload := c.Compress(data)
    // ...
}
```

In Go, this is so natural with interfaces that it barely feels like a "pattern". That's the point — Go's interface model bakes Strategy in.

## When to use

- Multiple interchangeable algorithms (compression, hashing, formatting, payment processing).
- Choice driven by configuration or input data.
- You want to test the consumer with a fake/mock strategy.

## When NOT to use

- Only one strategy exists "for now" — don't speculate. Add the interface when you actually have two implementations.
- The strategies share too much state — closure-based options (`patterns/01`) may be cleaner.

## Real-world

`io.Reader` / `io.Writer` are the king strategies of stdlib — `bytes.Buffer`, `os.File`, `net.Conn`, `strings.Reader` all interchangeable behind one tiny interface.
