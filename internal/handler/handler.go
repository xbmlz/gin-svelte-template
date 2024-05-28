package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (r Response) JSON(ctx *gin.Context) {
	if r.Message == "" || r.Message == nil {
		r.Message = http.StatusText(r.Code)
	}

	if err, ok := r.Message.(error); ok {
		r.Message = err.Error()
	}
	ctx.JSON(r.Code, r)
	return
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	Response{
		Code:    0,
		Message: "success",
		Data:    data,
	}.JSON(ctx)
}

func ResponseError(ctx *gin.Context, code int, msg interface{}) {
	Response{
		Code:    code,
		Message: msg,
	}.JSON(ctx)
}
