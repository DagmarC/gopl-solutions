// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

//!+client
type client struct {
	addr string
	name string
	ch   chan<- string // an outgoing message channel
}

func (c client) String() string {
	return c.name
}

//!-client

// connClients will outputs string of all connected clients from map
func connClients(clients map[client]bool) string {
	var buff bytes.Buffer
	if len(clients) == 0 {
		buff.WriteString("INFO: You are first connected.")
		return buff.String()

	}
	var i int // to prevent last comma ","
	buff.WriteString("INFO: Connected clients: ")
	for c, ok := range clients {
		if ok {
			buff.WriteString(c.String()) // Note: client implements String interface
			if i != len(clients)-1 {
				buff.WriteString(", ")
			}
			i++
		}
	}
	return buff.String()
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

// 8.12:  Make the broadcaster announce the current set of clients to each new arrival.
func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.ch <- msg
			}

		case cli := <-entering:
			cli.ch <- connClients(clients) // Send to client all connected clients.
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

//!-broadcaster

const MAX_CLIENT_RESPONSE int = 100

//!+handleConn
func handleConn(conn net.Conn) {
	timer := time.NewTimer(time.Duration(MAX_CLIENT_RESPONSE) * time.Second) // See https://gobyexample.com/timers
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	addr := conn.RemoteAddr().String()
	// 8.14 Set client name
	ch <- "Enter you name: " // Ask client its name
	buffer := make([]byte, 1024)
	nbytes, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	c := client{name: strings.TrimSpace(string(buffer[:nbytes])), addr: addr, ch: ch}

	ch <- "INFO: Your name was set to " + c.name // Inform client
	fmt.Printf("INFO: Client %s@%s connected and name was succesfully set.\n", c.name, c.addr)

	entering <- c                        // Attach to broadcaster
	messages <- c.name + " has arrived." // Boradcast the arrival to clients.

	// 8.13 Disconnect inactive client
	go func(timer *time.Timer, c *client) {
		<-timer.C
		fmt.Printf("INFO: Client %s not responding for %v seconds.\n", c.name, MAX_CLIENT_RESPONSE)
		timer.Stop()
		defer conn.Close()
	}(timer, &c)

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- c.name + ": " + input.Text()
		timer.Reset(time.Duration(MAX_CLIENT_RESPONSE) * time.Second) // Reset timer to again wait at most MAX_CLIENT_RESPONSE seconds
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- c
	messages <- c.name + " has left." // Boradcast the leaving to clients.
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}
//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
