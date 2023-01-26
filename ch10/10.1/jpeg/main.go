// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 287.

//!+main

// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

var outFmt = flag.String("output", "jpeg", "output formats: jpeg/png/gif")

func main() {
	if err := convert(os.Stdout, os.Stdin); err != nil {
		fmt.Fprintf(os.Stderr, "convert: %v\n", err)
		os.Exit(1)
	}
}

func convert(out io.Writer, in io.Reader) error {
	img, kind, err := image.Decode(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "decode: %v\n", err)
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)

	switch *outFmt {
	case "jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})

	case "png":
		return png.Encode(out, img)

	case "gif":
		return gif.Encode(out, img, &gif.Options{NumColors: 256})
	default:
		return errors.New("unsupported format")
	}
}

//!-main

/*
//!+with
$ go build gopl.io/ch3/mandelbrot
$ go build gopl.io/ch10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
Input format = png
//!-with

//!+without
$ go build gopl.io/ch10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
jpeg: image: unknown format
//!-without
*/
