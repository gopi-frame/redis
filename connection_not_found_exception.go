package redis

import (
	"fmt"

	"github.com/gopi-frame/exception"
)

// ConnectionNotFoundException connection not found exception
type ConnectionNotFoundException struct {
	*exception.Exception
}

// NewConnectionNotFoundException new connection not found exception
func NewConnectionNotFoundException(connection string) *ConnectionNotFoundException {
	return &ConnectionNotFoundException{
		Exception: exception.NewException(fmt.Sprintf("connection \"%s\" not found", connection)),
	}
}
