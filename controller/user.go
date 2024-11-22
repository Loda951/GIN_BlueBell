package controller

import (
	"bluebell/DAO/mysql"
	"bluebell/logic"
	"bluebell/models"
	"bluebell/utils"
	"bluebell/validate"
	"errors"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(ctx *gin.Context) {
	// 1 参数校验 (基础校验)
	p := new(models.ParamsSignUp)
	// 从 HTTP 请求中提取参数值，并将这些值绑定到开发者定义的结构体中
	if err := ctx.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Sign up with invalid param", zap.Error(err))
		// 判断error是不是validator的错误
		var validatorErr validator.ValidationErrors
		// 是否属于validatorErr
		ok := errors.As(err, &validatorErr)
		// 不是validator的错误类型返回原error
		if !ok {
			utils.ResponseError(ctx, utils.CodeInvalidParam)
			return
		}
		// 是validator的错误类型 翻译错误并返回
		utils.ResponseErrorWithMsg(ctx, utils.CodeInvalidParam, validate.RemoveTopStruct(validatorErr.Translate(validate.Trans)))
		return
	}

	// 2 业务处理
	if err := logic.SignUp(p); err != nil {
		if errors.Is(err, mysql.ErrorUserExists) {
			utils.ResponseError(ctx, utils.CodeUserExist)
			return
		}
		utils.ResponseError(ctx, utils.CodeServerBusy)
		return
	}

	// 3 返回响应
	utils.ResponseSuccess(ctx, nil)
}

func LogInHandler(ctx *gin.Context) {
	// 1 参数校验 (基础校验)
	p := new(models.ParamsLogIn)
	// 从 HTTP 请求中提取参数值，并将这些值绑定到开发者定义的结构体中
	if err := ctx.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Log in with invalid param", zap.Error(err))
		// 判断error是不是validator的错误
		var validatorErr validator.ValidationErrors
		// 是否属于validatorErr
		ok := errors.As(err, &validatorErr)
		// 不是validator的错误类型返回原error
		if !ok {
			utils.ResponseError(ctx, utils.CodeInvalidParam)
			return
		}
		// 是validator的错误类型 翻译错误并返回
		utils.ResponseErrorWithMsg(ctx, utils.CodeInvalidParam, validate.RemoveTopStruct(validatorErr.Translate(validate.Trans)))
		return
	}

	// 2 业务处理
	token, err := logic.LogIn(p)
	if err != nil {
		zap.L().Error("logic.LogIn failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExists) {
			utils.ResponseError(ctx, utils.CodeUserExist)
			return
		}
		utils.ResponseError(ctx, utils.CodeInvalidPassword)
		return
	}

	// 3 返回响应
	utils.ResponseSuccess(ctx, token)
}
