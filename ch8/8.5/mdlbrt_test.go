package main

import (
	"testing"
)

func Benchmark1_test(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buildMandelbrot()
	}
}
