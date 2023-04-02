package redis

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

const (
	threeMonthInSeconds = 90 * 86400
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
)

func CreatePost(postID int64) error {
	pipeline := client.TxPipeline()

	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  0,
		Member: postID,
	})
	_, err := pipeline.Exec()
	return err
}
func GetPostLikeNum(postID string) float64 {
	PostScore := client.ZScore(getRedisKey(KeyPostScoreZSet), postID).Val()

	
	return PostScore
}
func VoteForPost(userID, postID string) (float64, error) {

	//1. 判断投票时间限制
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > threeMonthInSeconds {
		return -1, ErrVoteTimeExpire
	}
	//2. 更新帖子的分数
	_, err := client.ZScore(getRedisKey(KeyPostVotedZSet+postID), userID).Result()

	pipeline := client.TxPipeline()
	if err == redis.Nil {

		pipeline.ZAdd(getRedisKey(KeyPostVotedZSet+postID), redis.Z{
			Score:  1,
			Member: userID,
		})
		pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), 1, postID)

	} else if err != nil {

		panic(err)

	} else {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSet+postID), userID)
		pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), -1, postID)

	}
	_, err = pipeline.Exec()
	PostScore := client.ZScore(getRedisKey(KeyPostScoreZSet), postID).Val()


	return PostScore, err

}
