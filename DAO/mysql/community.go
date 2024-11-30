package mysql

import (
	"bluebell/models"
	"database/sql"
	"errors"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			zap.L().Warn("no community list in DB")
			return nil, nil
		}
	}
	return communityList, nil
}

func GetCommunityDetailByID(id int64) (detail *models.CommunityDetail, err error) {
	detail = new(models.CommunityDetail)
	sqlStr := "select community_id, community_name, introduction, create_time from community where community_id = ?"
	if err := db.Get(detail, sqlStr, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			zap.L().Warn("no community detail in DB")
			err = ErrorInvalidID
			return nil, err
		}
		// 其他错误处理
		zap.L().Error("failed to query community detail", zap.Error(err))
		return nil, err
	}
	return detail, nil
}
