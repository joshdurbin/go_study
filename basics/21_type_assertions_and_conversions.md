# Type Assertions and Conversions

Two different operations that beginners confuse.

## Conversion

Changes the type representation. Always explicit, both types must be related (e.g., numeric, or based on the same underlying type).

```go
var i int = 42
var f float64 = float64(i)
var b []byte = []byte("hello")
```

## Type assertion

Extracts an interface's underlying concrete type. Only works on interface values.

```go
var x any = "hello"
s := x.(string)        // panics if x isn't a string
s, ok := x.(string)    // ok = false if not — safe form
```

**Always use the two-value form unless you can prove the type is correct.**

## Type switch

The multi-case form:

```go
switch v := x.(type) {
case int:    fmt.Println("int:", v)
case string: fmt.Println("string:", v)
default:     fmt.Println("other")
}
```

`v` has the concrete type inside each case branch.

## Worth knowing

- The interface→concrete assertion is **runtime work** — it does a type tag comparison. Cheap, but not free.
- Type switches are the idiomatic way to handle "this is one of N known types" (custom AST visitors, error variants, etc.).
