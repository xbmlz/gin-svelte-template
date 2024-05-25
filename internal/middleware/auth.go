package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/gin-svelte-template/internal/constant"
	"github.com/xbmlz/gin-svelte-template/internal/core"
	"github.com/xbmlz/gin-svelte-template/internal/handler"
	"github.com/xbmlz/gin-svelte-template/internal/service"
)

var _ IMiddleware = (*AuthMiddleware)(nil)

type AuthMiddleware struct {
	conf        core.Config
	handle      core.HTTPServer
	log         core.Logger
	authService service.AuthService
}

func NewAuthMiddleware(conf core.Config, handle core.HTTPServer, log core.Logger, authService service.AuthService) AuthMiddleware {
	return AuthMiddleware{
		conf:        conf,
		handle:      handle,
		log:         log,
		authService: authService,
	}
}

func (am AuthMiddleware) core() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ExtractToken(ctx)
		claims, err := am.authService.ParseToken(tokenString)
		if err != nil {
			handler.Response{Code: http.StatusUnauthorized, Message: err}.JSON(ctx)
			ctx.Abort()
		}
		ctx.Set(constant.CurrentUser, claims)
		ctx.Next()
	}
}

// ExtractToken extracts the token from the request Header.
func ExtractToken(ctx *gin.Context) (token string) {
	token = ctx.GetHeader("Authorization")
	if len(token) == 0 {
		token = ctx.Query("Authorization")
	}
	return strings.TrimPrefix(token, "Bearer ")
}

func (am AuthMiddleware) Setup() {
	am.handle.Engine.Use(am.core())
	am.log.Info("AuthMiddleware setup")
}
