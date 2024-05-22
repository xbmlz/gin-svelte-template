package bootstrap

import (
	"context"
	"fmt"

	"github.com/xbmlz/gin-svelte-template/internal/module"
	"go.uber.org/fx"
)

var Options = fx.Options(
	module.Options,

	// invoke
	fx.Invoke(bootstrap),
)

func bootstrap(lc fx.Lifecycle, config module.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// do something
			fmt.Println("app start")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// do something
			fmt.Println("app stop")
			return nil
		},
	})
}
