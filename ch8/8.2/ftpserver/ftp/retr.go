package ftp

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func (c *Conn) retr(args []string) {
	if len(args) != 1 {
		c.respond(SYNTAXERROR)
		return
	}

	path := filepath.Join(c.rootDir, c.workDir, args[0])
	file, err := os.Open(path)
	if err != nil {
		log.Print(err)
		c.respond(UNAVAILABLE)
	}
	c.respond(READY)
	// As with list, the client precedes every RETR request with a PORT command, 
	// so the ftp.Conn's dataPort field is already populated.
	dataConn, err := c.dataConnect()
	if err != nil {
		log.Print(err)
		c.respond(CONNERROR)
	}
	defer dataConn.Close()

	_, err = io.Copy(dataConn, file) // load the entire file into memory and copy it directly to the data connection
	if err != nil {
		log.Print(err)
		c.respond(ABORTED)
		return
	}
	io.WriteString(dataConn, c.EOL()) // terminate the transfer with the appropriate FTP line ending
	c.respond(SUCCESS)
}