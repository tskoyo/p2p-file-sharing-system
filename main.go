package main

import (
	"flag"
	"fmt"
	"log"
	"p2p-file-sharing-system/helper"
	"p2p-file-sharing-system/peer"
	"p2p-file-sharing-system/types"
	"strconv"
)

func main() {
	id := flag.String("id", "ABC", "Id of the node")
	peerPort := flag.String("peer-port", "9002", "Port of the peer to connect to")
	command := flag.String("command", "start-server", "Command to either start or connect to a node")
	peerAddr := flag.String("peer-addr", "0.0.0.0", "Peer address to connect to")

	flag.Parse()

	helper.PrintInfo(fmt.Sprintf("Peer port is: %s", *peerPort))

	port, err := strconv.Atoi(*peerPort)

	if err != nil {
		panic(err)
	}

	nodeAConfig := &types.NodeConfig{
		ID:   *id,
		Port: port,
	}

	nodeA, err := peer.NewNode(*nodeAConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Node A: %v", err)
	}

	log.Println("Node id: ", nodeA.Config.ID)

	switch *command {
	case "start-server":
		helper.PrintInfo(fmt.Sprintf("Starting server with ID %s on port %d", nodeA.Config.ID, nodeA.Config.Port))
		// nodeA.StartServer()
	case "connect-to-peer":
		err = nodeA.Connect(*peerAddr)
		if err != nil {
			log.Fatalf("Failed to connect to peer: %v", err)
		} else {
			log.Println("Successfully connected to peer")
		}
	default:
		log.Fatalf("Unknown command: %s. Use 'start-server' or 'connect-to-peer'.", *command)
	}

	select {}
}
