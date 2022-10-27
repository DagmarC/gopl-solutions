// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 227.

// Netcat is a simple read/write client for TCP servers.
package main

import (
	"io"
	"log"
	"net"
	"os"
)

//!+
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})

	go func() {
		_, err := io.Copy(os.Stdout, conn) // read connection (messages from server)
		if err != nil {
			log.Println("ERROR INFO:", err)
		}
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()

	go mustCopy(conn, os.Stdin) // write to connection (client side - send to the server side)

	// 8.23
	// conn.Close() // Closes both READ/WRITE sides of the connection.
	// NOTE: Closing the read side of the connectionn causes the background gorroutine's call to io.Copy (l.24)
	// to return a "read from closed connection" error
	// ERROR INFO: read tcp 127.0.0.1:53810->127.0.0.1:8000: use of closed network connection

	// Closes the write side of the connection causes the server to see EOF
	defer conn.(*net.TCPConn).CloseWrite() // net.TCPConn supports CloseWrite, so server could still be sending to the conn and wont close it with an error (l.24)
	<-done                                 // wait for background goroutine to finish

}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src) // read connection (messages from server)
	if err != nil {
		log.Fatal(err)
	}

}
