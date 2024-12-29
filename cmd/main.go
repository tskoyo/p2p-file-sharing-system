package main

import (
	"log"
	"p2p-file-sharing-system/pkg/peer"
	"time"
)

func main() {
	// mode := flag.String("mode", "server", "Start as 'server' or 'client'")
	// address := flag.String("address", "localhost:9000", "Address of the client")
	// peerAddress := flag.String("peer", "localhost:9002", "Address of the peer to connect to")
	// // filePath := flag.String("file", "", "Path to the file to upload (client mode only)")
	// flag.Parse()

	// switch *mode {
	// case "server":
	// 	readyChan := make(chan error, 1)
	// 	go peer.StartServer(*address, readyChan)

	// 	if err := <-readyChan; err != nil {
	// 		log.Fatalf("Failed to start server: %v", err)
	// 	}
	// case "client":
	// 	dialer := &client.NetDialer{}
	// 	client := client.NewClient(dialer, *address)
	// 	if err := client.Connect(*peerAddress); err != nil {
	// 		log.Fatalf("Failed to connect to server: %v", err)
	// 	}

	// 	// err := client.UploadFile(*filePath)

	// 	// if err != nil {
	// 	// 	panic(err)
	// 	// }
	// default:
	// 	log.Fatalf("Unknown mode: %s", *mode)
	// }

	// select {}

	nodeA, err := peer.NewNode("A", "stun.l.google.com:19302", "5000")
	if err != nil {
		log.Fatalf("Failed to initialize Node A: %v", err)
	}
	nodeA.StartServer()

	nodeB, err := peer.NewNode("B", "stun.l.google.com:19302", "5001")
	if err != nil {
		log.Fatalf("Failed to initialize Node B: %v", err)
	}
	nodeB.StartServer()

	time.Sleep(2 * time.Second)

	err = nodeB.ConnectToPeer(nodeA.PublicAddress)
	if err != nil {
		log.Fatalf("Failed to connect to peer: %v", err)
	}

	err = nodeA.ConnectToPeer(nodeB.PublicAddress)
	if err != nil {
		log.Fatalf("Failed to connect to peer: %v", err)
	}

	nodeA.ListConnections()
	nodeB.ListConnections()
}
