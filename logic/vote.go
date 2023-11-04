package logic

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"strconv"

	"go.uber.org/zap"
)

//投票功能

func VoteForPost(userID int64, p *models.ParamVoteDate) error {
	zap.L().Debug("VoteForPost", zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
