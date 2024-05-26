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

// @tags User
// @summary User Create
// @produce application/json
// @param data body model.User true "User"
// @success 200 {object} handler.Response "ok"
// @failure 400 {object} handler.Response "bad request"
// @failure 500 {object} handler.Response "internal error"
// @router /api/users [post]
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
