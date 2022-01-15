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
