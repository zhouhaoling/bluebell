package controller

import (
	"bluebell/define"
	"bluebell/logic"
	"bluebell/models"
	"fmt"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CreatePostHandler 创建帖子
func CreatePostHandler(c *gin.Context) {
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		//zap.L().Debug("c.ShouldBindJSON(p) error", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 从c取到当前用户id
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	//2.创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回响应-
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 根据id查询帖子详情
func GetPostDetailHandler(c *gin.Context) {
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostByID is failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// GetPostListHandler 分页获取帖子列表
func GetPostListHandler(c *gin.Context) {
	page, size := GetPageInfo(c)
	datas, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, datas)
}

// GetPostListHandler2 帖子列表接口2 按创建时间或分数排序
func GetPostListHandler2(c *gin.Context) {
	p := &models.ParamPostList{
		Page:  define.Page,
		Size:  define.Size,
		Order: define.OrderTime,
	}
	err := c.ShouldBindQuery(p)
	if err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	datas, err := logic.GetPostList2(p)
	if err != nil {
		zap.L().Error("logic.GetPostList2 failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, datas)
}

// GetCommunityPostListHandler 根据社区查询帖子列表
func GetCommunityPostListHandler(c *gin.Context) {
	p := &models.ParamCommunityPostList{
		ParamPostList: models.ParamPostList{
			Page:  define.Page,
			Size:  define.Size,
			Order: define.OrderTime,
		},
	}
	err := c.ShouldBindQuery(p)
	if err != nil {
		zap.L().Error("GetCommunityPostListHandler with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	datas, err := logic.GetCommunityPostList(p)
	if err != nil {
		zap.L().Error("logic.GetCommunityPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, datas)
}

func PostSearchHandler(c *gin.Context) {
	p := &models.ParamSearchList{}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("ParamSearchList with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	fmt.Println("Search", p.Search)
	fmt.Println("Order", p.Order)
	data, err := logic.PostSearch(p)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
