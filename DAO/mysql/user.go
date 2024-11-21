package mysql

import (
	"bluebell/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// 把每一个数据库操作封装成一个函数
// 待logic层根据业务需求调用

// CheckUserExist 检查指定用户名称的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username=?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return nil
}

// InsertUser 向数据库中插入一条指定user的记录
func InsertUser(user *models.User) (err error) {
	// 为密码加密
	user.Password = encryptPassword(user.Password)
	// 执行SQL入库
	sqlStr := `insert into user (user_id, username, password) values (?, ?, ?)`
	if _, err := db.Exec(sqlStr, user.UserID, user.Username, user.Password); err != nil {
		return err
	}
	return nil
}

func encryptPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hash)
}
