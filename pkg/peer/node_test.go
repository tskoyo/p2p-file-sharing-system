package peer

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
