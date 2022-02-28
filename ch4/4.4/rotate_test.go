package main

import "testing"

type data struct {
	expected [6]int // arr is for comparisons reasons, arrays can be compared via ==, slices not
	input    [6]int
	shift    int
}

var testData = []data{
	{expected: [6]int{3, 4, 5, 6, 1, 2}, input: [6]int{1, 2, 3, 4, 5, 6}, shift: 2},
	{expected: [6]int{2, 3, 4, 5, 6, 1}, input: [6]int{1, 2, 3, 4, 5, 6}, shift: 1},
	{expected: [6]int{1, 2, 3, 4, 5, 6}, input: [6]int{1, 2, 3, 4, 5, 6}, shift: 0},
	{expected: [6]int{1, 2, 3, 4, 5, 6}, input: [6]int{1, 2, 3, 4, 5, 6}, shift: 6},
	{expected: [6]int{3, 4, 5, 6, 1, 2}, input: [6]int{1, 2, 3, 4, 5, 6}, shift: 8},
}

func TestRotateLeft(t *testing.T) {
	for _, td := range testData {
		rotateL(td.input[:], td.shift)
		if td.input != td.expected {
			t.Errorf("error expected: %v, actual: %v and shift: %d", td.expected, td.input, td.shift)
		}
	}
}
