## Hint 1
Two pointers from opposite ends. Skip non-alphanumeric runes, case-fold before comparing.

```go
r := []rune(s)
i, j := 0, len(r)-1
for i < j {
    // skip non-alnum at either end, then compare case-folded
}
return true
```

## Hint 2
Use a switch to keep the cases tidy: skip left, skip right, mismatch, otherwise advance both.

```go
switch {
case !isAlnum(r[i]):                 i++
case !isAlnum(r[j]):                 j--
case unicode.ToLower(r[i]) != unicode.ToLower(r[j]):
    return false
default:
    i++; j--
}
```
