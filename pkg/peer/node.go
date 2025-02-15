package peer

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"

	libp2p "github.com/libp2p/go-libp2p"
	"github.com/multiformats/go-multiaddr"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
)

type Node struct {
	ID   string
	Host host.Host
	// DHT  *dht.IpfsDHT
	// Discovery *discovery.RoutingDiscovery
}

type NodeConfig struct {
	ID   string
	Port int
}

func NewNode(config NodeConfig) (*Node, error) {
	h, err := makeHost(config.Port, rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to create libp2p host: %s", err)
	}

	log.Printf("Node %s started with ID: %s\n", config.ID, h.ID().String())
	for _, addr := range h.Addrs() {
		log.Printf("Listening on: %s\n", addr)
	}

	return &Node{
		ID:   config.ID,
		Host: h,
	}, nil
}

func makeHost(port int, randomness io.Reader) (host.Host, error) {
	// Creates a new RSA key pair for this host.
	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, randomness)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	sourceMultiAddr, err := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return libp2p.New(
		libp2p.ListenAddrs(sourceMultiAddr),
		libp2p.Identity(prvKey),
	)
}
