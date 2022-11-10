package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	meals    int           = 3
	eating   time.Duration = 2
	thinking time.Duration = 5
)

func dining(ph string, lHand *sync.Mutex, rHand *sync.Mutex, wg *sync.WaitGroup) {

	for m := 0; m < meals; m++ {

		fmt.Printf("%*s is about to eat %d meal\n", m*8, ph, m+1)

		lHand.Lock()
		rHand.Lock()
		fmt.Printf("%*s is eating %d meal\n", m*8, ph, m+1)
		time.Sleep(eating * time.Second)
		lHand.Unlock()
		rHand.Unlock()

		fmt.Printf("%*s is thinking...\n", m*8, ph)
		time.Sleep(thinking * time.Second)
	}
	fmt.Printf("%s finished eating.\n", ph)
	wg.Done()
	fmt.Printf("%s left table.\n", ph)

}

func main() {
	var wg sync.WaitGroup

	lHand := &sync.Mutex{} // each right and left hand will be locked while philosopher is eating
	ph := []string{"Aron", "Bill", "Chris", "Dagmar", "Elen"}

	for _, name := range ph {
		rHand := &sync.Mutex{}
		wg.Add(1)
		go dining(name, lHand, rHand, &wg)
		lHand = rHand // switch hands to continue in a circle
	}
	wg.Wait()
}
