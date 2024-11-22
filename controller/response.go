package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	response := &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
	ctx.JSON(http.StatusOK, response)
}

// ResponseError 指定了错误代码的error返回
func ResponseError(ctx *gin.Context, code ResCode) {
	response := &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	ctx.JSON(http.StatusBadRequest, response)
}

// ResponseErrorWithMsg 没有指定错误代码的error返回
func ResponseErrorWithMsg(ctx *gin.Context, code ResCode, msg interface{}) {
	response := &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	ctx.JSON(http.StatusBadRequest, response)
}
