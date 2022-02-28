// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	categ := make(map[UniCategory]int) // counts categories of Unicode characters
	var utflen [utf8.UTFMax + 1]int    // count of lengths of UTF-8 encodings
	invalid := 0                       // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		categ[category(r)]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range categ {
		fmt.Printf("%s\t%d\n", c.String(), n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

type UniCategory int

const (
	letter UniCategory = iota
	digit
	number
	control
	space
	punctuation
	symbol
	mark
	undefined
)

func category(r rune) UniCategory {

	if unicode.IsLetter(r) {
		return letter
	}
	if unicode.IsControl(r) {
		return control
	}
	if unicode.IsNumber(r) {
		return number
	}
	if unicode.IsSymbol(r) {
		return symbol
	}
	if unicode.IsSpace(r) {
		return space
	}
	if unicode.IsPunct(r) {
		return punctuation
	}
	if unicode.IsMark(r) {
		return mark
	}
	return undefined
}

func (c UniCategory) String() string {
	switch c {
	case letter:
		return "Letter"
	case control:
		return "Control"
	case punctuation:
		return "Punctuation"
	case number:
		return "Number"
	case symbol:
		return "Symbol"
	case space:
		return "Space"
	case mark:
		return "Mark"
	default:
		return "undefined"
	}
}

//!-
