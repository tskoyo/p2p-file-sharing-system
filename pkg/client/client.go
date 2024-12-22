package client

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
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
