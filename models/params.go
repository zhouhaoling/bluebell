package models

const (
	Page       = 1
	Size       = 10
	OrderTime  = "time"
	OrderScore = "score"
)

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
	Page  int64  `json:"page" form:"page"`
	Size  int64  `json:"size" form:"size"`
	Order string `json:"order" form:"order"`
}
