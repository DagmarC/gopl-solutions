package main

import (
	"crypto/sha256"
	"fmt"
)

const byteLen = 8

func main() {
	sh1 := sha256.Sum256([]byte("x"))
	sh2 := sha256.Sum256([]byte("X"))

	d1, d2 := sha256cmp(&sh1, &sh2)
	fmt.Printf("Number of bits that are differrent is %d and %d.\n", d1, d2)
}

// sha256cmp counts the number of bits that are different.
func sha256cmp(sh1, sh2 *[32]byte) (int, int) {
	diff := 0
	diff2 := 0

	for i := 0; i < len(sh1); i++ {
		diff += bitDiff(sh1[i], sh2[i])

		bitXor := sh1[i] ^ sh2[i]
		diff2 += bitCount(bitXor)
	}
	return diff, diff2
}

func bitDiff(sh1, sh2 byte) int {

	diff := 0
	for i := 0; i < byteLen; i++ {
		if (sh1&(1<<i))^(sh2&(1<<i)) != 0 {
			diff++
		}
	}
	return diff
}

func bitCount(x uint8) int {
	c := 0
	for x != 0 {
		fmt.Printf("%b\n", x)
		fmt.Printf("%b\n", x-1)

		x = x & (x-1)
		c++
		
		fmt.Printf("%b\n", x)
		fmt.Println("===========")

	}
	return c
}
