package main

import (
	"log"
	"p2p-file-sharing-system/pkg/peer"
	"p2p-file-sharing-system/pkg/transfer"
)

func main() {
	go peer.StartServer("localhost:9000")

	client := peer.NewClient("localhost:9000")
	if err := client.Connect(); err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}

	// Example file transfer
	err := transfer.SendFile(client, "example.txt")

	if err != nil {
		panic(err)
	}

	select {}
}
