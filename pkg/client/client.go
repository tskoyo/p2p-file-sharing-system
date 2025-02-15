package client

import (
	"context"
	"fmt"
	"log"
	connectionpool "p2p-file-sharing-system/pkg/connection_pool"
	"time"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
)

const maxRetries = 3

type Client struct {
	Host           host.Host
	ConnectionPool *connectionpool.ConnectionPool
}

func NewClient(host host.Host, connectionPool *connectionpool.ConnectionPool) *Client {
	return &Client{
		Host:           host,
		ConnectionPool: connectionPool,
	}
}

// func (c *Client) Connect(peerAddress string, port string) error {
// 	log.Printf("trying to connect to %s on port %s", peerAddress, port)
// 	for i := 0; i < maxRetries; i++ {
// 		conn, err = net.Dial("tcp", peerAddress+":"+port)
// 		if err == nil {
// 			break
// 		}

// 		time.Sleep(1 * time.Second * time.Duration(i))
// 		log.Printf("Retrying connection to server: %v", err)
// 	}

// 	if err != nil {
// 		return fmt.Errorf("failed to connect to server XD: %w", err)
// 	}

// 	c.Conn = conn

// 	err = c.handshakeWithServer()
// 	if err != nil {
// 		return fmt.Errorf("failed to handshake with server: %w", err)
// 	}

// 	return nil
// }

func (c *Client) Connect(multiAddr string) error {
	peerMultiAddr, err := multiaddr.NewMultiaddr(multiAddr)
	if err != nil {
		return fmt.Errorf("Invalid multiaddress: %w", err)
	}

	peerInfo, err := peer.AddrInfoFromP2pAddr(peerMultiAddr)
	if err != nil {
		return fmt.Errorf("Failed to parse peer address: %w", err)
	}

	log.Printf("%s attempting to connet to %s", c.Host.ID(), multiAddr)

	var lastErr error
	// Connect with retries
	for i := 0; i < maxRetries; i++ {
		err := c.Host.Connect(context.Background(), *peerInfo)
		if err != nil {
			lastErr = err
			log.Printf("Connection attempt %d failed: %v", i+1, err)
			time.Sleep(time.Duration(i+1) * time.Second)
			continue
		}

		log.Printf("Successfully connected after %d attempts", i+1)
		return nil
	}

	return fmt.Errorf("Failed to connect to peer after %d attemtps: %w", maxRetries, lastErr)
}

// func (c *Client) UploadFile(filePath string) error {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return fmt.Errorf("Failed to open file: %w", err)
// 	}
// 	defer file.Close()

// 	_, err = io.Copy(c.Conn, file)
// 	if err != nil {
// 		return fmt.Errorf("failed to upload file: %w", err)
// 	}

// 	fmt.Printf("File '%s' uploaded successfully\n", filePath)
// 	return nil
// }

// func (c *Client) DownloadFile(filePath string) error {
// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		return fmt.Errorf("failed to create file: %w", err)
// 	}
// 	defer file.Close()

// 	_, err = io.Copy(file, c.Conn)
// 	if err != nil {
// 		return fmt.Errorf("failed to download file: %w", err)
// 	}

// 	fmt.Printf("File '%s' downloaded successfully\n", filePath)
// 	return nil
// }

// func (c *Client) handshakeWithServer() error {
// 	handshakeMsg := []byte("HELLO")
// 	_, err := c.Conn.Write(handshakeMsg)
// 	log.Println("Are we here?")
// 	if err != nil {
// 		return err
// 	}

// 	serverRespMsg := make([]byte, 2)
// 	_, err = c.Conn.Read(serverRespMsg)
// 	if err != nil {
// 		return err
// 	}

// 	if string(serverRespMsg) != "OK" {
// 		return fmt.Errorf("unexpected server response: %s", string(serverRespMsg))
// 	}

// 	return nil
// }
