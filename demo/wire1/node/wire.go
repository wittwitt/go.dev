//go:build wireinject
// +build wireinject

package node

import (
	"github.com/google/wire"

	"github.com/wittwitt/go.dev/demo/wire1/pool"
)

func IniPool() (*pool.Pool, error) {
	//	wire.Build(NewRedisConfig, NewPoolStore, pool.NewPool)
	wire.Build(db, pool.NewPool)
	return &pool.Pool{}, nil
}
