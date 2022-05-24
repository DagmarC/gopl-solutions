package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"

	outline "github.com/DagmarC/gopl-solutions/ch5/5.2"
)

type SimpleHtmlReader struct {
	s string
}

// Read implements the io.Reader interface. It reads all input at once and when it is done it signals io.EOF.
func (r *SimpleHtmlReader) Read(b []byte) (n int, err error) {
	n = copy(b, r.s)
	err = io.EOF //All is read, so signal it with EOF flag.
	return
}

// NewReader returns a new Reader reading from s.
// It is similar to bytes.NewBufferString but more efficient and read-only.
func NewReader(s string) *SimpleHtmlReader { return &SimpleHtmlReader{s} }

func main() {
	var sr SimpleHtmlReader
	sr.s = "<html><body>hello</body></html>"

	doc, err := html.Parse(&sr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	elements := make(map[string]int, 50)
	outline.Outline(elements, doc)
	fmt.Println(elements)
}
