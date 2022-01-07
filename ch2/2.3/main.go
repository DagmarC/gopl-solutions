// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package main

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountEx returns the population count (number of set bits) of x.
func PopCountEx(x uint64) int {
	// Ex 2.3
	// Ex 2.3
	r := byte(0)
	for i := 0; i < 8; i++ {
		r += pc[byte(x>>(i*8))]
	}
	return int(r)

}

//!-
