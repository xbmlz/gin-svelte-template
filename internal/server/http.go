package server

import (
	"context"
	"errors"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xbmlz/gin-svelte-template/docs"
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	"github.com/xbmlz/gin-svelte-template/internal/config"
	"github.com/xbmlz/gin-svelte-template/internal/logger"
)

const DefaultShutdownTimeout = time.Minute

var _ IServer = (*HTTPServer)(nil)

// HTTPServer Server is gin implementation.
type HTTPServer struct {
	srv *http.Server
}

func NewHTTPServer(log logger.Logger, conf config.Config) HTTPServer {
	zapLogger := log.GetZapLogger()
	// new engine
	engine := gin.New()

	engine.Use(ginzap.Ginzap(zapLogger, time.DateTime, true))

	engine.Use(ginzap.RecoveryWithZap(zapLogger, true))

	// swagger doc
	docs.SwaggerInfo.Host = conf.HTTP.ListenAddr()
	docs.SwaggerInfo.BasePath = "/v1"
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
	))

	engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	engine.Use(gin.Recovery())

	server := HTTPServer{
		srv: &http.Server{
			Addr:    conf.HTTP.ListenAddr(),
			Handler: engine,
		},
	}
	return server
}

// Start to start the server and wait for it to listen on the given address
func (s *HTTPServer) Start() (err error) {
	err = s.srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Shutdown shuts down the server and close with graceful shutdown duration
func (s *HTTPServer) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultShutdownTimeout)
	defer cancel()
	return s.srv.Shutdown(ctx)
}
