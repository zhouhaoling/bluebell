package redis

import (
	"bluebell/define"
	"bluebell/models"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func getIDsFormKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + size - 1
	//fmt.Println(result)
	return rdb.ZRevRange(key, start, end).Result()
}

// GetPostVoteByID 根据帖子id查询每篇帖子的投票数
func GetPostVoteByID(postID int64) (data int64, err error) {
	key := KeyPostVotedZSetPrefix + strconv.Itoa(int(postID))
	data = rdb.ZCount(key, "1", "1").Val()
	return data, nil
}

// GetPostIDsInOrder 根据order获取帖子列表 time score
func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	//默认是time,判断传入的order
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == define.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}

	return getIDsFormKey(key, p.Page, p.Size)
}

// GetCommunityPostIDsInOrder 按社区查询ids
func GetCommunityPostIDsInOrder(p *models.ParamCommunityPostList) ([]string, error) {
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == define.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}
	//生成分区set与帖子分数zset关联的新的zset
	cKey := getRedisKey(KeyCommunitySetPrefix + strconv.Itoa(int(p.CommunityID)))
	//利用缓存key减少zinterstore执行次数
	key := orderKey + strconv.Itoa(int(p.CommunityID))

	if rdb.Exists(key).Val() < 1 {
		//不存在，需要计算
		pipeline := rdb.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, cKey, orderKey) // zinterstore 计算
		pipeline.Expire(key, 60*time.Second)
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	fmt.Println("key:", key, ",page:", p.Page, ",size:", p.Size)

	return getIDsFormKey(key, p.Page, p.Size)
}

// CreatePost 创建帖子时间、分数
func CreatePost(postID, communityID int64) error {
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
	//帖子id加入到社区set
	cKey := getRedisKey(KeyCommunitySetPrefix + strconv.Itoa(int(communityID)))
	pipeline.SAdd(cKey, postID)
	_, err := pipeline.Exec()
	return err
}
