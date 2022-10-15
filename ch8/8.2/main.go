package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"path/filepath"

	"github.com/DagmarC/gopl-solutions/ch8/8.2/ftpserver/ftp"
)

// Source: https://medium.com/better-programming/how-to-write-a-concurrent-ftp-server-in-go-part-1-3904f2e3a9e5
// FTP Commands: https://www.serv-u.com/resources/tutorial/cwd-cdup-pwd-rmd-dele-smnt-site-ftp-command 

var portPtr = flag.Int("port", 8080, "port number")
var rootPtr = flag.String("roodDir", "ftpserver/public", "root directory")

func main() {
	flag.Parse()
	// localhost:port, which can be shortened to :port.
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *portPtr))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(listener.Addr())
	for {
		fmt.Println("D")
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

// net.Conn satisfies the io.ReadWriteCloser interface,
// we can read the request, write a response and, naturally, close the connection.
func handleConn(conn net.Conn) {
	defer conn.Close()
	absPath, err := filepath.Abs(*rootPtr)
	if err != nil {
		log.Fatal(err)
	}
	// ftp.Conn: many possible FTP actions achievable from a single touchpoint.
	ftp.Serve(ftp.NewConn(conn, absPath))
}
