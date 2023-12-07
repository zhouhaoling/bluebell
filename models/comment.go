package models

import (
	"encoding/json"
	"errors"
	"time"
)

type Comment struct {
	ParentID   int64     `json:"parent_id,string" gorm:"column:parent_id"` //评论的父id
	CommentID  int64     `json:"comment_id" gorm:"column:comment_id"`      //评论id
	AuthorID   int64     `json:"author_id" gorm:"column:author_id"`        //作者id
	PostID     int64     `json:"post_id,string" gorm:"column:post_id"`     //帖子id
	Content    string    `json:"content"`                                  //内容
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`    //创建时间
}

func (c *Comment) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		PostID  int64  `json:"post_id,string" gorm:"column:post_id"`
		Content string `json:"content"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.Content) == 0 {
		err = errors.New("评论内容不能为空")
	} else if required.PostID == 0 {
		err = errors.New("帖子id不能为空")
	} else {
		c.PostID = required.PostID
		c.Content = required.Content
	}
	return
}

type ApiCommentDetail struct {
	CommentID  int64
	PostID     int64
	AuthorName string
	Content    string
	CreateTime time.Time
}
