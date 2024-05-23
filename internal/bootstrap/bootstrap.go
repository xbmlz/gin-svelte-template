package bootstrap

import (
	"context"
	"github.com/xbmlz/gin-svelte-template/internal/dal"
	"github.com/xbmlz/gin-svelte-template/internal/logger"
	"github.com/xbmlz/gin-svelte-template/internal/server"

	"github.com/xbmlz/gin-svelte-template/internal/config"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		config.NewConfig,
		logger.NewLogger,
		server.NewHTTPServer,
		dal.NewDatabase,
	),
	// invoke
	fx.Invoke(bootstrap),
)

func bootstrap(
	lc fx.Lifecycle,
	config config.Config,
	log logger.Logger,
	httpSrv server.HTTPServer,
) {

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Starting app...")

			go func() {
				if err := httpSrv.Start(); err != nil {
					log.Errorf("Failed to start http server: %v", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Stopping app...")

			err := httpSrv.Shutdown()
			if err != nil {
				log.Errorf("Failed to shutdown http server: %v", err)
			}
			return nil
		},
	})
}
