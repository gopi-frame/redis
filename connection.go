package redis

import (
	"github.com/redis/go-redis/v9"
)

// NewConnection new connection
func NewConnection(client redis.UniversalClient) *Connection {
	return &Connection{
		UniversalClient: client,
	}
}

// NewLazyConnection new lazy connection
func NewLazyConnection(connector func() redis.UniversalClient) *Connection {
	return &Connection{
		connector: connector,
	}
}

// Resolve create connection from config
func Resolve(config ConnectionConfig) *Connection {
	return NewLazyConnection(func() redis.UniversalClient {
		return redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:            config.Addresses,
			MasterName:       config.MasterName,
			ClientName:       config.ClientName,
			DB:               config.Database,
			Username:         config.Username,
			Password:         config.Password,
			SentinelUsername: config.SentinelUsername,
			SentinelPassword: config.SentinelPassword,
		})
	})
}

// Connection connection
type Connection struct {
	redis.UniversalClient
	connector func() redis.UniversalClient
}

// Connect connect
func (c *Connection) Connect() {
	if c.connector != nil && c.UniversalClient == nil {
		c.UniversalClient = c.connector()
	}
}
