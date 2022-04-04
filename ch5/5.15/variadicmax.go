package main

import (
	"fmt"
	"math"
)

func main() {
	vals := []int{11, -11, 2323, 213, -323, 333}
	empty := []int{}
	fmt.Printf("max of %v is %d\n", vals, max(10, vals...))
	fmt.Printf("min of %v is %d\n", vals, min(vals...))
	fmt.Printf("min of empty is %d\n", min(empty...))
}

// max calculates the max number from given sequence of numbers and number and reguires at least one argument.
func max(number int, vals ...int) int {

	max := number
	for _, n := range vals {
		if n > int(max) {
			max = n
		}
	}
	return max
}

func min(vals ...int) int {

	min := math.MaxInt
	if len(vals) == 0 {
		return -min
	}
	for _, n := range vals {
		if n < int(min) {
			min = n
		}
	}
	return min
}
