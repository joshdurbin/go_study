//go:build ignore
// +build ignore

// TABLE-DRIVEN TESTS
// ==================
// Not runnable with `go run` — this is a test file pattern.
// In a real repo: save as foo_test.go and run with `go test ./...`
//
// Table-driven tests are THE idiomatic Go testing pattern.
// Every standard library test uses this approach.
// If you write tests in an interview, write table-driven tests.
//
// Benefits:
//   - Add new cases by adding a row — no new test function
//   - Test name describes what's being tested, not how
//   - Edge cases are visually obvious in the table
//   - t.Run() gives each case its own subtest (runnable individually)
//
// `go test -run TestDivide/divide_by_zero` runs just that subtest.
// `go test -v` shows each subtest name and result.

package main

import (
	"errors"
	"testing"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func TestDivide(t *testing.T) {
	// The table: a slice of anonymous structs, one per test case.
	// Fields: name (for t.Run), inputs, expected outputs.
	tests := []struct {
		name    string
		a, b    float64
		want    float64
		wantErr bool
	}{
		{
			name: "positive numbers",
			a:    10, b: 2,
			want: 5,
		},
		{
			name: "negative dividend",
			a:    -9, b: 3,
			want: -3,
		},
		{
			name: "fractional result",
			a:    1, b: 3,
			want: 0.3333333333333333,
		},
		{
			name:    "divide by zero",
			a:       5, b: 0,
			wantErr: true,
		},
		{
			name: "zero dividend",
			a:    0, b: 5,
			want: 0,
		},
	}

	for _, tc := range tests {
		tc := tc // capture loop variable for parallel subtests
		t.Run(tc.name, func(t *testing.T) {
			// t.Parallel() // uncomment to run subtests in parallel
			got, err := divide(tc.a, tc.b)

			if tc.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Use a small epsilon for float comparison
			const eps = 1e-9
			diff := got - tc.want
			if diff < -eps || diff > eps {
				t.Errorf("divide(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

// BENCHMARK: measure performance of a function
// Run with: go test -bench=. -benchmem
func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		divide(float64(i+1), 3.14)
	}
}

// HELPER PATTERN: testify-style helper with t.Helper()
// t.Helper() marks this as a helper so errors point to the call site, not here.
func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

// MOCK PATTERN: use an interface + test double
// In real code, define the interface in the package being tested.
// The test file provides a fake implementation.

type UserFetcher interface {
	FetchUser(id int) (string, error)
}

type MockUserFetcher struct {
	// What the mock should return — set per test
	ReturnUser string
	ReturnErr  error
	// Track calls for assertion
	CallCount int
	LastID    int
}

func (m *MockUserFetcher) FetchUser(id int) (string, error) {
	m.CallCount++
	m.LastID = id
	return m.ReturnUser, m.ReturnErr
}

func TestWithMock(t *testing.T) {
	tests := []struct {
		name      string
		mock      MockUserFetcher
		inputID   int
		wantUser  string
		wantErr   bool
	}{
		{
			name:     "successful fetch",
			mock:     MockUserFetcher{ReturnUser: "alice"},
			inputID:  1,
			wantUser: "alice",
		},
		{
			name:    "fetch error",
			mock:    MockUserFetcher{ReturnErr: errors.New("db down")},
			inputID: 2,
			wantErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			mock := tc.mock // copy
			user, err := mock.FetchUser(tc.inputID)

			assertEqual(t, mock.CallCount, 1)
			assertEqual(t, mock.LastID, tc.inputID)

			if tc.wantErr {
				if err == nil {
					t.Error("expected error")
				}
				return
			}
			assertNoError(t, err)
			assertEqual(t, user, tc.wantUser)
		})
	}
}
