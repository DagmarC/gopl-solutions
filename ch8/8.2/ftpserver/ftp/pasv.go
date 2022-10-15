package ftp

// This command tells the server to enter a passive FTP
// session rather than Active. This allows users behind
// routers/firewalls to connect over FTP when they might
// not be able to connect over an Active ( PORT ) FTP session.
// PASV mode has the server tell the client where to connect
// the data port on the server.

func (c *Conn) pasv(args []string) {
	c.respond(UKNOWN)
}
