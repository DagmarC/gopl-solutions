// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	want := "{1 9 144}"
	if want != x.String() {
		t.Errorf("got %s, want %s", x.String(), want)
	}
}

func TestHas(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)

	tc := []struct {
		tcase int
		want  bool
	}{
		{144, true},
		{11, false},
		{-1, false},
	}

	for _, tt := range tc {
		if ok := x.Has(tt.tcase); ok != tt.want {
			t.Errorf("x.Has(%v)=%t, want %t", tt.tcase, ok, tt.want)
		}
	}
}

func TestUnionWith(t *testing.T) {
	var x, y, z IntSet
	x.Add(1)
	x.Add(2)
	x.Add(222)

	y.Add(222)
	y.Add(1)
	y.Add(13)

	z.Add(0)

	tc := []struct {
		tcase *IntSet
		want  string
	}{
		{&y, "{1 2 13 222}"},   // xUy
		{&z, "{0 1 2 13 222}"}, // xUz Note: second x is union of xUy
	}

	for _, tt := range tc {
		if x.UnionWith(tt.tcase); x.String() != tt.want {
			t.Errorf("x.UnionWith(%v)=%s, want %s", tt.tcase.String(), x.String(), tt.want)
		}
	}
}

func benchmark(b *testing.B, num int) {
	var x IntSet
	for i := 0; i < 1000; i++ {
		x.Add(num)
	}
}

func BenchmarkAdd100(b *testing.B) {
	benchmark(b, 100)
}

func BenchmarkAdd10000(b *testing.B) {
	benchmark(b, 10000)
}

func BenchmarkAdd1000000(b *testing.B) {
	benchmark(b, 1000000)
}

func BenchmarkAdd100000000(b *testing.B) {
	benchmark(b, 100000000)
}