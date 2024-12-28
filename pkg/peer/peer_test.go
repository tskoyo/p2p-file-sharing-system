package peer

import (
	"log"
	"p2p-file-sharing-system/pkg/client"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultipleClientsConnectingToServer(t *testing.T) {
	serverAddress := "localhost:9000"
	numClients := 100

	readyChan := make(chan error, 1)
	go StartServer(serverAddress, readyChan)

	err := <-readyChan
	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}

	// time.Sleep(2 * time.Second) // Wait for the server to start

	var wg sync.WaitGroup
	clientErrors := make(chan error, numClients)

	for i := 0; i < numClients; i++ {
		wg.Add(1)
		go func(clientID int) {
			defer wg.Done()

			clientAddr := "localhost:0"
			clientNode := client.NewClient(clientAddr)
			err := clientNode.Connect(serverAddress)
			if err != nil {
				clientErrors <- err
				log.Printf("[Client %d] Failed tooo connect: %v", clientID, err)
				return
			}

			clientErrors <- nil
		}(i)
	}

	wg.Wait()
	close(clientErrors)

	for err := range clientErrors {
		if err != nil {
			t.Errorf("Client error: %v", err)
		}
	}

	numberOfConnections := len(connections.List())
	assert.Equal(t, numClients, numberOfConnections, "Number of connections should be %d", numClients)
}
