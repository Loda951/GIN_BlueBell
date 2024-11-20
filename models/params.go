package models

// 定义请求的参数结构体
type ParamsSignUp struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RePasswd string `json:"re_password"`
}
