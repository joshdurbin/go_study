# Testing

Go's `testing` package is built into the toolchain. Files ending in `_test.go` are compiled only by `go test`, so test code never ships in your binary.

## Worth knowing

- Tests are functions named `TestX(t *testing.T)`. They live in `*_test.go`, usually in the same package as the code (or `pkg_test` for black-box tests).
- Run with `go test ./...`. Useful flags: `-run TestX` (filter), `-v` (verbose), `-race` (data race detector), `-cover` (coverage).
- `t.Run("name", func(t *testing.T) { ... })` creates a subtest — required for table-driven tests with named cases.
- `t.Parallel()` marks a test as safe to run alongside others. Add it to both the parent and each `t.Run` for full parallelism.
- `t.Fatal` / `t.Fatalf` stop the current test immediately; `t.Error` / `t.Errorf` record a failure and continue. Use Fatal only when later assertions depend on the failed one.

## Benchmarks

`func BenchmarkX(b *testing.B)` runs the body `b.N` times; the runner auto-tunes `N`. Run with `go test -bench=. -benchmem`. Call `b.ResetTimer()` after setup so it isn't charged to the measurement.

## Common gotcha

Pre-Go 1.22, loop variables were shared across iterations — `for _, tc := range cases { t.Run(tc.name, func(t *testing.T) { ... }) }` would race on `tc` when subtests ran in parallel. The fix was `tc := tc` at the top of the loop body. Go 1.22+ scopes loop vars per iteration, so this is no longer needed on modern toolchains — but you'll see `tc := tc` everywhere in older code.

## Interview frame

Bring up table-driven tests when discussing testability. They're the idiomatic Go pattern for "given N inputs, expect N outputs" — concise, named subtests, easy to extend.
