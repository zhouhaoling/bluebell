package models

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	UserName   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVoteDate 投票数据
type ParamVoteDate struct {
	PostID    string `json:"post_id" binding:"required"`       //帖子id
	Direction int    `json:"direction" binding:"oneof=1 0 -1"` //点赞 中立 反对 （1 ，0 ,-1）
}

// ParamPostList 获取帖子列表参数
type ParamPostList struct {
	Page  int64  `json:"page" form:"page"`   //页码
	Size  int64  `json:"size" form:"size"`   //每页数量
	Order string `json:"order" form:"order"` //排序依据,example:time
}

type ParamSearchList struct {
	ParamPostList
	Search string `json:"search" form:"search"` //关键字搜索
}

// ParamCommunityPostList 根据社区获取帖子参数
type ParamCommunityPostList struct {
	ParamPostList
	CommunityID int64 `json:"community_id" form:"community_id"`
}

type ParamCommentList struct {
	PostID int64 `json:"post_id" form:"post_id"`
}
