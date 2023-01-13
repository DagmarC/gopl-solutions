package main

import (
	"fmt"
	"runtime"
	"time"
)

const maxStage = 10000000

func main() {
	ts := time.Now()
	start := producer()

	for i := 0; i < maxStage; i++ {
		start = transitter(start)
	}
	reciever(start)
	fmt.Println("N of GOROUTINES in MAIN that TOOK:", runtime.NumGoroutine(), time.Since(ts))
}

func producer() <-chan int {
	ch := make(chan int)
	go func() {
		ch <- 1
		close(ch)
	}()
	return ch
}

func reciever(in <-chan int) {
	fmt.Println("N of GOROUTINES in RECIEVER:", runtime.NumGoroutine())
	fmt.Println("N recieved: ", <-in)
}

func transitter(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		out <- (<-in + 1) // send from in to out
		close(out)
	}()
	return out
}
