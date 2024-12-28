package client

import "net"

type NetworkType string

const (
	TCP NetworkType = "tcp"
	UDP NetworkType = "udp"
)

type Dialer interface {
	Dial(network NetworkType, address string) (net.Conn, error)
}

type NetDialer struct{}

func (d *NetDialer) Dial(network NetworkType, address string) (net.Conn, error) {
	return net.Dial(string(network), address)
}
