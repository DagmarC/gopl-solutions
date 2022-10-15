package ftp

import (
	"fmt"
	"runtime"
)

// syst A client can issue this command to the server to determine the operating syst running
// on the server. Not all server responses are accurate in this regard, however,
// as some servers respond with the syst they emulate or may not respond at all due to
// potential security risks.
func (c *Conn) syst(args []string) {
	c.respond(fmt.Sprintf(OS, runtime.GOOS))
}
