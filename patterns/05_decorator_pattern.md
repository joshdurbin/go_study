# Decorator Pattern

Wrap an object to add behavior without changing the original. In Go, usually done via interface satisfaction + composition.

## The shape

```go
type Logger interface { Log(msg string) }

type baseLogger struct{}
func (baseLogger) Log(m string) { fmt.Println(m) }

type TimestampedLogger struct{ inner Logger }
func (t TimestampedLogger) Log(m string) {
    t.inner.Log(time.Now().Format(time.RFC3339) + " " + m)
}
```

Wrappers stack cleanly: `TimestampedLogger{Authenticated{Base{}}}`.

## Decorator vs. Middleware

They're the same pattern at heart. The difference is the unit:
- **Middleware** wraps a CALL (HTTP handler, gRPC method).
- **Decorator** wraps an OBJECT (a Logger, a Storage).

## When to use

- You can't (or shouldn't) modify the original type — third-party library, generated code, stable contract.
- Behaviors are independently composable (caching, logging, metrics, retry).

## When NOT to use

- The behavior should always apply — bake it into the base. Decoration is for OPTIONAL behavior.
- The wrapper needs access to internal state of the wrapped object — decoration only sees public methods.

## Real-world

`http.RoundTripper` decorators (auth, caching, retry), `io.Reader` wrappers (`bufio.Reader`, `gzip.Reader`, `cipher.StreamReader`).
