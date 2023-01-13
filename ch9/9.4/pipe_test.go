package main

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkPipeline(b *testing.B) {
	tstart := time.Now()
	start := producer()
	for i := 0; i < b.N; i++ {
		start = transitter(start)
	}
	reciever(start)
	fmt.Println("TIME TOOK :", time.Since(tstart))
	fmt.Println()
}
