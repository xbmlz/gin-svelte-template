package router

import (
	"github.com/xbmlz/gin-svelte-template/internal/controller"
	"github.com/xbmlz/gin-svelte-template/internal/core"
)

// UserRouter user router
type UserRouter struct {
	log            core.Logger
	srv            core.HTTPServer
	userController controller.UserController
}

// NewUserRouter new user router
func NewUserRouter(log core.Logger, srv core.HTTPServer, userController controller.UserController) UserRouter {
	return UserRouter{log, srv, userController}
}

// Setup setup user router
func (r UserRouter) Setup() {
	r.log.Info("user router setup")
	api := r.srv.RouterV1.Group("user")
	{
		api.POST("", r.userController.Create)
		api.POST("register", r.userController.Register)
		api.POST("login", r.userController.Login)
	}
}
