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
	"time"

	"github.com/DagmarC/gopl-solutions/ch8/8.14/client"
)

//!-client

// connClients will outputs string of all connected clients from map
func connClients(clients map[client.Client]bool) string {
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
	entering = make(chan client.Client)
	leaving  = make(chan client.Client)
	messages = make(chan string) // all incoming client messages
)

// 8.12:  Make the broadcaster announce the current set of clients to each new arrival.
func broadcaster() {
	clients := make(map[client.Client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all clients' outgoing message channels.
			for cli := range clients {
				select {
				case cli.Ch <- msg:
				default: // 8.15 unblocking msg to not to block others to accept message
				}
			}
		case cli := <-entering:
			select {
			case cli.Ch <- connClients(clients): // Send to client all connected clients.
				clients[cli] = true
			default:
				fmt.Println("Client was not connected successfully: ", cli.Name) // 8.15 if client got stuck and cannot be connected
			}
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.Ch)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	c, err := client.NewClient(&conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("INFO: Client %s@%s connected and name was succesfully set.\n", c.Name, c.Addr)

	entering <- *c                       // Attach to broadcaster
	messages <- c.Name + " has arrived." // Boradcast the arrival to clients, use client name.

	// 8.13 Disconnect inactive client
	go c.IdleClient()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- c.Name + ": " + input.Text()
		c.Timer.Reset(time.Duration(time.Duration(client.MAX_CLIENT_RESPONSE) * time.Second)) // Reset timer to again wait at most MAX_CLIENT_RESPONSE seconds
	}
	// NOTE: ignoring potential errors from input.Err()
	leaving <- *c
	messages <- c.Name + " has left." // Boradcast the leaving to clients.
	conn.Close()
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
