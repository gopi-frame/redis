package redis

import (
	"github.com/gopi-frame/support/maps"
	"github.com/redis/go-redis/v9"
)

// Manager manager
type Manager struct {
	defaultConnection string
	connections       *maps.Map[string, *Connection]
}

// SetDefaultConnection set default connection
func (m *Manager) SetDefaultConnection(name string) {
	m.defaultConnection = name
}

// DB get connection instance
func (m *Manager) DB(name string) *Connection {
	connection, ok := m.connections.Get(name)
	if ok {
		connection.Connect()
		return connection
	}
	panic(NewConnectionNotFoundException(name))
}

// AddConnection add a new connection
func (m *Manager) AddConnection(name string, client redis.UniversalClient) {
	m.connections.Set(name, NewConnection(client))
}

// AddLazyConnection add a new lazy connection
func (m *Manager) AddLazyConnection(name string, connector func() redis.UniversalClient) {
	m.connections.Set(name, NewLazyConnection(connector))
}
