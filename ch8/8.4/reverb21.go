package main

// Use async.WaitGroupper connection to count the number of active echo goroutines.
// When it falls to zero, close the write half of the TCP connection.

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done() // Signal gor is done, even if error occurs
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn) {
	// Close the write half of the TCP connection.
	var wg sync.WaitGroup

	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1) // Increment gor counter, right before gor starts
		go echo(c, input.Text(), 2*time.Second, &wg)

	}
	// NOTE: ignoring potential errors from input.Err()
	wg.Wait() // Wait until all echo gors will end.
	defer c.(*net.TCPConn).CloseWrite()
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
