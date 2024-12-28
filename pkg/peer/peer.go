package peer

import (
	"log"
	"net"
)

var connections = NewConnectionPool()

func StartServer(address string, readyChan chan<- error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		readyChan <- err
		return
	}

	readyChan <- nil // server is listening successfully

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
	buf := make([]byte, 5)

	_, err := conn.Read(buf)
	if err != nil {
		return
	}

	connections.Add(conn.RemoteAddr().String(), conn)

	handshakeRespMsg := []byte("OK")
	_, err = conn.Write(handshakeRespMsg)

	if err != nil {
		return
	}
}

// func receiveFile(conn net.Conn, filename string) error {
// 	file, err := os.Create(filename)
// 	if err != nil {
// 		return fmt.Errorf("failed to open file: %w", err)
// 	}
// 	defer file.Close()

// 	// Use io.Copy to copy remaining data from the conn
// 	_, err = io.Copy(file, conn)
// 	if err != nil {
// 		return fmt.Errorf("failed to copy data to file: %w", err)
// 	}

// 	return nil
// }

// func sendFile(conn net.Conn, filename string) error {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return fmt.Errorf("failed to open file: %w", err)
// 	}
// 	defer file.Close()

// 	_, err = io.Copy(conn, file)
// 	if err != nil {
// 		return fmt.Errorf("failed to send file: %w", err)
// 	}

// 	log.Printf("File '%s' sent successfully", filename)
// 	return nil
// }
