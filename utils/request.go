package utils

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	CtxUserIDKey   = "userID"
	CtxUserNameKey = "username"
)

var ErrorNotLogin = errors.New("not login")

func GetCurrentUserID(ctx *gin.Context) (userID int64, err error) {
	// uid类型是 interface{}
	uid, ok := ctx.Get(CtxUserIDKey)
	if !ok {
		err = ErrorNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorNotLogin
		return
	}
	return userID, nil
}

func GetPageInfo(ctx *gin.Context) (int64, int64, error) {
	// 获取分页参数
	pageStr := ctx.Query("page")
	sizeStr := ctx.Query("size")

	var (
		page, size int64
		err        error
	)

	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}

	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}

	return page, size, nil
}
