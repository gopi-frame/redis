// Package redis provides redis client.
// This package is a wrapper around [go-redis].
//
// [go-redis]: https://github.com/redis/go-redis
package redis

import (
	"fmt"

	"github.com/gopi-frame/collection/kv"
	rediscontract "github.com/gopi-frame/contract/redis"
	"github.com/gopi-frame/exception"
)

var drivers = kv.NewMap[string, rediscontract.Driver]()

// Register register driver
func Register(name string, driver rediscontract.Driver) {
	drivers.Lock()
	defer drivers.Unlock()
	if driver == nil {
		panic(exception.NewEmptyArgumentException("driver"))
	}
	if drivers.ContainsKey(name) {
		panic(exception.NewArgumentException("name", name, fmt.Sprintf("duplicate driver \"%s\"", name)))
	}
	drivers.Set(name, driver)
}

// Drivers list registered drivers
func Drivers() []string {
	drivers.RLock()
	defer drivers.RUnlock()
	list := drivers.Keys()
	return list
}

// Open opens redis client
func Open(name string, options map[string]any) (rediscontract.Client, error) {
	drivers.RLock()
	driver, ok := drivers.Get(name)
	drivers.RUnlock()
	if !ok {
		return nil, exception.NewArgumentException("name", name, fmt.Sprintf("unknown driver \"%s\"", name))
	}
	return driver.Open(options)
}
