// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package main

import (
	"bytes"
	"fmt"
)

//!+6.5
const UintSize = 32 << (^uint(0) >> 63) // uint type, which is the most efficient unsigned integer type for the platform
//!-6.5

//!+intset
// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/UintSize, uint(x%UintSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/UintSize, uint(x%UintSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-intset

//!+6.1
// Len returns the number of elements
func (s *IntSet) Len() int {
	var count int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < UintSize; j++ {
			if word&(1<<uint(j)) != 0 {
				count++
			}
		}
	}
	return count
}

// Remove removes x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/UintSize, uint(x%UintSize)

	for word >= len(s.words) {
		return // Not present
	}
	s.words[word] &^= 1 << bit // NOT AND is a Bit clear
}

// Clear removes all elements from the set
func (s *IntSet) Clear() {
	var c uint64 = uint64(0)
	for i := range s.words {
		s.words[i] &= c
	}
}

// Copy returns a copy of the set
func (s *IntSet) Copy() *IntSet {
	var cp IntSet
	cp.words = make([]uint64, len(s.words))
	copy(cp.words, s.words)
	return &cp
}

//!-6.1

//!+6.2
func (s *IntSet) AddAll(x ...int) {
	for _, xx := range x {
		s.Add(xx)
	}

}

//!-6.2

//!+6.3
// IntersectsWith returns elements that are in s and t at the same time
func (s *IntSet) IntersectsWith(t *IntSet) {
	for i := range s.words {
		if i < len(t.words) {
			s.words[i] &= t.words[i]
		} else {
			s.words[i] = 0 // t is no longer present
		}
	}
}

// DifferenceWith returns the elements that are in s but not in t
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := range s.words {
		if i >= len(t.words) {
			break // No need to continue in the difference
		}
		s.words[i] &^= t.words[i]
	}
}

// SymmetricDifference returns the elements that are present in s and t but not both
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] = (s.words[i] | tword) & ^(s.words[i] & tword)
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-6.3

//!+6.4
func (s *IntSet) Elems() []int {
	elems := make([]int, 0, s.Len())
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < UintSize; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, UintSize*i+j)
			}
		}
	}
	return elems
}

//!-6.4

//!+string
// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j <UintSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", UintSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)

	y.Add(9)
	y.Add(42)

	x.UnionWith(&y)

	x.SymmetricDifference(&y)
	fmt.Println(x.String())

}
