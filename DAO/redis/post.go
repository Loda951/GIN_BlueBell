package redis

import (
	"bluebell/models"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func CreatePost(postID, communityID int64) error {
	// 创建一个管道
	pipeline := client.Pipeline()

	// 帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Member: postID,
		Score:  float64(time.Now().Unix()),
	})

	// 帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Member: postID,
		Score:  float64(time.Now().Unix()),
	})

	cKey := getRedisKey(KeyCommunitySetPrefix + strconv.Itoa(int(communityID)))
	// 对于set而言 只有 key 和 member 也就是map的名称和k 没有v
	pipeline.SAdd(cKey, postID)

	// 执行管道中的所有命令
	_, err := pipeline.Exec()
	return err
}

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	return getIDsFromKey(key, p.Page, p.Size)
}

func GetPostVoteData(ids []string) (data []int64, err error) {
	pipeline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPrefix + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(ids))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}

func GetCommunityPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}
	// 使用zinterstore 把分区的帖子set 与 按分数的zset生成一个新的zset 针对新的zset 按之前逻辑取出
	// 社区的key
	cKey := getRedisKey(KeyCommunitySetPrefix + strconv.Itoa(int(p.CommunityID)))
	// 利用缓存key 减少zinterstore执行的次数
	key := orderKey + strconv.Itoa(int(p.CommunityID))
	if client.Exists(key).Val() < 1 {
		// 不存在 需要计算
		pipeline := client.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, cKey, orderKey)
		pipeline.Expire(key, 60*time.Second) //设置超时时间
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	return getIDsFromKey(key, p.Page, p.Size)
}

func getIDsFromKey(key string, page, size int64) ([]string, error) {
	// 索引从0开始 -1 减去的基本都是0这个考虑
	start := (page - 1) * size
	end := start + size - 1
	// zrevrange 查询 按分数从大到小
	// zrange 从小到大
	// 根据ZSet中的 score 值，从小到大排序后返回对应的 member 列表
	return client.ZRevRange(key, start, end).Result()
}
