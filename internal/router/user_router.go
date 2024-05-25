package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/gin-svelte-template/internal/core"
)

// UserRouter user router
type UserRouter struct {
	log core.Logger
	srv core.HTTPServer
}

// NewUserRouter new user router
func NewUserRouter(log core.Logger, srv core.HTTPServer) UserRouter {
	return UserRouter{log, srv}
}

// Setup setup user router
func (r UserRouter) Setup() {
	r.log.Info("user router setup")
	api := r.srv.Engine.Group("/users")
	{
		api.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "user info",
			})
		})
	}
}
