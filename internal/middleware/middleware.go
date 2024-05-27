package middleware

import "go.uber.org/fx"

// Module Middleware exported
var Module = fx.Options(
	fx.Provide(NewCorsMiddleware),
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
func NewMiddlewares(authMiddleware AuthMiddleware, corsMiddleware CorsMiddleware) Middlewares {
	return Middlewares{
		authMiddleware,
		corsMiddleware,
	}
}

// Setup setup middlewares
func (m Middlewares) Setup() {
	for _, middleware := range m {
		middleware.Setup()
	}
}

func isIgnorePath(path string, prefixes ...string) bool {
	pathLen := len(path)

	for _, p := range prefixes {
		if pl := len(p); pathLen >= pl && path[:pl] == p {
			return true
		}
	}

	return false
}
