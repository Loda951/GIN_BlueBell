package mysql

import (
	"bluebell/models"
	"database/sql"
	"errors"
	"strings"

	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post (post_id, title, content, author_id, community_id) values(?,?,?,?,?)`

	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)

	return
}

func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select 
    post_id, title, content, author_id, community_id, create_time 
	from post
	ORDER BY create_time
	DESC 
    limit ?, ?`

	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	if err != nil {
		return nil, err
	}
	return posts, err
}

func GetPostDetailByID(id int64) (detail *models.Post, err error) {
	detail = new(models.Post)
	sqlStr := "select post_id, title, content, author_id, community_id, create_time from post where post_id = ?"
	if err := db.Get(detail, sqlStr, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			zap.L().Warn("no post detail in DB")
			err = ErrorInvalidID
			return nil, err
		}
		// 其他错误处理
		zap.L().Error("failed to query post detail", zap.Error(err))
		return nil, err
	}
	return detail, nil
}

func GetPostListByIDs(ids []string) (postsList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	where post_id in (?)
	order by FIND_IN_SET(post_id, ?)
	`
	// 利用 sqlx.In 函数 动态生成一条 SQL 查询语句 解决 SQL 中动态 IN 子句
	// strings.Join(ids, ",") 可以把 ids := []string{"3", "1", "2", "4"} 这样的string切片 变成 "3,1,2" 单个的拼接字符串
	// args := []interface{}{"3", "1", "2", "3,1,2"}
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	// db.Rebind(query) 会根据数据库驱动自动将占位符调整为正确的格式
	query = db.Rebind(query)
	// db.Select(&postsList, query, "3", "1", "2", "3,1,2") 展开... 不展开 args := []interface{}{"3", "1", "2", "3,1,2"}
	db.Select(&postsList, query, args...) // !!! 这里会报错，因为 db.Select 并没有接受 []interface{} 类型切片作为单个参数的能力。
	return
}
