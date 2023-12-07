package redis

//redis key

const (
	Prefix                 = "bluebell:"
	KeyPostTimeZSet        = "post:time"   //zset：帖子及发贴时间
	KeyPostScoreZSet       = "post:score"  //zset: 帖子及投票的分数
	KeyPostVotedZSetPrefix = "post:voted:" //zset: 记录用户及投票的类型;参数：post_id
	KeyCommunitySetPrefix  = "community:"  //set;保存每个分区下帖子的id
)

func getRedisKey(key string) string {
	return Prefix + key
}
