//go:build ignore

package main

import (
	"fmt"
)

// STATE MACHINE
// =============
// Problem: model an entity (an order, a connection, a workflow) whose valid
// operations depend on its current state. A naive nested if/switch becomes
// unreadable and lets invalid transitions through.
//
// Pattern: enumerate states, enumerate events, encode the transition table
// explicitly. Two common encodings:
//
//   A) Transition map: map[State]map[Event]State — pure data, easy to inspect.
//   B) Handler-per-state: each state knows how to handle each event, may also
//      run side effects on entry/exit. Heavier but cleaner when transitions
//      have complex actions.
//
// This example uses (A) — the data-driven version — because the transition
// LOGIC is what reviewers want to see. Add per-transition callbacks only if
// you actually need entry/exit hooks.

type State int

const (
	StatePending State = iota
	StatePaid
	StateShipped
	StateDelivered
	StateCancelled
)

func (s State) String() string {
	return [...]string{"PENDING", "PAID", "SHIPPED", "DELIVERED", "CANCELLED"}[s]
}

type Event int

const (
	EventPay Event = iota
	EventShip
	EventDeliver
	EventCancel
)

func (e Event) String() string {
	return [...]string{"PAY", "SHIP", "DELIVER", "CANCEL"}[e]
}

// transitions[currentState][event] = nextState.
// Missing entry → invalid transition.
var transitions = map[State]map[Event]State{
	StatePending: {
		EventPay:    StatePaid,
		EventCancel: StateCancelled,
	},
	StatePaid: {
		EventShip:   StateShipped,
		EventCancel: StateCancelled, // refund handled in side-effect, not shown
	},
	StateShipped: {
		EventDeliver: StateDelivered,
		// note: no Cancel after shipped — business rule encoded in the table
	},
	// StateDelivered and StateCancelled are terminal — no outbound transitions.
}

type Order struct {
	ID    string
	State State
}

func (o *Order) Apply(e Event) error {
	next, ok := transitions[o.State][e]
	if !ok {
		return fmt.Errorf("order %s: cannot %s from %s", o.ID, e, o.State)
	}
	fmt.Printf("order %s: %s --%s--> %s\n", o.ID, o.State, e, next)
	o.State = next
	return nil
}

func main() {
	o := &Order{ID: "A1", State: StatePending}

	// Happy path.
	_ = o.Apply(EventPay)
	_ = o.Apply(EventShip)
	_ = o.Apply(EventDeliver)

	// Invalid: cancel after delivery (terminal state).
	if err := o.Apply(EventCancel); err != nil {
		fmt.Println("rejected:", err)
	}

	// Different path: pay then cancel before ship.
	o2 := &Order{ID: "B2", State: StatePending}
	_ = o2.Apply(EventPay)
	_ = o2.Apply(EventCancel)

	// Invalid: ship before pay.
	o3 := &Order{ID: "C3", State: StatePending}
	if err := o3.Apply(EventShip); err != nil {
		fmt.Println("rejected:", err)
	}
}
