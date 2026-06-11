# Functional Options

The canonical Go pattern for configurable constructors. Each option is a function that mutates the target.

## The problem

You have a constructor with many optional parameters. Three bad answers:
- Huge positional argument list — brittle and unreadable.
- Config struct — forces callers to know every field, and adding a field is a breaking change.
- Multiple constructors — combinatorial explosion.

## The solution

```go
type Option func(*Server)
func WithPort(p int) Option { return func(s *Server) { s.port = p } }

func NewServer(opts ...Option) *Server {
    s := &Server{port: 8080}
    for _, opt := range opts { opt(s) }
    return s
}
```

Callers pass only the options they care about. Adding new options is backward-compatible.

## When to use

- Public API for a configurable type (server, client, builder).
- Many optional fields, especially with sensible defaults.

## When NOT to use

- Two required arguments and nothing else — just take them as regular params.
- The constructor has complex VALIDATION across fields — prefer the builder pattern (`patterns/18`).

## Real-world

`grpc.Dial`, `http.NewServer`, `zap.NewLogger`, every major Go library. **Interview canon.**

## Interview frame

"How would you design a configurable API in Go?" → functional options.
