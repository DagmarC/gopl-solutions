// Package ftp provides structs and functions for running an FTP server.
package ftp

import (
	"bufio"
	"log"
	"strings"
)

// ftp.Serve is our application’s router
// Serve scans incoming requests for valid commands and routes them to handler functions.
func Serve(c *Conn) {
	// connection has been established successfully, server is ready to accept a user
	c.respond(OK)

	// listen for incoming commands
	s := bufio.NewScanner(c.conn)
	for s.Scan() {
		input := strings.Fields(s.Text())
		if len(input) == 0 {
			continue
		}
		// Note: cmd names mandated by FTP don’t always match
		// the commands you enter in the client to trigger them
		cmd, args := input[0], input[1:]
		// 2020/05/28 08:13:58 << PORT [127,0,0,1,245,1]
		log.Printf("<< %s %v", cmd, args)

		// The client sends certain commands without the user’s direct intervention.
		// For example, a PORT command is secretly sent before every get/RETR request.
		switch cmd {
		case "CWD": // cd
			c.cwd(args)
		case "LIST": // ls
			c.list(args)
		case "PORT":
			c.port(args)
		case "USER":
			c.user(args)
		case "QUIT": // close
			c.respond(CONTROLCLOSE)
			return
		case "RETR": // get
			c.retr(args)
		case "TYPE":
			c.setDataType(args)
		case "PASV": // close
			c.pasv(args)
			return
		case "PWD": // get
			c.pwd(args)
		case "SYST":
			c.syst(args)
		default:
			c.respond(UKNOWN)
		}
	}
	if s.Err() != nil {
		log.Print(s.Err())
	}
}
