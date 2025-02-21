package peer

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"p2p-file-sharing-system/helper"
	"p2p-file-sharing-system/types"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/p2p/muxer/yamux"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
)

type Node struct {
	Config types.NodeConfig
	Host   host.Host
}

const ProtocolID = "/file-transfer/1.0.0"

func NewNode(config types.NodeConfig) (*Node, error) {
	h, err := makeHost(config.Port, rand.Reader)
	if err != nil {
		helper.PrintError("Failed to create libp2p host:")
		return nil, fmt.Errorf("failed to create libp2p host: %s", err)
	}

	helper.PrintInfo(fmt.Sprintf("Node started with ID: %s\n", h.ID().String()))
	for _, addr := range h.Addrs() {
		log.Printf("Listening on: %s\n", addr)
	}

	h.SetStreamHandler(ProtocolID, func(s network.Stream) {
		buf := make([]byte, 1024)
		n, err := s.Read(buf)
		if err != nil {
			log.Printf("error reading from stream: %v\n", err)
			return
		}

		log.Printf("Host %v received: %v\n", h.ID(), string(buf[:n]))
	})

	node := &Node{
		Config: config,
		Host:   h,
	}

	return node, nil
}

func (n *Node) Connect(multiAddr string) error {
	peerMultiAddr, err := multiaddr.NewMultiaddr(multiAddr)
	if err != nil {
		return fmt.Errorf("Invalid multiaddress: %w", err)
	}

	peerInfo, err := peer.AddrInfoFromP2pAddr(peerMultiAddr)
	if err != nil {
		return fmt.Errorf("Failed to parse peer address: %w", err)
	}

	helper.PrintInfo(fmt.Sprintf("%s attempting to connet to %s", n.Host.ID(), multiAddr))

	if err := n.Host.Connect(context.Background(), *peerInfo); err != nil {
		return fmt.Errorf("Failed to connect to peer")
	}

	helper.PrintSuccess(fmt.Sprintf("%s successfully connected to: %s", n.Host.Network().LocalPeer(), multiAddr))

	stream, err := n.Host.NewStream(context.Background(), peerInfo.ID, protocol.ID(ProtocolID))
	if err != nil {
		return fmt.Errorf("Failed to open new stream: %w", err)
	}
	defer stream.Close()

	text := []byte("Hello from just a regular every day normal motherfucker")

	_, err = stream.Write(text)
	if err != nil {
		return fmt.Errorf("Error writing a message to stream: %w", err)
	}

	return nil
}

func makeHost(port int, randomness io.Reader) (host.Host, error) {
	// Creates a new RSA key pair for this host.
	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, randomness)
	if err != nil {
		return nil, err
	}

	sourceMultiAddr, err := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port))

	if err != nil {
		return nil, err
	}

	host, err := libp2p.New(
		libp2p.ListenAddrs(sourceMultiAddr),
		libp2p.Identity(prvKey),
		libp2p.Muxer("/yamux/1.0.0", yamux.DefaultTransport),
	)

	if err != nil {
		return nil, err
	}

	return host, nil
}
