package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	timer := time.NewTimer(time.Second) // max comunications in 1 second

	ping, pong := make(chan int), make(chan int)
	done := make(chan struct{}) // signal for closing pong

	wg.Add(1)
	go func() { // pong goroutine
		defer wg.Done()
	outer:
		for {
			select {
			case pin := <-ping:
				fmt.Println("PING REC: ", pin)
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
			fmt.Println("PONG REC: ", pon)
			ping <- pon + 1 // answer
		case <-timer.C:
			<-pong      // release pong recievement
			close(done) // send signal to goroutine
			close(ping)
			fmt.Println("Closing ping... ")
			break outer
		}
	}
	wg.Wait() // Wait for pong gor to finish
}
