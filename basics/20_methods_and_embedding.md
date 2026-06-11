# Methods and Embedding

Methods are functions with a typed receiver. Embedding is how Go does "inheritance" — except it's composition.

## Methods

```go
type Point struct{ X, Y int }
func (p Point) DistFromOrigin() float64 { ... } // value receiver
func (p *Point) Move(dx, dy int)        { ... } // pointer receiver
```

- **Value receiver**: method gets a copy. Can't mutate the original.
- **Pointer receiver**: method gets the address. Can mutate.

## Choosing receiver type

- Mutates the receiver → pointer.
- Type contains a mutex, channel, or other sync primitive → pointer (don't copy locks).
- Large struct (avoid copying) → pointer.
- Small, immutable type → value.
- **Be consistent within a type**: don't mix value and pointer receivers on the same type unless you really need to.

## Embedding

```go
type ReadWriter struct {
    io.Reader
    io.Writer
}
```

Methods of the embedded type are **promoted** to the outer type — `rw.Read(...)` works directly. This is Go's compositional alternative to inheritance.

## Worth knowing

- Embedding gives you method promotion, not "is-a" semantics. A `ReadWriter` is NOT an `io.Reader` for type-assertion purposes — but it satisfies the `io.Reader` interface.
- You can override a promoted method by defining one with the same name on the outer type.
