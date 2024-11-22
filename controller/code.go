package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeInvalidPassword
	CodeServerBusy
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户已存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",
}

func (code ResCode) Msg() string {
	msg, ok := codeMsgMap[code]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
