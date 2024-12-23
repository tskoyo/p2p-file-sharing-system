package peer

import (
	"net"
	"sync"
)

type ConnectionMap map[string]net.Conn

type ConnectionPool struct {
	sync.Mutex
	pool ConnectionMap
}

func NewConnectionPool() *ConnectionPool {
	return &ConnectionPool{
		pool: make(ConnectionMap),
	}
}

func (cp *ConnectionPool) Add(address string, conn net.Conn) {
	cp.Lock()
	defer cp.Unlock()
	cp.pool[address] = conn
}

func (cp *ConnectionPool) Remove(address string) {
	cp.Lock()
	defer cp.Unlock()
	delete(cp.pool, address)
}

func (cp *ConnectionPool) Get(address string) (net.Conn, bool) {
	cp.Lock()
	defer cp.Unlock()
	conn, exists := cp.pool[address]
	return conn, exists
}

func (cp *ConnectionPool) List() ConnectionMap {
	cp.Lock()
	defer cp.Unlock()

	copy := make(map[string]net.Conn, len(cp.pool))
	for address, conn := range cp.pool {
		copy[address] = conn
	}
	return copy
}
