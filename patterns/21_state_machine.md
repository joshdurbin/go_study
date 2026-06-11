# State Machine

Model an entity (order, connection, workflow) whose valid operations depend on its current state.

## Two encodings

- **A) Transition map**: `map[State]map[Event]State`. Pure data, easy to inspect, easy to test. Default choice.
- **B) Handler-per-state**: each state knows how to handle each event, may run entry/exit hooks. Heavier, but cleaner when transitions have complex side effects.

## When to use

- Order lifecycle: PENDING → PAID → SHIPPED → DELIVERED.
- Connection lifecycle: CONNECTING → CONNECTED → CLOSING → CLOSED.
- Protocol implementations: TCP, TLS handshakes, retry-with-backoff phases.

## When NOT to use

- Two states and one transition — that's a bool flag. Don't over-engineer.
- Behavior is determined by the EVENT alone, not the current state — that's a function dispatcher.

## Why the explicit table

A nested if/switch in the caller becomes unreadable past 3-4 states. The table lets you SEE all valid transitions at once — readers instantly know which transitions ARE NOT valid. Missing entry = invalid transition.

## Worth knowing

- Terminal states (DELIVERED, CANCELLED) have no outbound transitions. That's a feature, not an oversight — encode finality.
- For complex workflows, look at `github.com/looplab/fsm` — but most production state machines should NOT need a library. The 30-line `map[State]map[Event]State` is sharper and you understand exactly what it does.

## Interview frame

"Implement an order with valid status transitions" → state machine with explicit transition table. Bonus: explain how you'd add audit logging on every transition.
