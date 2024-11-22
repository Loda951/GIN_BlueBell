package models

// ParamsSignUp 定义注册请求的参数结构体
type ParamsSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamsLogIn 定义登陆请求的参数结构体
type ParamsLogIn struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
