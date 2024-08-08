package cluster

import (
	"github.com/go-viper/mapstructure/v2"
	rediscontract "github.com/gopi-frame/contract/redis"
	"github.com/gopi-frame/redis"
	redislib "github.com/redis/go-redis/v9"
)

// This variable can be replaced through `go build -ldflags=-X github.com/gopi-frame/redis/driver/cluster.driverName=custom`
var driverName = "cluster"

//goland:noinspection GoBoolExpressions
func init() {
	if driverName != "" {
		redis.Register(driverName, new(Driver))
	}
}

// Driver provides redis cluster driver
type Driver struct{}

// Open opens redis cluster.
// For more information about options, see https://pkg.go.dev/github.com/redis/go-redis/v9#ClusterOptions
func (d Driver) Open(options map[string]any) (rediscontract.Client, error) {
	cfg := new(redislib.ClusterOptions)
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook:       mapstructure.StringToTimeDurationHookFunc(),
		WeaklyTypedInput: true,
		Result:           cfg,
	})
	if err != nil {
		return nil, err
	}
	if err := decoder.Decode(options); err != nil {
		return nil, err
	}
	return redislib.NewClusterClient(cfg), nil
}

// Open is a convenience function that calls [Driver.Open].
func Open(options map[string]any) (rediscontract.Client, error) {
	return new(Driver).Open(options)
}
