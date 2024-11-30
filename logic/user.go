package logic

import (
	"bluebell/DAO/mysql"
	"bluebell/models"
	"bluebell/pkg/JWT"
	snowflake "bluebell/pkg/snowFlake"
)

// 存放业务逻辑的代码

func SignUp(p *models.ParamsSignUp) (err error) {
	// 1 检查用户是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 2 生成UID
	userID := snowflake.GenID()
	// 构造一个user实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 3 保存到数据库
	return mysql.InsertUser(user)
}

func LogIn(p *models.ParamsLogIn) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.LogIn(user); err != nil {
		return nil, err
	}
	// 生成 JWT
	token, err := JWT.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
