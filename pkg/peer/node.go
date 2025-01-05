package peer

import (
	"fmt"
	"log"
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

type NodeConfig struct {
	Id            string
	StunServer    string
	ServerAddress string
	ServerPort    string
	ClientPort    string
}

func NewNode(config NodeConfig) (*Node, error) {
	publicAddress, err := discoverPublicAddress(config.StunServer)
	if err != nil {
		log.Fatalf("Failed to discover public address: %v", err)
		return nil, err
	}

	log.Printf("Node %s public address: %s", config.Id, publicAddress)

	connectionPool := connectionpool.NewConnectionPool()
	server := server.NewServer(connectionPool, config.ServerAddress, config.ServerPort)
	client := client.NewClient(connectionPool)

	return &Node{
		ID:             config.Id,
		PublicAddress:  publicAddress,
		ConnectionPool: connectionPool,
		Server:         server,
		Client:         client,
	}, nil
}

func (n *Node) StartServer() {
	readyChan := make(chan error, 1)
	go n.Server.Start(readyChan)

	if err := <-readyChan; err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func (n *Node) ConnectToPeer(peerAddress string, port string) error {
	if err := n.Client.Connect(peerAddress, port); err != nil {
		return fmt.Errorf("Failed to connect to peer: %v", err)
	}

	return nil
}

func discoverPublicAddress(stunServer string) (string, error) {
	stunClient := stun.NewClient(stunServer)
	publicAddress, err := stunClient.DiscoverPublicAddress()
	if err != nil {
		return "", fmt.Errorf("error discovering public address: %w", err)
	}
	return publicAddress, nil
}
