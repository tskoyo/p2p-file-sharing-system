package peer

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
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

const maxRetries = 3

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

func (n *Node) Connect(multiAddr string) error {
	peerMultiAddr, err := multiaddr.NewMultiaddr(multiAddr)
	if err != nil {
		return fmt.Errorf("Invalid multiaddress: %w", err)
	}

	peerInfo, err := peer.AddrInfoFromP2pAddr(peerMultiAddr)
	if err != nil {
		return fmt.Errorf("Failed to parse peer address: %w", err)
	}

	log.Printf("%s attempting to connet to %s", n.Host.ID(), multiAddr)

	var lastErr error
	// Connect with retries
	for i := 0; i < maxRetries; i++ {
		err := n.Host.Connect(context.Background(), *peerInfo)
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
