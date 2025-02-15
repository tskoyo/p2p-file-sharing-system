package peer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewNode_Succes(t *testing.T) {
	config := NodeConfig{
		ID:   "test-node",
		Port: 9001,
	}

	node, err := NewNode(config)

	assert.NoError(t, err, "Connect should succeed")

	defer node.Host.Close()

	assert.Equal(t, node.ID, config.ID)
	assert.NotNil(t, node.Host, "Host should not be nil")
}

func TestNewNode_WithInvalidPort(t *testing.T) {
	config := NodeConfig{
		ID:   "test-node",
		Port: -1,
	}

	_, err := NewNode(config)

	assert.NotNil(t, err, "Expected error for invalid port")
}

func TestNewNode_PortConflict(t *testing.T) {
	config := NodeConfig{
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
	node1PortId := 9001
	node2PortId := 9003

	node1Config := NodeConfig{
		ID:   "peer-1",
		Port: node1PortId,
	}
	node2Config := NodeConfig{
		ID:   "peer-2",
		Port: node2PortId,
	}

	node1, err := NewNode(node1Config)
	require.NoError(t, err)

	node2, err := NewNode(node2Config)
	require.NoError(t, err)

	node1MultiAddr := node1.Host.Addrs()[0].String() + "/p2p/" + node1.Host.ID().String()
	require.NoError(t, err)

	err = node2.Connect(node1MultiAddr)
	require.NoError(t, err)

	connections := node1.Host.Network().ConnsToPeer(node2.Host.ID())

	for _, v := range connections {
		t.Logf("address id: %t", v.ID() == node2.Host.ID().String())
	}
}
