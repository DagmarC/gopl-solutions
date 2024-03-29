// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

//!+intset
type IntSet interface {
	Add(x int)
	AddAll(ints ...int)
	Has(x int) bool
	Ints() []int
	String() string
	UnionWith(t IntSet)
}

// An IntSet64 is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet64 struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet64) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet64) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll adds a variable number of non-negative integers to the set.
func (s *IntSet64) AddAll(ints ...int) {
	// Find the largest word and allocate sufficient memory to hold it and all other words in a
	// single operation. ~13% runtime reduction.
	max := 0
	for _, x := range ints {
		if x > max {
			max = x
		}
	}

	s.allocMaxLenRequired(max / 64)
	for _, x := range ints {
		s.Add(x)
	}
}

func (s *IntSet64) allocMaxLenRequired(word int) {
	if word >= len(s.words) {
		growth := make([]uint64, word-len(s.words)+1)
		s.words = append(s.words, growth...)
	}
}

// UnionWith sets s to the union of s and t.
func (s *IntSet64) UnionWith(t IntSet) {
	switch t := t.(type) {
	case *IntSet64:
		for i, tword := range t.words {
			if i < len(s.words) {
				s.words[i] |= tword
			} else {
				s.words = append(s.words, tword)
			}
		}
	default:
		s.AddAll(t.Ints()...)
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet64) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet64) Ints() []int {
	var ints []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				ints = append(ints, i*64+j)
			}
		}
	}
	return ints
}
