package main

import (
	"flag"
	"log"
	"p2p-file-sharing-system/pkg/client"
	"p2p-file-sharing-system/pkg/peer"
)

func main() {
	mode := flag.String("mode", "server", "Start as 'server' or 'client'")
	address := flag.String("address", "localhost:9000", "Address of the client")
	peerAddress := flag.String("peer", "localhost:9002", "Address of the peer to connect to")
	// filePath := flag.String("file", "", "Path to the file to upload (client mode only)")
	flag.Parse()

	switch *mode {
	case "server":
		// stopChan := make(chan struct{})
		go peer.StartServer(*address)
	case "client":
		client := client.NewClient(*address)
		if err := client.Connect(*peerAddress); err != nil {
			log.Fatalf("Failed to connect to server: %v", err)
		}

		// err := client.UploadFile(*filePath)

		// if err != nil {
		// 	panic(err)
		// }
	default:
		log.Fatalf("Unknown mode: %s", *mode)
	}

	select {}
}
