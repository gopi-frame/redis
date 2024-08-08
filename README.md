# Overview
[![Go Reference](https://pkg.go.dev/badge/github.com/gopi-frame/redis.svg)](https://pkg.go.dev/github.com/gopi-frame/redis)
[![Go report card](https://goreportcard.com/badge/github.com/gopi-frame/redis)](https://goreportcard.com/report/github.com/gopi-frame/redis)
[![Mit License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

Package redis provides a redis connection client.

This package is based on [go-redis](https://github.com/go-redis/redis)

## Installation

```shell
go get -u github.com/gopi-frame/redis
```

## Import

```go
import "github.com/gopi-frame/redis"
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "github.com/gopi-frame/redis"
    "time"
)

func main() {
    client, err := redis.Open("standalone", nil)
    if err != nil {
        panic(err)
    }
    client.Set(context.Background(), "key", "value", 10*time.Second)
    value := client.Get(context.Background(), "key")
    fmt.Println(value)
}
```

## Redis Cluster Client

```go
package main

import (
    "context"
    "fmt"
    "github.com/gopi-frame/redis"
    "time"
)

func main() {
    client, err := redis.Open("cluster", map[string]any{
        "addrs": []string{
            ":6379",
            ":6380",
            ":6381",
        },
    })
    if err != nil {
        panic(err)
    }
    if err := client.Set(context.Background(), "key", "value", 10*time.Second).Err(); err != nil {
        panic(err)
    }
    value := client.Get(context.Background(), "key")
    fmt.Println(value)
}
```

## Redis Sentinel Client

```go
package main

import (
    "context"
    "fmt"
    "github.com/gopi-frame/redis"
    "time"
)

func main() {
    client, err := redis.Open("sentinel", map[string]any{
        "masterName":    "master",
        "sentinelAddrs": []string{":6379", ":6380", ":6381"},
    })
    if err != nil {
        panic(err)
    }
    if err := client.Set(context.Background(), "key", "value", 10*time.Second).Err(); err != nil {
        panic(err)
    }
    value := client.Get(context.Background(), "key")
    fmt.Println(value)
}
```

## Redis Ring Client

```go
package main

import (
    "context"
    "fmt"
    "github.com/gopi-frame/redis"
    "time"
)

func main() {
    client, err := redis.Open("ring", map[string]any{
        "addrs": map[string]string{
            "shard1": ":7000",
            "shard2": ":7001",
            "shard3": ":7002",
        },
    })
    if err != nil {
        panic(err)
    }
    if err := client.Set(context.Background(), "key", "value", 10*time.Second).Err(); err != nil {
        panic(err)
    }
    value := client.Get(context.Background(), "key")
    fmt.Println(value)
}
```

## Options
This package uses [github.com/go-viper/mapstructure/v2](https://github.com/go-viper/mapstructure) to convert the options to the underlying
client.

For more information about the options, please see the [go-redis](https://github.com/go-redis/redis) documentation.
