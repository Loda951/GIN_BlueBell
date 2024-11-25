package redis

import (
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis"
)

const (
	oneWeekInSeconds = 7 * 24 * 60 * 60
	scorePerVote     = 432
)

var (
	ErrorVoteTimeout = errors.New("vote timeout")
)

/*
点赞
1.取消 -perScore
2.点踩 -2 * perScore

没点
1.点赞 + perScore
2.踩 - perScore

踩
1. 点赞 2 * perScore
2. 取消 + perScore

*/

func PostVote(userID, postID string, nextActionValue float64) error {
	// 1.判断投票限制
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrorVoteTimeout
	}
	// 2.更新分数
	// 先查之前 当前用户 是否对文章表过态(没表态0 点赞1 踩-1)
	preActionValue := client.ZScore(getRedisKey(KeyPostVotedZSetPrefix+postID), userID).Val()
	var dir float64
	// 接下来的 action 可能是(没表态0 点赞1 踩-1)
	if nextActionValue > preActionValue {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(preActionValue - nextActionValue) // 这里 pre - next or next - pre 都可以
	_, err := client.ZIncrBy(getRedisKey(KeyPostScoreZSet), dir*diff*scorePerVote, postID).Result()
	if err != nil {
		return err
	}
	if nextActionValue == 0 {
		_, err := client.ZRem(getRedisKey(KeyPostVotedZSetPrefix+postID), userID).Result()
		if err != nil {
			return err
		}
	} else {
		_, err = client.ZAdd(getRedisKey(KeyPostVotedZSetPrefix+postID), redis.Z{
			Score:  nextActionValue,
			Member: userID,
		}).Result()
	}
	return err
}
