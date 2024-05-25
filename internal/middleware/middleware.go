package middleware

import "go.uber.org/fx"

// Module Middleware exported
var Module = fx.Options(
	fx.Provide(NewAuthMiddleware),
	fx.Provide(NewMiddlewares),
)

// IMiddleware interface
type IMiddleware interface {
	Setup()
}

// Middlewares contains multiple middleware
type Middlewares []IMiddleware

// NewMiddlewares creates a new middlewares
func NewMiddlewares(authMiddleware AuthMiddleware) Middlewares {
	return Middlewares{
		authMiddleware,
	}
}

// Setup setup middlewares
func (m Middlewares) Setup() {
	for _, middleware := range m {
		middleware.Setup()
	}
}
