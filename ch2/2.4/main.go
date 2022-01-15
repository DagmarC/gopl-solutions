// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package main

// PopCountEx returns the population count (number of set bits) of x.
func PopCountBits(x uint64) int {
	res := 0
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			res++
		}
		x >>= 1
	}
	return res
}

//!-
