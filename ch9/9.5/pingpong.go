package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(4) // Ex 9.6 - measure how it affecct the program and what is the optimal

	timer := time.NewTimer(time.Second) // max comunications in 1 second

	ping, pong := make(chan int), make(chan int)
	done := make(chan struct{}) // signal for closing pong

	wg.Add(1)
	go func() { // pong goroutine
		defer wg.Done()
	outer:
		for {
			select {
			case <-done:
			default:
			}

			select {
			case pin := <-ping:
				// fmt.Println("PING REC: ", pin)
				pong <- pin + 1
			case <-done:
				fmt.Println("Closing Pong... ")
				close(pong)
				break outer
			}
		}
	}()

	ping <- 0 // START PING
outer:
	for {
		select {
		case pon := <-pong:
			// fmt.Println("PONG REC: ", pon)
			ping <- pon + 1 // answer
		case <-timer.C:
			close(done) // send signal to goroutine
			time.Sleep(1000 * time.Millisecond) // Give a time to recieve done signal
			for p := range pong {
				fmt.Println("LAST PONGs releasment: ", p) // release pong recievement
			}
			fmt.Println("Closing Ping... ")
			close(ping)
			break outer
		}
	}
	wg.Wait() // Wait for pong gor to finish
}
