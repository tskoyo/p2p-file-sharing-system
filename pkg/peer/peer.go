package peer

import (
	"fmt"
	"log"
	"net"
)

type Client struct {
	Conn net.Conn
}

func StartServer(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()

	log.Printf("Server listening on %s", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Connection error: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Fprintln(conn, "Hello, peer!")
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
	fmt.Println("Connected to server!")
	return nil
}
