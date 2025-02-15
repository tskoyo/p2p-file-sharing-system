package server

import (
	"log"
	"net"
	"strconv"
	"time"

	"github.com/multiformats/go-multiaddr"
)

type Server struct {
	Address multiaddr.Multiaddr
	Port    int
}

func NewServer(address multiaddr.Multiaddr, port int) *Server {
	return &Server{
		Address: address,
		Port:    port,
	}
}

func (s *Server) Start(readyChan chan<- error) {
	listener, err := net.Listen("tcp", s.Address.String()+":"+strconv.Itoa(s.Port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		readyChan <- err
		return
	}

	readyChan <- nil // server is listening successfully

	defer listener.Close()
	log.Printf("Server listening on %s", s.Address.String()+":"+strconv.Itoa(s.Port))

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Connection error: %v", err)
			continue
		}
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
