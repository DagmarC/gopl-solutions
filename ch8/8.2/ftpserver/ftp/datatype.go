package ftp

type dataType int

const (
	// every new ftp.Conn we create begins by default with dataType = 0 - ascii
	ascii dataType = iota
	binary
)

func (c *Conn) setDataType(args []string) {
	if len(args) == 0 {
		c.respond(SYNTAXERROR)
	}

	switch args[0] {
	case "A":
		c.dataType = ascii
	case "I": // image/binary
		c.dataType = binary
	default:
		c.respond(UNIMPLEMENTED)
		return
	}
	c.respond(OK)
}
