package main

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tc := []struct {
		s    string
		sep  string
		want int
	}{
		{"a:b:c", ":", 3},
		{"a:b:c", " ", 1},
		{"a,", ",", 2},
		{"a a", " ", 2},
	}
	for _, test := range tc {
		words := strings.Split(test.s, test.sep)
		if got := len(words); got != test.want {
			t.Errorf("Split(%q %q) returned %d words, want %d", test.s, test.sep, got, test.want)
		}
	}
}
