package redis

//redis key

const (
	KeyPrefix              = "bluebell:"
	KeyPostTimeZSet        = "post:time"   //zset：帖子及发贴时间
	KeyPostScoreZSet       = "post:score"  //zset: 帖子及投票的分数
	KeyPostVotedZSetPrefix = "post:voted:" //zset: 记录用户及投票的类型;参数：post_id
)
