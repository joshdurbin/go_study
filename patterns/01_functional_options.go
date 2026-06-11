//go:build ignore

package main

import (
	"fmt"
	"time"
)

// FUNCTIONAL OPTIONS PATTERN
// ==========================
// Problem: You have a struct with many optional configuration fields.
// Bad solutions:
//   - Huge constructor: NewServer(host, port, timeout, maxConns, ...) — unreadable, brittle
//   - Config struct: NewServer(Config{...}) — forces callers to know all fields upfront
// Good solution: functional options — each option is a function that mutates the config.
//
// This is THE canonical Go API design pattern. You'll see it everywhere:
// grpc.Dial, http.NewServeMux, zap.NewLogger, etc.
// If asked "how would you design a configurable API in Go?" — this is the answer.

type Server struct {
	host        string
	port        int
	timeout     time.Duration
	maxConns    int
	enableTLS   bool
}

// Option is a function that mutates a Server. The caller never sees *Server directly.
type Option func(*Server)

// Each exported function returns an Option closure.
// This is the key insight: the option *is* a function, not a value.

func WithHost(host string) Option {
	return func(s *Server) {
		s.host = host
	}
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

func WithTimeout(d time.Duration) Option {
	return func(s *Server) {
		s.timeout = d
	}
}

func WithMaxConns(n int) Option {
	return func(s *Server) {
		s.maxConns = n
	}
}

func WithTLS() Option {
	return func(s *Server) {
		s.enableTLS = true
	}
}

// NewServer applies defaults first, then applies each option in order.
// Callers only specify what they care about. New options are backward-compatible.
func NewServer(opts ...Option) *Server {
	// sensible defaults
	s := &Server{
		host:     "localhost",
		port:     8080,
		timeout:  30 * time.Second,
		maxConns: 100,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// VARIANT: Option that can return an error (for validation)
type ServerOption func(*Server) error

func WithPortValidated(port int) ServerOption {
	return func(s *Server) error {
		if port < 1 || port > 65535 {
			return fmt.Errorf("invalid port: %d", port)
		}
		s.port = port
		return nil
	}
}

func NewServerSafe(opts ...ServerOption) (*Server, error) {
	s := &Server{host: "localhost", port: 8080}
	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, err
		}
	}
	return s, nil
}

func main() {
	// Caller specifies only what they need. Clean, readable, extensible.
	s1 := NewServer()
	fmt.Printf("default: %s:%d timeout=%v\n", s1.host, s1.port, s1.timeout)

	s2 := NewServer(
		WithHost("0.0.0.0"),
		WithPort(9090),
		WithTLS(),
		WithMaxConns(500),
	)
	fmt.Printf("custom: %s:%d tls=%v maxConns=%d\n", s2.host, s2.port, s2.enableTLS, s2.maxConns)

	// Safe variant with validation
	_, err := NewServerSafe(WithPortValidated(99999))
	fmt.Println("validation error:", err)

	s3, _ := NewServerSafe(WithPortValidated(443))
	fmt.Println("valid port:", s3.port)
}
