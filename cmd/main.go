package main

import (
	"log"
	"p2p-file-sharing-system/pkg/client"
	"p2p-file-sharing-system/pkg/peer"
)

func main() {
	go peer.StartServer("localhost:9000")

	client := client.NewClient("localhost:9000")
	if err := client.Connect(); err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}

	err := client.UploadFile("example.txt")

	if err != nil {
		panic(err)
	}

	select {}
}
