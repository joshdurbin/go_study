# Repository Pattern

Abstract data access behind an interface. Business logic depends on the interface, not on the concrete storage (SQL, in-memory, mock).

## The shape

```go
type UserRepo interface {
    Get(ctx context.Context, id int) (User, error)
    Save(ctx context.Context, u User) error
}

type SQLUserRepo struct{ db *sql.DB }
type MemUserRepo struct{ data map[int]User }
```

Service layer accepts a `UserRepo`; tests inject `MemUserRepo`, production injects `SQLUserRepo`.

## When to use

- Business logic that needs to be testable without a real database.
- Multiple data sources for the same entity (cache + DB, primary + replica).
- Migration from one storage to another — the interface stays, the implementation swaps.

## When NOT to use

- Single-storage app that will never change — direct `*sql.DB` calls are fine. Don't add an interface for one implementation.
- The repository has 30 methods — that's a service masquerading as a data layer. Split it.

## Common pitfall

Don't leak SQL types or driver-specific errors through the interface. Translate to domain errors (`ErrNotFound`, `ErrConflict`) at the repository boundary.

## Real-world

Almost every "clean architecture" / "hexagonal architecture" Go codebase uses this. Combined with sqlc (which generates concrete impls from queries) it gives type-safe SQL and testable interfaces.
