package ftp

import "log"

func (c *Conn) port(args []string) {
	if len(args) != 1 {
		c.respond(SYNTAXERROR)
		return
	}
	// dataPortFromHostPort parses the six-byte IP address format into 
	// a struct of its parts that we store on the ftp.Conn instance
	dataPort, err := dataPortFromHostPort(args[0])
	if err != nil {
		log.Print(err)
		c.respond(SYNTAXERROR)
		return
	}
	c.dataPort = dataPort
	c.respond(OK)
}