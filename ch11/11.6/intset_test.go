// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"math/rand"
	"testing"
)

const randmax = 1 << 24
const count = 100

var (
	randInts1 = make([]int, 0)
	randInts2 = make([]int, 0)
)

func init() {
	for i := 0; i < count; i++ {
		randInts1 = append(randInts1, rand.Intn(randmax))
		randInts2 = append(randInts2, rand.Intn(randmax))
	}
}

func benchmarkAdd(b *testing.B, s IntSet) {
	for i := 0; i < 1000; i++ {
		s.Add(rand.Intn(randmax))
	}
}

func benchmarkAddAll(b *testing.B, s IntSet) {
	for i := 0; i < b.N; i++ {
		s.AddAll(randInts1...)
	}
}

func benchmarkHas(b *testing.B, set IntSet) {
	set.AddAll(randInts1...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, x := range randInts1 {
			set.Has(x)
		}
	}
}

func benchmarkUnionWith(b *testing.B, newSet func() IntSet) {
	set1, set2 := newSet(), newSet()
	set1.AddAll(randInts1...)
	set2.AddAll(randInts2...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set1.UnionWith(set2)
	}
}

func benchmarkString(b *testing.B, set IntSet) {
	set.AddAll(randInts1...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.String()
	}
}

// IntSet64
func BenchmarkAddIntSet64(b *testing.B) {
	benchmarkAdd(b, &IntSet64{})
}
func BenchmarkAddAllIntSet64(b *testing.B) { benchmarkAddAll(b, &IntSet64{}) }
func BenchmarkHasIntSet64(b *testing.B)    { benchmarkHas(b, &IntSet64{}) }
func BenchmarkUnionWithIntSet64(b *testing.B) {
	benchmarkUnionWith(b, func() IntSet { return &IntSet64{} })
}
func BenchmarkStringIntSet64(b *testing.B) { benchmarkString(b, &IntSet64{}) }

// IntSetMap
func BenchmarkAddIntSetMap(b *testing.B)    { benchmarkAdd(b, &IntSetMap{map[int]bool{}}) }
func BenchmarkAddAllIntSetMap(b *testing.B) { benchmarkAddAll(b, &IntSetMap{map[int]bool{}}) }
func BenchmarkHasIntSetMap(b *testing.B)    { benchmarkHas(b, &IntSetMap{map[int]bool{}}) }
func BenchmarkUnionWithIntSetMap(b *testing.B) {
	benchmarkUnionWith(b, func() IntSet { return &IntSetMap{map[int]bool{}} })
}
func BenchmarkStringIntSetMap(b *testing.B) { benchmarkString(b, &IntSetMap{map[int]bool{}}) }
