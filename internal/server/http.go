package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/gin-svelte-template/internal/config"
	"github.com/xbmlz/gin-svelte-template/internal/logger"
)

const DefaultShutdownTimeout = time.Minute

var _ IServer = (*HTTPServer)(nil)

// Server is gin implementation.
type HTTPServer struct {
	srv *http.Server
}

func NewHTTPServer(log logger.Logger, conf config.Config) HTTPServer {
	logger := log.GetZapLogger()
	// new engine
	engine := gin.New()

	engine.Use(ginzap.Ginzap(logger, time.DateTime, true))

	engine.Use(ginzap.RecoveryWithZap(logger, true))

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
