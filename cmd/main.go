package main

import (
	"flag"
	"log"
	"p2p-file-sharing-system/pkg/peer"
)

func main() {
	command := flag.String("command", "start-server", "Command to execute: start-server or connect-to-peer")
	id := flag.String("id", "ABC", "Id of the node")
	serverAddress := flag.String("server-address", "localhost", "Address of the server")
	serverPort := flag.String("server-port", "9000", "Port of the server")
	clientPort := flag.String("client-port", "9001", "Port of the client")
	peerAddress := flag.String("peer-address", "localhost", "Address of the peer to connect to")
	peerPort := flag.String("peer-port", "9002", "Port of the peer to connect to")

	flag.Parse()

	nodeAConfig := &peer.NodeConfig{
		Id:            *id,
		StunServer:    "stun.l.google.com:19302",
		ServerAddress: *serverAddress,
		ServerPort:    *serverPort,
		ClientPort:    *clientPort,
	}

	nodeA, err := peer.NewNode(*nodeAConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Node A: %v", err)
	}

	switch *command {
	case "start-server":
		log.Printf("Starting server with ID %s on port %s", nodeAConfig.Id, nodeAConfig.ServerPort)
		nodeA.StartServer()
	case "connect-to-peer":
		err = nodeA.ConnectToPeer(*peerAddress, *peerPort)
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
