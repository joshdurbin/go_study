//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Struct tags control field names and behavior on the wire.
// `omitempty` skips zero values; `-` excludes the field entirely.
type User struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"-"` // never serialize
	Created  time.Time `json:"created"`
}

// Custom marshaling: emit a Duration as a human string ("1h30m") instead of
// the default integer nanoseconds.
type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(d).String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	parsed, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	*d = Duration(parsed)
	return nil
}

func main() {
	// ─── Marshal: Go value -> JSON bytes ───
	u := User{ID: 1, Name: "Ada", Created: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)}
	out, _ := json.Marshal(u)
	fmt.Println(string(out)) // Email omitted, Password absent

	// MarshalIndent for pretty output (debugging only — slower, larger).
	pretty, _ := json.MarshalIndent(u, "", "  ")
	fmt.Println(string(pretty))

	// ─── Unmarshal: JSON bytes -> Go value ───
	var got User
	json.Unmarshal([]byte(`{"id":2,"name":"Lin","email":"lin@x.io"}`), &got)
	fmt.Printf("%+v\n", got)

	// ─── Decoding into map[string]any when the shape is unknown ───
	// Numbers become float64; nested objects become map[string]any.
	var generic map[string]any
	json.Unmarshal([]byte(`{"a":1,"b":{"c":[true,"x"]}}`), &generic)
	fmt.Println(generic) // map[a:1 b:map[c:[true x]]]

	// ─── Custom marshaling round-trip ───
	type Job struct {
		Timeout Duration `json:"timeout"`
	}
	j := Job{Timeout: Duration(90 * time.Second)}
	b, _ := json.Marshal(j)
	fmt.Println(string(b)) // {"timeout":"1m30s"}

	var back Job
	json.Unmarshal(b, &back)
	fmt.Println("parsed:", time.Duration(back.Timeout))

	// ─── Streaming with json.Decoder ───
	// Use Decoder for newline-delimited JSON / large files — it reads one
	// value at a time without buffering the whole stream.
	stream := `{"n":1}{"n":2}{"n":3}`
	dec := json.NewDecoder(strings.NewReader(stream))
	for {
		var v struct{ N int }
		if err := dec.Decode(&v); err != nil {
			break // io.EOF when done
		}
		fmt.Println("decoded:", v.N)
	}
}
