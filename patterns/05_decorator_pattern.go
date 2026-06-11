//go:build ignore

package main

import (
	"fmt"
	"strings"
	"time"
)

// DECORATOR PATTERN
// =================
// Intent: attach additional responsibilities to an object dynamically.
// In Go: wrap an interface implementation with another implementation of the same interface.
//
// This is how Go's standard library composes I/O:
//   bufio.NewWriter(file)     — adds buffering to any io.Writer
//   gzip.NewWriter(file)      — adds compression to any io.Writer
//   io.MultiWriter(w1, w2)    — fans out to multiple io.Writers
// They all implement io.Writer, so they stack infinitely.
//
// The pattern: wrapping struct holds a reference to the wrapped interface.
// This is subtly different from middleware (which transforms functions).
// Decorator transforms objects (interface implementations).

// Store is the interface we want to decorate
type Store interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Delete(key string) error
}

// --- Base implementation ---

type MemoryStore struct {
	data map[string]string
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{data: make(map[string]string)}
}

func (m *MemoryStore) Get(key string) (string, error) {
	v, ok := m.data[key]
	if !ok {
		return "", fmt.Errorf("key not found: %s", key)
	}
	return v, nil
}

func (m *MemoryStore) Set(key, value string) error {
	m.data[key] = value
	return nil
}

func (m *MemoryStore) Delete(key string) error {
	delete(m.data, key)
	return nil
}

// --- Decorator 1: Logging ---
// Wraps any Store, logs every operation. The wrapped store doesn't change.

type LoggingStore struct {
	inner  Store
	prefix string
}

func WithLogging(s Store, prefix string) Store {
	return &LoggingStore{inner: s, prefix: prefix}
}

func (l *LoggingStore) Get(key string) (string, error) {
	v, err := l.inner.Get(key)
	fmt.Printf("[%s] GET %q → %q err=%v\n", l.prefix, key, v, err)
	return v, err
}

func (l *LoggingStore) Set(key, value string) error {
	err := l.inner.Set(key, value)
	fmt.Printf("[%s] SET %q=%q err=%v\n", l.prefix, key, value, err)
	return err
}

func (l *LoggingStore) Delete(key string) error {
	err := l.inner.Delete(key)
	fmt.Printf("[%s] DELETE %q err=%v\n", l.prefix, key, err)
	return err
}

// --- Decorator 2: Metrics/Timing ---

type MetricsStore struct {
	inner Store
	calls map[string]int
	total time.Duration
}

func WithMetrics(s Store) *MetricsStore {
	// Returns *MetricsStore (not Store) so caller can also call Report()
	// This is a common real-world tradeoff: richer type vs pure interface
	return &MetricsStore{inner: s, calls: make(map[string]int)}
}

func (m *MetricsStore) Get(key string) (string, error) {
	start := time.Now()
	v, err := m.inner.Get(key)
	m.calls["GET"]++
	m.total += time.Since(start)
	return v, err
}

func (m *MetricsStore) Set(key, value string) error {
	start := time.Now()
	err := m.inner.Set(key, value)
	m.calls["SET"]++
	m.total += time.Since(start)
	return err
}

func (m *MetricsStore) Delete(key string) error {
	start := time.Now()
	err := m.inner.Delete(key)
	m.calls["DELETE"]++
	m.total += time.Since(start)
	return err
}

func (m *MetricsStore) Report() {
	fmt.Printf("calls: %v  total_time: %v\n", m.calls, m.total)
}

// --- Decorator 3: Key namespacing ---

type NamespacedStore struct {
	inner     Store
	namespace string
}

func WithNamespace(s Store, ns string) Store {
	return &NamespacedStore{inner: s, namespace: ns}
}

func (n *NamespacedStore) key(k string) string {
	return n.namespace + ":" + k
}

func (n *NamespacedStore) Get(key string) (string, error)         { return n.inner.Get(n.key(key)) }
func (n *NamespacedStore) Set(key, value string) error            { return n.inner.Set(n.key(key), value) }
func (n *NamespacedStore) Delete(key string) error                { return n.inner.Delete(n.key(key)) }

func main() {
	base := NewMemoryStore()

	// Stack decorators: namespace → metrics → logging → base
	// Order matters: logging sees namespaced keys, metrics measures everything including logging overhead
	metrics := WithMetrics(base)
	logged := WithLogging(metrics, "STORE")
	store := WithNamespace(logged, "app")

	store.Set("user", "josh")
	store.Set("role", "staff-ic")
	v, _ := store.Get("user")
	fmt.Println("got:", v)
	store.Delete("role")
	_, err := store.Get("missing")
	fmt.Println("missing:", err)

	fmt.Println(strings.Repeat("-", 40))
	metrics.Report()

	// Key insight: each decorator is unaware of the others.
	// You can reorder, add, or remove decorators at construction time.
	// The base MemoryStore never changed.
}
