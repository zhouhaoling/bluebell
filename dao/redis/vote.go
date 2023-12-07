package redis

import (
	"bluebell/define"
	"math"
	"time"

	"github.com/go-redis/redis"
)

// VoteForPost 投票
func VoteForPost(userID, postID string, direction float64) (err error) {
	// dir 方向（增加分数还是减少分数)
	var dir float64
	//1.判断投票限制
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > define.OneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	//2.更新帖子分数
	//检查当前用户给当前用户的投票记录
	key := getRedisKey(KeyPostVotedZSetPrefix + postID)
	uv := rdb.ZScore(key, userID).Val()
	//判断用户的操作
	if direction == uv {
		return ErrVoteRepeated
	}
	if direction > uv {
		dir = 1
	} else {
		dir = -1
	}

	diff := math.Abs(uv - direction) //计算两次投票的差值
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), dir*diff*define.ScorePerVote, postID) //更新分数

	//记录用户为该帖子投票的数据
	if direction == 0 {
		pipeline.ZRem(key, userID)
	} else {
		pipeline.ZAdd(key, redis.Z{
			Score:  direction,
			Member: userID,
		})
	}
	//pipeline.HIncrBy(K)
	_, err = pipeline.Exec()
	return err
}

// GetPostVoteData 根据id列表统计票数
func GetPostVoteData(ids []string) (data []int64, err error) {
	pipeline := rdb.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPrefix + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}
