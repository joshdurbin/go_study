# Middleware Chain

Wrap an HTTP handler with cross-cutting concerns (logging, auth, recovery, timing) without modifying the handler itself.

## The shape

```go
type Middleware func(http.Handler) http.Handler

func Logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println(r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

handler = Logging(Auth(Recover(handler)))
```

Each middleware is a function that takes a handler and returns a new handler with extra behavior wrapped around it.

## When to use

- HTTP request lifecycle hooks (the canonical case).
- Any pipeline where steps share an interface and can compose linearly.
- Database transactions, gRPC interceptors (same pattern, different signature).

## Order matters

Outermost wrappers run FIRST on the way in and LAST on the way out. Recovery should be outermost (catch panics anywhere). Timing should be near outermost (capture the whole request). Authorization usually middle.

## Real-world

`net/http` middleware libraries (chi, gorilla mux), gRPC interceptors, HTTP client `http.RoundTripper` chains.
