package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter struct {
	count int64
	w     io.Writer
}

func (b *ByteCounter) Write(p []byte) (int, error) {
	b.count = int64(len(p))
	return len(p), nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var bc ByteCounter
	bc.w = w
	return &bc, &bc.count
}

func main() {
	newWriter, pcLen := CountingWriter(os.Stdout)
	fmt.Fprintf(newWriter, "Hello")
	fmt.Println(*pcLen)
	fmt.Fprintf(newWriter, "Hello World")
	fmt.Println(*pcLen)
}
