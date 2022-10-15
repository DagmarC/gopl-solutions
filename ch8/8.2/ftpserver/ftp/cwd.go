package ftp

import (
	"log"
	"os"
	"path/filepath"
)

// If you’d like to challenge yourself by improving this naïve implementation, 
// consider how you would prevent the user from accessing files 
// above the server’s root directory. Currently, there’s nothing 
// to stop them entering cd ../../../../../../.. and getting access to 
// all sorts of things they shouldn’t.

// Explore how the standard library’s filepath.Clean can help solve this problem.


func (c *Conn) cwd(args []string) {
	if len(args) != 1 {
		c.respond(SYNTAXERROR)
		return
	}

	workDir := filepath.Join(c.workDir, args[0])
	absPath := filepath.Join(c.rootDir, workDir)
	_, err := os.Stat(absPath) // validate that the dir exists and is accessible
	if err != nil {
		log.Print(err)
		c.respond(UNAVAILABLE)
		return
	}
	c.workDir = workDir // update
	c.respond(OK)
}
