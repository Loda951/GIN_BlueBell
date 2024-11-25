package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"bluebell/utils"
	"bluebell/validate"
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func PostVoteController(ctx *gin.Context) {
	// 参数校验
	p := new(models.ParamVoteData)
	if err := ctx.ShouldBindJSON(&p); err != nil {
		zap.L().Error("vote with invalid param", zap.Error(err))
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

	userID, err := utils.GetCurrentUserID(ctx)
	if err != nil {
		utils.ResponseError(ctx, utils.CodeNeedLogin)
		return
	}

	if err := logic.PostVote(userID, p); err != nil {
		zap.L().Error("logic.PostVote() failed", zap.Error(err))
		utils.ResponseError(ctx, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(ctx, nil)
}
