package controller

import (
	"bluebell/logic"
	"bluebell/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(ctx *gin.Context) {
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		utils.ResponseError(ctx, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(ctx, data)
}

func CommunityDetailHandler(ctx *gin.Context) {
	// 获取参数
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("get community detail with invalid param", zap.Error(err))
		utils.ResponseError(ctx, utils.CodeInvalidParam)
		return
	}

	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail() failed", zap.Error(err))
		utils.ResponseError(ctx, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(ctx, data)
}
