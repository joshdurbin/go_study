# Table-Driven Tests

The idiomatic Go test pattern: a slice of cases, one assertion loop.

## The shape

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"zero", 0, 0, 0},
        {"positive", 2, 3, 5},
        {"negative", -1, -1, -2},
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            if got := Add(tc.a, tc.b); got != tc.want {
                t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.want)
            }
        })
    }
}
```

## Why this is idiomatic

- Each case is data, not code — easy to add, easy to read.
- `t.Run` gives each case its own name in failure output and lets you `go test -run TestAdd/positive`.
- Parallel-safe with `t.Parallel()` inside the subtest.

## When NOT to use

- The cases each need very different setup — separate tests are cleaner.
- The thing being tested is one-and-done, not a function with input/output shape.

## Worth knowing

- Use `t.Fatalf` (stops the subtest) for setup failures, `t.Errorf` (continues) for assertion failures so you see all failures in one run.
- `t.Helper()` in helper functions makes failure traces point at the caller, not the helper.

This file is in `patterns/` because table-driven tests are the de-facto Go test architecture — fluency with this is interview-table-stakes.
