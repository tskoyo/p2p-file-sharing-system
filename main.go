package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"p2p-file-sharing-system/helper"
	"p2p-file-sharing-system/peer"
	"p2p-file-sharing-system/types"
	"strconv"
)

func main() {
	id := flag.String("id", "ABC", "Id of the node")
	peerPort := flag.String("peer-port", "9002", "Port of the peer to connect to")
	command := flag.String("command", "start-server", "Command to either start or connect to a node")
	remotePeerAddr := flag.String("remote-peer-addr", "0.0.0.0", "Peer address to connect to")
	remotePeerPort := flag.String("remote-peer-port", "0.0.0.0", "Peer port to connect to")
	remotePeerID := flag.String("remote-peer-id", "0.0.0.0", "Peer id to connect to")

	flag.Parse()

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

	switch *command {
	case "start-server":
	case "connect-to-peer":
		multiAddr := fmt.Sprintf("/ip4/%s/tcp/%s/p2p/%s", *remotePeerAddr, *remotePeerPort, *remotePeerID)
		err = nodeA.Connect(multiAddr)
		if err != nil {
			helper.PrintError(fmt.Sprintf("Failed to connect to peer: %v", err))
		} else {
			helper.PrintSuccess("Successfully connected to peer")
		}
	default:
		log.Fatalf("Unknown command: %s. Use 'start-server' or 'connect-to-peer'.", *command)
		os.Exit(1)
	}

	select {}
}
