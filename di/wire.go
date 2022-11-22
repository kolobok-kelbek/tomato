//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/kolobok-kelbek/go-example-service/infra"
	"github.com/kolobok-kelbek/go-example-service/infra/config"
	"github.com/kolobok-kelbek/go-example-service/static"
)

func Init() *infra.App {
	wire.Build(
		config.Load,
		wire.Value(static.Snapshot),
		infra.NewApp,
	)

	return &infra.App{}
}
