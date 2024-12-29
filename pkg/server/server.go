package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	connectionpool "p2p-file-sharing-system/pkg/connection_pool"
	"time"
)

type Server struct {
	connectionPool *connectionpool.ConnectionPool
	Port           string
}

func NewServer(connectionPool *connectionpool.ConnectionPool, serverPort string) *Server {
	return &Server{
		connectionPool: connectionPool,
		Port:           serverPort,
	}
}

// var connections = connectionpool.NewConnectionPool()

func (s *Server) Start(address string, readyChan chan<- error) {
	listener, err := net.Listen("tcp", ":"+s.Port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		readyChan <- err
		return
	}

	readyChan <- nil // server is listening successfully

	defer listener.Close()
	log.Printf("Server listening on %s", ":"+s.Port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Connection error: %v", err)
			continue
		}
		s.connectionPool.Add(conn.RemoteAddr().String(), conn)
		go handleConnection(conn)
	}
}

func SyncConnection(localAddress, peerAddress string) error {
	conn, err := net.ListenPacket("udp", localAddress)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Send a packet to the peer to punch a hole in the NAT
	peerAddr, err := net.ResolveUDPAddr("udp", peerAddress)
	if err != nil {
		return err
	}

	log.Printf("Sending synchronization packet to: %s", peerAddress)
	_, err = conn.WriteTo([]byte("sync"), peerAddr)
	if err != nil {
		return err
	}

	// Wait for a response from the peer
	buffer := make([]byte, 1024)
	err = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		return err
	}

	n, addr, err := conn.ReadFrom(buffer)
	if err != nil {
		return err
	}

	log.Printf("Received synchronization packet from: %s", addr.String())
	log.Printf("Data: %s", string(buffer[:n]))

	return nil
}

func handleConnection(conn net.Conn) {
	buf := make([]byte, 5)

	_, err := conn.Read(buf)
	if err != nil {
		return
	}

	handshakeRespMsg := []byte("OK")
	_, err = conn.Write(handshakeRespMsg)

	if err != nil {
		return
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
