package main

import (
	"strings"
	"testing"
)

func TestCharcountValid(t *testing.T) {
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)
	m3 := make(map[rune]int)

	m1['@'] = 1
	m1['1'] = 2
	m2['#'] = 1
	m2['/'] = 2
	m3['a'] = 3

	var tc = []struct {
		s       string
		want    map[rune]int
		invalid int
	}{
		{"@11", m1, 0},
		{"#//", m2, 0},
		{"aaa", m3, 0},
	}
	for _, tt := range tc {
		got, invalid, err := Charcount(strings.NewReader(tt.s))
		if err != nil || tt.invalid != invalid || len(got) != len(tt.want) {
			t.Errorf("Charcount(%s)=%v, invalid=%d, want %v, invalid=%d", tt.s, got, invalid, tt.want, tt.invalid)
		}
		for k, v := range got {
			if v != tt.want[k] {
				t.Errorf("Charcount(%s)=%v, invalid=%d, want %v, invalid=%d", tt.s, got, invalid, tt.want, tt.invalid)
			}
		}
	}
}

func TestCharcountInvalid(t *testing.T) {
	m1 := make(map[rune]int)
	m1[127] = 1

	var tc = []struct {
		s       string
		want    map[rune]int
		invalid int
	}{
		{"\xC2\x7F\x80\x80\xC2\xC0\x80\x80", m1, 7},
	}
	for _, tt := range tc {
		got, invalid, err := Charcount(strings.NewReader(tt.s))
		if err != nil || tt.invalid != invalid || len(got) != len(tt.want) {
			t.Errorf("Charcount(%s)=%v, invalid=%d, want %v, want invalid=%d", tt.s, got, invalid, tt.want, tt.invalid)

		}
		for k, v := range got {
			if v != tt.want[k] {
				t.Errorf("Charcount(%s)=%v, invalid=%d, want %v, want invalid=%d", tt.s, got, invalid, tt.want, tt.invalid)
			}
		}
	}
}
