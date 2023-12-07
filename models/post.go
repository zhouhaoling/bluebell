package models

import (
	"encoding/json"
	"errors"
	"time"
)

// Post 帖子结构体
type Post struct {
	ID          int64     `json:"id,string" gorm:"column:post_id"`
	AuthorID    int64     `json:"author_id" gorm:"column:author_id"`
	CommunityID int64     `json:"community_id" gorm:"column:community_id" binding:"required"`
	Status      int32     `json:"status" gorm:"column:status"`
	Title       string    `json:"title" gorm:"column:title" binding:"required"`
	Content     string    `json:"content" gorm:"column:content" binding:"required"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`
}

// UnmarshalJSON 为Post类型实现自定义的UnmarshalJSON方法
func (p *Post) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		Title       string `json:"title" db:"title"`
		Content     string `json:"content" db:"content"`
		CommunityID int64  `json:"community_id" db:"community_id"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.Title) == 0 {
		err = errors.New("帖子标题不能为空")
	} else if len(required.Content) == 0 {
		err = errors.New("帖子内容不能为空")
	} else if required.CommunityID == 0 {
		err = errors.New("未指定版块")
	} else {
		p.Title = required.Title
		p.Content = required.Content
		p.CommunityID = required.CommunityID
	}
	return
}

// ApiPostDetail 帖子详情接口的结构体
type ApiPostDetail struct {
	AuthorName string `json:"author_name"` //作者名
	VoteNum    int64  `json:"vote_num"`    //投票数量
	CommentNum int64  `json:"comment_num"` //评论数量
	*Post
	*CommunityDetail `json:"community"`
}

type Page struct {
	Total int64 `json:"total"`
	Page  int64 `json:"page"`
	Size  int64 `json:"size"`
}

type ApiPostDetailRes struct {
	Page Page             `json:"page"`
	List []*ApiPostDetail `json:"list"`
}
