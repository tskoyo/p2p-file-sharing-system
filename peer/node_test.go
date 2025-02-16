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
		ID:   "test-node",
		Port: 9001,
	}

	node, err := NewNode(config)

	assert.NoError(t, err, "Connect should succeed")

	defer node.Host.Close()

	assert.Equal(t, node.Config.ID, config.ID)
	assert.NotNil(t, node.Host, "Host should not be nil")
}

func TestNewNode_WithInvalidPort(t *testing.T) {
	config := types.NodeConfig{
		ID:   "test-node",
		Port: -1,
	}

	_, err := NewNode(config)

	assert.NotNil(t, err, "Expected error for invalid port")
}

func TestNewNode_PortConflict(t *testing.T) {
	config := types.NodeConfig{
		ID:   "test-node",
		Port: 9001,
	}

	node1, err := NewNode(config)

	assert.NoError(t, err, "Connect should succeed")

	defer node1.Host.Close()

	_, err = NewNode(config)

	assert.NoError(t, err, "Connect should succeed")
}

func TestConnect_Success(t *testing.T) {
	node1Config := helper.BuildNodeConfig("peer-1", 9001)
	node2Config := helper.BuildNodeConfig("peer-2", 9002)

	node1, err := NewNode(node1Config)
	require.NoError(t, err)

	node2, err := NewNode(node2Config)
	require.NoError(t, err)

	node1MultiAddr := node1.Host.Addrs()[0].String() + "/p2p/" + node1.Host.ID().String()

	err = node2.Connect(node1MultiAddr)
	require.NoError(t, err)
}

func TestConnect_MultipleClients(t *testing.T) {
	numClients := 20
	var wg sync.WaitGroup
	clientErrors := make(chan error, numClients)

	node1Config := helper.BuildNodeConfig("peer-1", 9001)
	node1, err := NewNode(node1Config)
	require.NoError(t, err)

	node1MultiAddr := node1.Host.Addrs()[0].String() + "/p2p/" + node1.Host.ID().String()

	for i := 1; i <= 8; i++ {
		wg.Add(1)

		go func(clientId int) {
			defer wg.Done()

			clientNodeConfig := helper.BuildNodeConfig(fmt.Sprintf("client-node-%d", clientId), 9001+clientId)
			clientNode, err := NewNode(clientNodeConfig)
			require.NoError(t, err)

			err = clientNode.Connect(node1MultiAddr)
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
