package ftp

// The first thing an FTP client does when it establishes a connection to a server 
// is provide information about the user.

// From there, it’s easy to imagine how the details would be checked 
// against a database of known users and their access permissions. 
// Once authenticated, the username and privileges could be stored 
// as additional fields on the ftp.Conn instance.

import (
	"fmt"
	"strings"
)

func (c *Conn) user(args []string) {
	c.respond(fmt.Sprintf(LOGGED, strings.Join(args, " ")))
}
