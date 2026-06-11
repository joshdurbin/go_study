//go:build ignore

package main

import (
	"fmt"
	"strings"
)

// BUILDER PATTERN
// ===============
// Problem: complex object with many required AND optional fields, plus
// validation rules that depend on combinations of fields.
//
// Builder vs. Functional Options (see patterns/01):
//   - Functional options: best when fields are independent and options are
//     "set this one thing". The Go idiom in stdlib and most libraries.
//   - Builder: best when the object has a multi-step construction with
//     intermediate state, or when fields interact (e.g., a query DSL where
//     order/composition matters).
//
// Interview question: "When would you pick builder over functional options?"
// Answer: when build steps need to be chainable in a specific order, when
// validation must happen at Build() time across the whole config, or when
// you want to enforce required fields at compile time via type states.

type Query struct {
	table   string
	wheres  []string
	limit   int
	orderBy string
}

// QueryBuilder accumulates state. Each method returns *QueryBuilder for chaining.
type QueryBuilder struct {
	q   Query
	err error // first error wins; final Build() reports it
}

func From(table string) *QueryBuilder {
	if table == "" {
		return &QueryBuilder{err: fmt.Errorf("table required")}
	}
	return &QueryBuilder{q: Query{table: table}}
}

func (b *QueryBuilder) Where(cond string) *QueryBuilder {
	if b.err != nil {
		return b
	}
	b.q.wheres = append(b.q.wheres, cond)
	return b
}

func (b *QueryBuilder) Limit(n int) *QueryBuilder {
	if b.err != nil {
		return b
	}
	if n < 0 {
		b.err = fmt.Errorf("limit must be >= 0")
		return b
	}
	b.q.limit = n
	return b
}

func (b *QueryBuilder) OrderBy(col string) *QueryBuilder {
	if b.err != nil {
		return b
	}
	b.q.orderBy = col
	return b
}

// Build runs final cross-field validation, returns the built object or the
// first error encountered along the way.
func (b *QueryBuilder) Build() (string, error) {
	if b.err != nil {
		return "", b.err
	}
	var sb strings.Builder
	fmt.Fprintf(&sb, "SELECT * FROM %s", b.q.table)
	if len(b.q.wheres) > 0 {
		fmt.Fprintf(&sb, " WHERE %s", strings.Join(b.q.wheres, " AND "))
	}
	if b.q.orderBy != "" {
		fmt.Fprintf(&sb, " ORDER BY %s", b.q.orderBy)
	}
	if b.q.limit > 0 {
		fmt.Fprintf(&sb, " LIMIT %d", b.q.limit)
	}
	return sb.String(), nil
}

func main() {
	q, err := From("users").
		Where("active = true").
		Where("age > 18").
		OrderBy("created_at").
		Limit(50).
		Build()
	fmt.Println(q, err)

	// Error captured at the offending step, surfaced at Build().
	_, err = From("orders").Limit(-1).Where("x").Build()
	fmt.Println("err:", err)

	// Missing required field caught at the entry point.
	_, err = From("").Build()
	fmt.Println("err:", err)
}
