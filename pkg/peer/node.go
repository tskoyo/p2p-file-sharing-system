package peer

import (
	"context"
	"fmt"
	"log"

	libp2p "github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"

	// discovery "github.com/libp2p/go-libp2p/p2p/discovery"
	"github.com/libp2p/go-libp2p/core/host"
)

type Node struct {
	ID   string
	Host host.Host
	DHT  *dht.IpfsDHT
	// Discovery *discovery.RoutingDiscovery
}

type NodeConfig struct {
	ID string
}

func NewNode(ctx context.Context, config NodeConfig) (*Node, error) {
	h, err := libp2p.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create libp2p host: %s", err)
	}

	log.Printf("Node %s started with ID: %s\n", config.ID, h.ID().String())
	for _, addr := range h.Addrs() {
		log.Printf("Listening on: %s\n", addr)
	}

	kadDHT, err := dht.New(ctx, h)
	if err != nil {
		return nil, fmt.Errorf("failed to create DHT: %w", err)
	}

	// Bootstrap the DHT
	if err := kadDHT.Bootstrap(ctx); err != nil {
		return nil, fmt.Errorf("failed to bootstrap DHT: %w", err)
	}

	return &Node{
		ID:   config.ID,
		Host: h,
		DHT:  kadDHT,
	}, nil
}

// func NewNode(config NodeConfig) (*Node, error) {
// 	publicAddress, err := discoverPublicAddress(config.StunServer)
// 	if err != nil {
// 		log.Fatalf("Failed to discover public address: %v", err)
// 		return nil, err
// 	}

// 	log.Printf("Node %s public address: %s", config.Id, publicAddress)

// 	connectionPool := connectionpool.NewConnectionPool()
// 	server := server.NewServer(connectionPool, config.ServerAddress, config.ServerPort)
// 	client := client.NewClient(connectionPool)

// 	return &Node{
// 		ID:             config.Id,
// 		PublicAddress:  publicAddress,
// 		ConnectionPool: connectionPool,
// 		Server:         server,
// 		Client:         client,
// 	}, nil
// }

// func (n *Node) StartServer() {
// 	readyChan := make(chan error, 1)
// 	go n.Server.Start(readyChan)

// 	if err := <-readyChan; err != nil {
// 		log.Fatalf("Failed to start server: %v", err)
// 	}
// }

// func (n *Node) ConnectToPeer(peerAddress string, port string) error {
// 	if err := n.Client.Connect(peerAddress, port); err != nil {
// 		return fmt.Errorf("Failed to connect to peer: %v", err)
// 	}

// 	return nil
// }

// func discoverPublicAddress(stunServer string) (string, error) {
// 	stunClient := stun.NewClient(stunServer)
// 	publicAddress, err := stunClient.DiscoverPublicAddress()
// 	if err != nil {
// 		return "", fmt.Errorf("error discovering public address: %w", err)
// 	}
// 	return publicAddress, nil
// }
