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
	sqlStr := "select post_id, title, content, author_id, community_id, create_time from post where id = ?"
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

	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	db.Select(&postsList, query, args...) // !!!
	return
}
