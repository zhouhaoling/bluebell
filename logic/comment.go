package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"time"

	"go.uber.org/zap"
)

// CreateComment 创建评论逻辑处理
func CreateComment(comment *models.Comment) error {

	commentID := snowflake.GenID()
	comment.CommentID = commentID
	comment.CreateTime = time.Now()
	err := mysql.CreateComment(comment)

	return err
}

func GetCommentList(postID int64) ([]*models.ApiCommentDetail, error) {
	//从mysql中查询出评论信息
	comments, err := mysql.GetCommentByPostID(postID)
	if err != nil {
		zap.L().Error("GetCommentByPostID failed", zap.Error(err))
		return nil, err
	}
	datas := make([]*models.ApiCommentDetail, 0, len(comments))
	for i, comment := range comments {
		user, err := mysql.GetUserByID(comment.AuthorID)
		if err != nil {
			zap.L().Error("GetCommentByPostID failed", zap.Error(err))
			return nil, err
		}
		dataDetail := &models.ApiCommentDetail{
			AuthorName: user.UserName,
			PostID:     postID,
			Content:    comments[i].Content,
			CreateTime: comments[i].CreateTime,
			CommentID:  comments[i].CommentID,
		}
		datas = append(datas, dataDetail)
	}
	return datas, nil
}
