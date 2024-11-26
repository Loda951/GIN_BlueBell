package logic

import (
	"bluebell/DAO/mysql"
	"bluebell/DAO/redis"
	"bluebell/models"
	snowflake "bluebell/pkg/snowFlake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	// 1 生成Post ID
	p.ID = snowflake.GenID()
	// 2 保存到数据库
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.ID, p.CommunityID)
	if err != nil {
		return err
	}
	return nil
}

func GetPostDetail(id int64) (data *models.APIPostDetial, err error) {
	data = new(models.APIPostDetial)
	post, err := mysql.GetPostDetailByID(id)
	if err != nil {
		zap.L().Error("mysql.GetPostDetailByID()", zap.Error(err))
		return nil, err
	}

	// 根据作者id查询作者信息
	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID()", zap.Error(err))
		return nil, err
	}

	// 根据社区id查询社区信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID()", zap.Error(err))
		return nil, err
	}

	data = &models.APIPostDetial{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return
}

func GetPostList(page, size int64) (data []*models.APIPostDetial, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		zap.L().Error("mysql.GetPostList()", zap.Error(err))
		return nil, err
	}
	data = make([]*models.APIPostDetial, 0, len(posts))
	for _, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID()", zap.Error(err))
			return nil, err
		}

		// 根据社区id查询社区信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID()", zap.Error(err))
			return nil, err
		}
		postDetail := &models.APIPostDetial{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

func GetPostList2(p *models.ParamPostList) (data []*models.APIPostDetial, err error) {
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder get 0 data")
		return nil, nil
	}
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		zap.L().Error("mysql.GetPostList()", zap.Error(err))
		return nil, err
	}
	// 获得投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return nil, err
	}

	data = make([]*models.APIPostDetial, 0, len(posts))
	for index, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID()", zap.Error(err))
			return nil, err
		}
		// 根据社区id查询社区信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID()", zap.Error(err))
			return nil, err
		}
		postDetail := &models.APIPostDetial{
			AuthorName:      user.Username,
			VotesNumber:     voteData[index],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

func GetCommunityPostList(p *models.ParamPostList) (data []*models.APIPostDetial, err error) {
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder get 0 data")
		return nil, nil
	}
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		zap.L().Error("mysql.GetPostList()", zap.Error(err))
		return nil, err
	}
	// 获得投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return nil, err
	}

	data = make([]*models.APIPostDetial, 0, len(posts))
	for index, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID()", zap.Error(err))
			return nil, err
		}
		// 根据社区id查询社区信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID()", zap.Error(err))
			return nil, err
		}
		postDetail := &models.APIPostDetial{
			AuthorName:      user.Username,
			VotesNumber:     voteData[index],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

func GetPostListNew(p *models.ParamPostList) (data []*models.APIPostDetial, err error) {
	if p.CommunityID == 0 {
		data, err = GetPostList2(p)
	} else {
		data, err = GetCommunityPostList(p)
	}
	if err != nil {
		zap.L().Error("logic.GetPostListNew()", zap.Error(err))
		return nil, err
	}
	return
}
