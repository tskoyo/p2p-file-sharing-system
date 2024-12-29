package peer

import (
	"fmt"
	"log"
	"net"
	"p2p-file-sharing-system/pkg/client"
	connectionpool "p2p-file-sharing-system/pkg/connection_pool"
	"p2p-file-sharing-system/pkg/server"
	"p2p-file-sharing-system/stun"
)

type Node struct {
	ID             string
	PublicAddress  string
	ConnectionPool *connectionpool.ConnectionPool
	Server         *server.Server
	Client         *client.Client
}

func NewNode(id, stunServer string, serverPort string) (*Node, error) {
	publicAddress, err := discoverPublicAddress(stunServer)
	if err != nil {
		log.Fatalf("Failed to discover public address: %v", err)
		return nil, err
	}

	formattedPublicAddress, err := formatPublicAddress(publicAddress, serverPort)
	if err != nil {
		return nil, fmt.Errorf("failed to format public address: %w", err)
	}

	log.Printf("Node %s public address: %s", id, formattedPublicAddress)

	connectionPool := connectionpool.NewConnectionPool()
	server := server.NewServer(connectionPool, serverPort)
	client := client.NewClient(connectionPool)

	return &Node{
		ID:             id,
		PublicAddress:  formattedPublicAddress,
		ConnectionPool: connectionPool,
		Server:         server,
		Client:         client,
	}, nil
}

func (n *Node) StartServer() {
	readyChan := make(chan error, 1)
	go n.Server.Start("localhost:9000", readyChan)

	if err := <-readyChan; err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func (n *Node) ConnectToPeer(peerAddress string) error {
	if err := n.Client.Connect(n.PublicAddress); err != nil {
		return fmt.Errorf("Failed to connect to peer: %v", err)
	}

	return nil
}

func (n *Node) ListConnections() {
	log.Printf("Connections for %v: %v", n.ID, n.ConnectionPool.List())
}

func discoverPublicAddress(stunServer string) (string, error) {
	stunClient := stun.NewClient(stunServer)
	publicAddress, err := stunClient.DiscoverPublicAddress()
	if err != nil {
		return "", fmt.Errorf("error discovering public address: %w", err)
	}
	return publicAddress, nil
}

func formatPublicAddress(publicAddress, serverPort string) (string, error) {
	host, _, err := net.SplitHostPort(publicAddress)
	if err != nil {
		return "", fmt.Errorf("error splitting public address: %w", err)
	}
	return fmt.Sprintf("%s:%s", host, serverPort), nil
}
