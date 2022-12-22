package client

import (
	"fmt"
	"net"
	"strings"
	"time"
)

const MAX_CLIENT_RESPONSE int = 100

//!+Client
type Client struct {
	Addr  string
	Name  string
	Ch    chan string // an outgoing message channel
	conn  *net.Conn
	Timer *time.Timer
}

func NewClient(conn *net.Conn) (*Client, error) {
	ch := make(chan string) // outgoing client messages
	addr := (*conn).RemoteAddr().String()
	timer := time.NewTimer(time.Duration(MAX_CLIENT_RESPONSE) * time.Second) // See https://gobyexample.com/timers

	c := &Client{conn: conn, Addr: addr, Ch: ch, Timer: timer}
	
	go c.Writer() // Accepts messages via ch channel and write it to the connection

	err := c.setName()
	if err != nil {
		return &Client{}, err
	}
	return c, nil
}

// setName 8.14 exercise: Ask client his name and use it instead of remote addr.
func (c *Client) setName() error {
	// 8.14 Set client name
	c.Ch <- "Enter you name: " // Ask client its name
	buffer := make([]byte, 1024)
	nbytes, err := (*c.conn).Read(buffer)
	if err != nil {
		return err
	}
	c.Name = strings.TrimSpace(string(buffer[:nbytes]))
	return nil
}

func (c *Client) String() string {
	return c.Name
}

func (c *Client) Writer() {
	for msg := range c.Ch {
		fmt.Fprintln(*c.conn, msg) // NOTE: ignoring network errors
	}
}

func (c *Client) IdleClient() {
	<-c.Timer.C
	fmt.Printf("INFO: Client %s not responding for %v seconds.\n", c.Name, MAX_CLIENT_RESPONSE)
	c.Timer.Stop()
	defer (*c.conn).Close()
}
