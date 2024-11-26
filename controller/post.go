package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"bluebell/utils"
	"strconv"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

func CreatPostHandler(ctx *gin.Context) {
	// 1 参数校验 (基础校验)
	p := new(models.Post)
	// 从 HTTP 请求中提取参数值，并将这些值绑定到开发者定义的结构体中
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("create post with invalid param", zap.Error(err))
		utils.ResponseError(ctx, utils.CodeInvalidParam)
		return
	}

	// 从ctx获取authorID
	userID, err := utils.GetCurrentUserID(ctx)
	if err != nil {
		utils.ResponseError(ctx, utils.CodeNeedLogin)
		return
	}
	p.AuthorID = userID

	// 2 业务处理
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost() failed", zap.Error(err))
		utils.ResponseError(ctx, utils.CodeServerBusy)
		return
	}

	// 3 返回响应
	utils.ResponseSuccess(ctx, nil)
}

func GetPostDetailHandler(ctx *gin.Context) {
	// 1 参数校验 (基础校验)
	// 获取参数
	pidStr := ctx.Param("id")
	id, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		utils.ResponseError(ctx, utils.CodeInvalidParam)
		return
	}

	data, err := logic.GetPostDetail(id)
	if err != nil {
		zap.L().Error("logic.GetPostDetail() failed", zap.Error(err))
		utils.ResponseError(ctx, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(ctx, data)
}

func GetPostListHandler(ctx *gin.Context) {
	// 获取分页信息
	page, size, err := utils.GetPageInfo(ctx)
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		utils.ResponseError(ctx, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(ctx, data)
}

func GetPostListHandler2(ctx *gin.Context) {
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	if err := ctx.ShouldBindQuery(p); err != nil {
		zap.L().Error("get post list with invalid param", zap.Error(err))
		utils.ResponseError(ctx, utils.CodeInvalidParam)
		return
	}
	// 获取分页信息
	data, err := logic.GetPostListNew(p)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		utils.ResponseError(ctx, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(ctx, data)
}
