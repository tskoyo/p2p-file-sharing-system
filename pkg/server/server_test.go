package server

// func TestMultipleClientsConnectingToServer(t *testing.T) {
// 	serverAddress := "localhost:9000"
// 	numClients := 100

// 	connectionPool := NewConnectionPool()
// 	server := NewServer(connectionPool, "9000")

// 	readyChan := make(chan error, 1)
// 	go StartServer(serverAddress, readyChan)

// 	err := <-readyChan
// 	if err != nil {
// 		t.Fatalf("Failed to start server: %v", err)
// 	}

// 	var wg sync.WaitGroup
// 	clientErrors := make(chan error, numClients)

// 	for i := 0; i < numClients; i++ {
// 		wg.Add(1)
// 		go func(clientID int) {
// 			defer wg.Done()

// 			clientAddr := "localhost:0"
// 			dialer := &client.NetDialer{}

// 			clientNode := client.NewClient(dialer, clientAddr)
// 			err := clientNode.Connect(serverAddress)
// 			if err != nil {
// 				clientErrors <- err
// 				log.Printf("[Client %d] Failed tooo connect: %v", clientID, err)
// 				return
// 			}

// 			clientErrors <- nil
// 		}(i)
// 	}

// 	wg.Wait()
// 	close(clientErrors)

// 	for err := range clientErrors {
// 		if err != nil {
// 			t.Errorf("Client error: %v", err)
// 		}
// 	}

// 	numberOfConnections := len(connections.List())
// 	assert.Equal(t, numClients, numberOfConnections, "Number of connections should be %d", numClients)
// }
