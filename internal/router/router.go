package router

import "go.uber.org/fx"

// Module is the router module.
var Module = fx.Options(
	fx.Provide(NewUserRouter),
	fx.Provide(NewRouter),
)

// Router is the interface for router.
type IRoute interface {
	Setup()
}

// Routes is the list of routes.
type Routes []IRoute

// New returns a new router.
func NewRouter(userRouter UserRouter) Routes {
	return Routes{
		userRouter,
	}
}

// Setup set all routes.
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
