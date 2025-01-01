package client

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

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
