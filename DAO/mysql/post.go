package mysql

import (
	"bluebell/models"
	"database/sql"
	"errors"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post (post_id, title, content, author_id, community_id) values(?,?,?,?,?)`

	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)

	return
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
