package bootstrap

import (
	"context"

	"github.com/xbmlz/gin-svelte-template/pkg/logger"
	"go.uber.org/fx"
)

func registerHooks(lc fx.Lifecycle, log logger.Logger) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				log.Info("on start")
				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Info("on stop")
				return nil
			},
		},
	)
}

var Modules = fx.Options(
	// modules export
	fx.Provide(
		logger.NewLogger(logger.DebugLevel),
	),

	fx.Invoke(registerHooks),
)
