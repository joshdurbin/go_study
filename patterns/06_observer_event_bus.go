//go:build ignore

package main

import (
	"fmt"
	"sync"
)

// OBSERVER / EVENT BUS PATTERN
// =============================
// Intent: when one object changes state, dependents are notified automatically.
// In Go, this is typically implemented with channels or callbacks — not the
// classic OOP Subject/Observer class hierarchy.
//
// Two implementations shown:
//   1. Callback-based (simpler, synchronous, good for in-process events)
//   2. Channel-based (async, good for fan-out to goroutines, supports backpressure)
//
// Real-world use: your AlertManager routing, Redis keyspace notifications,
// Prometheus metrics propagation — all observer variants.

// ── Implementation 1: Callback-based EventBus ───────────────────────────────

type EventType string

const (
	EventUserCreated EventType = "user.created"
	EventUserDeleted EventType = "user.deleted"
	EventOrderPlaced EventType = "order.placed"
)

type Event struct {
	Type    EventType
	Payload interface{}
}

type Handler func(Event)

type EventBus struct {
	mu       sync.RWMutex
	handlers map[EventType][]Handler
}

func NewEventBus() *EventBus {
	return &EventBus{handlers: make(map[EventType][]Handler)}
}

// Subscribe registers a handler for an event type. Returns an unsubscribe func.
func (eb *EventBus) Subscribe(eventType EventType, h Handler) func() {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.handlers[eventType] = append(eb.handlers[eventType], h)

	// Return unsubscribe closure — cleaner than an Unsubscribe(id) API
	idx := len(eb.handlers[eventType]) - 1
	return func() {
		eb.mu.Lock()
		defer eb.mu.Unlock()
		handlers := eb.handlers[eventType]
		eb.handlers[eventType] = append(handlers[:idx], handlers[idx+1:]...)
	}
}

// Publish notifies all subscribers synchronously. For async, wrap in go func().
func (eb *EventBus) Publish(e Event) {
	eb.mu.RLock()
	handlers := make([]Handler, len(eb.handlers[e.Type]))
	copy(handlers, eb.handlers[e.Type]) // copy to release lock before calling handlers
	eb.mu.RUnlock()

	for _, h := range handlers {
		h(e)
	}
}

// ── Implementation 2: Channel-based pub/sub (async, with backpressure) ──────

type PubSub struct {
	mu          sync.RWMutex
	subscribers map[EventType][]chan Event
}

func NewPubSub() *PubSub {
	return &PubSub{subscribers: make(map[EventType][]chan Event)}
}

// Subscribe returns a channel that receives events of the given type.
// bufSize controls backpressure — slow subscribers won't block the publisher.
func (ps *PubSub) Subscribe(eventType EventType, bufSize int) <-chan Event {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ch := make(chan Event, bufSize)
	ps.subscribers[eventType] = append(ps.subscribers[eventType], ch)
	return ch
}

// Publish sends to all subscriber channels. Non-blocking: drops if buffer full.
func (ps *PubSub) Publish(e Event) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	for _, ch := range ps.subscribers[e.Type] {
		select {
		case ch <- e:
		default:
			fmt.Printf("WARNING: subscriber buffer full, dropping event %s\n", e.Type)
		}
	}
}

func (ps *PubSub) Close() {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	for _, subs := range ps.subscribers {
		for _, ch := range subs {
			close(ch)
		}
	}
}

func main() {
	fmt.Println("=== Callback-based EventBus ===")
	bus := NewEventBus()

	// Multiple handlers for the same event
	bus.Subscribe(EventUserCreated, func(e Event) {
		fmt.Printf("audit: user created: %v\n", e.Payload)
	})
	bus.Subscribe(EventUserCreated, func(e Event) {
		fmt.Printf("email: welcome email queued for %v\n", e.Payload)
	})

	unsubOrder := bus.Subscribe(EventOrderPlaced, func(e Event) {
		fmt.Printf("inventory: order placed: %v\n", e.Payload)
	})

	bus.Publish(Event{Type: EventUserCreated, Payload: "alice"})
	bus.Publish(Event{Type: EventOrderPlaced, Payload: map[string]int{"item-42": 3}})

	unsubOrder() // dynamic unsubscription
	bus.Publish(Event{Type: EventOrderPlaced, Payload: "ignored after unsub"})

	fmt.Println("\n=== Channel-based PubSub ===")
	ps := NewPubSub()
	var wg sync.WaitGroup

	ch1 := ps.Subscribe(EventUserDeleted, 5)
	ch2 := ps.Subscribe(EventUserDeleted, 5)

	// Goroutine consumers
	for i, ch := range []<-chan Event{ch1, ch2} {
		wg.Add(1)
		go func(id int, c <-chan Event) {
			defer wg.Done()
			for e := range c {
				fmt.Printf("consumer-%d received: %v\n", id, e.Payload)
			}
		}(i+1, ch)
	}

	ps.Publish(Event{Type: EventUserDeleted, Payload: "bob"})
	ps.Publish(Event{Type: EventUserDeleted, Payload: "carol"})
	ps.Close()
	wg.Wait()
}
