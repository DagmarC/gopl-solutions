package main

import (
	"fmt"
	"io"
	"strings"
)

type LimitReaderer struct {
	r     io.Reader
	limit int64
	i     int64
}

func (lr *LimitReaderer) Read(b []byte) (n int, err error) {
	if lr.i >= lr.limit {
		return 0, io.EOF
	}
	m, err := lr.r.Read(b)
	lr.i += int64(m)
	return
}

func NewLimitReaderer(r io.Reader, limit int64) *LimitReaderer { return &LimitReaderer{r, limit, 0} }

// LimitReader function in the io package accepts an io.Reader r and a number of bytesn,
// and returns another Reader that reads from r but reports an end-of-file condition after n bytes.
func LimitReader(r io.Reader, n int64) io.Reader {
	lr := NewLimitReaderer(r, n)
	b := make([]byte, n)
	lr.Read(b)
	return lr
}

func main() {
	myReader := LimitReader(strings.NewReader("Hello World"), 6)
	fmt.Println(myReader)
}
