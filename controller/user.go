package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(ctx *gin.Context) {
	// 1 参数校验 (基础校验)
	p := new(models.ParamsSignUp)
	// 从 HTTP 请求中提取参数值，并将这些值绑定到开发者定义的结构体中
	if err := ctx.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断error是不是validator的错误 是的话翻译 不是不翻译
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": removeTopStruct(errs.Translate(trans)),
		})
		return

	}

	// 2 业务处理
	if err := logic.SignUp(p); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "注册失败",
		})
		return
	}

	// 3 返回响应
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func LogInHandler(ctx *gin.Context) {
	// 1 参数校验 (基础校验)
	p := new(models.ParamsLogIn)
	// 从 HTTP 请求中提取参数值，并将这些值绑定到开发者定义的结构体中
	if err := ctx.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Log in with invalid param", zap.Error(err))
		// 判断error是不是validator的错误 是的话翻译 不是不翻译
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": removeTopStruct(errs.Translate(trans)),
		})
		return

	}

	// 2 业务处理
	if err := logic.LogIn(p); err != nil {
		zap.L().Error("logic.LogIn failed", zap.String("username", p.Username), zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户名或密码错误",
		})
		return
	}

	// 3 返回响应
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "登陆成功",
	})
}
