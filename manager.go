package redis

import (
	"sync"

	"github.com/gopi-frame/collection/kv"
	"github.com/gopi-frame/contract/redis"
)

type RedisManager struct {
	once sync.Once
	redis.Client

	defaultConnection string
	connections       *kv.Map[string, redis.Client]
	deferConnections  *kv.Map[string, func() (redis.Client, error)]
}

func NewRedisManager() *RedisManager {
	return &RedisManager{
		connections: kv.NewMap[string, redis.Client](),
	}
}

func (m *RedisManager) SetDefaultConnection(name string) {
	m.defaultConnection = name
	m.once.Do(func() {
		m.Client = m.GetConnection(name)
	})
}

func (m *RedisManager) AddConnection(name string, client redis.Client) {
	m.connections.Lock()
	defer m.connections.Unlock()
	m.connections.Set(name, client)
}

func (m *RedisManager) AddDeferConnection(name string, config map[string]any) {
	m.connections.Lock()
	defer m.connections.Unlock()
	m.deferConnections.Set(name, func() (redis.Client, error) {
		driver := config["driver"].(string)
		return Open(driver, config)
	})
}

func (m *RedisManager) HasConnection(name string) bool {
	m.connections.RLock()
	if m.connections.ContainsKey(name) {
		m.connections.RUnlock()
		return true
	}
	m.connections.RUnlock()
	m.deferConnections.RLock()
	if m.deferConnections.ContainsKey(name) {
		m.deferConnections.RUnlock()
		return true
	}
	m.deferConnections.RUnlock()
	return false
}

func (m *RedisManager) TryGetConnection(name string) (redis.Client, error) {
	m.connections.RLock()
	if conn, ok := m.connections.Get(name); ok {
		m.connections.RUnlock()
		return conn, nil
	}
	m.connections.RUnlock()
	m.deferConnections.RLock()
	if conn, ok := m.deferConnections.Get(name); ok {
		m.deferConnections.RUnlock()
		if conn, err := conn(); err != nil {
			return nil, err
		} else {
			m.connections.Lock()
			defer m.connections.Unlock()
			m.connections.Set(name, conn)
			return conn, nil
		}
	}
	m.deferConnections.RUnlock()
	return nil, NewConnectionNotConfiguredException(name)
}

func (m *RedisManager) GetConnection(name string) redis.Client {
	if conn, err := m.TryGetConnection(name); err != nil {
		panic(err)
	} else {
		return conn
	}
}

func (m *RedisManager) ConnectionOrDefault(name string) redis.Client {
	if conn, err := m.TryGetConnection(name); err != nil {
		return m.GetConnection(m.defaultConnection)
	} else {
		return conn
	}
}
