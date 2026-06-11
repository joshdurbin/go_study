# encoding/json

Standard library JSON support. Reflection-based, decent performance, zero dependencies. The interview-relevant surface is small.

## Worth knowing

- `json.Marshal(v)` returns `([]byte, error)`. `json.Unmarshal(data, &v)` parses into a pointer — pass by pointer or it can't write back.
- Struct tags drive serialization: `json:"name"` renames, `json:"name,omitempty"` skips zero values, `json:"-"` excludes the field.
- Only exported (capitalized) fields are serialized. A lowercase field is invisible to encoding/json.
- When the shape is unknown, decode into `map[string]any` or `any`. Numbers become `float64` by default — use `json.Decoder.UseNumber()` if precision matters.
- `json.NewDecoder(r).Decode(&v)` streams one value at a time. Use it for newline-delimited JSON, large files, or HTTP bodies — never `io.ReadAll` then `Unmarshal` when you can avoid it.

## Custom serialization

Implement `MarshalJSON() ([]byte, error)` on a value receiver and `UnmarshalJSON([]byte) error` on a pointer receiver. Use this for types whose default representation is wrong (durations, enums, opaque IDs).

## Common gotcha

`omitempty` checks Go's zero value, not "missing." A `bool` field with `omitempty` will be omitted when `false` — even if the caller explicitly set it to false. For nullable optionals, use a pointer (`*bool`) or `sql.NullBool`-style wrapper so absent and zero are distinguishable.

## Interview frame

If asked about API design: mention struct tags, `omitempty`, and the pointer-for-nullable pattern. If asked about performance: streaming with `json.Decoder` over buffer-then-decode, and reach for `easyjson` / `jsoniter` only when profiling demands it.
