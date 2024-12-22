package peer

import (
	"log"
	"net"
)

type Client struct {
	Conn net.Conn
}

func NewClient(address string) *Client {
	return &Client{}
}

func (c *Client) Connect() error {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		return err
	}
	c.Conn = conn
	log.Println("[CLIENT]: Connected to server!")
	return nil
}
