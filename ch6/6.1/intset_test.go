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
	set IntSet
	len int
}

// var testCases = []TestValues{
// 	{set: IntSet{words: []uint64{1}}, len: 1},        // representation of words: [...0000 0001,] means only one element {0} (64*i+j)[i=0, j=0], where i is the words index in []uint64 slice and j is the index inside uint64 word
// 	{set: IntSet{words: []uint64{16, 1, 2}}, len: 3}, // representation of words: [...0001 0000, ...0000 0001,...0000 0010] -> {4, 64 (64*i+j)[i=1, j=0], 129(64*i+j)[i=]}
// }

func InitTestCases() []TestValues {
	var x, y IntSet

	x.Add(0)

	y.Add(4)
	y.Add(64)
	y.Add(129)

	return []TestValues{
		{set: x, len: 1},
		{set: y, len: 3},
	}
}

func TestLen(t *testing.T) {
	testCases := InitTestCases()
	for _, tc := range testCases {
		fmt.Println(tc.set.Len(), tc.set.String())
		if tc.set.Len() != tc.len {
			t.Fail()
		}
	}
}

func TestRemove(t *testing.T) {
	testCases := InitTestCases()
	for _, tc := range testCases {
		tc.set.Remove(4)
		if tc.set.Has(4) {
			t.FailNow()
		}
	}
}

func TestClear(t *testing.T) {
	testCases := InitTestCases()

	for _, tc := range testCases {
		tc.set.Clear()
		for _, w := range tc.set.words {
			if w != 0 {
				t.Fail()
			}
		}
	}

}

func TestCopy(t *testing.T) {
	testCases := InitTestCases()

	for _, tc := range testCases {
		cp := tc.set.Copy()
		// fmt.Println("TEST COPY", tc.set.String(), cp) // Note that tc.has is not a pointer but String() method is defined on the pointer receiver *IntSet and that is why we need to either pass String() method explicitely that will make the conversion automatically or pass the pointer type as cp.
		if tc.set.String() != cp.String() {
			t.FailNow()
		}
	}
}

func TestAddAll(t *testing.T) {
	testCases := InitTestCases()

	expected := "{0 1 2 3}"
	testCases[0].set.AddAll(1, 2, 3)
	if testCases[0].set.String() != expected {
		t.Log("expected and actual", expected, testCases[0].set.String())
		t.Fail()
	}
}

func TestIntersectsWith(t *testing.T) {
	testCases := InitTestCases()

	expected := "{}"

	t1 := testCases[0].set
	t2 := testCases[1].set

	t1.IntersectsWith(&t2)
	if t1.String() != expected {
		t.Fail()
	}

	var t3 IntSet
	t3.AddAll(11, 129)
	expected = "{129}"

	t2.IntersectsWith(&t3)
	if t2.String() != expected {
		t.Log("expected and actual", expected, t2.String())
		t.Fail()
	}
}

func TestDifferenceWith(t *testing.T) {
	testCases := InitTestCases()

	expected := "{129}"

	t1 := testCases[1].set
	var t3 IntSet
	t3.AddAll(4, 64)

	t1.DifferenceWith(&t3)
	if t1.String() != expected {
		t.Fail()
	}
}

func TestSymmetricDifference(t *testing.T) {
	testCases := InitTestCases()

	expected := "{4 32 33 129 233 431}"

	t1 := testCases[1].set
	var t3 IntSet
	t3.AddAll(64, 32, 33, 233, 431)

	t1.SymmetricDifference(&t3)
	if t1.String() != expected {
		t.Fail()
	}

	expected = "{}"
	t4 := testCases[0].set

	var t5 IntSet
	t5.AddAll(0)

	t4.SymmetricDifference(&t5)
	if t4.String() != expected {
		t.Fail()
	}
}

func TestElems(t *testing.T) {
	testCases := InitTestCases()

	expected := []int{4, 64, 129}
	for i, el := range testCases[1].set.Elems() {
		if el != expected[i] {
			t.Fail()
		}
	}
}
