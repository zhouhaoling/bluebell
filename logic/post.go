package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"time"

	"go.uber.org/zap"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	p.ID = snowflake.GenID()
	p.CreateTime = time.Now()
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.ID)
	return err
}

// GetPostById 根据帖子id获取帖子数据
func GetPostById(postID int64) (data *models.ApiPostDetail, err error) {
	post, err := mysql.GetPostById(postID)
	if err != nil {
		zap.L().Error("GetPostByID failed", zap.Error(err))
		return
	}
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("GetUserById failed", zap.Error(err))
		return
	}
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("GetCommunityDetailByID failed", zap.Error(err))
		return
	}
	//根据帖子id查询帖子投票数
	voteNum, err := redis.GetPostVoteByID(postID)
	data = &models.ApiPostDetail{
		AuthorName:      user.UserName,
		Post:            post,
		CommunityDetail: community,
		VoteNum:         voteNum,
	}

	return
}

// GetPostList 获取帖子列表
func GetPostList(page, size int) (datas []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		zap.L().Error("mysql.GetPostList() failed")
		return nil, err
	}
	datas = make([]*models.ApiPostDetail, 0, len(posts))

	for _, post := range posts {
		//根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("GetUserById failed", zap.Error(err))
			continue
		}
		//根据社区id查询社区信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("GetCommunityDetailByID failed", zap.Error(err))
			continue
		}
		//接口数据拼接
		dataDetail := &models.ApiPostDetail{
			AuthorName:      user.UserName,
			Post:            post,
			CommunityDetail: community,
		}
		datas = append(datas, dataDetail)
	}
	return datas, nil
}

// GetPostList2 获取帖子列表 根据score time
func GetPostList2(p *models.ParamPostList) (datas []*models.ApiPostDetail, err error) {
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder return 0 data")
		return
	}
	//查询帖子的详细数据
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return nil, err
	}
	//根据帖子id统计帖子投票数
	voteDatas, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}
	//查询其他数据，补充帖子
	for index, post := range posts {
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("GetUserById failed", zap.Error(err))
			return nil, err
		}
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("GetCommunityDetailByID failed", zap.Error(err))
			return nil, err
		}
		//根据帖子id查询帖子投票数
		//voteNum, err := redis.GetPostVoteByID(post.ID)
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.UserName,
			Post:            post,
			CommunityDetail: community,
			VoteNum:         voteDatas[index],
		}
		datas = append(datas, postDetail)
	}
	return
}
