package main

import (
	"fmt"
)

func main() {
	b := [...]int{1, 2, 3, 4, 5, 6}

	rotateL(b[:], 4)
	fmt.Println(b) // "[5 4 3 2 1 0]"
}
// rotateLeft rotates the slice by n times to the left
func rotateL(s []int, n int) {
	if n < 0 {
		return
	}
	if n >= len(s) {
		n %= len(s)
	}
	
	tmp := make([]int, n)
	copy(tmp, s[:n])
	copy(s, s[n:])
	copy(s[len(s)-n:], tmp)
}
