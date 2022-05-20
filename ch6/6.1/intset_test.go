// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"fmt"
	"testing"
)

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

type TestValues struct {
	has  IntSet
	want int
}

var testCases = []TestValues{
	{has: IntSet{words: []uint64{1}}, want: 1},        // representation of words: [...0000 0001,] means only one element {0} (64*i+j)[i=0, j=0], where i is the words index in []uint64 slice and j is the index inside uint64 word
	{has: IntSet{words: []uint64{16, 1, 2}}, want: 3}, // representation of words: [...0001 0000, ...0000 0001,...0000 0010] -> {4, 64 (64*i+j)[i=1, j=0], 129(64*i+j)[i=]}
}

func TestLen(t *testing.T) {

	for _, tc := range testCases {
		fmt.Println(tc.has.Len(), tc.has.String())
		if tc.has.Len() != tc.want {
			t.Fail()
		}
	}
}

func TestRemove(t *testing.T) {

	for _, tc := range testCases {
		tc.has.Remove(4)
		if tc.has.Has(4) {
			t.FailNow()
		}
	}
}

func TestClear(t *testing.T) {
	for _, tc := range testCases {
		tc.has.Clear()
		for _, w := range tc.has.words {
			if w != 0 {
				t.Fail()
			}
		}
	}
}

func TestCopy(t *testing.T) {
	for _, tc := range testCases {
		cp := tc.has.Copy()
		fmt.Println(tc.has.String(), cp) // Note that tc.has is not a pointer but String() method is defined on the pointer receiver *IntSet and that is why we need to either pass String() method explicitely that will make the conversion automatically or pass the pointer type as cp.
		if tc.has.String() != cp.String() {
			t.FailNow()
		}
	}
}

func TestAddAll(t *testing.T) {
	expected := "{0 1 2 3}"

	testCases[0].has.AddAll(1, 2, 3)
	if testCases[0].has.String() != expected {
		t.Fail()
	}
}

func TestIntersectsWith(t *testing.T) {

	expected := "{}"

	t1 := testCases[0].has
	t2 := testCases[1].has

	t1.IntersectsWith(&t2)
	if t1.String() != expected {
		t.Fail()
	}

	var t3 IntSet
	t3.AddAll(11, 129)
	expected = "{129}"

	t2.IntersectsWith(&t3)
	if t2.String() != expected {
		t.Fail()
	}
}

func TestDifferenceWith(t *testing.T) {
	expected := "{129}"

	t1 := testCases[1].has
	var t3 IntSet
	t3.AddAll(4, 64)

	t1.DifferenceWith(&t3)
	fmt.Println(t1.String())
	if t1.String() != expected {
		t.Fail()
	}
}

func TestSymmetricDifference(t *testing.T) {
	expected := "{4 32 33 129 233 431}"

	t1 := testCases[1].has
	var t3 IntSet
	t3.AddAll(64, 32, 33, 233, 431)

	t1.SymmetricDifference(&t3)
	fmt.Println(t1.String())
	if t1.String() != expected {
		t.Fail()
	}

	expected = "{}"
	t4 := testCases[0].has

	var t5 IntSet
	t5.AddAll(0)

	t4.SymmetricDifference(&t5)
	fmt.Println(t4.String())
	if t4.String() != expected {
		t.Fail()
	}
}
