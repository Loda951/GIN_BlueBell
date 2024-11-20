package logic

import (
	"bluebell/DAO/mysql"
	"bluebell/models"
	snowflake "bluebell/pkg/snowFlake"
)

// 存放业务逻辑的代码

func SignUp(p *models.ParamsSignUp) {
	// 检查用户是否存在
	mysql.QueryUserByUserName()
	// 生成UID
	snowflake.GenID()
	// 保存到数据库
	mysql.InsertUser()
}
