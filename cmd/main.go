package main

import (
	"flag"
	"log"
	"p2p-file-sharing-system/pkg/peer"
)

func main() {
	id := flag.String("id", "ABC", "Id of the node")
	serverAddress := flag.String("server-address", "localhost", "Address of the server")
	serverPort := flag.String("server-port", "9000", "Port of the server")
	clientPort := flag.String("client-port", "9001", "Port of the client")
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
	nodeA.StartServer()

	select {}
}
