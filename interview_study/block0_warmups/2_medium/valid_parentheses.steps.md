## Hint 1
"Innermost match" → stack. Push openers, pop on closers, verify the popped opener matches.

```go
stack := make([]rune, 0, len(s))
for _, c := range s {
    // push openers, on closer pop and match
}
return len(stack) == 0
```

## Hint 2
Use a `closer → opener` map so the match check is one lookup. Bail early on mismatch or empty stack.

```go
pair := map[rune]rune{')': '(', ']': '[', '}': '{'}
switch c {
case '(', '[', '{':
    stack = append(stack, c)
case ')', ']', '}':
    if len(stack) == 0 || stack[len(stack)-1] != pair[c] {
        return false
    }
    stack = stack[:len(stack)-1]
}
```

## Hint 3
Final state matters: leftover openers also mean invalid, so return `len(stack) == 0`, not just `true`.

```go
return len(stack) == 0
```
