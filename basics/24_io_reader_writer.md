# io.Reader & io.Writer

Two interfaces, four lines of code, the spine of Go's I/O. Everything that produces or consumes bytes implements them — files, network connections, buffers, compressors, hashes — so generic code works on all of them.

## The interfaces

```go
type Reader interface { Read(p []byte) (n int, err error) }
type Writer interface { Write(p []byte) (n int, err error) }
```

## Worth knowing

- `Read` fills the caller's buffer; it returns `(n, nil)` for partial reads and `(0, io.EOF)` when done. A short read is NOT an error.
- `io.EOF` is a sentinel value, not a failure — always check it explicitly.
- `bytes.Buffer` implements both interfaces; it's the go-to scratch space in tests.
- `strings.NewReader(s)` is the cheapest way to turn a string into a Reader.
- `io.Copy(dst, src)` pumps bytes from any Reader to any Writer with a small internal buffer — never load the whole stream into memory.
- `io.TeeReader(r, w)` returns a Reader that mirrors every byte it reads into `w`. Perfect for hashing or logging a stream as it's consumed.

## Composition is the point

Wrap a Reader with `bufio.NewReader`, `gzip.NewReader`, `io.LimitReader`, or `io.TeeReader` — each layer adds behavior without changing the interface. The same `io.Copy` call works no matter how deep the stack is.

## Common gotcha

Implementing `Read`: don't return `(0, nil)` — that's an infinite loop for callers. If you have no data and no error, you must block, return EOF, or return at least one byte.
