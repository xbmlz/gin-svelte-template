package router

import "go.uber.org/fx"

// Module is the router module.
var Module = fx.Options(
	fx.Provide(NewAuthRouter),
	fx.Provide(NewUIRouter),
	fx.Provide(NewRouter),
)

// Router is the interface for router.
type IRoute interface {
	Setup()
}

// Routes is the list of routes.
type Routes []IRoute

// New returns a new router.
func NewRouter(authRouter AuthRouter, uiRouter UIRouter) Routes {
	return Routes{
		authRouter,
		uiRouter,
	}
}

// Setup set all routes.
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
