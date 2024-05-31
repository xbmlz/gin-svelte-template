package bootstrap

import (
	"context"

	"github.com/xbmlz/gin-svelte-template/internal/controller"
	"github.com/xbmlz/gin-svelte-template/internal/core"
	"github.com/xbmlz/gin-svelte-template/internal/middleware"
	"github.com/xbmlz/gin-svelte-template/internal/repo"
	"github.com/xbmlz/gin-svelte-template/internal/router"
	"github.com/xbmlz/gin-svelte-template/internal/service"

	"go.uber.org/fx"
)

var Module = fx.Options(
	controller.Module,
	router.Module,
	core.Module,
	service.Module,
	middleware.Module,
	repo.Module,
	// invoke
	fx.Invoke(bootstrap),
)

func bootstrap(
	lc fx.Lifecycle,
	conf core.Config,
	log core.Logger,
	httpSrv core.HTTPServer,
	routes router.Routes,
	middlewares middleware.Middlewares,
	database core.Database,
) {
	db, err := database.DB.DB()
	if err != nil {
		log.Errorf("Failed to connect to database: %v", err)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Starting app...")

			go func() {
				middlewares.Setup()
				routes.Setup()

				if err := httpSrv.Start(); err != nil {
					log.Errorf("Failed to start http server: %v", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Stopping app...")

			err := db.Close()
			if err != nil {
				return err
			}

			err = httpSrv.Shutdown()
			if err != nil {
				log.Errorf("Failed to shutdown http server: %v", err)
			}
			return nil
		},
	})
}
