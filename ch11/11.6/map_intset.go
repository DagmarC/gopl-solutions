package intset

import (
	"bytes"
	"fmt"
)

type IntSetMap struct {
	m map[int]bool
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSetMap) Has(x int) bool {
	_, ok := s.m[x]
	return ok
}

// Add adds the non-negative value x to the set.
func (s *IntSetMap) Add(x int) {
	s.m[x] = true
}

// AddAll adds a variable number of non-negative integers to the set.
func (s *IntSetMap) AddAll(ints ...int) {
	for _, x := range ints {
		s.Add(x)
	}
}
// UnionWith sets s to the union of s and t.
func (s *IntSetMap) UnionWith(t IntSet) {
	switch t := t.(type) {
	case *IntSetMap:
		for k := range t.m {
			s.m[k] = true
		}
	default:
		for _, i := range t.Ints() {
			s.m[i] = true
		}
	}
}

//!-intset

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSetMap) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for k := range s.m {
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", k)
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSetMap) Ints() []int {
	ints := make([]int, 0, len(s.m))
	for k := range s.m {
		ints = append(ints, k)
	}
	return ints
}
