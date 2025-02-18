package peer

import (
	"context"
	"fmt"
	"net"
	"os"
	"p2p-file-sharing-system/helper"
	"time"

	"github.com/libp2p/go-libp2p/p2p/muxer/yamux"
)

func OpenStream(peerMultiAddr string) error {
	transport := buildTransport()

	conn, err := net.Dial("tcp", "127.0.0.1:9001")
	if err != nil {
		return fmt.Errorf("Failed to dial: %v", err)
	}

	muxedConn, err := transport.NewConn(conn, false, nil)
	if err != nil {
		return fmt.Errorf("Failed to create MuxedConn: %v", err)
	}

	stream, err := muxedConn.OpenStream(context.Background())
	if err != nil {
		return fmt.Errorf("failed to open stream: %v", err)
	}

	_, err = stream.Write([]byte("Hello peer!"))
	if err != nil {
		return fmt.Errorf("Failed to write to stream: %v", err)
	}

	helper.PrintSuccess(fmt.Sprintf("Created MuxedConn: %v", muxedConn))
	return nil
}

func buildTransport() yamux.Transport {
	transport := &yamux.Transport{}
	transport.AcceptBacklog = 256
	transport.PingBacklog = 32
	transport.EnableKeepAlive = true
	transport.KeepAliveInterval = 30 * time.Second
	transport.MeasureRTTInterval = 30 * time.Second
	transport.ConnectionWriteTimeout = 10 * time.Second
	transport.MaxIncomingStreams = 1000
	transport.InitialStreamWindowSize = yamux.DefaultTransport.InitialStreamWindowSize
	transport.MaxStreamWindowSize = yamux.DefaultTransport.MaxStreamWindowSize
	transport.LogOutput = os.Stderr
	transport.ReadBufSize = 4096
	transport.MaxMessageSize = 64 * 1024
	transport.WriteCoalesceDelay = 100 * time.Microsecond

	return *transport
}
