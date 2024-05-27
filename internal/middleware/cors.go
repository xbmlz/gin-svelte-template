package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/xbmlz/gin-svelte-template/internal/core"
)

// CorsMiddleware configures CORS middleware.
type CorsMiddleware struct {
	log    core.Logger
	handle core.HTTPServer
}

// NewCorsMiddleware creates new cors middleware
func NewCorsMiddleware(log core.Logger, handle core.HTTPServer) CorsMiddleware {
	return CorsMiddleware{
		log,
		handle,
	}
}

func (a CorsMiddleware) Setup() {
	cors := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
	a.handle.Engine.Use(cors)
	a.log.Info("Cors middleware is setup")
}
