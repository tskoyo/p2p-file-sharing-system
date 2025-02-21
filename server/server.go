package server

import (
	"fmt"
	"p2p-file-sharing-system/helper"
	"p2p-file-sharing-system/peer"
	"p2p-file-sharing-system/types"
)

// type Server struct {
// 	Address multiaddr.Multiaddr
// 	Port    int
// }

// func NewServer(address multiaddr.Multiaddr, port int, config types.NodeConfig) *Server {
// 	return &Server{
// 		Address: address,
// 		Port:    port,
// 	}
// }

func Start(config types.NodeConfig) error {
	_, err := peer.NewNode(config)
	if err != nil {
		helper.PrintError(fmt.Sprintf("Failed to create node: %s", err))
		return fmt.Errorf("Failed to create node: %w", err)
	}

	return nil
}

func Connect(config types.NodeConfig, multiAddr string) error {
	node, err := peer.NewNode(config)
	if err != nil {
		return fmt.Errorf("Failed to create node: %w", err)
	}

	if err = node.Connect(multiAddr); err != nil {
		return fmt.Errorf("Failed to connect to node: %w", err)
	}

	return nil
}

// func SyncConnection(localAddress, peerAddress string) error {
// 	conn, err := net.ListenPacket("udp", localAddress)
// 	if err != nil {
// 		return err
// 	}
// 	defer conn.Close()

// 	// Send a packet to the peer to punch a hole in the NAT
// 	peerAddr, err := net.ResolveUDPAddr("udp", peerAddress)
// 	if err != nil {
// 		return err
// 	}

// 	log.Printf("Sending synchronization packet to: %s", peerAddress)
// 	_, err = conn.WriteTo([]byte("sync"), peerAddr)
// 	if err != nil {
// 		return err
// 	}

// 	// Wait for a response from the peer
// 	buffer := make([]byte, 1024)
// 	err = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
// 	if err != nil {
// 		return err
// 	}

// 	n, addr, err := conn.ReadFrom(buffer)
// 	if err != nil {
// 		return err
// 	}

// 	log.Printf("Received synchronization packet from: %s", addr.String())
// 	log.Printf("Data: %s", string(buffer[:n]))

// 	return nil
// }

// func handleConnection(conn net.Conn) {

// 	buf := make([]byte, 5)

// 	_, err := conn.Read(buf)
// 	if err != nil {
// 		return
// 	}

// 	handshakeRespMsg := []byte("OK")
// 	_, err = conn.Write(handshakeRespMsg)

// 	if err != nil {
// 		return
// 	}
// }
