//go:build ignore

package main

import (
	"fmt"
	"strings"
)

// Real tests live in *_test.go files and run via `go test`. This file demos
// the SHAPE of a test function and the conventions; main() runs equivalent
// assertions inline so you can see the behavior without `go test`.

// Reverse — the function under test.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

/*
A real _test.go file would look like this:

	package main

	import "testing"

	func TestReverse(t *testing.T) {
	    got := Reverse("abc")
	    if got != "cba" {
	        t.Errorf("Reverse(\"abc\") = %q; want %q", got, "cba")
	    }
	}

	// Table-driven: one TestX, many subtests via t.Run.
	func TestReverseTable(t *testing.T) {
	    cases := []struct{
	        name, in, want string
	    }{
	        {"empty",   "",     ""},
	        {"ascii",   "abc",  "cba"},
	        {"unicode", "héllo","olléh"},
	    }
	    for _, tc := range cases {
	        tc := tc                      // capture loop var (pre-Go 1.22)
	        t.Run(tc.name, func(t *testing.T) {
	            t.Parallel()              // run subtests concurrently
	            if got := Reverse(tc.in); got != tc.want {
	                t.Errorf("Reverse(%q) = %q; want %q", tc.in, got, tc.want)
	            }
	        })
	    }
	}

	// Fatal stops the current test immediately; Error records and continues.
	// Use Fatal when later assertions depend on the failed one.
	func TestSetup(t *testing.T) {
	    db, err := openDB()
	    if err != nil {
	        t.Fatalf("setup: %v", err)    // can't continue without db
	    }
	    defer db.Close()
	    // ... assertions that use db ...
	}

	// Benchmarks run via `go test -bench=.`. b.N is auto-tuned by the runner.
	func BenchmarkReverse(b *testing.B) {
	    s := strings.Repeat("a", 1024)
	    b.ResetTimer()                    // exclude setup from timing
	    for i := 0; i < b.N; i++ {
	        _ = Reverse(s)
	    }
	}
*/

// assertEq is a tiny stand-in to print pass/fail without the testing package.
func assertEq(name, got, want string) {
	if got == want {
		fmt.Printf("PASS %s\n", name)
	} else {
		fmt.Printf("FAIL %s: got %q want %q\n", name, got, want)
	}
}

func main() {
	// ─── Inline assertions (simulating what TestReverseTable would check) ───
	cases := []struct{ name, in, want string }{
		{"empty", "", ""},
		{"ascii", "abc", "cba"},
		{"unicode", "héllo", "olléh"},
	}
	for _, tc := range cases {
		assertEq(tc.name, Reverse(tc.in), tc.want)
	}

	// ─── Why table-driven? ───
	// One function, many inputs. Adding a case is one struct literal — not a
	// new test function. t.Run gives each row its own name in test output.
	fmt.Println("see comments above for real _test.go shape")
	fmt.Println("test names:", strings.Join([]string{"empty", "ascii", "unicode"}, ", "))
}
