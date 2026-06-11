## Hint 1
Keep a monotonic **increasing** stack of indices. While the stack stays increasing, every bar's rectangle is still "open to the right." A taller-then-shorter break is what triggers a pop and area calculation.

```go
stack := make([]int, 0, len(heights)+1)
best := 0
for i := 0; i <= len(heights); i++ {
    // sentinel 0 at i == len(heights) to drain stack
}
```

## Hint 2
Sentinel trick: virtually append a 0 at the end so every bar still on the stack gets popped and measured. Without it, the tail of an increasing histogram is never evaluated.

```go
var h int
if i == len(heights) {
    h = 0
} else {
    h = heights[i]
}
```

## Hint 3
On pop, the popped bar's rectangle spans from **just after the new top of the stack** to **just before i**. That's `i - newTop - 1`. If the stack is empty after popping, the bar extends from index 0, so width = i.

```go
for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
    top := stack[len(stack)-1]
    stack = stack[:len(stack)-1]
    width := i
    if len(stack) > 0 {
        width = i - stack[len(stack)-1] - 1
    }
    // area = heights[top] * width
}
```

## Hint 4
Track the running max and push the current index after the inner pop-loop drains. Easy off-by-one: don't push the sentinel index? Actually it's fine — the sentinel only ever pops, never blocks future iterations since the outer loop ends.

```go
if area := heights[top] * width; area > best {
    best = area
}
// after inner loop:
stack = append(stack, i)
// ...
return best
```
