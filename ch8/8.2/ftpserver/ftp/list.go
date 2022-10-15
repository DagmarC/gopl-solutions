package ftp

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func (c *Conn) list(args []string) {
	var target string
	if len(args) > 0 {
		target = filepath.Join(c.rootDir, c.workDir, args[0])
	} else {
		target = filepath.Join(c.rootDir, c.workDir) // current dir
	}

	files, err := ioutil.ReadDir(target)
	if err != nil {
		log.Print(err)
		c.respond(UNAVAILABLE)
		return
	}
	c.respond(READY) // 150 File status okay; about to open data connection

	// When sending anything other than statuses,
	// the server must establish a second, temporary
	// connection to the client, known as the data connection
	// ...to a specific port that the FTP client has selected in advance.
	// ...client sends another command behind the scenes: PORT.
	//
	// Note that dataConnect returns a struct satisfying the
	// net.Conn interface (net.TCPConn), not our custom ftp.Conn.
	dataConn, err := c.dataConnect()
	if err != nil {
		log.Print(err)
		c.respond(CONNERROR)
		return
	}
	defer dataConn.Close()

	for _, file := range files {
		_, err := fmt.Fprint(dataConn, file.Name(), c.EOL())
		if err != nil {
			log.Print(err)
			c.respond(ABORTED)
		}
	}
	_, err = fmt.Fprintf(dataConn, c.EOL())
	if err != nil {
		log.Print(err)
		c.respond(ABORTED)
	}

	c.respond(SUCCESS)
}
