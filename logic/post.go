package logic

import (
	"bluebell/DAO/mysql"
	"bluebell/models"
	snowflake "bluebell/pkg/snowFlake"
)

func CreatePost(p *models.Post) (err error) {
	// 1 生成Post ID
	p.ID = snowflake.GenID()
	// 2 保存到数据库
	return mysql.CreatePost(p)
}

func GetPostDetail(id int64) (data *models.Post, err error) {
	return mysql.GetPostDetailByID(id)
}
