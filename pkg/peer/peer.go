package peer

import (
	"io"
	"log"
	"net"
	"os"
)

func StartServer(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()

	log.Printf("Server listening on %s", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Connection error: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	log.Println("Server handling connection")

	file, err := os.Create("received_file.txt")
	if err != nil {
		log.Fatal("Error creating a file: ", err)
	}

	defer file.Close()

	_, err = io.Copy(file, conn)
	if err != nil {
		log.Fatal("Failed to send file: ", err)
	}
}
