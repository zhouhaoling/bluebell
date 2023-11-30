package redis

import "errors"

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrorVoted        = errors.New("已经投过票了")
	ErrVoteRepeated   = errors.New("不允许重复投票")
)
