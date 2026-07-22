package openaifx

import (
	"github.com/go-core-fx/logger"
	"go.uber.org/fx"
)

const ModuleName = "openaifx"

func Module() fx.Option {
	return fx.Module(
		ModuleName,
		logger.WithNamedLogger(ModuleName),
		fx.Provide(New),
	)
}
