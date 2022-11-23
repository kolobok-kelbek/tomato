//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/kolobok-kelbek/go-example-service/infra"
	"github.com/kolobok-kelbek/go-example-service/infra/config"
	"github.com/kolobok-kelbek/go-example-service/infra/logger"
	"github.com/kolobok-kelbek/go-example-service/static"
	"go.uber.org/zap"
)

func Init() *infra.App {
	wire.Build(
		config.Load,
		wire.Value(static.Snapshot),
		logger.NewLogger,
		wire.Bind(new(logger.Logger), new(*zap.SugaredLogger)),
		infra.NewApp,
	)

	return &infra.App{}
}
