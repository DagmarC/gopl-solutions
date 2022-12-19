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
)

//!+broadcaster
type client struct {
	name string
	ch   chan<- string // an outgoing message channel
}

func (c client) String() string {
	return c.name
}

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

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	c := client{name: who, ch: ch}

	ch <- "You are " + c.name
	messages <- c.name + " has arrived." // Boradcast the arrival to clients.
	entering <- c                        // Attach to broadcaster

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- c.name + ": " + input.Text()
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
