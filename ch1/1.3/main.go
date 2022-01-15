// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var s string
	sep := " "
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += strconv.Itoa(i) + sep + os.Args[i] + "\n"
	}
	fmt.Println(s)
	fmt.Printf("%d s elapsed\n", time.Since(start).Microseconds())
}