# Pointers

Address-of (`&x`) and dereference (`*p`). No pointer arithmetic — Go pointers are safer than C's.

## Worth knowing

- A nil pointer dereference is a runtime **panic**, not undefined behavior.
- `new(T)` allocates a zero-valued T and returns `*T`. Rarely needed — `&T{}` literal is more common.
- Pointers are how you let a function MUTATE its argument. Otherwise Go passes by value.
- Pointer to a struct lets methods modify the struct: `func (p *Point) Move()`. Value receiver makes a copy.

## When to use pointer vs value

- Struct is large (avoid copying): pointer.
- Method needs to mutate: pointer.
- Type holds a mutex, channel, or sync primitive: pointer (don't accidentally copy locks).
- Otherwise: value is fine — and is the default in idiomatic code for small, immutable structs.
