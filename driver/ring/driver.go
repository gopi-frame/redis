package ring

import (
	"github.com/go-viper/mapstructure/v2"
	rediscontract "github.com/gopi-frame/contract/redis"
	"github.com/gopi-frame/redis"
	redislib "github.com/redis/go-redis/v9"
)

// This variable can be replaced through `go build -ldflags=-X github.com/gopi-frame/redis/driver/ring.driverName=custom`
var driverName = "ring"

//goland:noinspection GoBoolExpressions
func init() {
	if driverName != "" {
		redis.Register(driverName, new(Driver))
	}
}

// Driver provides redis ring driver
type Driver struct{}

// Open opens redis ring client.
//
// For more information about options, see https://pkg.go.dev/github.com/redis/go-redis/v9#RingOptions
func (d Driver) Open(options map[string]any) (rediscontract.Client, error) {
	cfg := new(redislib.RingOptions)
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
	return redislib.NewRing(cfg), nil
}

// Open is a convenience function that calls [Driver.Open].
func Open(options map[string]any) (rediscontract.Client, error) {
	return new(Driver).Open(options)
}
