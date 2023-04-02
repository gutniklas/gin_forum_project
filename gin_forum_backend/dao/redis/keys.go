package redis

const (
	Prefix           = "gin_forum:"
	KeyPostTimeZSet  = "post:time"  // zset;发帖贴子
	KeyPostScoreZSet = "post:score" // zset;贴子投票的分数
	KeyPostVotedZSet = "post:voted" // zset;记录投票过的用户
)

func getRedisKey(key string) string {
	return Prefix + key
}
