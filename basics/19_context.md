# Context

`context.Context` carries **cancellation signals**, **deadlines**, and **request-scoped values** across API boundaries and goroutines.

## The four constructors

- `context.Background()` — root context, never cancels. Use at the top of `main` or a request handler.
- `context.TODO()` — placeholder when you'll add a real context later.
- `context.WithCancel(parent)` — manual cancel via the returned `cancel()` function.
- `context.WithTimeout(parent, d)` / `context.WithDeadline(parent, t)` — cancels automatically.

## How to use it

- Accept `ctx context.Context` as the **first parameter** of any function that does I/O or can take a while.
- Pass the same ctx down. Don't create new background contexts in lower layers.
- Check cancellation: `select { case <-ctx.Done(): return ctx.Err() }`.
- Always call `cancel()` (usually via `defer cancel()`), even if the timeout will trigger — it releases resources immediately.

## Don't

- Don't store a context in a struct field. Pass it through call args.
- Don't use `ctx.Value()` for arguments that should be regular parameters. Reserve it for true request-scoped data (request ID, auth, trace span).

## Worth knowing

Almost every stdlib I/O function (HTTP, DB, file) accepts a context — propagating timeout/cancellation works "for free" as long as you pass ctx through.
