package client

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	connectionpool "p2p-file-sharing-system/pkg/connection_pool"
	"time"
)

const maxRetries = 3

var conn net.Conn
var err error

type Client struct {
	Conn           net.Conn
	ConnectionPool *connectionpool.ConnectionPool
}

func NewClient(cp *connectionpool.ConnectionPool) *Client {
	return &Client{
		ConnectionPool: cp,
	}
}

func (c *Client) Connect(peerAddress string) error {
	for i := 0; i < maxRetries; i++ {
		conn, err = net.Dial("tcp", peerAddress) // TODO: Add support for UDP
		if err == nil {
			break
		}

		time.Sleep(1 * time.Second * time.Duration(i))
		log.Printf("Retrying connection to server: %v", err)
	}

	c.Conn = conn

	err = c.handshakeWithServer()
	if err != nil {
		return fmt.Errorf("failed to handshake with server: %w", err)
	}

	return nil
}

func (c *Client) UploadFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Failed to open file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(c.Conn, file)
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}

	fmt.Printf("File '%s' uploaded successfully\n", filePath)
	return nil
}

func (c *Client) DownloadFile(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, c.Conn)
	if err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}

	fmt.Printf("File '%s' downloaded successfully\n", filePath)
	return nil
}

func (c *Client) handshakeWithServer() error {
	handshakeMsg := []byte("HELLO")
	_, err := c.Conn.Write(handshakeMsg)
	if err != nil {
		return err
	}

	serverRespMsg := make([]byte, 2)
	_, err = c.Conn.Read(serverRespMsg)
	if err != nil {
		return err
	}

	if string(serverRespMsg) != "OK" {
		return fmt.Errorf("unexpected server response: %s", string(serverRespMsg))
	}

	return nil
}
