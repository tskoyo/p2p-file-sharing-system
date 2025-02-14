package main

import (
	"context"
	"flag"
	"log"
	"p2p-file-sharing-system/pkg/peer"
)

func main() {
	// command := flag.String("command", "start-server", "Command to execute: start-server or connect-to-peer")
	id := flag.String("id", "ABC", "Id of the node")
	// peerAddress := flag.String("peer-address", "localhost", "Address of the peer to connect to")
	// peerPort := flag.String("peer-port", "9002", "Port of the peer to connect to")

	flag.Parse()

	nodeAConfig := &peer.NodeConfig{
		ID: *id,
	}

	ctx := context.Background()
	nodeA, err := peer.NewNode(ctx, *nodeAConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Node A: %v", err)
	}

	log.Println("Node id: ", nodeA.ID)

	// switch *command {
	// case "start-server":
	// 	log.Printf("Starting server with ID %s on port %s", nodeAConfig.Id, nodeAConfig.ServerPort)
	// 	// nodeA.StartServer()
	// case "connect-to-peer":
	// 	// err = nodeA.ConnectToPeer(*peerAddress, *peerPort)
	// 	if err != nil {
	// 		log.Fatalf("Failed to connect to peer: %v", err)
	// 	} else {
	// 		log.Println("Successfully connected to peer")
	// 	}
	// default:
	// 	log.Fatalf("Unknown command: %s. Use 'start-server' or 'connect-to-peer'.", *command)
	// }

	select {}
}
