# Hello World

Every Go program is a package; the executable one is `main`. `func main()` is the entry point — no return value, no arguments.

## Worth knowing

- `package main` makes the binary; any other name makes a library.
- `import "fmt"` brings in the standard formatting package; unused imports are a **compile error** (not a warning).
- `go run file.go` compiles+runs in one step; `go build` produces a binary you ship.
