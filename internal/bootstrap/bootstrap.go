package bootstrap

import (
	"context"
	"fmt"
	"github.com/xbmlz/gin-svelte-template/internal/logger"

	"github.com/xbmlz/gin-svelte-template/internal/config"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(config.LoadConfig),
	fx.Provide(logger.NewLogger),
	// invoke
	fx.Invoke(bootstrap),
)

func bootstrap(lc fx.Lifecycle, config config.Config, log logger.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// do something

			return nil
		},
		OnStop: func(ctx context.Context) error {
			// do something
			fmt.Println("app stop")
			return nil
		},
	})
}
