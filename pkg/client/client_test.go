package client

import (
	"p2p-file-sharing-system/pkg/peer"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConnect_Success(t *testing.T) {
	node1PortId := 9001
	node2PortId := 9003

	node1Config := peer.NodeConfig{
		ID:   "peer-1",
		Port: node1PortId,
	}
	node2Config := peer.NodeConfig{
		ID:   "peer-2",
		Port: node2PortId,
	}

	node1, err := peer.NewNode(node1Config)
	require.NoError(t, err)

	node2, err := peer.NewNode(node2Config)
	require.NoError(t, err)

	client2 := NewClient(node2.Host, nil)

	node1MultiAddr := node1.Host.Addrs()[0].String() + "/p2p/" + node1.Host.ID().String()
	err = client2.Connect(node1MultiAddr)

	require.NoError(t, err)

	connections := node1.Host.Network().ConnsToPeer(client2.Host.ID())

	for _, v := range connections {
		t.Logf("address id: %t", v.ID() == node2.Host.ID().String())
	}
}

// func TestConnect_Success(t *testing.T) {
// 	t.Run("success on first dial", func(t *testing.T) {
// 		mockDialer := &MockDialer{}
// 		mockConn, serverConn := NewMockConn()

// 		// 3. Simulate server behavior in a goroutine
// 		//    If your handshakeWithServer() writes "HELLO" and expects "OK",
// 		//    you must actually read "HELLO" from serverConn and then write "OK".
// 		go func() {
// 			// Read "HELLO" from client
// 			buf := make([]byte, 5)
// 			n, err := serverConn.Read(buf)
// 			if err != nil {
// 				// handle error or just log
// 				return
// 			}
// 			// We expect "HELLO"
// 			// if needed, verify the string: string(buf[:n]) == "HELLO"

// 			// Write "OK" back so the client sees it
// 			_, _ = serverConn.Write([]byte("OK"))

// 			// Close the server side when done
// 			serverConn.Close()
// 		}()

// 		peerAddress := "localhost:9000"
// 		mockDialer.On("Dial", TCP, peerAddress).Return(mockConn, nil).Once()

// 		client := NewClient(mockDialer, "localhost:9001")
// 		err := client.Connect(peerAddress)

// 		assert.NoError(t, err, "Connect should succeed")
// 		mockDialer.AssertExpectations(t)
// 	})
// }

// // func TestDialer_Failure(t *testing.T) {
// // 	mockDialer := &MockDialer{
// // 		MockBehavior: func(network NetworkType, peerAddress string) (net.Conn, error) {
// // 			return nil, errors.New("mock connection error")
// // 		},
// // 	}

// // 	conn, err := mockDialer.Dial(TCP, "127.0.0.1:8080")
// // 	assert.Error(t, err)
// // 	assert.Nil(t, conn)
// // }
