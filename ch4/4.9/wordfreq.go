package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	words := make(map[string]int) // counts words

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words[scanner.Text()]++
	}
	if scanner.Err() != nil {
		log.Fatalf("Error %s", scanner.Err().Error())
	}

	for word, count := range words {
		fmt.Printf("%s\t%d\n", word, count)
	}
}
