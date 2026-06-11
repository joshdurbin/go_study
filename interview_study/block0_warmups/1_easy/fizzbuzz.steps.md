## Hint 1
Only two divisors. "FizzBuzz" is "Fizz"+"Buzz" — build the label by concatenation, not three if-else branches.

```go
for i := 1; i <= n; i++ {
    var b strings.Builder
    // ... append Fizz / Buzz / number
    out = append(out, b.String())
}
```

## Hint 2
Append "Fizz" if i%3==0, "Buzz" if i%5==0. If nothing was appended, write the number.

```go
if i%3 == 0 { b.WriteString("Fizz") }
if i%5 == 0 { b.WriteString("Buzz") }
if b.Len() == 0 {
    fmt.Fprintf(&b, "%d", i)
}
```
