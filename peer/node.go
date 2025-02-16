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

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
)

type Node struct {
	Config types.NodeConfig
	Host   host.Host
}

func NewNode(config types.NodeConfig) (*Node, error) {
	h, err := makeHost(config.Port, rand.Reader)
	if err != nil {
		helper.PrintError("Failed to create libp2p host:")
		return nil, fmt.Errorf("failed to create libp2p host: %s", err)
	}

	helper.PrintInfo(fmt.Sprintf("Node %s started with ID: %s\n", config.ID, h.ID().String()))
	for _, addr := range h.Addrs() {
		log.Printf("Listening on: %s\n", addr)
	}

	return &Node{
		Config: config,
		Host:   h,
	}, nil
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
		helper.PrintError(fmt.Sprintf("Failed to connect to peer: %s", err))
		return fmt.Errorf("Failed to connect to peer")
	}

	helper.PrintSuccess(fmt.Sprintf("%s successfully connected to: %s", n.Host.Network().LocalPeer(), multiAddr))
	return nil
}

func makeHost(port int, randomness io.Reader) (host.Host, error) {
	// Creates a new RSA key pair for this host.
	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, randomness)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	sourceMultiAddr, err := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return libp2p.New(
		libp2p.ListenAddrs(sourceMultiAddr),
		libp2p.Identity(prvKey),
	)
}
