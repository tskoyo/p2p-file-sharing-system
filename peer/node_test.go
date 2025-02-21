package peer

import (
	"fmt"
	"p2p-file-sharing-system/helper"
	"p2p-file-sharing-system/types"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewNode_Succes(t *testing.T) {
	config := types.NodeConfig{
		ID:   "success-node",
		Port: 9001,
	}

	node := createTestNode(t, config)

	assert.Equal(t, node.Config.ID, config.ID)
	assert.NotNil(t, node.Host, "Host should not be nil")
}

func TestNewNode_WithInvalidPort(t *testing.T) {
	config := types.NodeConfig{
		ID:   "invalid-node",
		Port: -1,
	}

	_, err := NewNode(config)
	assert.NotNil(t, err, "Expected error for invalid port")
}

func TestConnect_MultipleClients(t *testing.T) {
	// # TODO: Find a way to handle this more effectively
	// currently only 8 nodes can connect simulatenously to one node
	numClients := 8
	var wg sync.WaitGroup
	clientErrors := make(chan error, numClients)

	node1Config := helper.BuildNodeConfig("peer-1", 9001)
	node1 := createTestNode(t, node1Config)

	node1MultiAddr := node1.Host.Addrs()[0].String() + "/p2p/" + node1.Host.ID().String()

	for i := 1; i <= numClients; i++ {
		wg.Add(1)

		go func(clientId int) {
			defer wg.Done()

			clientNodeConfig := helper.BuildNodeConfig(fmt.Sprintf("client-node-%d", clientId), 9001+clientId)
			clientNode := createTestNode(t, clientNodeConfig)

			err := clientNode.Connect(node1MultiAddr)
			require.NoError(t, err)

			clientErrors <- nil
		}(i)
	}

	wg.Wait()
	close(clientErrors)

	for err := range clientErrors {
		require.NoError(t, err)
	}
}

func createTestNode(t *testing.T, config types.NodeConfig) *Node {
	node, err := NewNode(config)
	require.NoError(t, err)

	t.Cleanup(func() {
		node.Host.Close()
	})

	return node
}
