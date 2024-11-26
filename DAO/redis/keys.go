package redis

// 说白了key就是这些map的名称
const (
	KeyPrefix              = "bluebell:"
	KeyPostTimeZSet        = "post:time"   // zset; 帖子及发帖时间
	KeyPostScoreZSet       = "post:score"  // zset; 帖子及投票分数
	KeyPostVotedZSetPrefix = "post:voted:" // zset; 记录用户及投票是0 1 or -1; 参数是post id(会拼接postID组成map名字)

	KeyCommunitySetPrefix = "community:" // set; 保存每个分区下帖子id set 只有key和member 没有 score
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
