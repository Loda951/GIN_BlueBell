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
	like             = 1
	dislike          = -1
	noAction         = 0
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

func PostVote(userID, postID string, nextActionValue float64) error {
	// 1.判断投票限制
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrorVoteTimeout
	}
	// 2.更新分数
	// 先查之前 当前用户 是否对文章表过态(没表态0 点赞1 踩-1)
	// KeyPostVotedZSetPrefix+postID 为 key(也就是ZSet的名称) userID是 member .Val()可以返回 score (就是 1 0 -1)
	// 维护了一个对当前post的map 区分post用postID 里面k为userID是 v为 1 0 -1
	preActionValue := client.ZScore(getRedisKey(KeyPostVotedZSetPrefix+postID), userID).Val()
	// 只允许用户投一次 1 0 -1 三选一
	if preActionValue == nextActionValue {
		return ErrorVoteRepeated
	}
	var dir float64
	// 接下来的 action 可能是(没表态0 点赞1 踩-1)
	if nextActionValue > preActionValue {
		dir = like
	} else {
		dir = dislike
	}
	diff := math.Abs(preActionValue - nextActionValue) // 这里 pre - next or next - pre 都可以
	// 使用管道减少 RTT（Round-Trip Time）
	pipeline := client.Pipeline()
	// client.ZIncrBy(key, increment, member) 让分数增加
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), dir*diff*scorePerVote, postID)
	if nextActionValue == noAction {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPrefix+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPrefix+postID), redis.Z{
			Member: userID,
			Score:  nextActionValue,
		})
	}
	// 执行管道中的所有命令
	_, err := pipeline.Exec()
	return err
}
