package ftp

import (
	"fmt"
	"net"
)

func (c *Conn) dataConnect() (net.Conn, error) {
	conn, err := net.Dial("tcp", c.dataPort.toAddress())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

type dataPort struct {
	h1, h2, h3, h4 int // host
	p1, p2         int // port
}

//dataPortFromHostPort parses the six-byte IP address format into a struct of its parts that we store on the ftp.Conn
func dataPortFromHostPort(hostPort string) (*dataPort, error) {
	var dp dataPort
	_, err := fmt.Sscanf(hostPort, "%d,%d,%d,%d,%d,%d",
		&dp.h1, &dp.h2, &dp.h3, &dp.h4, &dp.p1, &dp.p2)
	if err != nil {
		return nil, err
	}
	return &dp, nil
}

// toAddress converts that struct to a traditionally formatted IP-address-plus-port
// that the server can connect to with net.Dial
//
// If p1 = 00011011, p1<<8 = 0001101100000000. When you add p2, 
// it fills the eight bits left empty by the shift. 
// p2 = 11111111; p1 + p2 = 0001101111111111 = 7167.
func (d *dataPort) toAddress() string {
	if d == nil {
		return ""
	}
	// convert hex port bytes to decimal port
	port := d.p1<<8 + d.p2
	return fmt.Sprintf("%d.%d.%d.%d:%d", d.h1, d.h2, d.h3, d.h4, port)
}
