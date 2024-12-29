package stun

import (
	"fmt"
	"log"

	"github.com/pion/stun"
)

type Client struct {
	serverAddress string
}

func NewClient(serverAddress string) *Client {
	return &Client{
		serverAddress: serverAddress,
	}
}

func (c *Client) DiscoverPublicAddress() (string, error) {
	// Create a UDP connection to the STUN server
	conn, err := stun.Dial("udp", c.serverAddress)
	if err != nil {
		return "", fmt.Errorf("failed to connect to STUN server: %w", err)
	}

	// Build and send the STUN Binding Request
	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)
	var publicAddr string
	err = conn.Do(message, func(res stun.Event) {
		if res.Error != nil {
			log.Printf("STUN response error: %v", res.Error)
			return
		}

		var xorAddr stun.XORMappedAddress
		if err := xorAddr.GetFrom(res.Message); err != nil {
			log.Printf("Failed to parse XOR-Mapped-Address: %v", err)
			return
		}
		publicAddr = fmt.Sprintf("%s:%d", xorAddr.IP, xorAddr.Port)
	})
	if err != nil {
		return "", fmt.Errorf("failed to perform STUN transaction: %w", err)
	}

	return publicAddr, nil
}
