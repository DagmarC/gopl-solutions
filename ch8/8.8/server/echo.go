package main

// Use async.WaitGroupper connection to count the number of active echo goroutines.
// When it falls to zero, close the write half of the TCP connection.

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

const MAX_CLIENT_RESPONSE int = 10
const MAX_BUFFER_MSG int = 10

var portPtr = flag.String("port", "8000", "port number on clock wall")

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

	// Timers represent a single event in the future. You tell the timer how long you want to wait,
	// and it provides a channel that will be notified at that time. This timer will wait N seconds.
	timer := time.NewTimer(time.Duration(MAX_CLIENT_RESPONSE) * time.Second) // See https://gobyexample.com/timers

	msgCh := make(chan string, MAX_BUFFER_MSG) // Buffer to accept more than 1 message at a time to not to bloc communication

	// Receiving goroutine (message)
	go func() {
		for {
			msg, ok := <-msgCh
			if ok {
				wg.Add(1)
				go echo(c, msg, 6*time.Second, &wg)
				timer.Reset(time.Duration(MAX_CLIENT_RESPONSE) * time.Second) // Reset timer.
			}
		}

	}()

	// Channel closing principle: Dont close the channel from the receiver side. 1 sender can can close data channel anytime, receiver cannot.
	// Sending 'main' goroutine.

	input := bufio.NewScanner(c)
outer:
	for input.Scan() {
		select {
		case <-timer.C:
			fmt.Printf("INFO: Client not responding for %v seconds. Waiting for GORS to finish...\n", MAX_CLIENT_RESPONSE)
			wg.Wait()

			fmt.Printf("INFO: Closing channels.\n")
			timer.Stop() // Stop timer
			close(msgCh)
			break outer
		case msgCh <- input.Text():
			fmt.Printf("INFO: MSG SENT...\n") // Sending text
		default:
		}
	}

	if err := input.Err(); err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("INFO: Closing conection.\n")
	defer c.(*net.TCPConn).CloseWrite()
}

//!-

func main() {
	flag.Parse()
	address := "localhost:" + *portPtr
	l, err := net.Listen("tcp", address)
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
