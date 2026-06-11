# Structs

Aggregations of named fields. Go has no classes — structs + methods do everything classes do, with composition instead of inheritance.

## Worth knowing

- Struct literal: `Point{X: 1, Y: 2}` (named) or `Point{1, 2}` (positional, fragile — avoid in production).
- Comparable if all fields are comparable. `==` does field-wise equality.
- Embedded fields promote: `type Animal struct { Name string }; type Dog struct { Animal; Breed string }` — `d.Name` works directly.
- Field tags (`` `json:"name"` ``) attach metadata used by reflection (json, sqlc, validate, etc.).

## Common gotcha

Unexported (lowercase) fields aren't visible to other packages — including encoding packages. `json.Marshal` silently skips them.
