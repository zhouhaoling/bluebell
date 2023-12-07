package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"fmt"
	"strconv"
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
	err = redis.CreatePost(p.ID, p.CommunityID)
	return err
}

// GetPostById 根据帖子id获取帖子数据
func GetPostById(postID int64) (data *models.ApiPostDetail, err error) {
	post, err := mysql.GetPostById(postID)
	if err != nil {
		zap.L().Error("GetPostByID failed", zap.Error(err))
		return
	}
	user, err := mysql.GetUserByID(post.AuthorID)
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
	if err != nil {
		zap.L().Error("redis.GetPostVoteByID failed", zap.Error(err))
		return
	}
	//根据帖子id查询帖子评论数
	commentNum, err := mysql.GetCommentByPostIdCount(postID)
	if err != nil {
		zap.L().Error("mysql.GetCommentByPostId", zap.Error(err))
		return
	}
	data = &models.ApiPostDetail{
		AuthorName:      user.UserName,
		Post:            post,
		CommunityDetail: community,
		VoteNum:         voteNum,
		CommentNum:      commentNum,
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
		user, err := mysql.GetUserByID(post.AuthorID)
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
		//

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
		user, err := mysql.GetUserByID(post.AuthorID)
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

// GetCommunityPostList 根据社区查询帖子列表 order = time / score expire:time
func GetCommunityPostList(p *models.ParamCommunityPostList) (datas []*models.ApiPostDetail, err error) {
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	fmt.Println("ids", ids)
	if err != nil {
		zap.L().Error("redis.GetCommunityPostIDsInOrder failed", zap.Error(err))
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder return 0 data")
		return
	}
	zap.L().Debug("GetCommunityPostIDsInOrder", zap.Any("ids", ids))
	//查询帖子的详细数据
	posts, err := mysql.GetPostListByIDs(ids)
	fmt.Println("posts:", posts)
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
		user, err := mysql.GetUserByID(post.AuthorID)
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

func PostSearch(p *models.ParamSearchList) (*models.ApiPostDetailRes, error) {
	var res models.ApiPostDetailRes
	count, err := mysql.GetPostListTotalCount(p)
	if err != nil {
		zap.L().Error("GetPostListTotalCount failed", zap.Error(err))
		return nil, err
	}
	//相关帖子的数量
	res.Page.Total = count
	posts, err := mysql.GetPostListByKeywords(p)
	if err != nil {
		zap.L().Error("GetPostListByKeywords failed", zap.Error(err))
		return nil, err
	}
	//查询结果为空返回空
	if len(posts) == 0 {
		return &models.ApiPostDetailRes{}, nil
	}
	ids := make([]string, 0, len(posts))
	for _, post := range posts {
		ids = append(ids, strconv.Itoa(int(post.ID)))
	}
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return nil, err
	}
	res.Page.Size = p.Size
	res.Page.Page = p.Page
	//拼接数据
	res.List = make([]*models.ApiPostDetail, 0, len(posts))
	for idx, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID() failed", zap.Error(err))
			user = nil
		}
		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityByID() failed", zap.Error(err))
			community = nil
		}
		//根据帖子id查询帖子评论数
		commentNum, err := mysql.GetCommentByPostIdCount(post.ID)
		if err != nil {
			zap.L().Error("mysql.GetCommentByPostId", zap.Error(err))
			return nil, err
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.UserName,
			VoteNum:         voteData[idx],
			CommentNum:      commentNum,
			Post:            post,
			CommunityDetail: community,
		}
		res.List = append(res.List, postDetail)
	}
	return &res, nil
}
