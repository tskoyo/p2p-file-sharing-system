package client

import (
	"net"

	"github.com/stretchr/testify/mock"
)

type MockDialer struct {
	mock.Mock
}

type MockConn struct {
	net.Conn
}

func NewMockConn() (*MockConn, net.Conn) {
	conn1, serverConn := net.Pipe()

	mockConn := &MockConn{
		Conn: conn1,
	}

	return mockConn, serverConn
}

func (m *MockDialer) Dial(network NetworkType, peerAddress string) (net.Conn, error) {
	args := m.Called(network, peerAddress)
	conn := args.Get(0)
	err := args.Error(1)

	if conn != nil {
		return conn.(net.Conn), err
	}

	return nil, err
}
