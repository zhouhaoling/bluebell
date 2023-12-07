package controller

import (
	"bluebell/logic"
	"bluebell/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CommentHandler 创建评论
func CommentHandler(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		zap.L().Error("ShouldBindJSON comment failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID failed", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}
	comment.AuthorID = userID

	err = logic.CreateComment(&comment)
	if err != nil {
		zap.L().Error("CreateComment failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// CommentListHandler 评论列表，根据帖子id查询该帖子下的评论列表
func CommentListHandler(c *gin.Context) {
	var p models.ParamCommentList
	if err := c.ShouldBindQuery(&p); err != nil {
		zap.L().Error("ShouldBindQuery ParamCommentList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	commentList, err := logic.GetCommentList(p.PostID)
	if err != nil {
		zap.L().Error("GetCommentList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, commentList)
}
