package node

import (
	"github.com/google/wire"

	"github.com/wittwitt/go.dev/demo/wire1/pool"
	"github.com/wittwitt/go.dev/demo/wire1/redis"
)

//var db = wire.NewSet(NewRedisConfig, NewPoolStore)

var db = wire.NewSet(
	NewRedisConfig,
	redis.NewClient,
	wire.Bind(new(pool.Store), new(*redis.Client)),
)

func NewRedisConfig() *redis.Config {
	return &redis.Config{}
}

// use wire.Bind
//func NewPoolStore(cfg *redis.Config) (pool.Store, error) {
//	return redis.NewClient(cfg)
//}
