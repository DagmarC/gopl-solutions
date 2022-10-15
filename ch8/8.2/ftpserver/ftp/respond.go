package ftp

import (
	"fmt"
	"log"
)

const (
	READY         = "150 File status okay; about to open data connection."
	OK            = "200 Command okay."
	OS            = "215 NAME %s system type"
	NEWUSER       = "220 Service ready for new user."
	CONTROLCLOSE  = "221 Service closing control connection."
	SUCCESS       = "226 Closing data connection. Requested file action successful."
	LOGGED        = "230 User %s logged in, proceed."
	CREATED       = "257 %s created."
	CONNERROR     = "425 Can't open data connection."
	ABORTED       = "426 Connection closed; transfer aborted."
	SYNTAXERROR   = "501 Syntax error in parameters or arguments."
	UKNOWN        = "502 Command not implemented."
	UNIMPLEMENTED = "504 Cammand not implemented for that parameter."
	UNAVAILABLE   = "550 Requested action not taken. File unavailable."
)

// respond copies a string to the client and terminates it with the appropriate FTP line terminator
// for the datatype.
func (c *Conn) respond(s string) {
	log.Print(">> ", s)
	// sending the data to the client
	_, err := fmt.Fprint(c.conn, s, c.EOL())
	if err != nil {
		log.Print(err)
	}
}

// EOL returns the line terminator matching the FTP standard for the datatype.
func (c *Conn) EOL() string {
	switch c.dataType {
	case ascii:
		return "\r\n"
	case binary:
		return "\n"
	default:
		return "\n"
	}
}
