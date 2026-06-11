//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// io.Reader and io.Writer are the two most important interfaces in the stdlib.
// Anything that produces bytes implements Reader; anything that consumes them
// implements Writer. Files, sockets, buffers, HTTP bodies, and hashes all
// satisfy these interfaces — so the same code works on all of them.

// chunkReader returns its data in fixed-size chunks. Real readers do this too:
// a TCP socket gives you whatever bytes have arrived so far, not a full message.
type chunkReader struct {
	data []byte
	size int
}

func (c *chunkReader) Read(p []byte) (int, error) {
	if len(c.data) == 0 {
		return 0, io.EOF // signal "no more bytes ever"
	}
	n := c.size
	if n > len(c.data) {
		n = len(c.data)
	}
	if n > len(p) {
		n = len(p)
	}
	copy(p, c.data[:n])
	c.data = c.data[n:]
	return n, nil
}

func main() {
	// ─── strings.NewReader: cheapest way to wrap a string as a Reader ───
	r := strings.NewReader("hello world")
	buf := make([]byte, 5)
	n, _ := r.Read(buf)
	fmt.Printf("read %d bytes: %q\n", n, buf[:n]) // read 5 bytes: "hello"

	// ─── bytes.Buffer: implements BOTH Reader and Writer ───
	// Handy for building up bytes in tests or composing transformations.
	var b bytes.Buffer
	b.WriteString("foo")
	b.WriteString("-bar")
	fmt.Println("buffer:", b.String()) // foo-bar

	// ─── io.Copy: pump bytes from any Reader to any Writer ───
	// The whole stdlib I/O story is built on this one function.
	src := strings.NewReader("copied via io.Copy\n")
	var dst bytes.Buffer
	io.Copy(&dst, src)
	fmt.Print(dst.String())

	// ─── io.TeeReader: read from src, simultaneously write to a side channel ───
	// Useful for hashing/logging a stream as it's consumed.
	src2 := strings.NewReader("teed bytes")
	var spy bytes.Buffer
	tee := io.TeeReader(src2, &spy)
	consumed, _ := io.ReadAll(tee)
	fmt.Printf("consumed=%q spy=%q\n", consumed, spy.String())

	// ─── Custom Reader: returns bytes in 3-byte chunks ───
	// io.ReadAll loops until EOF, so chunked behavior is transparent to callers.
	cr := &chunkReader{data: []byte("abcdefghi"), size: 3}
	all, _ := io.ReadAll(cr)
	fmt.Printf("chunked read: %q\n", all) // abcdefghi
}
