package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/gin-svelte-template/internal/handler"
	"github.com/xbmlz/gin-svelte-template/internal/model"
	"github.com/xbmlz/gin-svelte-template/internal/service"
)

type AuthController struct {
	authService service.AuthService
	userService service.UserService
}

func NewAuthController(authService service.AuthService, userService service.UserService) AuthController {
	return AuthController{
		authService,
		userService,
	}
}

// func (c *AuthController) Create(ctx *gin.Context) {
// 	req := &model.User{}
// 	if err := ctx.ShouldBindJSON(req); err != nil {
// 		handler.Response{Code: http.StatusBadRequest}.JSON(ctx)
// 		return
// 	}

// 	err := c.userService.Create(req)

// 	if err != nil {
// 		handler.Response{Code: http.StatusInternalServerError}.JSON(ctx)
// 		return
// 	}

// 	handler.Response{Code: http.StatusOK}.JSON(ctx)
// }

func (c *AuthController) Register(ctx *gin.Context) {
	req := &model.UserRegister{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		handler.Response{Code: http.StatusBadRequest}.JSON(ctx)
		return
	}

	saveUser := &model.User{}
	saveUser.Username = req.Username
	saveUser.Password = req.Password

	err := c.userService.Create(saveUser)

	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, nil)
}

func (c *AuthController) Login(ctx *gin.Context) {
	req := &model.UserLogin{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		handler.Response{Code: http.StatusBadRequest}.JSON(ctx)
		return
	}

	user, err := c.userService.Verify(req.Username, req.Password)

	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := c.authService.GenerateToken(user)

	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, gin.H{
		"token": token,
	})
}
