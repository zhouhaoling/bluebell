package redis

import (
	"bluebell/models"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

// GetPostVoteByID 根据帖子id查询每篇帖子的投票数
func GetPostVoteByID(postID int64) (data int64, err error) {
	key := KeyPostVotedZSetPrefix + strconv.Itoa(int(postID))
	data = rdb.ZCount(key, "1", "1").Val()
	return data, nil
}

// GetPostIDsInOrder 根据order获取帖子列表 tiem score
func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	result, err := rdb.ZRevRange(key, start, end).Result()
	//fmt.Println(result)
	return result, err
}

// CreatePost 创建帖子时间、分数
func CreatePost(postID int64) error {
	pipeline := rdb.TxPipeline()
	//帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	//帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	_, err := pipeline.Exec()
	return err
}
