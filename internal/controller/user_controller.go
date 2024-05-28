package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/gin-svelte-template/internal/handler"
	"github.com/xbmlz/gin-svelte-template/internal/model"
	"github.com/xbmlz/gin-svelte-template/internal/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{
		userService,
	}
}

func (c *UserController) Create(ctx *gin.Context) {
	req := &model.User{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		handler.Response{Code: http.StatusBadRequest}.JSON(ctx)
		return
	}

	err := c.userService.Create(req)

	if err != nil {
		handler.Response{Code: http.StatusInternalServerError}.JSON(ctx)
		return
	}

	handler.Response{Code: http.StatusOK}.JSON(ctx)
}

func (c *UserController) Register(ctx *gin.Context) {
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
		handler.Response{Code: http.StatusInternalServerError, Message: err.Error()}.JSON(ctx)
		return
	}

	handler.ResponseSuccess(ctx, nil)
}

func (c *UserController) Login(ctx *gin.Context) {
	req := &model.UserLogin{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		handler.Response{Code: http.StatusBadRequest}.JSON(ctx)
		return
	}

	user, err := c.userService.Login(req.Username, req.Password)

	if err != nil {
		handler.Response{Code: http.StatusInternalServerError}.JSON(ctx)
		return
	}

	handler.ResponseSuccess(ctx, user)
}
