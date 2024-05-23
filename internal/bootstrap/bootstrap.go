package bootstrap

import (
	"context"

	"github.com/xbmlz/gin-svelte-template/internal/database"
	"github.com/xbmlz/gin-svelte-template/internal/logger"
	"github.com/xbmlz/gin-svelte-template/internal/server"
	"github.com/xbmlz/gin-svelte-template/internal/service"

	"github.com/xbmlz/gin-svelte-template/internal/config"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		config.NewConfig,
		logger.NewLogger,
		server.NewHTTPServer,
		database.NewDatabase,
	),
	service.Module,
	// invoke
	fx.Invoke(bootstrap),
)

func bootstrap(
	lc fx.Lifecycle,
	config config.Config,
	log logger.Logger,
	httpSrv server.HTTPServer,
	database database.Database,
) {
	db, err := database.DB.DB()
	if err != nil {
		log.Errorf("Failed to connect to database: %v", err)
	}

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

			db.Close()

			err := httpSrv.Shutdown()
			if err != nil {
				log.Errorf("Failed to shutdown http server: %v", err)
			}
			return nil
		},
	})
}
