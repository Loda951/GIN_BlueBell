package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const (
	CtxUserIDKey   = "userID"
	CtxUsernameKey = "username"
)

var ErrorNotLogin = errors.New("not login")

func GetCurrentUser(ctx *gin.Context) (userID int64, err error) {
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
