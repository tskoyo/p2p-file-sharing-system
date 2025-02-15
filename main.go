package main

import (
	"flag"
	"log"
	"p2p-file-sharing-system/pkg/peer"
	"strconv"
)

func main() {
	id := flag.String("id", "ABC", "Id of the node")
	peerPort := flag.String("peer-port", "9002", "Port of the peer to connect to")

	flag.Parse()

	port, err := strconv.Atoi(*peerPort)

	if err != nil {
		panic(err)
	}

	nodeAConfig := &peer.NodeConfig{
		ID:   *id,
		Port: port,
	}

	nodeA, err := peer.NewNode(*nodeAConfig)
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
