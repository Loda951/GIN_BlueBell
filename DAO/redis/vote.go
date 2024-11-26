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
	ErrorVoteTimeout  = errors.New("vote timeout")
	ErrorVoteRepeated = errors.New("repeated vote are not allowed")
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

func CreatePost(postID int64) error {
	// 创建一个管道
	pipeline := client.Pipeline()

	// 帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 执行管道中的所有命令
	_, err := pipeline.Exec()
	return err
}

func PostVote(userID, postID string, nextActionValue float64) error {
	// 1.判断投票限制
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrorVoteTimeout
	}
	// 2.更新分数
	// 先查之前 当前用户 是否对文章表过态(没表态0 点赞1 踩-1)
	preActionValue := client.ZScore(getRedisKey(KeyPostVotedZSetPrefix+postID), userID).Val()
	if preActionValue == nextActionValue {
		return ErrorVoteRepeated
	}
	var dir float64
	// 接下来的 action 可能是(没表态0 点赞1 踩-1)
	if nextActionValue > preActionValue {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(preActionValue - nextActionValue) // 这里 pre - next or next - pre 都可以

	pipeline := client.Pipeline()

	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), dir*diff*scorePerVote, postID)

	if nextActionValue == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPrefix+postID), userID)

	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPrefix+postID), redis.Z{
			Score:  nextActionValue,
			Member: userID,
		})
	}
	// 执行管道中的所有命令
	_, err := pipeline.Exec()
	return err
}
