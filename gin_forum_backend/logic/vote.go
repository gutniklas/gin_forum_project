package logic

import (
	"gin_forum/models"
	"strconv"

	"gin_forum/dao/redis"
)

func VoteForPost(userID int64, p *models.ParamVoteData) (float64, error) {

	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID)

}
