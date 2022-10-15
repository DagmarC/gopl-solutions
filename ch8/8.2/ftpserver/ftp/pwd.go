package ftp

import (
	"fmt"
	"path/filepath"
)

func (c *Conn) pwd(args []string) {
	if len(args) != 0 {
		c.respond(SYNTAXERROR)
		return
	}
	absPath := filepath.Join(c.rootDir, c.workDir)

	c.respond(fmt.Sprintf(CREATED, absPath))
}
