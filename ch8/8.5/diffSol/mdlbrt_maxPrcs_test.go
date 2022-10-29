package main

import (
	"fmt"
	"runtime"
	"testing"
)

const maxGortns = 100

func benchmarkMaxProcs(b *testing.B, maxProcs int) {
	fmt.Printf("%d procs\n", maxProcs)
	runtime.GOMAXPROCS(maxProcs)
	for i := 0; i < b.N; i++ {
		buildImg(maxGortns)
	}
}

func Benchmark1_test(b *testing.B) { benchmarkMaxProcs(b, 1) }
func Benchmark2_test(b *testing.B) { benchmarkMaxProcs(b, 2) }
func Benchmark3_test(b *testing.B) { benchmarkMaxProcs(b, 3) }
func Benchmark4_test(b *testing.B) { benchmarkMaxProcs(b, 4) }
func Benchmark5_test(b *testing.B) { benchmarkMaxProcs(b, 5) }
func Benchmark6_test(b *testing.B) { benchmarkMaxProcs(b, 6) }
func Benchmark7_test(b *testing.B) { benchmarkMaxProcs(b, 7) }
func Benchmark8_test(b *testing.B) { benchmarkMaxProcs(b, 8) }
