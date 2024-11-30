package logic

import (
	"bluebell/DAO/redis"
	"bluebell/models"
	"strconv"
)

func PostVote(userID int64, p *models.ParamVoteData) error {
	return redis.PostVote(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
