package models

import "time"

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

// ApiPostDetail 帖子详情接口的结构体
type ApiPostDetail struct {
	AuthorName string `json:"author_name"`
	VoteNum    int64  `json:"vote_num"` //投票数量
	*Post
	*CommunityDetail `json:"community"`
}
