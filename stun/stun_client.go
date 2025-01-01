package stun

import (
	"fmt"

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
	conn, err := stun.Dial("udp", "stun.l.google.com:19302")
	if err != nil {
		panic(err)
	}

	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)
	var publicAddr string
	if err := conn.Do(message, func(res stun.Event) {
		if res.Error != nil {
			panic(res.Error)
		}

		var xorAddr stun.XORMappedAddress
		if err := xorAddr.GetFrom(res.Message); err != nil {
			panic(err)
		}
		publicAddr = fmt.Sprintf("%s:%d", xorAddr.IP, xorAddr.Port)
	}); err != nil {
		panic(err)
	}

	return publicAddr, nil
}
