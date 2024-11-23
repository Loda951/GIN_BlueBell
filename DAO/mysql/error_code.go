package mysql

import "errors"

// 错误变量
var (
	ErrorUserExists      = errors.New("user already exists")
	ErrorUserNotFound    = errors.New("user not found")
	ErrorInvalidPassword = errors.New("invalid password")
	ErrorInvalidID       = errors.New("invalid ID")
)
