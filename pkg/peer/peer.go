package peer

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var connections = NewConnectionPool()

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
		go handleConnection(conn, 0)
	}
}

func handleConnection(conn net.Conn, connectionType int) {
	log.Println("Handling incoming connection...")
	defer conn.Close()

	address := conn.RemoteAddr().String()
	connections.Add(address, conn)

	log.Printf("Connection added: %s", address)

	switch connectionType {
	case 0: // upload
		err := receiveFile(conn, "received_file.txt")
		if err != nil {
			log.Printf("Error receiving file: %v", err)
		}
	case 1: // download
		err := sendFile(conn, "received_file.txt")
		if err != nil {
			log.Printf("Error sending file: %v", err)
		}
	default:
		log.Printf("Unknown command: %d", connectionType)
	}
}

func receiveFile(conn net.Conn, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Use io.Copy to copy remaining data from the conn
	_, err = io.Copy(file, conn)
	if err != nil {
		return fmt.Errorf("failed to copy data to file: %w", err)
	}

	return nil
}

func sendFile(conn net.Conn, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(conn, file)
	if err != nil {
		return fmt.Errorf("failed to send file: %w", err)
	}

	log.Printf("File '%s' sent successfully", filename)
	return nil
}
