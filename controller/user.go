package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(ctx *gin.Context) {
	// 1 参数校验 (基础校验)
	p := new(models.ParamsSignUp)
	// 从 HTTP 请求中提取参数值，并将这些值绑定到开发者定义的结构体中
	if err := ctx.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数缺失!",
		})
		return
	}
	// 2 手动对请求参数进行详细的业务规则校验 永远不要相信前端校验 (业务逻辑校验)
	if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePasswd) == 0 || p.RePasswd != p.Password {
		zap.L().Error("SignUp with invalid param")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数有误!",
		})
		return
	}

	// 2 业务处理
	logic.SignUp(p)
	// 3 返回响应
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
