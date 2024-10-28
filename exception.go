package redis

import (
	. "github.com/gopi-frame/contract/exception"
	"github.com/gopi-frame/exception"
)

type ConnectionNotConfiguredException struct {
	connection string
	Throwable
}

func NewConnectionNotConfiguredException(connection string) *ConnectionNotConfiguredException {
	return &ConnectionNotConfiguredException{
		connection: connection,
		Throwable:  exception.New("connection [%s] not configured", connection),
	}
}

func (e *ConnectionNotConfiguredException) Connection() string {
	return e.connection
}
