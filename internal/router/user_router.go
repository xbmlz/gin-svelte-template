package router

import (
	"github.com/xbmlz/gin-svelte-template/internal/controller"
	"github.com/xbmlz/gin-svelte-template/internal/core"
)

// AuthRouter user router
type AuthRouter struct {
	log            core.Logger
	srv            core.HTTPServer
	authController controller.AuthController
}

// NewAuthRouter new user router
func NewAuthRouter(log core.Logger, srv core.HTTPServer, userController controller.AuthController) AuthRouter {
	return AuthRouter{log, srv, userController}
}

// Setup setup user router
func (r AuthRouter) Setup() {
	api := r.srv.RouterV1
	{
		api.POST("register", r.authController.Register)
		api.POST("login", r.authController.Login)
		api.POST("logout", r.authController.Logout)
	}
}
